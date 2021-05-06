// Copyright 2016-2021, Pulumi Corporation.

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
	"github.com/pulumi/pulumi-google-native/provider/pkg/resources"
	"github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	"google.golang.org/api/storage/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	"net/url"
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
	client      *googleHttpClient
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
func (p *googleCloudProvider) Configure(ctx context.Context,
	req *rpc.ConfigureRequest) (*rpc.ConfigureResponse, error) {
	for key, val := range req.GetVariables() {
		p.config[strings.TrimPrefix(key, "google-native:config:")] = val
	}

	p.setLoggingContext(ctx)

	impersonateServiceAccountDelegatesString := p.getConfig("impersonateServiceAccountDelegates", "")
	var impersonateServiceAccountDelegates []string
	if impersonateServiceAccountDelegatesString != "" {
		err := json.Unmarshal([]byte(impersonateServiceAccountDelegatesString), &impersonateServiceAccountDelegates)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal %q as Impersonate Service Account Delegates", impersonateServiceAccountDelegatesString)
		}
	}

	scopesString := p.getConfig("scopes", "")
	var scopes []string
	if scopesString != "" {
		err := json.Unmarshal([]byte(scopesString), &scopes)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal %q as Scopes", scopesString)
		}
	}

	config := credentialsConfig{
		Credentials:                        p.getConfig("credentials", "GOOGLE_CREDENTIALS"),
		AccessToken:                        p.getConfig("accessToken", "GOOGLE_OAUTH_ACCESS_TOKEN"),
		ImpersonateServiceAccount:          p.getConfig("impersonateServiceAccount", "GOOGLE_IMPERSONATE_SERVICE_ACCOUNT"),
		ImpersonateServiceAccountDelegates: impersonateServiceAccountDelegates,
		Scopes:                             scopes,
	}

	client, err := newGoogleHttpClient(ctx, config)
	if err != nil {
		return nil, err
	}
	p.client = client

	return &rpc.ConfigureResponse{
		AcceptSecrets: true,
	}, nil
}

// Invoke dynamically executes a built-in function in the provider.
func (p *googleCloudProvider) Invoke(_ context.Context, _ *rpc.InvokeRequest) (*rpc.InvokeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Invoke is not yet implemented")
}

// StreamInvoke dynamically executes a built-in function in the provider. The result is streamed
// back as a series of messages.
func (p *googleCloudProvider) StreamInvoke(_ *rpc.InvokeRequest, _ rpc.ResourceProvider_StreamInvokeServer) error {
	return status.Error(codes.Unimplemented, "StreamInvoke is not yet implemented")
}

// Check validates that the given property bag is valid for a resource of the given type and returns
// the inputs that should be passed to successive calls to Diff, Create, or Update for this
// resource. As a rule, the provider inputs returned by a call to Check should preserve the original
// representation of the properties as present in the program inputs. Though this rule is not
// required for correctness, violations thereof can negatively impact the end-user experience, as
// the provider inputs are using for detecting and rendering diffs.
func (p *googleCloudProvider) Check(_ context.Context, req *rpc.CheckRequest) (*rpc.CheckResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Check(%s)", p.name, urn)
	glog.V(9).Infof("%s executing", label)

	// Deserialize RPC inputs.
	newResInputs := req.GetNews()

	return &rpc.CheckResponse{Inputs: newResInputs}, nil
}

func (p *googleCloudProvider) GetSchema(_ context.Context, req *rpc.GetSchemaRequest) (*rpc.GetSchemaResponse, error) {
	if v := req.GetVersion(); v != 0 {
		return nil, fmt.Errorf("unsupported schema version %d", v)
	}

	return &rpc.GetSchemaResponse{Schema: string(p.schemaBytes)}, nil
}

// CheckConfig validates the configuration for this provider.
func (p *googleCloudProvider) CheckConfig(_ context.Context, req *rpc.CheckRequest) (*rpc.CheckResponse, error) {
	return &rpc.CheckResponse{Inputs: req.GetNews()}, nil
}

// DiffConfig diffs the configuration for this provider.
func (p *googleCloudProvider) DiffConfig(context.Context, *rpc.DiffRequest) (*rpc.DiffResponse, error) {
	return &rpc.DiffResponse{
		Changes:             0,
		Replaces:            []string{},
		Stables:             []string{},
		DeleteBeforeReplace: false,
	}, nil
}

