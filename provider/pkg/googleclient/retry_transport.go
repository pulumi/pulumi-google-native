// Copyright 2016-2022, Pulumi Corporation.

// Adapted from https://github.com/GoogleCloudPlatform/magic-modules/blob/e9a0e8f3ca4efe711c2db0b12846623822120de0/mmv1/third_party/terraform/utils/retry_transport.go

package googleclient

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/hashicorp/errwrap"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	"google.golang.org/api/googleapi"
)

// A http.RoundTripper that retries common errors, with convenience constructors.
//
// NOTE: This meant for TEMPORARY, TRANSIENT ERRORS.
// Do not use for waiting on operations or polling of resource state,
// especially if the expected state (operation done, resource ready, etc)
// takes longer to reach than the default client Timeout.
// In those cases, retryTimeDuration(...)/Retry with appropriate timeout
// and error predicates/handling should be used as a wrapper around the request
// instead.
//
// Example Usage:
// For handwritten/Go clients, the retry transport should be provided via
// the main client or a shallow copy of the HTTP resources, depending on the
// API-specific retry predicates.
// Example Usage in Terraform Config:
//	client := oauth2.NewClient(ctx, tokenSource)
//	// Create with default retry predicates
//	client.Transport := NewTransportWithDefaultRetries(client.Transport, defaultTimeout)
//
//	// If API uses just default retry predicates:
//	c.clientCompute, err = compute.NewService(ctx, option.WithHTTPClient(client))
//	...
//	// If API needs custom additional retry predicates:
//	sqlAdminHttpClient := ClientWithAdditionalRetries(client, retryTransport,
//			isTemporarySqlError1,
//			isTemporarySqlError2)
//	c.clientSqlAdmin, err = compute.NewService(ctx, option.WithHTTPClient(sqlAdminHttpClient))
// ...

const defaultRetryTransportTimeoutSec = 90

// NewTransportWithDefaultRetries constructs a default retryTransport that will retry common temporary errors
func NewTransportWithDefaultRetries(t http.RoundTripper) *retryTransport {
	return &retryTransport{
		retryPredicates: defaultErrorRetryPredicates,
		internal:        t,
	}
}

// ClientWithAdditionalRetries is a helper method to create a shallow copy of an HTTP client with a shallow-copied
// retryTransport s.t. the base HTTP transport is the same (i.e. client connection pools are shared,
// retryPredicates are different)
func ClientWithAdditionalRetries(baseClient *http.Client, predicates ...RetryErrorPredicateFunc) *http.Client {
	copied := *baseClient
	baseRetryTransport := NewTransportWithDefaultRetries(baseClient.Transport)
	copied.Transport = baseRetryTransport.WithAddedPredicates(predicates...)
	return &copied
}

// WithAddedPredicates returns a shallow copy of the retry transport with additional retry
// predicates but same wrapped http.RoundTripper
func (t *retryTransport) WithAddedPredicates(predicates ...RetryErrorPredicateFunc) *retryTransport {
	copyT := *t
	copyT.retryPredicates = append(t.retryPredicates, predicates...)
	return &copyT
}

type retryTransport struct {
	retryPredicates []RetryErrorPredicateFunc
	internal        http.RoundTripper
}

// RoundTrip implements the RoundTripper interface method.
// It retries the given HTTP request based on the retry predicates
// registered under the retryTransport.
func (t *retryTransport) RoundTrip(req *http.Request) (resp *http.Response, respErr error) {
	// Set timeout to default value.
	ctx := req.Context()
	var ccancel context.CancelFunc
	if _, ok := ctx.Deadline(); !ok {
		ctx, ccancel = context.WithTimeout(ctx, defaultRetryTransportTimeoutSec*time.Second)
		defer func() {
			if ctx.Err() == nil {
				// Cleanup child context created for retry loop if ctx not done.
				ccancel()
			}
		}()
	}

	attempts := 0
	backoff := time.Millisecond * 500
	nextBackoff := time.Millisecond * 500

	// VCR depends on the original request body being consumed, so
	// consume here. Since this won't affect the request itself,
	// we do this before the actual Retry loop so we can consume the request Body as needed
	// e.g. if the request couldn't be retried, we use the original request
	if _, err := httputil.DumpRequestOut(req, true); err != nil {
		logging.V(3).Infof("[WARN] Retry Transport: Consuming original request body failed: %v", err)
	}

	logging.V(9).Infof("Retry Transport: starting RoundTrip retry loop")
Retry:
	for {
		// RoundTrip contract says request body can/will be consumed, so we need to
		// copy the request body for each attempt.
		// If we can't copy the request, we run as a single request.
		newRequest, copyErr := copyHttpRequest(req)
		if copyErr != nil {
			logging.V(3).Infof("[WARN] Retry Transport: Unable to copy request body: %v.", copyErr)
			logging.V(3).Infof("[WARN] Retry Transport: Running request as non-retryable")
			resp, respErr = t.internal.RoundTrip(req)
			break Retry
		}

		logging.V(9).Infof("[DEBUG] Retry Transport: request attempt %d", attempts)
		// Do the wrapped Roundtrip. This is one request in the retry loop.
		resp, respErr = t.internal.RoundTrip(newRequest)
		attempts++

		retryErr := t.checkForRetryableError(resp, respErr)
		if retryErr == nil {
			logging.V(9).Info("[DEBUG] Retry Transport: Stopping retries, last request was successful")
			break Retry
		}
		if !retryErr.Retryable {
			logging.V(9).Infof("[DEBUG] Retry Transport: Stopping retries, "+
				"last request failed with non-retryable error: %s", retryErr.Err)
			break Retry
		}

		logging.V(9).Infof("[DEBUG] Retry Transport: Waiting %s before trying request again", backoff)
		select {
		case <-ctx.Done():
			logging.V(9).Infof("[DEBUG] Retry Transport: Stopping retries, context done: %v", ctx.Err())
			break Retry
		case <-time.After(backoff):
			logging.V(9).Infof("[DEBUG] Retry Transport: Finished waiting %s before next retry", backoff)

			// Fibonnaci backoff - 0.5, 1, 1.5, 2.5, 4, 6.5, 10.5, ...
			lastBackoff := backoff
			backoff = backoff + nextBackoff
			nextBackoff = lastBackoff
			continue
		}
	}
	logging.V(9).Infof("[DEBUG] Retry Transport: Returning after %d attempts", attempts)
	return resp, respErr
}

