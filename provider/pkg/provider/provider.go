// Copyright 2016-2020, Pulumi Corporation.

package provider

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jpillora/backoff"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-google-cloud/provider/pkg/resources"
	"github.com/pulumi/pulumi/pkg/v2/resource/provider"
	"github.com/pulumi/pulumi/sdk/v2/go/common/diag"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"
	rpc "github.com/pulumi/pulumi/sdk/v2/proto/go"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	"google.golang.org/api/storage/v1"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type googleCloudProvider struct {
	host        *provider.HostClient
	name        string
	version     string
	config      map[string]string
	schemaBytes []byte
	resourceMap *resources.CloudAPIMetadata
}

func makeProvider(host *provider.HostClient, name, version string, schemaBytes []byte,
	cloudAPIResourcesBytes []byte) (rpc.ResourceProviderServer, error) {
	resourceMap, err := loadMetadata(cloudAPIResourcesBytes)
	if err != nil {
		return nil, err
	}

	// Return the new provider
	return &googleCloudProvider{
		host:        host,
		name:        name,
		version:     version,
		config:      map[string]string{},
		schemaBytes: schemaBytes,
		resourceMap: resourceMap,
	}, nil
}

// loadMetadata deserializes the provided compressed json byte array into a CloudAPIMetadata struct.
func loadMetadata(azureAPIResourcesBytes []byte) (*resources.CloudAPIMetadata, error) {
	var resourceMap resources.CloudAPIMetadata

	uncompressed, err := gzip.NewReader(bytes.NewReader(azureAPIResourcesBytes))
	if err != nil {
		return nil, errors.Wrap(err, "expand compressed metadata")
	}
	if err = json.NewDecoder(uncompressed).Decode(&resourceMap); err != nil {
		return nil, errors.Wrap(err, "unmarshalling resource map")
	}
	if err = uncompressed.Close(); err != nil {
		return nil, errors.Wrap(err, "closing uncompress stream for metadata")
	}
	return &resourceMap, nil
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
	res, ok := k.resourceMap.Resources[resourceKey]
	if !ok {
		return nil, errors.Errorf("resource '%s' not found", resourceKey)
	}

	id := res.IdPath
	idParams := res.IdParams
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
		obj, err := insertCall.Do()
		if err != nil {
			return nil, err
		}

		resp = map[string]interface{}{
			"mediaLink": obj.MediaLink,
			"name":      obj.Name,
			"selfLink":  obj.SelfLink,
		}
	default:
		uri = fmt.Sprintf("%s%s", res.BaseUrl, res.CreatePath)
		for _, param := range res.CreateParams {
			key := resource.PropertyKey(param)
			value := inputs[key].StringValue()
			uri = strings.Replace(uri, fmt.Sprintf("{%s}", param), value, 1)
			delete(inputs, key)
		}
		for _, param := range res.IdParams {
			key := resource.PropertyKey(param)
			delete(inputs, key)
		}

		inputsMap := inputs.Mappable()
		body := map[string]interface{}{}
		for name, value := range res.CreateProperties {
			parent := body
			if value.Container != "" {
				if v, has := body[value.Container].(map[string]interface{}); has {
					parent = v
				} else {
					parent = map[string]interface{}{}
					body[value.Container] = parent
				}
			}
			parent[name] = inputsMap[name]
		}

		op, err := sendRequestWithTimeout(ctx, "POST", uri, body, 0)
		if err != nil {
			return nil, fmt.Errorf("error sending request: %s: %q %+v", err, uri, inputs.Mappable())
		}

		resp, err = k.waitForResourceOpCompletion(ctx, res.BaseUrl, op)
		if err != nil {
			return nil, errors.Wrapf(err, "waiting for completion")
		}
	}

	// Serialize and return RPC outputs
	outputs, err := plugin.MarshalProperties(
		resource.NewPropertyMapFromMap(resp),
		plugin.MarshalOptions{Label: fmt.Sprintf("%s.response", label), SkipNulls: true},
	)

	return &rpc.CreateResponse{
		Id:         id,
		Properties: outputs,
	}, nil
}

