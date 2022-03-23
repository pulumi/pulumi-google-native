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

package provider

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
	"time"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/golang/protobuf/ptypes/empty"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/iancoleman/strcase"
	"github.com/jpillora/backoff"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-google-native/provider/pkg/gen"
	"github.com/pulumi/pulumi-google-native/provider/pkg/googleclient"
	"github.com/pulumi/pulumi-google-native/provider/pkg/resources"
	"github.com/pulumi/pulumi-google-native/provider/pkg/version"
	"github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/rpcutil/rpcerror"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"google.golang.org/api/googleapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type googleCloudProvider struct {
	host        *provider.HostClient
	name        string
	version     string
	config      map[string]string
	schemaBytes []byte
	client      *googleclient.GoogleClient
	resourceMap *resources.CloudAPIMetadata
	converter   *resources.SdkShapeConverter
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
		converter:   &resources.SdkShapeConverter{Types: resourceMap.Types},
	}, nil
}

// loadMetadata deserializes the provided compressed json byte array into a CloudAPIMetadata struct.
func loadMetadata(metadataBytes []byte) (*resources.CloudAPIMetadata, error) {
	var resourceMap resources.CloudAPIMetadata

	uncompressed, err := gzip.NewReader(bytes.NewReader(metadataBytes))
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

	impersonateServiceAccountDelegatesString := p.getConfig(
		"impersonateServiceAccountDelegates", "")
	var impersonateServiceAccountDelegates []string
	if impersonateServiceAccountDelegatesString != "" {
		err := json.Unmarshal([]byte(impersonateServiceAccountDelegatesString), &impersonateServiceAccountDelegates)
		if err != nil {
			return nil, errors.Wrapf(err,
				"failed to unmarshal %q as Impersonate Service Account Delegates",
				impersonateServiceAccountDelegatesString)
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

	appendUserAgent := p.getConfig("appendUserAgent", "GOOGLE_APPEND_USER_AGENT")
	impersonateServiceAccount := p.getConfig("impersonateServiceAccount", "GOOGLE_IMPERSONATE_SERVICE_ACCOUNT")

	config := googleclient.Config{
		Credentials:                        p.getConfig("credentials", "GOOGLE_CREDENTIALS"),
		AccessToken:                        p.getConfig("accessToken", "GOOGLE_OAUTH_ACCESS_TOKEN"),
		ImpersonateServiceAccount:          impersonateServiceAccount,
		ImpersonateServiceAccountDelegates: impersonateServiceAccountDelegates,
		Scopes:                             scopes,
		PulumiVersion:                      getPulumiVersion(),
		ProviderVersion:                    version.Version,
		PartnerName:                        p.getPartnerName(),
		AppendUserAgent:                    appendUserAgent,
	}

	client, err := googleclient.New(ctx, config)
	if err != nil {
		return nil, err
	}
	p.client = client

	return &rpc.ConfigureResponse{
		AcceptSecrets: true,
	}, nil
}

// Invoke dynamically executes a built-in function in the provider.
func (p *googleCloudProvider) Invoke(_ context.Context, req *rpc.InvokeRequest) (*rpc.InvokeResponse, error) {
	label := fmt.Sprintf("%s.Invoke(%s)", p.name, req.Tok)
	inv, ok := p.resourceMap.Functions[req.Tok]
	if !ok {
		return nil, errors.Errorf("invoke %q not found", req.Tok)
	}

	args, err := plugin.UnmarshalProperties(req.GetArgs(), plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.args", label), SkipNulls: true,
	})
	if err != nil {
		return nil, err
	}

	// Apply default config values.
	for _, param := range inv.URL.Values {
		sdkName := param.Name
		if param.SdkName != "" {
			sdkName = param.SdkName
		}
		switch sdkName {
		case "project":
			key := resource.PropertyKey(sdkName)
			if value, ok := p.getDefaultValue(key, sdkName, args); ok {
				args[key] = *value
			}
		}
	}

	uri, err := buildFunctionURL(inv, args)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.RequestWithTimeout(inv.Verb, uri, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %s", err)
	}

	// Serialize and return outputs.
	result, err := plugin.MarshalProperties(
		resource.NewPropertyMapFromMap(resp),
		plugin.MarshalOptions{Label: fmt.Sprintf("%s.response", label), SkipNulls: true},
	)
	if err != nil {
		return nil, err
	}

	return &rpc.InvokeResponse{Return: result}, nil
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
	logging.V(9).Infof("%s executing", label)

	// Deserialize RPC inputs.
	olds, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.olds", label), SkipNulls: true,
	})
	if err != nil {
		return nil, err
	}
	news, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.news", label), KeepUnknowns: true, SkipNulls: true,
	})
	if err != nil {
		return nil, err
	}

	resourceKey := string(urn.Type())
	res, ok := p.resourceMap.Resources[resourceKey]
	if !ok {
		return nil, errors.Errorf("resource %q not found", resourceKey)
	}

	// Apply default config values.
	var failures []*rpc.CheckFailure
	for _, param := range res.Create.Endpoint.Values {
		sdkName := param.Name
		if param.SdkName != "" {
			sdkName = param.SdkName
		}
		switch sdkName {
		case "project", "location", "zone":
			key := resource.PropertyKey(sdkName)
			configName := sdkName
			if sdkName == "location" {
				configName = "region"
			}
			if _, has := news[key]; has {
				continue
			}
			if value, ok := p.getDefaultValue(key, configName, olds); ok {
				news[key] = *value
			} else {
				reason := fmt.Sprintf(
					"missing required property '%s'. Either set it explicitly or configure it with "+
						"'pulumi config set google-native:%s <value>'.", sdkName, configName)
				failures = append(failures, &rpc.CheckFailure{
					Reason: reason,
				})
			}
		}
	}

	// Auto-naming.
	var nameKey resource.PropertyKey
	if autonameFieldRegex.MatchString(res.Create.Autoname.FieldName) {
		nameKey = resource.PropertyKey(res.Create.Autoname.FieldName)
	} else {
		nameKey = "name"
	}
	if res.Create.Autoname.FieldName != "" && !news.HasValue(nameKey) {
		news[nameKey] = getDefaultName(urn, res.Create.Autoname.FieldName, olds, news)
	}

	// Apply property patterns.
	for name, prop := range res.Create.SDKProperties {
		key := resource.PropertyKey(name)
		if prop.SdkName != "" {
			key = resource.PropertyKey(prop.SdkName)
		}
		if value, ok := applyPropertyPattern(key, prop, news); ok {
			news[key] = *value
		}
	}

	resInputs, err := plugin.MarshalProperties(news, plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.resInputs", label), KeepUnknowns: true})
	if err != nil {
		return nil, err
	}

	return &rpc.CheckResponse{Inputs: resInputs, Failures: failures}, nil
}

