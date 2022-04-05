// Copyright 2016-2022, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package googleclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	"golang.org/x/oauth2"
	"google.golang.org/api/googleapi"
)

// Config is a bag of authentication related options used to instantiate a GoogleClient
type Config struct {
	Credentials                        string
	AccessToken                        string
	ImpersonateServiceAccount          string
	ImpersonateServiceAccountDelegates []string
	Scopes                             []string
	PulumiVersion                      string
	ProviderVersion                    string
	PartnerName                        string
	AppendUserAgent                    string
}

// GoogleClient is an appropriately initialized HTTP client used by the Google Native provider.
// The client checks if its credentials are valid and automatically refreshes if they have expired.
type GoogleClient struct {
	config        Config
	http          *http.Client
	credRetriever credentialRetriever
	credValidator credentialValidator
	userAgent     string
}

// New returns a new GoogleClient with credentials based on the provided Config.
func New(ctx context.Context, config Config) (*GoogleClient, error) {
	partnerString := ""
	if config.PartnerName != "" {
		partnerString = fmt.Sprintf("GPN:%s; ", config.PartnerName)
	}

	userAgent := fmt.Sprintf("Pulumi/%s (%shttps://www.pulumi.com) pulumi-google-native/%s",
		config.PulumiVersion, partnerString, config.ProviderVersion)

	if config.AppendUserAgent != "" {
		userAgent = fmt.Sprintf("%s %s", userAgent, config.AppendUserAgent)
	}

	googleClient := &GoogleClient{config: config, userAgent: userAgent, credRetriever: defaultCredentialRetriever()}
	err := googleClient.refreshClientCredentials(ctx)
	if err != nil {
		return nil, err
	}

	return googleClient, nil
}

// HTTPClient returns an initialized HTTP client for Google Cloud.
func (c *GoogleClient) HTTPClient() *http.Client { return c.http }

