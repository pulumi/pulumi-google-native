// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The HTTP method to use for the request. The default is POST. The app's request handler for the task's target URL must be able to handle HTTP requests with this http_method, otherwise the task attempt fails with error code 405 (Method Not Allowed). See [Writing a push task request handler](https://cloud.google.com/appengine/docs/java/taskqueue/push/creating-handlers#writing_a_push_task_request_handler) and the App Engine documentation for your runtime on [How Requests are Handled](https://cloud.google.com/appengine/docs/standard/python3/how-requests-are-handled).
type AppEngineHttpRequestHttpMethod string

const (
	// HTTP method unspecified
	AppEngineHttpRequestHttpMethodHttpMethodUnspecified = AppEngineHttpRequestHttpMethod("HTTP_METHOD_UNSPECIFIED")
	// HTTP POST
	AppEngineHttpRequestHttpMethodPost = AppEngineHttpRequestHttpMethod("POST")
	// HTTP GET
	AppEngineHttpRequestHttpMethodGet = AppEngineHttpRequestHttpMethod("GET")
	// HTTP HEAD
	AppEngineHttpRequestHttpMethodHead = AppEngineHttpRequestHttpMethod("HEAD")
	// HTTP PUT
	AppEngineHttpRequestHttpMethodPut = AppEngineHttpRequestHttpMethod("PUT")
	// HTTP DELETE
	AppEngineHttpRequestHttpMethodDelete = AppEngineHttpRequestHttpMethod("DELETE")
	// HTTP PATCH
	AppEngineHttpRequestHttpMethodPatch = AppEngineHttpRequestHttpMethod("PATCH")
	// HTTP OPTIONS
	AppEngineHttpRequestHttpMethodOptions = AppEngineHttpRequestHttpMethod("OPTIONS")
)

func (AppEngineHttpRequestHttpMethod) ElementType() reflect.Type {
	return reflect.TypeOf((*AppEngineHttpRequestHttpMethod)(nil)).Elem()
}

func (e AppEngineHttpRequestHttpMethod) ToAppEngineHttpRequestHttpMethodOutput() AppEngineHttpRequestHttpMethodOutput {
	return pulumi.ToOutput(e).(AppEngineHttpRequestHttpMethodOutput)
}

func (e AppEngineHttpRequestHttpMethod) ToAppEngineHttpRequestHttpMethodOutputWithContext(ctx context.Context) AppEngineHttpRequestHttpMethodOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AppEngineHttpRequestHttpMethodOutput)
}

func (e AppEngineHttpRequestHttpMethod) ToAppEngineHttpRequestHttpMethodPtrOutput() AppEngineHttpRequestHttpMethodPtrOutput {
	return e.ToAppEngineHttpRequestHttpMethodPtrOutputWithContext(context.Background())
}

func (e AppEngineHttpRequestHttpMethod) ToAppEngineHttpRequestHttpMethodPtrOutputWithContext(ctx context.Context) AppEngineHttpRequestHttpMethodPtrOutput {
	return AppEngineHttpRequestHttpMethod(e).ToAppEngineHttpRequestHttpMethodOutputWithContext(ctx).ToAppEngineHttpRequestHttpMethodPtrOutputWithContext(ctx)
}

func (e AppEngineHttpRequestHttpMethod) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AppEngineHttpRequestHttpMethod) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AppEngineHttpRequestHttpMethod) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AppEngineHttpRequestHttpMethod) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AppEngineHttpRequestHttpMethodOutput struct{ *pulumi.OutputState }

func (AppEngineHttpRequestHttpMethodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppEngineHttpRequestHttpMethod)(nil)).Elem()
}

func (o AppEngineHttpRequestHttpMethodOutput) ToAppEngineHttpRequestHttpMethodOutput() AppEngineHttpRequestHttpMethodOutput {
	return o
}

func (o AppEngineHttpRequestHttpMethodOutput) ToAppEngineHttpRequestHttpMethodOutputWithContext(ctx context.Context) AppEngineHttpRequestHttpMethodOutput {
	return o
}

func (o AppEngineHttpRequestHttpMethodOutput) ToAppEngineHttpRequestHttpMethodPtrOutput() AppEngineHttpRequestHttpMethodPtrOutput {
	return o.ToAppEngineHttpRequestHttpMethodPtrOutputWithContext(context.Background())
}