// Get a default project name for the given inputs.
func (p *googleCloudProvider) getDefaultValue(
	key resource.PropertyKey,
	configName string,
	olds resource.PropertyMap,
) (*resource.PropertyValue, bool) {
	// 1. Check if old inputs define the value.
	if v, ok := olds[key]; ok {
		return &v, true
	}

	// 2. Check if the config has a corresponding value.
	if cv, ok := p.config[configName]; ok {
		v := resource.NewStringProperty(cv)
		return &v, true
	}

	return nil, false
}

func (p *googleCloudProvider) GetSchema(_ context.Context, req *rpc.GetSchemaRequest) (*rpc.GetSchemaResponse, error) {
	if v := req.GetVersion(); v != 0 {
		return nil, fmt.Errorf("unsupported schema version %d", v)
	}

	decompressed, err := gen.DecompressSchema(p.schemaBytes)
	if err != nil {
		return nil, errors.New("failure loading schema")
	}
	return &rpc.GetSchemaResponse{Schema: string(decompressed)}, nil
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
	logging.V(9).Infof("%s executing", label)

	resourceKey := string(urn.Type())
	_, ok := p.resourceMap.Resources[resourceKey]
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

	return &rpc.DiffResponse{Changes: rpc.DiffResponse_DIFF_UNKNOWN, DeleteBeforeReplace: true}, nil
}