// Diff checks what impacts a hypothetical update will have on the resource's properties.
func (p *googleCloudProvider) Diff(_ context.Context, req *rpc.DiffRequest) (*rpc.DiffResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Diff(%s)", p.name, urn)
	glog.V(9).Infof("%s executing", label)

	resourceKey := string(urn.Type())
	res, ok := p.resourceMap.Resources[resourceKey]
	if !ok {
		return nil, errors.Errorf("resource %q not found", resourceKey)
	}

	oldState, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.oldState", label),
		KeepUnknowns: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "diff failed because malformed resource inputs")
	}

	// Extract old inputs from the `__inputs` field of the old state.
	oldInputs := parseCheckpointObject(oldState)

	newInputs, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.newInputs", label),
		KeepUnknowns: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "diff failed because malformed resource inputs")
	}

	diff := oldInputs.Diff(newInputs)
	if diff == nil {
		return &rpc.DiffResponse{Changes: rpc.DiffResponse_DIFF_NONE}, nil
	}

	var replaces []string
	for name := range res.CreateProperties {
		if _, ok := diff.Updates[resource.PropertyKey(name)]; ok {
			if _, has := res.UpdateProperties[name]; !has {
				replaces = append(replaces, name)
			}
		}
	}

	return &rpc.DiffResponse{Changes: rpc.DiffResponse_DIFF_UNKNOWN, Replaces: replaces, DeleteBeforeReplace: true}, nil
}

// Create allocates a new instance of the provided resource and returns its unique ID afterwards.
func (p *googleCloudProvider) Create(ctx context.Context, req *rpc.CreateRequest) (*rpc.CreateResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Create(%s)", p.name, urn)
	glog.V(9).Infof("%s executing", label)

	// Deserialize RPC inputs
	inputs, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.properties", label), SkipNulls: true,
	})
	if err != nil {
		return nil, err
	}

	resourceKey := string(urn.Type())
	res, ok := p.resourceMap.Resources[resourceKey]
	if !ok {
		return nil, errors.Errorf("resource %q not found", resourceKey)
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
	case "google-native:storage/v1:BucketObject":
		// This is a very sketchy implementation based on the Go SDK client that exposes media upload.
		// TODO: We may be able to do that based purely on Discovery API.
		opts := []option.ClientOption{option.WithScopes(defaultClientScopes...)}
		opts = append(opts, internaloption.WithDefaultEndpoint(res.BaseUrl))
		clientStorage, err := storage.NewService(ctx, option.WithHTTPClient(p.client.http))
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
		path := res.CreatePath
		querySep := "?"
		if strings.Contains(path, "?") {
			querySep = "&"
		}
		for _, param := range res.CreateParams {
			sdkName := param.Name
			if param.SdkName != "" {
				sdkName = param.SdkName
			}
			key := resource.PropertyKey(sdkName)
			if !inputs[key].HasValue() {
				continue
			}

			value := inputs[key].StringValue()
			switch param.Location {
			case "path":
				path = strings.Replace(path, fmt.Sprintf("{%s}", param.Name), url.PathEscape(value), 1)
			case "query":
				path = fmt.Sprintf("%s%s%s=%s", path, querySep, key, url.QueryEscape(value))
				querySep = "&"
			default:
				return nil, errors.Errorf("unknown param location %q", param.Location)
			}
		}
		uri = res.RelativePath(path)

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

		op, err := p.client.sendRequestWithTimeout(res.CreateVerb, uri, body, 0)
		if err != nil {
			return nil, fmt.Errorf("error sending request: %s: %q %+v", err, uri, inputs.Mappable())
		}

		resp, err = p.waitForResourceOpCompletion(res.BaseUrl, op)
		if err != nil {
			return nil, errors.Wrapf(err, "waiting for completion")
		}
	}

	// Store both outputs and inputs into the state.
	checkpoint, err := plugin.MarshalProperties(
		checkpointObject(inputs, resp),
		plugin.MarshalOptions{Label: fmt.Sprintf("%s.checkpoint", label), KeepSecrets: true, SkipNulls: true},
	)

	return &rpc.CreateResponse{
		Id:         id,
		Properties: checkpoint,
	}, nil
}

func (p *googleCloudProvider) waitForResourceOpCompletion(baseUrl string, resp map[string]interface{}) (map[string]interface{}, error) {
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
				return p.client.sendRequestWithTimeout("GET", targetLink, nil, 0)
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

		op, err := p.client.sendRequestWithTimeout("GET", pollUri, nil, 0)
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
	panic(fmt.Sprintf("%q not found", param))
}