func (o AppEngineHttpRequestHttpMethodOutput) ToAppEngineHttpRequestHttpMethodPtrOutputWithContext(ctx context.Context) AppEngineHttpRequestHttpMethodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppEngineHttpRequestHttpMethod) *AppEngineHttpRequestHttpMethod {
		return &v
	}).(AppEngineHttpRequestHttpMethodPtrOutput)
}

func (o AppEngineHttpRequestHttpMethodOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AppEngineHttpRequestHttpMethodOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AppEngineHttpRequestHttpMethod) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AppEngineHttpRequestHttpMethodOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AppEngineHttpRequestHttpMethodOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AppEngineHttpRequestHttpMethod) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AppEngineHttpRequestHttpMethodPtrOutput struct{ *pulumi.OutputState }

func (AppEngineHttpRequestHttpMethodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppEngineHttpRequestHttpMethod)(nil)).Elem()
}

func (o AppEngineHttpRequestHttpMethodPtrOutput) ToAppEngineHttpRequestHttpMethodPtrOutput() AppEngineHttpRequestHttpMethodPtrOutput {
	return o
}

func (o AppEngineHttpRequestHttpMethodPtrOutput) ToAppEngineHttpRequestHttpMethodPtrOutputWithContext(ctx context.Context) AppEngineHttpRequestHttpMethodPtrOutput {
	return o
}

func (o AppEngineHttpRequestHttpMethodPtrOutput) Elem() AppEngineHttpRequestHttpMethodOutput {
	return o.ApplyT(func(v *AppEngineHttpRequestHttpMethod) AppEngineHttpRequestHttpMethod {
		if v != nil {
			return *v
		}
		var ret AppEngineHttpRequestHttpMethod
		return ret
	}).(AppEngineHttpRequestHttpMethodOutput)
}

func (o AppEngineHttpRequestHttpMethodPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AppEngineHttpRequestHttpMethodPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AppEngineHttpRequestHttpMethod) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AppEngineHttpRequestHttpMethodInput is an input type that accepts AppEngineHttpRequestHttpMethodArgs and AppEngineHttpRequestHttpMethodOutput values.
// You can construct a concrete instance of `AppEngineHttpRequestHttpMethodInput` via:
//
//          AppEngineHttpRequestHttpMethodArgs{...}
type AppEngineHttpRequestHttpMethodInput interface {
	pulumi.Input

	ToAppEngineHttpRequestHttpMethodOutput() AppEngineHttpRequestHttpMethodOutput
	ToAppEngineHttpRequestHttpMethodOutputWithContext(context.Context) AppEngineHttpRequestHttpMethodOutput
}

var appEngineHttpRequestHttpMethodPtrType = reflect.TypeOf((**AppEngineHttpRequestHttpMethod)(nil)).Elem()

type AppEngineHttpRequestHttpMethodPtrInput interface {
	pulumi.Input

	ToAppEngineHttpRequestHttpMethodPtrOutput() AppEngineHttpRequestHttpMethodPtrOutput
	ToAppEngineHttpRequestHttpMethodPtrOutputWithContext(context.Context) AppEngineHttpRequestHttpMethodPtrOutput
}

type appEngineHttpRequestHttpMethodPtr string

func AppEngineHttpRequestHttpMethodPtr(v string) AppEngineHttpRequestHttpMethodPtrInput {
	return (*appEngineHttpRequestHttpMethodPtr)(&v)
}

func (*appEngineHttpRequestHttpMethodPtr) ElementType() reflect.Type {
	return appEngineHttpRequestHttpMethodPtrType
}

func (in *appEngineHttpRequestHttpMethodPtr) ToAppEngineHttpRequestHttpMethodPtrOutput() AppEngineHttpRequestHttpMethodPtrOutput {
	return pulumi.ToOutput(in).(AppEngineHttpRequestHttpMethodPtrOutput)
}

func (in *appEngineHttpRequestHttpMethodPtr) ToAppEngineHttpRequestHttpMethodPtrOutputWithContext(ctx context.Context) AppEngineHttpRequestHttpMethodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AppEngineHttpRequestHttpMethodPtrOutput)
}

// The HTTP method to use for the request. The default is POST.
type HttpRequestHttpMethod string

