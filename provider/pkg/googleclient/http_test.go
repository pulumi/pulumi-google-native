// Copyright 2021, Pulumi Corporation.  All rights reserved.

package googleclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func TestRequestWithTimeout(t *testing.T) {
	body := map[string]interface{}{"body": "abc"}

	server := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if got, want := r.Header.Get("Authorization"), "Bearer valid"; got != want {
			t.Errorf("Authorization header = %q; want %q", got, want)
		}
		content := map[string]interface{}{}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&content))
		assert.Equal(t, body, content)
		require.NoError(t, json.NewEncoder(w).Encode(map[string]interface{}{"status": "OK"}))
	})
	defer server.Close()

	for _, test := range []struct {
		name        string
		tokenSource func() *MockTokenSource
		err         error
	}{
		{
			name: "Valid Token",
			tokenSource: func() *MockTokenSource {
				mt := &MockTokenSource{}
				mt.On("Token").Return(&oauth2.Token{
					AccessToken: "valid",
					Expiry:      time.Now().Add(24 * time.Hour),
				}, nil)
				return mt
			},
		},
		{
			name: "Expired",
			tokenSource: func() *MockTokenSource {
				mt := &MockTokenSource{}
				mt.On("Token").Return(&oauth2.Token{
					AccessToken: "expired",
					Expiry:      time.Now().Add(-24 * time.Hour),
				}, nil)
				return mt
			},
			err: fmt.Errorf("Get %q: expired token provided", server.URL+"?alt=json"),
		},
		{
			name: "Renewed",
			tokenSource: func() *MockTokenSource {
				mt := &MockTokenSource{}
				mt.On("Token").Return(&oauth2.Token{
					AccessToken: "expired",
					Expiry:      time.Now().Add(-24 * time.Hour),
				}, nil).Once()
				mt.On("Token").Return(&oauth2.Token{
					AccessToken: "valid",
					Expiry:      time.Now().Add(24 * time.Hour),
				}, nil).Once()
				return mt
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			c := &GoogleClient{config: Config{}, userAgent: "test"}
			ts := test.tokenSource()
			c.credRetriever = getCredentialsFunc(func(ctx context.Context, _ Config) (*google.Credentials, error) {
				return &google.Credentials{
					TokenSource: ts,
				}, nil
			})
			c.credValidator = credentialValidatorFunc(func() (bool, error) {
				t, err := ts.Token()
				if err != nil {
					return false, err
				}
				return t.Valid(), nil
			})
			c.http = oauth2.NewClient(context.Background(), ts)
			_, err := c.RequestWithTimeout("GET", server.URL, body, 1*time.Second)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				assert.NoError(t, err)
			}
			mock.AssertExpectationsForObjects(t, ts)
		})
	}
}

func newMockServer(handler func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(handler))
}
