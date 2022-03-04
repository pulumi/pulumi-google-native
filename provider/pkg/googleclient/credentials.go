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
	"context"
	"fmt"

	"github.com/golang/glog"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
)

type tokenSourceFunc func() (*oauth2.Token, error)

func (t tokenSourceFunc) Token() (*oauth2.Token, error) {
	return t()
}

// validatedTokenSource wraps a source TokenSource such that if
// the source returns a token that is invalid (e.g. expired),
// an error is returned.
func validatedTokenSource(src oauth2.TokenSource) oauth2.TokenSource {
	return tokenSourceFunc(func() (*oauth2.Token, error) {
		t, err := src.Token()
		if err != nil {
			return nil, err
		}

		if !t.Valid() {
			return nil, fmt.Errorf("expired token provided")
		}

		return t, nil
	})
}

// newReusedValidatedTokenSource returns a validatedTokenSource which uses an
// oauth2.ReuseTokenSource as its source.
// This is used because while ReuseTokenSource will refresh a token if it detects it
// is invalid, the new token is not itself validated.
func newReuseValidatedTokenSource(t *oauth2.Token, src oauth2.TokenSource) *reuseValidatedTokenSource {
	return &reuseValidatedTokenSource{
		src: validatedTokenSource(oauth2.ReuseTokenSource(t, src)),
	}
}

type reuseValidatedTokenSource struct {
	src oauth2.TokenSource
}

func (r *reuseValidatedTokenSource) Token() (*oauth2.Token, error) {
	return r.src.Token()
}

type credentialValidator interface {
	// isValid detects if the current active credential/token is valid.
	isValid() (bool, error)
}

type credentialValidatorFunc func() (bool, error)

func (c credentialValidatorFunc) isValid() (bool, error) {
	return c()
}

type credentialRetriever interface {
	// getCredentials uses the provided Config to create appropriate credentials
	// to pass to Google cloud clients.
	getCredentials(ctx context.Context, c Config) (*google.Credentials, error)
}

type getCredentialsFunc func(context.Context, Config) (*google.Credentials, error)

func (g getCredentialsFunc) getCredentials(ctx context.Context, c Config) (*google.Credentials, error) {
	return g(ctx, c)
}

func defaultCredentialRetriever() credentialRetriever {
	return getCredentialsFunc(getCredentials)
}

func getCredentials(ctx context.Context, c Config) (*google.Credentials, error) {
	if c.AccessToken != "" {
		contents, _, err := pathOrContents(c.AccessToken)
		if err != nil {
			return nil, errors.Wrapf(err, "Error loading access token")
		}
		token := &oauth2.Token{AccessToken: contents}

		if c.ImpersonateServiceAccount != "" {
			opts := []option.ClientOption{
				option.WithTokenSource(oauth2.StaticTokenSource(token)),
				option.ImpersonateCredentials(c.ImpersonateServiceAccount, c.ImpersonateServiceAccountDelegates...),
				option.WithScopes(c.Scopes...),
			}
			creds, err := transport.Creds(ctx, opts...)
			if err != nil {
				return nil, err
			}
			return creds, nil
		}

		glog.V(9).Info("Authenticating using configured Google JSON Access Token...")
		glog.V(9).Infof("  -- Scopes: %s", c.Scopes)

		return &google.Credentials{
			TokenSource: oauth2.StaticTokenSource(token),
		}, nil
	}

	if c.Credentials != "" {
		contents, _, err := pathOrContents(c.Credentials)
		if err != nil {
			return nil, errors.Wrapf(err, "error loading credentials")
		}
		if c.ImpersonateServiceAccount != "" {
			opts := []option.ClientOption{
				option.WithCredentialsJSON([]byte(contents)),
				option.ImpersonateCredentials(c.ImpersonateServiceAccount, c.ImpersonateServiceAccountDelegates...),
				option.WithScopes(c.Scopes...),
			}
			creds, err := transport.Creds(ctx, opts...)
			if err != nil {
				return nil, err
			}

			glog.V(9).Info(
				"Authenticating using configured Google JSON Credentials and Impersonate Service Account...")
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
		return nil, fmt.Errorf("Attempted to load application default credentials since neither "+
			"`credentials` nor `access_token` was set in the provider block.  No credentials loaded. To use your "+
			"gcloud credentials, run 'gcloud auth application-default login'.  Original error: %w", err)
	}
	return &google.Credentials{
		TokenSource: defaultTS,
	}, err
}
