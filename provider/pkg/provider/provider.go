// Copyright 2016-2020, Pulumi Corporation.

package provider

import (
	"bytes"
	"context"
	"fmt"
	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-google-cloud/provider/pkg/resources"
	"github.com/pulumi/pulumi/pkg/v2/resource/provider"
	"github.com/pulumi/pulumi/sdk/v2/go/common/diag"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"
	rpc "github.com/pulumi/pulumi/sdk/v2/proto/go"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	"google.golang.org/api/storage/v1"
	"log"
	"os"
	"strings"
)

type googleCloudProvider struct {
	host        *provider.HostClient
	name        string
	version     string
	config      map[string]string
	schemaBytes []byte
}

func makeProvider(host *provider.HostClient, name, version string, schemaBytes []byte) (rpc.ResourceProviderServer, error) {
	// Return the new provider
	return &googleCloudProvider{
		host:        host,
		name:        name,
		version:     version,
		config:      map[string]string{},
		schemaBytes: schemaBytes,
	}, nil
}

// Configure configures the resource provider with "globals" that control its behavior.
func (k *googleCloudProvider) Configure(ctx context.Context,
	req *rpc.ConfigureRequest) (*rpc.ConfigureResponse, error) {
	for key, val := range req.GetVariables() {
		k.config[strings.TrimPrefix(key, "google-cloud:config:")] = val
	}

	k.setLoggingContext(ctx)

	return &rpc.ConfigureResponse{}, nil
}

// Invoke dynamically executes a built-in function in the provider.
func (k *googleCloudProvider) Invoke(_ context.Context, _ *rpc.InvokeRequest) (*rpc.InvokeResponse, error) {
	panic("Invoke not implemented")
}

// StreamInvoke dynamically executes a built-in function in the provider. The result is streamed
// back as a series of messages.
func (k *googleCloudProvider) StreamInvoke(_ *rpc.InvokeRequest, _ rpc.ResourceProvider_StreamInvokeServer) error {
	panic("StreamInvoke not implemented")
}

// Check validates that the given property bag is valid for a resource of the given type and returns
// the inputs that should be passed to successive calls to Diff, Create, or Update for this
// resource. As a rule, the provider inputs returned by a call to Check should preserve the original
// representation of the properties as present in the program inputs. Though this rule is not
// required for correctness, violations thereof can negatively impact the end-user experience, as
// the provider inputs are using for detecting and rendering diffs.
func (k *googleCloudProvider) Check(_ context.Context, req *rpc.CheckRequest) (*rpc.CheckResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Check(%s)", k.name, urn)
	glog.V(9).Infof("%s executing", label)

	// Deserialize RPC inputs.
	newResInputs := req.GetNews()

	return &rpc.CheckResponse{Inputs: newResInputs}, nil
}

func (k *googleCloudProvider) GetSchema(_ context.Context, req *rpc.GetSchemaRequest) (*rpc.GetSchemaResponse, error) {
	if v := req.GetVersion(); v != 0 {
		return nil, fmt.Errorf("unsupported schema version %d", v)
	}

	return &rpc.GetSchemaResponse{Schema: string(k.schemaBytes)}, nil
}

// CheckConfig validates the configuration for this provider.
func (k *googleCloudProvider) CheckConfig(_ context.Context, req *rpc.CheckRequest) (*rpc.CheckResponse, error) {
	return &rpc.CheckResponse{Inputs: req.GetNews()}, nil
}

// DiffConfig diffs the configuration for this provider.
func (k *googleCloudProvider) DiffConfig(context.Context, *rpc.DiffRequest) (*rpc.DiffResponse, error) {
	return &rpc.DiffResponse{
		Changes:             0,
		Replaces:            []string{},
		Stables:             []string{},
		DeleteBeforeReplace: false,
	}, nil
}

// Diff checks what impacts a hypothetical update will have on the resource's properties.
func (k *googleCloudProvider) Diff(_ context.Context, req *rpc.DiffRequest) (*rpc.DiffResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Diff(%s)", k.name, urn)
	glog.V(9).Infof("%s executing", label)

	return &rpc.DiffResponse{
		Changes: rpc.DiffResponse_DIFF_UNKNOWN,
	}, nil
}