// Create allocates a new instance of the provided resource and returns its unique ID afterwards.
func (p *googleCloudProvider) Create(ctx context.Context, req *rpc.CreateRequest) (*rpc.CreateResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Create(%s)", p.name, urn)
	logging.V(9).Infof("%s executing", label)

	// Deserialize RPC inputs
	inputs, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.properties", label), SkipNulls: true,
	})
	if err != nil {
		return nil, err
	}
	inputsMap := inputs.Mappable()

	resourceKey := string(urn.Type())
	res, ok := p.resourceMap.Resources[resourceKey]
	if !ok {
		return nil, errors.Errorf("resource %q not found", resourceKey)
	}
	logging.V(9).Infof("Looked up metadata for %q: %+v", resourceKey, res)

	uri, err := buildCreateURL(res, inputs)
	if err != nil {
		return nil, err
	}
	body := p.prepareAPIInputs(inputs, nil, res.Create.SDKProperties)

	var op map[string]interface{}
	if res.AssetUpload {
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

		op, err = p.client.UploadWithTimeout(res.Create.Verb, uri, body, content, 0)
		if err != nil {
			return nil, fmt.Errorf("error sending upload request: %s: %q %+v %d", err, uri, inputs.Mappable(), len(content))
		}
	} else {
		op, err = p.client.RequestWithTimeout(res.Create.Verb, uri, body, 0)
		if err != nil {
			return nil, fmt.Errorf("error sending request: %s: %q %+v", err, uri, inputs.Mappable())
		}
	}

	resp, err := p.waitForResourceOpCompletion(res.RootURL, res.Create.Polling.Strategy, op)
	if err != nil {
		if resp == nil {
			return nil, errors.Wrapf(err, "waiting for completion")
		}
		// A partial failure may have occurred because we got an error and a response.
		// Try reading the resource state and return a partial error if there is some.
		id, idErr := calculateResourceID(res, inputsMap, resp)
		if idErr != nil {
			return nil, errors.Wrapf(err, "waiting for completion / calculate ID %s", idErr)
		}
		readResp, getErr := p.client.RequestWithTimeout("GET", id, nil, 0)
		if getErr != nil {
			return nil, errors.Wrapf(err, "waiting for completion / read state %s", getErr)
		}
		checkpoint, cpErr := plugin.MarshalProperties(
			checkpointObject(inputs, readResp),
			plugin.MarshalOptions{Label: fmt.Sprintf("%s.partialCheckpoint", label), KeepSecrets: true, SkipNulls: true},
		)
		if cpErr != nil {
			return nil, errors.Wrapf(err, "waiting for completion / checkpoint %s", cpErr)
		}
		return nil, partialError(id, err, checkpoint, req.GetProperties())
	}

	// Store both outputs and inputs into the state.
	checkpoint, err := plugin.MarshalProperties(
		checkpointObject(inputs, resp),
		plugin.MarshalOptions{Label: fmt.Sprintf("%s.checkpoint", label), KeepSecrets: true, SkipNulls: true},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal checkpoint: %w", err)
	}

	id, err := calculateResourceID(res, inputsMap, resp)
	if err != nil {
		return nil, errors.Wrapf(err, "calculating resource ID")
	}

	return &rpc.CreateResponse{
		Id:         id,
		Properties: checkpoint,
	}, nil
}

func (p *googleCloudProvider) prepareAPIInputs(
	inputs, state resource.PropertyMap,
	properties map[string]resources.CloudAPIProperty,
) map[string]interface{} {
	return p.converter.SdkPropertiesToRequestBody(properties, inputs.Mappable(), state.Mappable())
}

