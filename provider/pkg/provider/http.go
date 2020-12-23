// Copyright 2016-2020, Pulumi Corporation.

package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/googleapi"
	"net/http"
	"net/url"
	"time"
)

// TODO: This is taken from the TF provider (cut down to a minimal viable thing). We need to make it "good".
func sendRequestWithTimeout(ctx context.Context, method, rawurl string, body map[string]interface{}, timeout time.Duration) (map[string]interface{}, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", "pulumi") // TODO: Pulumi UA
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
	client := newClient(ctx)
	res, err = client.Do(req)
	if err != nil {
		return nil, err
	}

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

func GetCredentials(clientScopes []string) (google.Credentials, error) {

	//if c.AccessToken != "" {
	//	contents, _, err := pathOrContents(c.AccessToken)
	//	if err != nil {
	//		return googleoauth.Credentials{}, fmt.Errorf("Error loading access token: %s", err)
	//	}
	//	token := &oauth2.Token{AccessToken: contents}
	//
	//	if c.ImpersonateServiceAccount != "" {
	//		opts := []option.ClientOption{option.WithTokenSource(oauth2.StaticTokenSource(token)), option.ImpersonateCredentials(c.ImpersonateServiceAccount, c.ImpersonateServiceAccountDelegates...), option.WithScopes(clientScopes...)}
	//		creds, err := transport.Creds(context.TODO(), opts...)
	//		if err != nil {
	//			return googleoauth.Credentials{}, err
	//		}
	//		return *creds, nil
	//	}
	//
	//	log.Printf("[INFO] Authenticating using configured Google JSON 'access_token'...")
	//	log.Printf("[INFO]   -- Scopes: %s", clientScopes)
	//
	//	return googleoauth.Credentials{
	//		TokenSource: staticTokenSource{oauth2.StaticTokenSource(token)},
	//	}, nil
	//}
	//
	//if c.Credentials != "" {
	//	contents, _, err := pathOrContents(c.Credentials)
	//	if err != nil {
	//		return googleoauth.Credentials{}, fmt.Errorf("error loading credentials: %s", err)
	//	}
	//	if c.ImpersonateServiceAccount != "" {
	//		opts := []option.ClientOption{option.WithCredentialsJSON([]byte(contents)), option.ImpersonateCredentials(c.ImpersonateServiceAccount, c.ImpersonateServiceAccountDelegates...), option.WithScopes(clientScopes...)}
	//		creds, err := transport.Creds(context.TODO(), opts...)
	//		if err != nil {
	//			return googleoauth.Credentials{}, err
	//		}
	//		return *creds, nil
	//	}
	//	creds, err := googleoauth.CredentialsFromJSON(c.context, []byte(contents), clientScopes...)
	//	if err != nil {
	//		return googleoauth.Credentials{}, fmt.Errorf("unable to parse credentials from '%s': %s", contents, err)
	//	}
	//
	//	log.Printf("[INFO] Authenticating using configured Google JSON 'credentials'...")
	//	log.Printf("[INFO]   -- Scopes: %s", clientScopes)
	//	return *creds, nil
	//}
	//
	//if c.ImpersonateServiceAccount != "" {
	//	opts := option.ImpersonateCredentials(c.ImpersonateServiceAccount, c.ImpersonateServiceAccountDelegates...)
	//	creds, err := transport.Creds(context.TODO(), opts, option.WithScopes(clientScopes...))
	//	if err != nil {
	//		return googleoauth.Credentials{}, err
	//	}
	//	return *creds, nil
	//}

	defaultTS, err := google.DefaultTokenSource(context.Background(), clientScopes...)
	if err != nil {
		return google.Credentials{}, fmt.Errorf("Attempted to load application default credentials since neither `credentials` nor `access_token` was set in the provider block.  No credentials loaded. To use your gcloud credentials, run 'gcloud auth application-default login'.  Original error: %w", err)
	}
	return google.Credentials{
		TokenSource: defaultTS,
	}, err
}

func getTokenSource(clientScopes []string) (oauth2.TokenSource, error) {
	creds, err := GetCredentials(clientScopes)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	return creds.TokenSource, nil
}

func newClient(ctx context.Context) *http.Client {
	scopes := defaultClientScopes

	tokenSource, err := getTokenSource(scopes)
	if err != nil {
		panic(err)
	}

	cleanCtx := context.WithValue(ctx, oauth2.HTTPClient, cleanhttp.DefaultClient())

	// 1. OAUTH2 TRANSPORT/CLIENT - sets up proper auth headers
	client := oauth2.NewClient(cleanCtx, tokenSource)

	// 2. Logging Transport - ensure we log HTTP requests to GCP APIs.
	loggingTransport := logging.NewTransport("Google", client.Transport)

	// 3. Retry Transport - retries common temporary errors
	// Keep order for wrapping logging so we log each retried request as well.
	// This value should be used if needed to create shallow copies with additional retry predicates.
	// See ClientWithAdditionalRetries
	//retryTransport := &retryTransport{
	//	retryPredicates: defaultErrorRetryPredicates,
	//	internal:        loggingTransport,
	//}

	// Set final transport value.
	client.Transport = loggingTransport

	// This timeout is a timeout per HTTP request, not per logical operation.
	client.Timeout = 30 * time.Second

	return client
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