func (p *googleCloudProvider) waitForResourceOpCompletion(ctx context.Context, baseUrl string, resp map[string]interface{}) (map[string]interface{}, error) {
	retryPolicy := backoff.Backoff{
		Min:    1 * time.Second,
		Max:    15 * time.Second,
		Factor: 1.5,
		Jitter: true,
	}
	for {
		glog.V(9).Infof("waiting for completion: %+v", resp)

		// There are two styles of operations: one returns a 'done' boolean flag, another one returns status='DONE'.
		done, hasDone := resp["done"].(bool)
		status, hasStatus := resp["status"].(string)
		if completed := (hasDone && done) || (hasStatus && status == "DONE"); completed {
			if failure, has := resp["error"]; has {
				return nil, errors.Errorf("operation errored with %+v", failure)
			}
			if statusMessage, has := resp["statusMessage"]; has {
				return nil, errors.Errorf("operation failed with %q", statusMessage)
			}
			if response, has := resp["response"].(map[string]interface{}); has {
				return response, nil
			}
			if operationType, has := resp["operationType"].(string); has && strings.Contains(strings.ToLower(operationType), "delete") {
				return resp, nil
			}
			if targetLink, has := resp["targetLink"].(string); has {
				return sendRequestWithTimeout(ctx, "GET", targetLink, nil, 0)
			}
			return resp, nil
		}

		var pollUri string
		if selfLink, has := resp["selfLink"].(string); has && hasStatus {
			pollUri = selfLink
		} else {
			if name, has := resp["name"].(string); has && strings.HasPrefix(name, "operations/") {
				pollUri = fmt.Sprintf("%s/v1/%s", baseUrl, name)
			}
		}

		if pollUri == "" {
			return resp, nil
		}

		time.Sleep(retryPolicy.Duration())

		op, err := sendRequestWithTimeout(ctx, "GET", pollUri, nil, 0)
		if err != nil {
			return nil, errors.Wrapf(err, "polling operation status")
		}

		resp = op
	}
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
func (k *googleCloudProvider) Read(ctx context.Context, req *rpc.ReadRequest) (*rpc.ReadResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Read(%s)", k.name, urn)
	resourceKey := string(urn.Type())
	res, ok := k.resourceMap.Resources[resourceKey]
	if !ok {
		return nil, errors.Errorf("resource '%s' not found", resourceKey)
	}

	id := req.GetId()
	if res.NoGet {
		return &rpc.ReadResponse{Id: id}, nil
	}

	uri := fmt.Sprintf("%s%s", res.BaseUrl, id)

	resp, err := sendRequestWithTimeout(ctx, "GET", uri, nil, 0)
	if err != nil {
		if reqErr, ok := err.(*googleapi.Error); ok && reqErr.Code == http.StatusNotFound {
			// 404 means that the resource was deleted.
			return &rpc.ReadResponse{Id: ""}, nil
		}
		return nil, fmt.Errorf("error sending request: %s", err)
	}

	// Serialize and return RPC outputs.
	checkpoint, err := plugin.MarshalProperties(
		resource.NewPropertyMapFromMap(resp),
		plugin.MarshalOptions{Label: fmt.Sprintf("%s.checkpoint", label), SkipNulls: true},
	)
	if err != nil {
		return nil, err
	}

	return &rpc.ReadResponse{Id: id, Properties: checkpoint}, nil
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
	res, ok := k.resourceMap.Resources[resourceKey]
	if !ok {
		return nil, errors.Errorf("resource '%s' not found", resourceKey)
	}

	if res.NoDelete {
		// TODO: delete with Set (e.g. Policy resources)
		return &empty.Empty{}, nil
	}

	id := req.GetId()
	uri := fmt.Sprintf("%s/%s", res.BaseUrl, id)

	resp, err := sendRequestWithTimeout(ctx, "DELETE", uri, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %s", err)
	}

	_, err = k.waitForResourceOpCompletion(ctx, res.BaseUrl, resp)
	if err != nil {
		return nil, errors.Wrapf(err, "waiting for completion")
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