// Create allocates a new instance of the provided resource and returns its unique ID afterwards.
func (k *googleCloudProvider) Create(ctx context.Context, req *rpc.CreateRequest) (*rpc.CreateResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Create(%s)", k.name, urn)
	glog.V(9).Infof("%s executing", label)

	// Deserialize RPC inputs
	inputs, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.properties", label), SkipNulls: true,
	})
	if err != nil {
		return nil, err
	}

	resourceKey := string(urn.Type())
	res, ok := resources.CloudAPIResources[resourceKey]
	if !ok {
		return nil, errors.Errorf("resource '%s' not found", resourceKey)
	}

	id := res.DeletePath
	idParams := res.DeleteParams
	if id == "" {
		id = res.CreatePath
		idParams = res.CreateParams
	}
	for _, param := range idParams {
		value := evalParam(inputs, param)
		id = strings.Replace(id, fmt.Sprintf("{%s}", param), value, 1)
	}

	var uri string
	var resp map[string]interface{}
	switch resourceKey {
	case "google-cloud:storage/v1:BucketObject":
		// This is a very sketchy implementation based on the Go SDK client that exposes media upload.
		// TODO: We may be able to do that based purely on Discovery API.
		opts := []option.ClientOption{option.WithScopes(defaultClientScopes...)}
		opts = append(opts, internaloption.WithDefaultEndpoint(res.BaseUrl))
		client := newClient(ctx)
		clientStorage, err := storage.NewService(ctx, option.WithHTTPClient(client))
		if err != nil {
			return nil, err
		}
		objectsService := storage.NewObjectsService(clientStorage)
		bucket := inputs["bucket"].StringValue()
		object := &storage.Object{Bucket: bucket}
		insertCall := objectsService.Insert(bucket, object)
		insertCall.Name(inputs["name"].StringValue())
		// TODO: use other properties
		var content []byte
		source := inputs["source"]
		if source.IsAsset() {
			content, err = source.AssetValue().Bytes()
		} else if source.IsArchive() {
			content, err = source.ArchiveValue().Bytes(resource.ZIPArchive)
		}
		if err != nil {
			return nil, err
		}
		insertCall.Media(bytes.NewReader(content))
		_, err = insertCall.Do()
		if err != nil {
			return nil, err
		}

		resp = map[string]interface{}{}
	default:
		uri = fmt.Sprintf("%s/%s", res.BaseUrl, res.CreatePath)
		for _, param := range res.CreateParams {
			key := resource.PropertyKey(param)
			value := inputs[key].StringValue()
			uri = strings.Replace(uri, fmt.Sprintf("{%s}", param), value, 1)
			delete(inputs, key)
		}

		resp, err = sendRequestWithTimeout(ctx, "POST", uri, inputs.Mappable(), 0)
		if err != nil {
			return nil, fmt.Errorf("error sending request: %s", err)
		}
	}

	// Serialize and return RPC outputs
	outputs, err := plugin.MarshalProperties(
		resource.NewPropertyMapFromMap(resp),
		plugin.MarshalOptions{Label: fmt.Sprintf("%s.response", label), SkipNulls: true},
	)

	// TODO: Await. Successful response doesn't mean that operation is complete.

	return &rpc.CreateResponse{
		Id:         id,
		Properties: outputs,
	}, nil
}

func evalParam(inputs resource.PropertyMap, param string) string {
	// TODO: make it more robust if this function proves useful
	current := inputs
	parts := strings.Split(param, ".")
	for _, part := range parts {
		value := current[resource.PropertyKey(part)]
		switch {
		case value.IsString():
			return value.StringValue()
		case value.IsObject():
			current = value.ObjectValue()
		}
	}
	panic(fmt.Sprintf("'%s' not found", param))
}

// Read the current live state associated with a resource.
func (k *googleCloudProvider) Read(_ context.Context, _ *rpc.ReadRequest) (*rpc.ReadResponse, error) {
	panic("Read not implemented")
}

// Update updates an existing resource with new values.
func (k *googleCloudProvider) Update(_ context.Context, _ *rpc.UpdateRequest) (*rpc.UpdateResponse, error) {
	panic("Update not implemented")
}

// Delete tears down an existing resource with the given ID. If it fails, the resource is assumed
// to still exist.
func (k *googleCloudProvider) Delete(ctx context.Context, req *rpc.DeleteRequest) (*empty.Empty, error) {
	urn := resource.URN(req.GetUrn())
	resourceKey := string(urn.Type())
	res, ok := resources.CloudAPIResources[resourceKey]
	if !ok {
		return nil, errors.Errorf("resource '%s' not found", resourceKey)
	}

	if res.DeletePath == "" {
		// TODO: delete with Set (e.g. Policy resources)
		return &empty.Empty{}, nil
	}

	id := req.GetId()
	uri := fmt.Sprintf("%s/%s", res.BaseUrl, id)

	_, err := sendRequestWithTimeout(ctx, "DELETE", uri, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %s", err)
	}

	return &empty.Empty{}, nil
}

// Construct creates a new component resource.
func (k *googleCloudProvider) Construct(_ context.Context, _ *rpc.ConstructRequest) (*rpc.ConstructResponse, error) {
	panic("Construct not implemented")
}

// GetPluginInfo returns generic information about this plugin, like its version.
func (k *googleCloudProvider) GetPluginInfo(context.Context, *empty.Empty) (*rpc.PluginInfo, error) {
	return &rpc.PluginInfo{
		Version: k.version,
	}, nil
}

// Cancel signals the provider to gracefully shut down and abort any ongoing resource operations.
// Operations aborted in this way will return an error (e.g., `Update` and `Create` will either a
// creation error or an initialization error). Since Cancel is advisory and non-blocking, it is up
// to the host to decide how long to wait after Cancel is called before (e.g.)
// hard-closing any gRPC connection.
func (k *googleCloudProvider) Cancel(context.Context, *empty.Empty) (*empty.Empty, error) {
	// TODO
	return &empty.Empty{}, nil
}

func (k *googleCloudProvider) setLoggingContext(ctx context.Context) {
	log.SetOutput(&LogRedirector{
		writers: map[string]func(string) error{
			tfTracePrefix: func(msg string) error { return k.host.Log(ctx, diag.Debug, "", msg) },
			tfDebugPrefix: func(msg string) error { return k.host.Log(ctx, diag.Debug, "", msg) },
			tfInfoPrefix:  func(msg string) error { return k.host.Log(ctx, diag.Info, "", msg) },
			tfWarnPrefix:  func(msg string) error { return k.host.Log(ctx, diag.Warning, "", msg) },
			tfErrorPrefix: func(msg string) error { return k.host.Log(ctx, diag.Error, "", msg) },
		},
	})
}

func (k *googleCloudProvider) getConfig(configName, envName string) string {
	if val, ok := k.config[configName]; ok {
		return val
	}

	return os.Getenv(envName)
}