// Read the current live state associated with a resource.
func (p *googleCloudProvider) Read(_ context.Context, req *rpc.ReadRequest) (*rpc.ReadResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Read(%s)", p.name, urn)
	resourceKey := string(urn.Type())
	res, ok := p.resourceMap.Resources[resourceKey]
	if !ok {
		return nil, errors.Errorf("resource %q not found", resourceKey)
	}

	id := req.GetId()
	uri := res.RelativePath(id)

	// Retrieve the old state.
	oldState, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.olds", label), KeepUnknowns: true, SkipNulls: true, KeepSecrets: true,
	})
	if err != nil {
		return nil, err
	}

	// Read the current state of the resource from the API.
	newState, err := p.client.sendRequestWithTimeout("GET", uri, nil, 0)
	if err != nil {
		if reqErr, ok := err.(*googleapi.Error); ok && reqErr.Code == http.StatusNotFound {
			// 404 means that the resource was deleted.
			return &rpc.ReadResponse{Id: ""}, nil
		}
		return nil, fmt.Errorf("error sending request: %s", err)
	}

	// Extract old inputs from the `__inputs` field of the old state.
	inputs := parseCheckpointObject(oldState)
	newStateProps := resource.NewPropertyMapFromMap(newState)
	if inputs == nil {
		return nil, status.Error(codes.Unimplemented, "Import is not yet implemented")
	} else {
		// It's hard to infer the changes in the inputs shape based on the outputs without false positives.
		// The current approach is complicated but it's aimed to minimize the noise while refreshing:
		// 0. We have "old" inputs and outputs before refresh and "new" outputs read from API.
		// 1. Project old outputs to their corresponding input shape (exclude read-only properties).
		oldInputProjection := getInputsFromState(res, oldState)
		// 2. Project new outputs to their corresponding input shape (exclude read-only properties).
		newInputProjection := getInputsFromState(res, newStateProps)
		// 3. Calculate the difference between two projections. This should give us actual significant changes
		// that happened in Google Cloud between the last resource update and its current state.
		diff := oldInputProjection.Diff(newInputProjection)
		// 4. Apply this difference to the actual inputs (not a projection) that we have in state.
		inputs = applyDiff(inputs, diff)
	}

	// Store both outputs and inputs into the state checkpoint.
	checkpoint, err := plugin.MarshalProperties(
		checkpointObject(inputs, newState),
		plugin.MarshalOptions{Label: fmt.Sprintf("%s.checkpoint", label), KeepSecrets: true, SkipNulls: true},
	)
	if err != nil {
		return nil, err
	}

	inputsRecord, err := plugin.MarshalProperties(
		inputs,
		plugin.MarshalOptions{Label: fmt.Sprintf("%s.inputs", label), KeepUnknowns: true, SkipNulls: true},
	)
	if err != nil {
		return nil, err
	}

	return &rpc.ReadResponse{Id: id, Properties: checkpoint, Inputs: inputsRecord}, nil
}

// Update updates an existing resource with new values.
func (p *googleCloudProvider) Update(_ context.Context, req *rpc.UpdateRequest) (*rpc.UpdateResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Update(%s)", p.name, urn)
	glog.V(9).Infof("%s executing", label)

	// Deserialize RPC inputs
	inputs, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.properties", label), SkipNulls: true,
	})
	if err != nil {
		return nil, err
	}

	resourceKey := string(urn.Type())
	res, ok := p.resourceMap.Resources[resourceKey]
	if !ok {
		return nil, errors.Errorf("resource %q not found", resourceKey)
	}

	inputsMap := inputs.Mappable()
	body := map[string]interface{}{}
	for name, value := range res.UpdateProperties {
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

	uri := res.RelativePath(req.GetId())
	if strings.HasSuffix(uri, ":getIamPolicy") {
		uri = strings.ReplaceAll(uri, ":getIamPolicy", ":setIamPolicy")
	}

	op, err := p.client.sendRequestWithTimeout(res.UpdateVerb, uri, body, 0)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %s: %q %+v", err, uri, body)
	}

	resp, err := p.waitForResourceOpCompletion(res.BaseUrl, op)
	if err != nil {
		return nil, errors.Wrapf(err, "waiting for completion")
	}

	// Read the inputs to persist them into state.
	newInputs, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.newInputs", label),
		KeepUnknowns: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "diff failed because malformed resource inputs")
	}

	// Store both outputs and inputs into the state and return RPC checkpoint.
	outputs, err := plugin.MarshalProperties(
		checkpointObject(newInputs, resp),
		plugin.MarshalOptions{Label: fmt.Sprintf("%s.response", label), KeepSecrets: true, SkipNulls: true},
	)

	return &rpc.UpdateResponse{
		Properties: outputs,
	}, nil
}