// waitForResourceOpCompletion keeps polling the resource or operation URL until it gets
// a success or a failure of provisioning.
// Note that both a response and an error can be returned in case of a partially-failed deployment
// (e.g., resource is created but failed to initialize to completion).
func (p *googleCloudProvider) waitForResourceOpCompletion(
	rootURL string,
	pollingStrategy resources.PollingStrategy,
	resp map[string]interface{},
) (map[string]interface{}, error) {
	retryPolicy := backoff.Backoff{
		Min:    1 * time.Second,
		Max:    15 * time.Second,
		Factor: 1.5,
		Jitter: true,
	}
	for {
		logging.V(9).Infof("waiting for completion: polling strategy: %q: %+v", pollingStrategy, resp)

		var pollURI string

		// if the resource has a custom polling strategy, use that.
		switch pollingStrategy {
		case resources.KNativeStatusPoll:
			logging.V(9).Info("Getting self link URL")
			sl, err := getKNativeSelfLinkURL(resp)
			if err != nil {
				return nil, err
			}
			logging.V(9).Infof("selfLink: %q from response: %+v", sl, resp)
			pollURI = sl
			completed, err := knativeStatusCheck(resp)
			if err != nil {
				return nil, err
			}
			if completed {
				return resp, nil
			}
		default:
			// Otherwise there are two styles of operations: one returns a 'done' boolean flag,
			// another one returns status='DONE'.
			done, hasDone := resp["done"].(bool)
			status, hasStatus := resp["status"].(string)
			if completed := (hasDone && done) || (hasStatus && status == "DONE"); completed {
				// Extract an error message from the response, if any.
				var err error
				if failure, has := resp["error"]; has {
					err = errors.Errorf("operation errored with %+v", failure)
				} else if statusMessage, has := resp["statusMessage"]; has {
					err = errors.Errorf("operation failed with %q", statusMessage)
				}
				// Extract the resource response, if any.
				// A partial error could happen, so both response and error could be available.
				if response, has := resp["response"].(map[string]interface{}); has {
					return response, err
				}
				if operationType, has := resp["operationType"].(string); has &&
					strings.Contains(strings.ToLower(operationType), "delete") {
					return resp, err
				}
				// Check if there's a target link.
				if targetLink, has := resp["targetLink"].(string); has {
					// Try reading resource state.
					state, getErr := p.client.RequestWithTimeout("GET", targetLink, nil, 0)
					if getErr != nil {
						if err != nil {
							// Return the original creation error if resource read failed.
							return nil, err
						}
						return nil, getErr
					}
					// A partial error could happen, so both response and error could be available.
					return state, err
				}
				// At this point, we assume either a complete failure or a clean response.
				if err != nil {
					return nil, err
				}
				return resp, nil
			}

			if selfLink, has := resp["selfLink"].(string); has && hasStatus {
				pollURI = selfLink
			} else {
				if name, has := resp["name"].(string); has && strings.HasPrefix(name, "operations/") {
					pollURI = fmt.Sprintf("%s/v1/%s", rootURL, name)
				}
			}
		}

		if pollURI == "" {
			logging.V(3).Infof("No pollURI found for rootURL: %q", rootURL)
			return resp, nil
		}

		time.Sleep(retryPolicy.Duration())

		logging.V(9).Infof("Polling URL: %q", pollURI)
		op, err := p.client.RequestWithTimeout("GET", pollURI, nil, 0)
		if err != nil {
			return nil, errors.Wrapf(err, "polling operation status")
		}

		resp = op
	}
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

	// Retrieve the old state.
	oldState, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.olds", label), KeepUnknowns: true, SkipNulls: true, KeepSecrets: true,
	})
	if err != nil {
		return nil, err
	}

	// Extract old inputs from the `__inputs` field of the old state.
	inputs := parseCheckpointObject(oldState)

	uri, err := res.Read.Endpoint.URI(inputs.Mappable(), oldState.Mappable())
	if err != nil {
		return nil, err
	}

	// Read the current state of the resource from the API.
	newState, err := p.client.RequestWithTimeout(res.Read.Verb, uri, nil, 0)
	if err != nil {
		if reqErr, ok := err.(*googleapi.Error); ok && reqErr.Code == http.StatusNotFound {
			// 404 means that the resource was deleted.
			return &rpc.ReadResponse{Id: ""}, nil
		}
		return nil, fmt.Errorf("error sending request: %s", err)
	}

	newStateProps := resource.NewPropertyMapFromMap(newState)
	if inputs == nil {
		return nil, status.Error(codes.Unimplemented, "Import is not yet implemented")
	}
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
	logging.V(9).Infof("%s executing", label)

	// Deserialize RPC inputs
	inputs, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.properties", label), SkipNulls: true,
	})
	if err != nil {
		return nil, err
	}

	// Deserialize the last known state.
	oldState, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.oldState", label), SkipNulls: true,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "reading resource state")
	}

	resourceKey := string(urn.Type())
	res, ok := p.resourceMap.Resources[resourceKey]
	if !ok {
		return nil, errors.Errorf("resource %q not found", resourceKey)
	}

	if res.Update.Undefined() {
		logging.V(1).Infof("update is currently undefined for resource: %q. %s will not be updated",
			resourceKey, req.GetId())
		return &rpc.UpdateResponse{}, nil
	}

	body := p.prepareAPIInputs(inputs, oldState, res.Update.SDKProperties)
	if res.Update.UpdateMask.BodyPropertyName != "" || res.Update.UpdateMask.QueryParamName != "" {
		newJson, err := json.Marshal(inputs)
		if err != nil {
			return nil, errors.Errorf("failed to serialize new inputs as json")
		}
		// Extract old inputs from the `__inputs` field of the old state.
		oldInputs := parseCheckpointObject(oldState)
		oldJson, err := json.Marshal(oldInputs)
		if err != nil {
			return nil, errors.Errorf("failed to serialize existing inputs as json")
		}
		patch, err := jsonpatch.CreateMergePatch(oldJson, newJson)
		if err != nil {
			return nil, errors.Errorf("failed to generate patch")
		}
		var keys []string
		patchObj := map[string]interface{}{}
		if err = json.Unmarshal(patch, &patchObj); err != nil {
			return nil, errors.Errorf("failed to unmarshal patch object")
		}
		for k := range patchObj {
			found := false
			for name, sdkProp := range res.Update.SDKProperties {
				if sdkProp.SdkName == k {
					keys = append(keys, strcase.ToSnake(name))
					found = true
					break
				}
			}
			if !found {
				keys = append(keys, strcase.ToSnake(k))
			}
		}
		if res.Update.UpdateMask.QueryParamName != "" {
			inputs[resource.PropertyKey(res.Update.UpdateMask.QueryParamName)] =
				resource.NewStringProperty(strings.Join(keys, ","))
		}
		if res.Update.UpdateMask.BodyPropertyName != "" {
			body[res.Update.UpdateMask.BodyPropertyName] = map[string]interface{}{"paths": keys}
		}
	}

	uri, err := res.Update.Endpoint.URI(inputs.Mappable(), oldState.Mappable())
	if err != nil {
		return nil, err
	}

	op, err := p.client.RequestWithTimeout(res.Update.Verb, uri, body, 0)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %s: %q %+v", err, uri, body)
	}

	resp, err := p.waitForResourceOpCompletion(res.RootURL, res.Update.Polling.Strategy, op)
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
	if err != nil {
		return nil, fmt.Errorf("failed to marshal checkpoint: %w", err)
	}

	return &rpc.UpdateResponse{
		Properties: outputs,
	}, nil
}