// copyHttpRequest provides an copy of the given HTTP request for one RoundTrip.
// If the request has a non-empty body (io.ReadCloser), the body is deep copied
// so it can be consumed.
func copyHttpRequest(req *http.Request) (*http.Request, error) {
	newRequest := *req
	if req.Body == nil || req.Body == http.NoBody {
		return &newRequest, nil
	}
	// Helpers like http.NewRequest add a GetBody for copying.
	// If not given, we should reject the request.
	if req.GetBody == nil {
		return nil, errors.New("request.GetBody is not defined for non-empty Body")
	}

	bd, err := req.GetBody()
	if err != nil {
		return nil, err
	}

	newRequest.Body = bd
	return &newRequest, nil
}

// checkForRetryableError uses the googleapi.CheckResponse util to check for
// errors in the response, and determines whether there is a retryable error.
// in response/response error.
func (t *retryTransport) checkForRetryableError(resp *http.Response, respErr error) *retryError {
	var errToCheck error

	if respErr != nil {
		errToCheck = respErr
	} else {
		respToCheck := *resp
		// The RoundTrip contract states that the HTTP response/response error
		// returned cannot be edited. We need to consume the Body to check for
		// errors, so we need to create a copy if the Response has a body.
		if resp.Body != nil && resp.Body != http.NoBody {
			// Use httputil.DumpResponse since the only important info is
			// error code and messages in the response body.
			dumpBytes, err := httputil.DumpResponse(resp, true)
			if err != nil {
				return nonRetryableError(fmt.Errorf("unable to check response for error: %v", err))
			}
			respToCheck.Body = ioutil.NopCloser(bytes.NewReader(dumpBytes))
		}
		errToCheck = googleapi.CheckResponse(&respToCheck)
	}

	if errToCheck == nil {
		return nil
	}
	if isRetryableError(errToCheck, t.retryPredicates...) {
		return retryableError(errToCheck)
	}
	return nonRetryableError(errToCheck)
}

func isRetryableError(topErr error, customPredicates ...RetryErrorPredicateFunc) bool {
	if topErr == nil {
		return false
	}

	retryPredicates := append(
		// Global error retry predicates are registered in this default list.
		defaultErrorRetryPredicates,
		customPredicates...)

	// Check all wrapped errors for a retryable error status.
	isRetryable := false
	errwrap.Walk(topErr, func(werr error) {
		for _, pred := range retryPredicates {
			if predRetry, predReason := pred(werr); predRetry {
				logging.V(9).Infof("[DEBUG] Dismissed an error as retryable. %s - %s", predReason, werr)
				isRetryable = true
				return
			}
		}
	})
	return isRetryable
}

// retryError is the required return type of RetryFunc. It forces client code
// to choose whether or not a given error is retryable.
type retryError struct {
	Err       error
	Retryable bool
}

func (e *retryError) Unwrap() error {
	return e.Err
}

// retryableError is a helper to create a retryError that's retryable from a
// given error. To prevent logic errors, will return an error when passed a
// nil error.
func retryableError(err error) *retryError {
	if err == nil {
		return &retryError{
			Err: errors.New("empty retryable error received. " +
				"This is a bug with the google-native provider and should be " +
				"reported as a GitHub issue in the provider repository"),
			Retryable: false,
		}
	}
	return &retryError{Err: err, Retryable: true}
}

// nonRetryableError is a helper to create a retryError that's _not_ retryable
// from a given error. To prevent logic errors, will return an error when
// passed a nil error.
func nonRetryableError(err error) *retryError {
	if err == nil {
		return &retryError{
			Err: errors.New("empty non-retryable error received. " +
				"This is a bug with the google-native provider and should be " +
				"reported as a GitHub issue in the provider repository"),
			Retryable: false,
		}
	}
	return &retryError{Err: err, Retryable: false}
}