// Delete tears down an existing resource with the given ID. If it fails, the resource is assumed
// to still exist.
func (p *googleCloudProvider) Delete(_ context.Context, req *rpc.DeleteRequest) (*empty.Empty, error) {
	urn := resource.URN(req.GetUrn())
	resourceKey := string(urn.Type())
	res, ok := p.resourceMap.Resources[resourceKey]
	if !ok {
		return nil, errors.Errorf("resource %q not found", resourceKey)
	}

	uri := res.RelativePath(req.GetId())

	if strings.HasSuffix(uri, ":getIamPolicy") {
		uri = strings.ReplaceAll(uri, ":getIamPolicy", ":setIamPolicy")

		resp, err := p.client.sendRequestWithTimeout(res.UpdateVerb, uri, map[string]interface{}{}, 0)
		if err != nil {
			return nil, fmt.Errorf("error sending request: %s", err)
		}

		_, err = p.waitForResourceOpCompletion(res.BaseUrl, resp)
		if err != nil {
			return nil, errors.Wrapf(err, "waiting for completion")
		}

		return &empty.Empty{}, nil
	}

	if res.NoDelete {
		// At the time of writing, the classic GCP provider has the same behavior and warning for 10 resources.
		glog.V(4).Infof("%q resources"+
			" cannot be deleted from Google Cloud. The resource %s will be removed from Pulumi"+
			" state, but will still be present on Google Cloud.", resourceKey, uri)
		return &empty.Empty{}, nil
	}

	resp, err := p.client.sendRequestWithTimeout("DELETE", uri, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %s", err)
	}

	_, err = p.waitForResourceOpCompletion(res.BaseUrl, resp)
	if err != nil {
		return nil, errors.Wrapf(err, "waiting for completion")
	}

	return &empty.Empty{}, nil
}

// Construct creates a new component resource.
func (p *googleCloudProvider) Construct(_ context.Context, _ *rpc.ConstructRequest) (*rpc.ConstructResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Construct is not yet implemented")
}

// GetPluginInfo returns generic information about this plugin, like its version.
func (p *googleCloudProvider) GetPluginInfo(context.Context, *empty.Empty) (*rpc.PluginInfo, error) {
	return &rpc.PluginInfo{
		Version: p.version,
	}, nil
}

// Cancel signals the provider to gracefully shut down and abort any ongoing resource operations.
// Operations aborted in this way will return an error (e.g., `Update` and `Create` will either a
// creation error or an initialization error). Since Cancel is advisory and non-blocking, it is up
// to the host to decide how long to wait after Cancel is called before (e.g.)
// hard-closing any gRPC connection.
func (p *googleCloudProvider) Cancel(context.Context, *empty.Empty) (*empty.Empty, error) {
	// TODO
	return &empty.Empty{}, nil
}

func (p *googleCloudProvider) setLoggingContext(ctx context.Context) {
	log.SetOutput(&LogRedirector{
		writers: map[string]func(string) error{
			tfTracePrefix: func(msg string) error { return p.host.Log(ctx, diag.Debug, "", msg) },
			tfDebugPrefix: func(msg string) error { return p.host.Log(ctx, diag.Debug, "", msg) },
			tfInfoPrefix:  func(msg string) error { return p.host.Log(ctx, diag.Info, "", msg) },
			tfWarnPrefix:  func(msg string) error { return p.host.Log(ctx, diag.Warning, "", msg) },
			tfErrorPrefix: func(msg string) error { return p.host.Log(ctx, diag.Error, "", msg) },
		},
	})
}

func (p *googleCloudProvider) getConfig(configName, envName string) string {
	if val, ok := p.config[configName]; ok {
		return val
	}

	return os.Getenv(envName)
}

// checkpointObject puts inputs in the `__inputs` field of the state.
func checkpointObject(inputs resource.PropertyMap, outputs map[string]interface{}) resource.PropertyMap {
	object := resource.NewPropertyMapFromMap(outputs)
	object["__inputs"] = resource.MakeSecret(resource.NewObjectProperty(inputs))
	return object
}

// parseCheckpointObject returns inputs that are saved in the `__inputs` field of the state.
func parseCheckpointObject(obj resource.PropertyMap) resource.PropertyMap {
	if inputs, ok := obj["__inputs"]; ok {
		return inputs.SecretValue().Element.ObjectValue()
	}

	return nil
}
