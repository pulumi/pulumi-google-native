package googleclient

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (r roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return r(req)
}

type wrapperFunc func(rt http.RoundTripper) http.RoundTripper

func traceInjector() wrapperFunc {
	return func(rt http.RoundTripper) http.RoundTripper {
		return roundTripFunc(func(req *http.Request) (*http.Response, error) {
			ctx := req.Context()
			if ctx == nil {
				return rt.RoundTrip(req)
			}

			name := fmt.Sprintf("%s@%s", req.Method, req.URL.Path)
			span, ctx := opentracing.StartSpanFromContext(ctx, name)
			defer span.Finish()
			req2 := req.WithContext(ctx)
			body, _ := httputil.DumpRequestOut(req2, true)
			span.SetTag("type", "network")
			ext.HTTPUrl.Set(span, req2.URL.String())
			ext.HTTPMethod.Set(span, req2.Method)
			span.LogKV("body", string(body))

			res, err := rt.RoundTrip(req2)
			if err != nil {
				ext.LogError(span, err)
			}
			if res != nil {
				debugResp, _ := httputil.DumpResponse(res, true)
				ext.HTTPStatusCode.Set(span, uint16(res.StatusCode))
				span.LogKV("response-body", string(debugResp))
			}
			return res, err
		})
	}
}