// RequestWithTimeout performs the specified request using the specified HTTP method and with the specified timeout.
// TODO: This is taken from the TF provider (cut down to a minimal viable thing). We need to make it "good".
func (c *GoogleClient) RequestWithTimeout(
	method, rawurl string,
	body map[string]interface{},
	timeout time.Duration,
) (map[string]interface{}, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.userAgent)
	reqHeaders.Set("Content-Type", "application/json")

	if timeout == 0 {
		timeout = time.Duration(1) * time.Hour
	}

	var res *http.Response
	var buf bytes.Buffer
	if body != nil {
		err := json.NewEncoder(&buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	if err := c.refreshClientCredentials(context.Background()); err != nil {
		return nil, err
	}

	u, err := addQueryParams(rawurl, map[string]string{"alt": "json"})
	if err != nil {
		return nil, err
	}
	// TODO: request not handling timeout
	req, err := http.NewRequest(method, u, &buf)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders

	debugBody, _ := httputil.DumpRequestOut(req, true)
	logging.V(9).Infof(requestFormat, req.Method, req.URL, prettyPrintJSONLines(debugBody))

	res, err = c.http.Do(req)
	if err != nil {
		return nil, err
	}

	debugResp, _ := httputil.DumpResponse(res, true)
	logging.V(9).Infof(responseFormat, res.Request.Method, res.Request.URL, res.ContentLength, debugResp)

	if err := googleapi.CheckResponse(res); err != nil {
		googleapi.CloseBody(res)
		return nil, err
	}

	if res == nil {
		return nil, fmt.Errorf("unable to parse server response")
	}

	// The defer call must be made outside of the retryFunc otherwise it's closed too soon.
	defer googleapi.CloseBody(res)

	// 204 responses will have no body, so we're going to error with "EOF" if we
	// try to parse it. Instead, we can just return nil.
	if res.StatusCode == 204 {
		return nil, nil
	}
	result := make(map[string]interface{})
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

// multipartBoundary is a random string used to separate parts of multi-part request bodies.
const multipartBoundary = "boundary-fa78ad331d"

// UploadWithTimeout performs a multi-part upload using the specified HTTP method, rawurl, etc. using the specified
// timeout.
func (c *GoogleClient) UploadWithTimeout(
	method, rawurl string,
	metadata map[string]interface{},
	binary []byte,
	timeout time.Duration,
) (map[string]interface{}, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.userAgent)
	reqHeaders.Set("Content-Type", fmt.Sprintf("multipart/related; boundary=%s", multipartBoundary))

	if timeout == 0 {
		timeout = time.Duration(1) * time.Hour
	}

	// See the docs on multipart uploads: https://cloud.google.com/storage/docs/uploading-objects
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("--%s\r\n", multipartBoundary))
	buf.WriteString("Content-Type: application/json; charset=UTF-8\r\n\r\n")
	err := json.NewEncoder(&buf).Encode(metadata)
	if err != nil {
		return nil, err
	}
	buf.WriteString(fmt.Sprintf("\r\n--%s\r\n", multipartBoundary))
	contentType := "application/octet-stream"
	if ct, has := metadata["contentType"].(string); has {
		contentType = ct
	}
	buf.WriteString(fmt.Sprintf("Content-Type: %s\r\n\r\n", contentType))
	buf.Write(binary)
	buf.WriteString(fmt.Sprintf("\r\n--%s--\r\n", multipartBoundary))

	if err := c.refreshClientCredentials(context.Background()); err != nil {
		return nil, err
	}

	u, err := addQueryParams(rawurl, map[string]string{"alt": "json", "uploadType": "multipart"})
	if err != nil {
		return nil, err
	}

	// TODO: request not handling timeout
	req, err := http.NewRequest(method, u, &buf)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders

	debugBody, _ := httputil.DumpRequest(req, true)
	logging.V(9).Infof(requestFormat, req.Method, req.URL, prettyPrintJSONLines(debugBody))

	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	debugResp, _ := httputil.DumpResponse(res, true)
	logging.V(9).Infof(responseFormat, res.Request.Method, res.Request.URL, res.ContentLength,
		debugResp)

	if err := googleapi.CheckResponse(res); err != nil {
		googleapi.CloseBody(res)
		return nil, err
	}

	if res == nil {
		return nil, fmt.Errorf("unable to parse server response")
	}

	// The defer call must be made outside of the retryFunc otherwise it's closed too soon.
	defer googleapi.CloseBody(res)

	// 204 responses will have no body, so we're going to error with "EOF" if we
	// try to parse it. Instead, we can just return nil.
	if res.StatusCode == 204 {
		return nil, nil
	}
	result := make(map[string]interface{})
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

var defaultClientScopes = []string{
	"https://www.googleapis.com/auth/compute",
	"https://www.googleapis.com/auth/cloud-platform",
	"https://www.googleapis.com/auth/cloud-identity",
	"https://www.googleapis.com/auth/ndev.clouddns.readwrite",
	"https://www.googleapis.com/auth/devstorage.full_control",
	"https://www.googleapis.com/auth/userinfo.email",
}

// refreshClientCredentials first checks if the current active credentials in use are
// valid. If so it simply reuses the existing HTTP client.
// Otherwise it refreshes the credentials using the Config used to create the GoogleClient.
// It then caches the credentials and the client for subsequent calls.
func (c *GoogleClient) refreshClientCredentials(ctx context.Context) error {
	if c.credValidator != nil {
		valid, err := c.credValidator.isValid()
		if err != nil {
			return err
		}
		if valid {
			return nil
		}
	}

	if len(c.config.Scopes) == 0 {
		c.config.Scopes = defaultClientScopes
	}

	credentials, err := c.credRetriever.getCredentials(ctx, c.config)
	if err != nil {
		return errors.Wrapf(err, "getting credentials")
	}

	ts := newReuseValidatedTokenSource(nil, credentials.TokenSource)
	c.credValidator = credentialValidatorFunc(func() (bool, error) {
		t, err := ts.Token()
		if err != nil {
			return false, err
		}
		return t.Valid(), nil
	})

	cleanCtx := context.WithValue(ctx, oauth2.HTTPClient, cleanhttp.DefaultClient())

	// OAUTH2 TRANSPORT/CLIENT - sets up proper auth headers
	client := oauth2.NewClient(cleanCtx, ts)

	// Set final transport value.
	// This timeout is a timeout per HTTP request, not per logical operation.
	client.Timeout = 30 * time.Second
	currTransport := client.Transport
	if currTransport == nil {
		currTransport = http.DefaultTransport
	}
	client.Transport = NewTransportWithDefaultRetries(currTransport).WithAddedPredicates(
		/* Adding some known retryable 409 errors */
		isSqlOperationInProgressError,
		isMonitoringConcurrentEditError,
		isAppEngineRetryableError,
		datastoreIndex409Contention,
		iapClient409Operation,
		isCloudRunCreationConflict)
	c.http = client
	return nil
}

func addQueryParams(rawurl string, params map[string]string) (string, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}

// If the argument is a path, pathOrContents loads it and returns the contents,
// otherwise the argument is assumed to be the desired contents and is simply
// returned.
//
// The boolean second return value can be called `wasPath` - it indicates if a
// path was detected and a file loaded.
func pathOrContents(poc string) (string, bool, error) {
	if len(poc) == 0 {
		return poc, false, nil
	}

	path := poc
	if path[0] == '~' {
		var err error
		path, err = homedir.Expand(path)
		if err != nil {
			return "", true, err
		}
	}

	if _, err := os.Stat(path); err == nil {
		contents, err := ioutil.ReadFile(path)
		if err != nil {
			return string(contents), true, err
		}
		return string(contents), true, nil
	}

	return poc, false, nil
}

// prettyPrintJSONLines iterates through a []byte line-by-line,
// transforming any lines that are complete json into pretty-printed json.
func prettyPrintJSONLines(b []byte) string {
	parts := strings.Split(string(b), "\n")
	for i, p := range parts {
		if b := []byte(p); json.Valid(b) {
			var out bytes.Buffer
			err := json.Indent(&out, b, "", " ")
			contract.IgnoreError(err)
			parts[i] = out.String()
		}
	}
	return strings.Join(parts, "\n")
}

const (
	requestFormat = `HTTP Request Begin %[1]s %[2]s ===================================================
%[3]s
===================================================== HTTP Request End %[1]s %[2]s
`
	responseFormat = `HTTP Response Begin %[1]s [%[2]s] [Length: %[3]d] ===================================================
%[4]s
===================================================== HTTP Response End %[1]s %[2]s
`
)