// Delete tears down an existing resource with the given ID. If it fails, the resource is assumed
// to still exist.
func (p *googleCloudProvider) Delete(_ context.Context, req *rpc.DeleteRequest) (*empty.Empty, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Delete(%s)", p.name, urn)
	logging.V(9).Infof("%s executing", label)

	resourceKey := string(urn.Type())
	res, ok := p.resourceMap.Resources[resourceKey]
	if !ok {
		return nil, errors.Errorf("resource %q not found", resourceKey)
	}

	// Deserialize RPC inputs
	oldState, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.properties", label), SkipNulls: true,
	})
	if err != nil {
		return nil, err
	}

	// Extract old inputs from the `__inputs` field of the old state.
	inputs := parseCheckpointObject(oldState)

	// Look up operation functions by resourceKey
	// Example: google-native:storage/v1:Bucket
	// If defined, run that function instead of a single HTTP call
	if f, ok := ResourceDeleteOverrides[resourceKey]; ok {
		err := f(p, res, inputs.Mappable(), oldState.Mappable())
		if err != nil {
			return nil, err
		}

		return &empty.Empty{}, nil
	}

	uri, err := res.Delete.Endpoint.URI(inputs.Mappable(), oldState.Mappable())
	if err != nil {
		return nil, err
	}

	if res.Delete.Undefined() {
		// At the time of writing, the classic GCP provider has the same behavior and warning for 10 resources.
		logging.V(1).Infof("%q resources"+
			" cannot be deleted from Google Cloud. The resource %s will be removed from Pulumi"+
			" state, but will still be present on Google Cloud.", resourceKey, req.GetId())
		return &empty.Empty{}, nil
	}

	resp, err := p.client.RequestWithTimeout(res.Delete.Verb, uri, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %s", err)
	}

	_, err = p.waitForResourceOpCompletion(res.RootURL, resources.DefaultPoll, resp)
	if err != nil {
		return nil, errors.Wrapf(err, "waiting for completion")
	}

	return &empty.Empty{}, nil
}