const (
	// HTTP method unspecified
	HttpRequestHttpMethodHttpMethodUnspecified = HttpRequestHttpMethod("HTTP_METHOD_UNSPECIFIED")
	// HTTP POST
	HttpRequestHttpMethodPost = HttpRequestHttpMethod("POST")
	// HTTP GET
	HttpRequestHttpMethodGet = HttpRequestHttpMethod("GET")
	// HTTP HEAD
	HttpRequestHttpMethodHead = HttpRequestHttpMethod("HEAD")
	// HTTP PUT
	HttpRequestHttpMethodPut = HttpRequestHttpMethod("PUT")
	// HTTP DELETE
	HttpRequestHttpMethodDelete = HttpRequestHttpMethod("DELETE")
	// HTTP PATCH
	HttpRequestHttpMethodPatch = HttpRequestHttpMethod("PATCH")
	// HTTP OPTIONS
	HttpRequestHttpMethodOptions = HttpRequestHttpMethod("OPTIONS")
)

func (HttpRequestHttpMethod) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpRequestHttpMethod)(nil)).Elem()
}

func (e HttpRequestHttpMethod) ToHttpRequestHttpMethodOutput() HttpRequestHttpMethodOutput {
	return pulumi.ToOutput(e).(HttpRequestHttpMethodOutput)
}

func (e HttpRequestHttpMethod) ToHttpRequestHttpMethodOutputWithContext(ctx context.Context) HttpRequestHttpMethodOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HttpRequestHttpMethodOutput)
}

func (e HttpRequestHttpMethod) ToHttpRequestHttpMethodPtrOutput() HttpRequestHttpMethodPtrOutput {
	return e.ToHttpRequestHttpMethodPtrOutputWithContext(context.Background())
}

func (e HttpRequestHttpMethod) ToHttpRequestHttpMethodPtrOutputWithContext(ctx context.Context) HttpRequestHttpMethodPtrOutput {
	return HttpRequestHttpMethod(e).ToHttpRequestHttpMethodOutputWithContext(ctx).ToHttpRequestHttpMethodPtrOutputWithContext(ctx)
}

func (e HttpRequestHttpMethod) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HttpRequestHttpMethod) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HttpRequestHttpMethod) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HttpRequestHttpMethod) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HttpRequestHttpMethodOutput struct{ *pulumi.OutputState }

func (HttpRequestHttpMethodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpRequestHttpMethod)(nil)).Elem()
}

func (o HttpRequestHttpMethodOutput) ToHttpRequestHttpMethodOutput() HttpRequestHttpMethodOutput {
	return o
}

func (o HttpRequestHttpMethodOutput) ToHttpRequestHttpMethodOutputWithContext(ctx context.Context) HttpRequestHttpMethodOutput {
	return o
}

func (o HttpRequestHttpMethodOutput) ToHttpRequestHttpMethodPtrOutput() HttpRequestHttpMethodPtrOutput {
	return o.ToHttpRequestHttpMethodPtrOutputWithContext(context.Background())
}

func (o HttpRequestHttpMethodOutput) ToHttpRequestHttpMethodPtrOutputWithContext(ctx context.Context) HttpRequestHttpMethodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HttpRequestHttpMethod) *HttpRequestHttpMethod {
		return &v
	}).(HttpRequestHttpMethodPtrOutput)
}

func (o HttpRequestHttpMethodOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HttpRequestHttpMethodOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HttpRequestHttpMethod) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HttpRequestHttpMethodOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HttpRequestHttpMethodOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HttpRequestHttpMethod) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HttpRequestHttpMethodPtrOutput struct{ *pulumi.OutputState }

func (HttpRequestHttpMethodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpRequestHttpMethod)(nil)).Elem()
}

func (o HttpRequestHttpMethodPtrOutput) ToHttpRequestHttpMethodPtrOutput() HttpRequestHttpMethodPtrOutput {
	return o
}

func (o HttpRequestHttpMethodPtrOutput) ToHttpRequestHttpMethodPtrOutputWithContext(ctx context.Context) HttpRequestHttpMethodPtrOutput {
	return o
}

func (o HttpRequestHttpMethodPtrOutput) Elem() HttpRequestHttpMethodOutput {
	return o.ApplyT(func(v *HttpRequestHttpMethod) HttpRequestHttpMethod {
		if v != nil {
			return *v
		}
		var ret HttpRequestHttpMethod
		return ret
	}).(HttpRequestHttpMethodOutput)
}

func (o HttpRequestHttpMethodPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HttpRequestHttpMethodPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HttpRequestHttpMethod) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// HttpRequestHttpMethodInput is an input type that accepts HttpRequestHttpMethodArgs and HttpRequestHttpMethodOutput values.
// You can construct a concrete instance of `HttpRequestHttpMethodInput` via:
//
//          HttpRequestHttpMethodArgs{...}
type HttpRequestHttpMethodInput interface {
	pulumi.Input

	ToHttpRequestHttpMethodOutput() HttpRequestHttpMethodOutput
	ToHttpRequestHttpMethodOutputWithContext(context.Context) HttpRequestHttpMethodOutput
}

var httpRequestHttpMethodPtrType = reflect.TypeOf((**HttpRequestHttpMethod)(nil)).Elem()

type HttpRequestHttpMethodPtrInput interface {
	pulumi.Input

	ToHttpRequestHttpMethodPtrOutput() HttpRequestHttpMethodPtrOutput
	ToHttpRequestHttpMethodPtrOutputWithContext(context.Context) HttpRequestHttpMethodPtrOutput
}

type httpRequestHttpMethodPtr string

func HttpRequestHttpMethodPtr(v string) HttpRequestHttpMethodPtrInput {
	return (*httpRequestHttpMethodPtr)(&v)
}

func (*httpRequestHttpMethodPtr) ElementType() reflect.Type {
	return httpRequestHttpMethodPtrType
}

func (in *httpRequestHttpMethodPtr) ToHttpRequestHttpMethodPtrOutput() HttpRequestHttpMethodPtrOutput {
	return pulumi.ToOutput(in).(HttpRequestHttpMethodPtrOutput)
}

func (in *httpRequestHttpMethodPtr) ToHttpRequestHttpMethodPtrOutputWithContext(ctx context.Context) HttpRequestHttpMethodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HttpRequestHttpMethodPtrOutput)
}

// The response_view specifies which subset of the Task will be returned. By default response_view is BASIC; not all information is retrieved by default because some data, such as payloads, might be desirable to return only when needed because of its large size or because of the sensitivity of data that it contains. Authorization for FULL requires `cloudtasks.tasks.fullView` [Google IAM](https://cloud.google.com/iam/) permission on the Task resource.
type TaskResponseView string

const (
	// Unspecified. Defaults to BASIC.
	TaskResponseViewViewUnspecified = TaskResponseView("VIEW_UNSPECIFIED")
	// The basic view omits fields which can be large or can contain sensitive data. This view does not include the body in AppEngineHttpRequest. Bodies are desirable to return only when needed, because they can be large and because of the sensitivity of the data that you choose to store in it.
	TaskResponseViewBasic = TaskResponseView("BASIC")
	// All information is returned. Authorization for FULL requires `cloudtasks.tasks.fullView` [Google IAM](https://cloud.google.com/iam/) permission on the Queue resource.
	TaskResponseViewFull = TaskResponseView("FULL")
)

func (TaskResponseView) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskResponseView)(nil)).Elem()
}

func (e TaskResponseView) ToTaskResponseViewOutput() TaskResponseViewOutput {
	return pulumi.ToOutput(e).(TaskResponseViewOutput)
}

func (e TaskResponseView) ToTaskResponseViewOutputWithContext(ctx context.Context) TaskResponseViewOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TaskResponseViewOutput)
}

func (e TaskResponseView) ToTaskResponseViewPtrOutput() TaskResponseViewPtrOutput {
	return e.ToTaskResponseViewPtrOutputWithContext(context.Background())
}

func (e TaskResponseView) ToTaskResponseViewPtrOutputWithContext(ctx context.Context) TaskResponseViewPtrOutput {
	return TaskResponseView(e).ToTaskResponseViewOutputWithContext(ctx).ToTaskResponseViewPtrOutputWithContext(ctx)
}

func (e TaskResponseView) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TaskResponseView) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TaskResponseView) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TaskResponseView) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TaskResponseViewOutput struct{ *pulumi.OutputState }

func (TaskResponseViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskResponseView)(nil)).Elem()
}

func (o TaskResponseViewOutput) ToTaskResponseViewOutput() TaskResponseViewOutput {
	return o
}

