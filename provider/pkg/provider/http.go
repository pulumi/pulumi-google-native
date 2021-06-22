// Copyright 2016-2021, Pulumi Corporation.

package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	requestFormat   = `HTTP Request Begin %[1]s %[2]s ===================================================
%[3]s
===================================================== HTTP Request End %[1]s %[2]s
`
	responseFormat = `HTTP Response Begin %[1]s [%[2]s ===================================================
%[3]s
===================================================== HTTP Response End %[1]s %[2]s
`
)

// prettyPrintJsonLines iterates through a []byte line-by-line,
// transforming any lines that are complete json into pretty-printed json.
func prettyPrintJsonLines(b []byte) string {
	parts := strings.Split(string(b), "\n")
	for i, p := range parts {
		if b := []byte(p); json.Valid(b) {
			var out bytes.Buffer
			json.Indent(&out, b, "", " ")
			parts[i] = out.String()
		}
	}
	return strings.Join(parts, "\n")
}


func newGoogleHttpClient(ctx context.Context, config httpClientConfig) (*googleHttpClient, error) {
	client, err := newClient(ctx, config)
	if err != nil {
		return nil, err
	}

	partnerString := ""
	if config.PartnerName != "" {
		partnerString = fmt.Sprintf("GPN:%s; ", config.PartnerName)
	}

	userAgent := fmt.Sprintf("Pulumi/%s (%shttps://www.pulumi.com) pulumi-google-native/%s",
		config.PulumiVersion, partnerString, config.ProviderVersion)

	if config.AppendUserAgent != "" {
		userAgent = fmt.Sprintf("%s %s", userAgent, config.AppendUserAgent)
	}

	return &googleHttpClient{http: client, userAgent: userAgent}, nil
}

type httpClientConfig struct {
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

type googleHttpClient struct {
	http *http.Client
	userAgent string
}

// TODO: This is taken from the TF provider (cut down to a minimal viable thing). We need to make it "good".
func (c *googleHttpClient) sendRequestWithTimeout(method, rawurl string, body map[string]interface{}, timeout time.Duration) (map[string]interface{}, error) {
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

	u, err := addQueryParams(rawurl, map[string]string{"alt": "json"})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, u, &buf)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders

	debugBody, _ := httputil.DumpRequest(req, true)
	logging.V(9).Infof(requestFormat, req.Method, req.URL, prettyPrintJsonLines(debugBody))
	res, err = c.http.Do(req)
	if err != nil {
		return nil, err
	}

	debugResp, _ := httputil.DumpResponse(res, true)
	logging.V(9).Infof(responseFormat, res.Request.Method, res.Request.URL, prettyPrintJsonLines(debugResp))

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

func newClient(ctx context.Context, config httpClientConfig) (*http.Client, error) {
	if len(config.Scopes) == 0 {
		config.Scopes = defaultClientScopes
	}

	credentials, err := getCredentials(ctx, config)
	if err != nil {
		return nil, errors.Wrapf(err, "getting credentials")
	}

	cleanCtx := context.WithValue(ctx, oauth2.HTTPClient, cleanhttp.DefaultClient())

	// 1. OAUTH2 TRANSPORT/CLIENT - sets up proper auth headers
	client := oauth2.NewClient(cleanCtx, credentials.TokenSource)

	// 3. Retry Transport - retries common temporary errors
	// Keep order for wrapping logging so we log each retried request as well.
	// This value should be used if needed to create shallow copies with additional retry predicates.
	// See ClientWithAdditionalRetries
	//retryTransport := &retryTransport{
	//	retryPredicates: defaultErrorRetryPredicates,
	//	internal:        loggingTransport,
	//}

	// Set final transport value.
	// This timeout is a timeout per HTTP request, not per logical operation.
	client.Timeout = 30 * time.Second

	return client, nil
}

func getCredentials(ctx context.Context, c httpClientConfig) (*google.Credentials, error) {
	if c.AccessToken != "" {
		contents, _, err := pathOrContents(c.AccessToken)
		if err != nil {
			return nil, errors.Wrapf(err, "Error loading access token")
		}
		token := &oauth2.Token{AccessToken: contents}

		if c.ImpersonateServiceAccount != "" {
			opts := []option.ClientOption{option.WithTokenSource(oauth2.StaticTokenSource(token)), option.ImpersonateCredentials(c.ImpersonateServiceAccount, c.ImpersonateServiceAccountDelegates...), option.WithScopes(c.Scopes...)}
			creds, err := transport.Creds(ctx, opts...)
			if err != nil {
				return nil, err
			}
			return creds, nil
		}

		glog.V(9).Info("Authenticating using configured Google JSON Access Token...")
		glog.V(9).Infof("  -- Scopes: %s", c.Scopes)

		return &google.Credentials{
			TokenSource: staticTokenSource{oauth2.StaticTokenSource(token)},
		}, nil
	}

	if c.Credentials != "" {
		contents, _, err := pathOrContents(c.Credentials)
		if err != nil {
			return nil, errors.Wrapf(err, "error loading credentials")
		}
		if c.ImpersonateServiceAccount != "" {
			opts := []option.ClientOption{option.WithCredentialsJSON([]byte(contents)), option.ImpersonateCredentials(c.ImpersonateServiceAccount, c.ImpersonateServiceAccountDelegates...), option.WithScopes(c.Scopes...)}
			creds, err := transport.Creds(ctx, opts...)
			if err != nil {
				return nil, err
			}

			glog.V(9).Info("Authenticating using configured Google JSON Credentials and Impersonate Service Account...")
			glog.V(9).Infof("   -- Scopes: %s", c.Scopes)
			return creds, nil
		}
		creds, err := google.CredentialsFromJSON(ctx, []byte(contents), c.Scopes...)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to parse credentials from %q", contents)
		}

		glog.V(9).Info("Authenticating using configured Google JSON Credentials...")
		glog.V(9).Infof("   -- Scopes: %s", c.Scopes)
		return creds, nil
	}

	if c.ImpersonateServiceAccount != "" {
		opts := option.ImpersonateCredentials(c.ImpersonateServiceAccount, c.ImpersonateServiceAccountDelegates...)
		creds, err := transport.Creds(ctx, opts, option.WithScopes(c.Scopes...))
		if err != nil {
			return nil, err
		}

		glog.V(9).Info("Authenticating using configured Impersonate Service Account...")
		glog.V(9).Infof("   -- Scopes: %s", c.Scopes)
		return creds, nil
	}

	glog.V(9).Info("Authenticating using DefaultClient...")
	glog.V(9).Infof("   -- Scopes: %s", c.Scopes)

	defaultTS, err := google.DefaultTokenSource(context.Background(), c.Scopes...)
	if err != nil {
		return nil, fmt.Errorf("Attempted to load application default credentials since neither `credentials` nor `access_token` was set in the provider block.  No credentials loaded. To use your gcloud credentials, run 'gcloud auth application-default login'.  Original error: %w", err)
	}
	return &google.Credentials{
		TokenSource: defaultTS,
	}, err
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

// staticTokenSource is used to be able to identify static token sources without reflection.
type staticTokenSource struct {
	oauth2.TokenSource
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