// Construct creates a new component resource.
func (p *googleCloudProvider) Construct(_ context.Context, _ *rpc.ConstructRequest) (*rpc.ConstructResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Construct is not yet implemented")
}

// Call dynamically executes a method in the provider associated with a component resource.
func (p *googleCloudProvider) Call(_ context.Context, _ *rpc.CallRequest) (*rpc.CallResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Call is not yet implemented")
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

func (p *googleCloudProvider) getPartnerName() string {
	result := p.getConfig("partnerName", "GOOGLE_PARTNER_NAME")
	if result != "" {
		return result
	}
	disablePartner := p.getConfig("disablePartnerName", "GOOGLE_DISABLE_PARTNER_NAME")
	if disablePartner == "true" {
		return ""
	}
	return "Pulumi"
}

func getPulumiVersion() string {
	if bi, ok := debug.ReadBuildInfo(); ok {
		for _, dep := range bi.Deps {
			if strings.HasPrefix(dep.Path, "github.com/pulumi/pulumi/pkg") {
				return strings.TrimPrefix(dep.Version, "v")
			}
		}
	}
	// We should never get here but let's not panic and return something sensible if we do.
	logging.V(4).Info("No Pulumi package version found, using '3' as the default version for user-agent")
	return "3"
}

// partialError creates an error for resources that did not complete an operation in progress.
// The last known state of the object is included in the error so that it can be checkpointed.
func partialError(id string, err error, state *structpb.Struct, inputs *structpb.Struct) error {
	detail := rpc.ErrorResourceInitFailed{
		Id:         id,
		Properties: state,
		Reasons:    []string{err.Error()},
		Inputs:     inputs,
	}
	return rpcerror.WithDetails(rpcerror.New(codes.Unknown, err.Error()), &detail)
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
		if inputs.IsSecret() {
			return inputs.SecretValue().Element.ObjectValue()
		}
		return inputs.ObjectValue()
	}

	return nil
}