func (o TaskResponseViewOutput) ToTaskResponseViewOutputWithContext(ctx context.Context) TaskResponseViewOutput {
	return o
}

func (o TaskResponseViewOutput) ToTaskResponseViewPtrOutput() TaskResponseViewPtrOutput {
	return o.ToTaskResponseViewPtrOutputWithContext(context.Background())
}

func (o TaskResponseViewOutput) ToTaskResponseViewPtrOutputWithContext(ctx context.Context) TaskResponseViewPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TaskResponseView) *TaskResponseView {
		return &v
	}).(TaskResponseViewPtrOutput)
}

func (o TaskResponseViewOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TaskResponseViewOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TaskResponseView) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TaskResponseViewOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TaskResponseViewOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TaskResponseView) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TaskResponseViewPtrOutput struct{ *pulumi.OutputState }

func (TaskResponseViewPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TaskResponseView)(nil)).Elem()
}

func (o TaskResponseViewPtrOutput) ToTaskResponseViewPtrOutput() TaskResponseViewPtrOutput {
	return o
}

func (o TaskResponseViewPtrOutput) ToTaskResponseViewPtrOutputWithContext(ctx context.Context) TaskResponseViewPtrOutput {
	return o
}

func (o TaskResponseViewPtrOutput) Elem() TaskResponseViewOutput {
	return o.ApplyT(func(v *TaskResponseView) TaskResponseView {
		if v != nil {
			return *v
		}
		var ret TaskResponseView
		return ret
	}).(TaskResponseViewOutput)
}

func (o TaskResponseViewPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TaskResponseViewPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TaskResponseView) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// TaskResponseViewInput is an input type that accepts TaskResponseViewArgs and TaskResponseViewOutput values.
// You can construct a concrete instance of `TaskResponseViewInput` via:
//
//          TaskResponseViewArgs{...}
type TaskResponseViewInput interface {
	pulumi.Input

	ToTaskResponseViewOutput() TaskResponseViewOutput
	ToTaskResponseViewOutputWithContext(context.Context) TaskResponseViewOutput
}

var taskResponseViewPtrType = reflect.TypeOf((**TaskResponseView)(nil)).Elem()

type TaskResponseViewPtrInput interface {
	pulumi.Input

	ToTaskResponseViewPtrOutput() TaskResponseViewPtrOutput
	ToTaskResponseViewPtrOutputWithContext(context.Context) TaskResponseViewPtrOutput
}

type taskResponseViewPtr string

func TaskResponseViewPtr(v string) TaskResponseViewPtrInput {
	return (*taskResponseViewPtr)(&v)
}

func (*taskResponseViewPtr) ElementType() reflect.Type {
	return taskResponseViewPtrType
}

func (in *taskResponseViewPtr) ToTaskResponseViewPtrOutput() TaskResponseViewPtrOutput {
	return pulumi.ToOutput(in).(TaskResponseViewPtrOutput)
}

func (in *taskResponseViewPtr) ToTaskResponseViewPtrOutputWithContext(ctx context.Context) TaskResponseViewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TaskResponseViewPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppEngineHttpRequestHttpMethodInput)(nil)).Elem(), AppEngineHttpRequestHttpMethod("HTTP_METHOD_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*AppEngineHttpRequestHttpMethodPtrInput)(nil)).Elem(), AppEngineHttpRequestHttpMethod("HTTP_METHOD_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*HttpRequestHttpMethodInput)(nil)).Elem(), HttpRequestHttpMethod("HTTP_METHOD_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*HttpRequestHttpMethodPtrInput)(nil)).Elem(), HttpRequestHttpMethod("HTTP_METHOD_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*TaskResponseViewInput)(nil)).Elem(), TaskResponseView("VIEW_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*TaskResponseViewPtrInput)(nil)).Elem(), TaskResponseView("VIEW_UNSPECIFIED"))
	pulumi.RegisterOutputType(AppEngineHttpRequestHttpMethodOutput{})
	pulumi.RegisterOutputType(AppEngineHttpRequestHttpMethodPtrOutput{})
	pulumi.RegisterOutputType(HttpRequestHttpMethodOutput{})
	pulumi.RegisterOutputType(HttpRequestHttpMethodPtrOutput{})
	pulumi.RegisterOutputType(TaskResponseViewOutput{})
	pulumi.RegisterOutputType(TaskResponseViewPtrOutput{})
}