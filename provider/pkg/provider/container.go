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
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	pulumiprov "github.com/pulumi/pulumi/sdk/v3/go/pulumi/provider"

	"github.com/imdario/mergo"

	"github.com/pulumi/pulumi/pkg/v3/codegen"

	"github.com/jpillora/backoff"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-google-native/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

// Container cluster support
type clusterMutations struct {
	updateHandlers map[string]containerCustomUpdateHandlerFunc
}

func (c clusterMutations) mutablePropertyPaths() []string {
	return codegen.SortedKeys(c.updateHandlers)
}

func (c clusterMutations) update(
	providerInstance *googleCloudProvider,
	urn resource.URN,
	label string,
	res *resources.CloudAPIResource,
	inputs resource.PropertyMap,
	oldState resource.PropertyMap,
) (appliedInputs resource.PropertyMap, resp map[string]interface{}, err error) {
	// Extract old inputs from the `__inputs` field of the old state.
	oldInputs := parseCheckpointObject(oldState)
	// Delete the auto-injected __autonamed tag in old and new inputs to avoid spurious diffs.
	if v, ok := oldInputs[autonamed]; ok && v.IsBool() {
		delete(oldInputs, autonamed)
	}

	if v, ok := inputs[autonamed]; ok && v.IsBool() {
		delete(inputs, autonamed)
	}

	diff := oldInputs.Diff(inputs)
	batchedInputs := resource.NewObjectProperty(oldInputs)

	if diff == nil {
		logging.V(9).Infof("[%s] no diff found in update!", label)
	} else {
		// Calculate the detailed diff object containing information about replacements.
		detailedDiff := calculateDetailedDiff(res, providerInstance.resourceMap.Types, diff)
		logging.V(9).Infof("[%s] detailedDiff: %+v", label, detailedDiff)
		logging.V(9).Infof("[%s] updateCluster oldState: %+v", label, oldState)

		for _, key := range clusterUpdateOrder {
			updateHandler, ok := c.updateHandlers[key]
			if !ok {
				logging.V(9).Infof("[%s] No handler found for field: %q", label, key)
				// TODO This should be an error once all the targetted mutations are in place!
				continue
			}
			logging.V(9).Infof("[%s] Invoking update for field: %q", label, key)
			for k := range detailedDiff {
				logging.V(9).Infof("[%s] processing diff on field: %q", label, k)
				var diffPropPath, propPath resource.PropertyPath
				diffPropPath, err = resource.ParsePropertyPath(k)
				if err != nil {
					err = fmt.Errorf("failed to parse property path: %q: %w", k, err)
					break
				}
				propPath, err = resource.ParsePropertyPath(key)
				if err != nil {
					err = fmt.Errorf("failed to parse property path: %q: %w", k, err)
					break
				}
				if propPath.Contains(diffPropPath) {
					_ = providerInstance.host.LogStatus(context.Background(), diag.Info, urn,
						fmt.Sprintf("Performing update for field: %q", key))
					logging.V(9).Infof("[%s] processing diff for %q with handler for %q",
						label, diffPropPath.String(), propPath.String())
					logging.V(9).Infof("[%s] waiting for cluster to reach resting state", label)
					_, err = waitForClusterRestingState(providerInstance, urn, res, inputs, oldState)
					if err != nil {
						break
					}
					newVal, _ := diffPropPath.Get(resource.NewObjectProperty(inputs))
					logging.V(9).Infof("[%s] calling update handler", label)
					err = updateHandler(providerInstance, urn, label, res, inputs, oldState)
					if err != nil {
						logging.V(9).Infof("[%s] failure updating: %+v", label, err)
						break
					}
					batchedInputs, ok = diffPropPath.Add(batchedInputs, newVal)
					if !ok {
						err = fmt.Errorf("failed to record update at field: %q", k)
						break
					}
				}
			}
			if err != nil {
				break
			}
		}
	}
	appliedInputs = batchedInputs.ObjectValue()
	uri, readErr := res.Read.Endpoint.URI(inputs.Mappable(), oldState.Mappable())
	if readErr != nil {
		return nil, nil, readErr
	}
	resp, err2 := readContainerResourceStatus(providerInstance, uri)
	if err2 != nil {
		logging.V(9).Infof("[%s] Failed to read cluster status: %v", label, err2)
		return nil, nil, fmt.Errorf("failed to retrieve cluster status: %v. Previous error: %w", err2, err)
	}

	return appliedInputs, resp, err
}

func updateClusterMapping(
	diffPropPath resource.PropertyPath,
	apiFieldName string,
	defaultIfMissing propValueIfMissingFunc,
	operationSuffix string,
	httpMethod string,
	additionalMetadata map[string]resources.CloudAPIProperty,
) containerCustomUpdateHandlerFunc {
	return func(p *googleCloudProvider,
		urn resource.URN,
		label string,
		res *resources.CloudAPIResource,
		newInputs,
		oldState resource.PropertyMap) error {
		tok := urn.Type()
		version := strings.Split(tok.Module().String(), "/")[1]
		endpoint := resources.CloudAPIEndpoint{
			Template: fmt.Sprintf(`https://container.googleapis.com/%s/`+
				`projects/{projectsId}/locations/{locationsId}/clusters/{clustersId}%s`, version, operationSuffix),
			Values: append(res.Create.Endpoint.Values,
				resources.CloudAPIResourceParam{
					Name:    "clustersId",
					SdkName: "name",
					Kind:    "path",
				}),
		}
		uri, err := endpoint.URI(newInputs.Mappable(), oldState.Mappable())
		if err != nil {
			return err
		}

		// If key is omitted, use default.
		_, ok := diffPropPath.Get(resource.NewObjectProperty(newInputs))
		if !ok {
			updated, _ := diffPropPath.Add(resource.NewObjectProperty(newInputs), defaultIfMissing(oldState, diffPropPath))
			newInputs = updated.ObjectValue()
		}

		properties := map[string]resources.CloudAPIProperty{}
		err = mergo.Merge(&properties, additionalMetadata)
		if err != nil {
			return fmt.Errorf("failed to initialize metadata: %w", err)
		}
		properties[apiFieldName] = resources.CloudAPIProperty{SdkName: diffPropPath.String()}
		body := p.prepareAPIInputs(newInputs, oldState, properties)
		op, err := retryRequest(p.client, httpMethod, uri, "", body)
		if err != nil {
			return fmt.Errorf("error sending request: %s: %q %+v", err, uri, body)
		}

		_, err = p.waitForResourceOpCompletion(urn, res.Update.CloudAPIOperation, op, nil)
		if err != nil {
			return errors.Wrapf(err, "waiting for completion")
		}
		return nil
	}
}

func updateClusterNestedField(diffPropPath, targetPropPath resource.PropertyPath,
	defaultIfMissing propValueIfMissingFunc, operationSuffix, httpMethod string) containerCustomUpdateHandlerFunc {
	return func(p *googleCloudProvider,
		urn resource.URN,
		label string,
		res *resources.CloudAPIResource,
		newInputs,
		oldState resource.PropertyMap) error {
		logging.V(9).Infof("[%s] updateClusterNestedField called for property path: %q with input: %+v", label,
			diffPropPath.String(), newInputs)
		tok := urn.Type()
		version := strings.Split(tok.Module().String(), "/")[1]
		endpoint := resources.CloudAPIEndpoint{
			Template: fmt.Sprintf(`https://container.googleapis.com/%s/`+
				`projects/{projectsId}/locations/{locationsId}/clusters/{clustersId}%s`, version, operationSuffix),
			Values: append(res.Create.Endpoint.Values,
				resources.CloudAPIResourceParam{
					Name:    "clustersId",
					SdkName: "name",
					Kind:    "path",
				}),
		}
		uri, err := endpoint.URI(newInputs.Mappable(), oldState.Mappable())
		if err != nil {
			return err
		}
		// Get the new config object to copy the value for fieldName from
		newConfigField, ok := diffPropPath.Get(resource.NewObjectProperty(newInputs))
		if !ok {
			logging.V(9).Infof("[%s] no update found at path: %q. Assuming the field is deleted.", label,
				diffPropPath.String())
			newConfigField = defaultIfMissing(oldState, diffPropPath)
		}
		// Overwrite the config object where the one value for fieldName has been updated.
		newConfig := resource.NewObjectProperty(resource.NewPropertyMapFromMap(nil))
		if targetPropPath == nil {
			targetPropPath = diffPropPath
		}
		// Now updatedConfigField essentially is a config object with just one field - `fieldName`
		updatedConfigField, ok := targetPropPath.Add(newConfig, newConfigField)
		if !ok {
			return fmt.Errorf("failed to inject config at property path: %q", targetPropPath.String())
		}
		logging.V(9).Infof("[%s] staged property value: %+v", label, updatedConfigField.ObjectValue().Mappable())
		body := p.prepareAPIInputs(updatedConfigField.ObjectValue(), oldState, map[string]resources.CloudAPIProperty{
			"update": {},
		})
		logging.V(9).Infof("[%s] cluster update request body: %v", label, body)
		op, err := retryRequest(p.client, "PUT", uri, "", body)
		if err != nil {
			return fmt.Errorf("error sending request: %s: %q %+v", err, uri, body)
		}

		_, err = p.waitForResourceOpCompletion(urn, res.Update.CloudAPIOperation, op, nil)
		if err != nil {
			return errors.Wrapf(err, "waiting for completion")
		}
		return nil
	}
}

// The ordering here loosely follows what is in the terraform provider which is itself somewhat arbitrary.
// There is some underlying value since there are a lot overlapping dependencies on mutations (e.g. setting
// network policies require the network policy add on etc.). This may need to change over time.
var clusterUpdateOrder = []string{
	"masterAuthorizedNetworksConfig",          //
	"addonsConfig",                            //
	"autoscaling",                             //
	"binaryAuthorization",                     //
	"shieldedNodes",                           //
	"releaseChannel",                          //
	"networkConfig.enableIntraNodeVisibility", //
	"networkConfig.privateIpv6GoogleAccess",   //
	"networkConfig.enableL4ilbSubsetting",     //
	"networkConfig.defaultSnatStatus",         //
	"maintenancePolicy",                       //
	"locations",                               //
	"legacyAbac.enabled",                      //
	"monitoringService",                       //
	"loggingService",                          //
	"networkPolicy",                           //
	//"nodePools",                               // TODO: Don't believe we can allow this to be changed?
	"initialClusterVersion", //
	// nodeVersion, // nodePools[*].version
	"nodeConfig.imageType",   //
	"notificationConfig",     //
	"verticalPodAutoscaling", //
	"meshCertificates",       //
	"databaseEncryption",     //
	"workloadIdentifyConfig", //
	"identityServiceConfig",  //
	// masterAuth has a pretty imperative API for mutations.
	// Its mostly deprecated anyway so lets mark it replace-on-change.
	// "masterAuth",
	// TODO: Enable these later. See note in overrides.go.
	//"loggingConfig",             //
	//"monitoringConfig",          //
	"resourceLabels",            //
	"resourceUsageExportConfig", //
	// Only beta
	"podSecurityPolicyConfig", //
	"clusterTelemetry",        //
	// removeDefaultNodepool,
	// Additional mutations:
	"authenticatorGroupsConfig",              //
	"networkConfig.datapathProvider",         //
	"networkConfig.dnsConfig",                //
	"nodeConfig.gcfsConfig",                  //
	"privateClusterConfig",                   //
	"networkConfig.serviceExternalIpsConfig", //
}

func waitForClusterRestingState(
	p *googleCloudProvider,
	urn resource.URN,
	res *resources.CloudAPIResource,
	inputs,
	oldState resource.PropertyMap,
) (map[string]interface{}, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancelFunc()
	retryPolicy := backoff.Backoff{
		Min:    1 * time.Second,
		Max:    15 * time.Second,
		Factor: 1.5,
		Jitter: true,
	}

	uri, err := res.Read.Endpoint.URI(inputs.Mappable(), oldState.Mappable())
	if err != nil {
		return nil, err
	}

	for {
		resp, err := readContainerResourceStatus(p, uri)
		if err != nil {
			return nil, err
		}

		status, hasStatus := resp["status"].(string)
		if hasStatus && isClusterInRestingStatus(status) {
			return resp, nil
		} else {
			_ = p.host.LogStatus(context.Background(), diag.Info, urn,
				fmt.Sprintf("Waiting for cluster to reach resting state. "+
					"Current status: %q", status))
		}
		select {
		case <-time.After(retryPolicy.Duration()):
		case <-ctx.Done():
			return nil, fmt.Errorf("timeout polling for nodepool resting state: %q", uri)
		}
	}
}

func isClusterInRestingStatus(status string) bool {
	switch status {
	case "RUNNING", "DEGRADED", "ERROR":
		return true
	}
	return false
}

// == Container Nodepool support ==

// nodepoolMutations models the custom mutations supported for Nodepool resource.
type nodepoolMutations struct {
	updateHandlers map[string]containerCustomUpdateHandlerFunc
}

func (n nodepoolMutations) mutablePropertyPaths() []string {
	return codegen.SortedKeys(n.updateHandlers)
}

func (n nodepoolMutations) update(
	providerInstance *googleCloudProvider,
	urn resource.URN,
	label string,
	res *resources.CloudAPIResource,
	inputs resource.PropertyMap,
	oldState resource.PropertyMap,
) (appliedInputs resource.PropertyMap, resp map[string]interface{}, err error) {
	// TODO wait for the cluster to be in resting state

	// Extract old inputs from the `__inputs` field of the old state.
	oldInputs := parseCheckpointObject(oldState)
	// Delete the auto-injected __autonamed tag in old and new inputs to avoid spurious diffs.
	if v, ok := oldInputs[autonamed]; ok && v.IsBool() {
		delete(oldInputs, autonamed)
	}

	if v, ok := inputs[autonamed]; ok && v.IsBool() {
		delete(inputs, autonamed)
	}

	diff := oldInputs.Diff(inputs)
	batchedInputs := resource.NewObjectProperty(oldInputs)

	if diff == nil {
		logging.V(9).Infof("[%s] no diff found in update!", label)
	} else {
		// Calculate the detailed diff object containing information about replacements.
		detailedDiff := calculateDetailedDiff(res, providerInstance.resourceMap.Types, diff)
		logging.V(9).Infof("[%s] detailedDiff: %+v", label, detailedDiff)
		logging.V(9).Infof("[%s] updateNodepool oldState: %+v", label, oldState)

		for _, key := range nodepoolUpdateOrder {
			updateHandler, ok := n.updateHandlers[key]
			if !ok {
				logging.V(9).Infof("[%s] No handler found for field: %q", label, key)
				// TODO This should be an error once all the targetted mutations are in place!
				continue
			}
			logging.V(9).Infof("[%s] Invoking update for field: %q", label, key)
			for k := range detailedDiff {
				logging.V(9).Infof("[%s] processing diff on field: %q", label, k)
				var diffPropPath, propPath resource.PropertyPath
				diffPropPath, err = resource.ParsePropertyPath(k)
				if err != nil {
					err = fmt.Errorf("failed to parse property path: %q: %w", k, err)
					break
				}
				propPath, err = resource.ParsePropertyPath(key)
				if err != nil {
					err = fmt.Errorf("failed to parse property path: %q: %w", k, err)
					break
				}
				if propPath.Contains(diffPropPath) {
					_ = providerInstance.host.LogStatus(context.Background(), diag.Info, urn,
						fmt.Sprintf("Performing update for field: %q", key))
					logging.V(9).Infof("[%s] processing diff for %q with handler for %q",
						label, diffPropPath.String(), propPath.String())
					logging.V(9).Infof("[%s] waiting for nodepool to reach resting state", label)
					_, err = waitForNodepoolRestingState(providerInstance, urn, res, inputs, oldState)
					if err != nil {
						break
					}
					newVal, _ := diffPropPath.Get(resource.NewObjectProperty(inputs))
					logging.V(9).Infof("[%s] calling update handler", label)
					err = updateHandler(providerInstance, urn, label, res, inputs, oldState)
					if err != nil {
						logging.V(9).Infof("[%s] failure updating: %+v", label, err)
						break
					}
					batchedInputs, ok = diffPropPath.Add(batchedInputs, newVal)
					if !ok {
						err = fmt.Errorf("failed to record update at field: %q", k)
						break
					}
				}
			}
			if err != nil {
				break
			}
		}
	}
	appliedInputs = batchedInputs.ObjectValue()
	uri, readErr := res.Read.Endpoint.URI(inputs.Mappable(), oldState.Mappable())
	if readErr != nil {
		return nil, nil, readErr
	}
	resp, err2 := readContainerResourceStatus(providerInstance, uri)
	if err2 != nil {
		logging.V(9).Infof("[%s] Failed to read nodepool status: %v", label, err2)
		return nil, nil, fmt.Errorf("failed to retrieve nodepool status: %v. Previous error: %w", err2, err)
	}

	return appliedInputs, resp, err
}

type containerCustomUpdateHandlerFunc func(
	p *googleCloudProvider,
	urn resource.URN,
	label string,
	res *resources.CloudAPIResource,
	newInputs,
	oldState resource.PropertyMap,
) error

var emptyObjVal = resource.NewObjectProperty(resource.NewPropertyMapFromMap(nil))
var emptyArrayVal = resource.NewArrayProperty([]resource.PropertyValue{})

type propValueIfMissingFunc func(oldState resource.PropertyMap, propPath resource.PropertyPath) resource.PropertyValue

func extractFromDefaults(valIfMissing resource.PropertyValue) propValueIfMissingFunc {
	return func(oldState resource.PropertyMap, propPath resource.PropertyPath) resource.PropertyValue {
		defaults := parseDefaultsObject(oldState)
		if val, ok := propPath.Get(resource.NewObjectProperty(defaults)); ok {
			return val
		}
		return valIfMissing
	}
}

func defaultValue(defVal resource.PropertyValue) propValueIfMissingFunc {
	return func(_ resource.PropertyMap, _ resource.PropertyPath) resource.PropertyValue {
		return defVal
	}
}

// Following an ordering here that the terraform provider performs updates in.
// We have a lot more mutations defined than the terraform providers does though.
var nodepoolUpdateOrder = []string{
	"autoscaling",
	"config.imageType",
	"config.workloadMetadataConfig",
	"config.kubeletConfig",
	"config.linuxNodeConfig",
	"config.confidentialNodes",
	"config.gcfsConfig",
	"config.gvnic",
	"config.labels",
	"config.tags",
	"config.taints",
	"networkConfig",
	"initialNodeCount",
	"management",
	"version",
	"locations",
	"upgradeSettings"}

func mustParsePropertyPath(pattern string) resource.PropertyPath {
	pp, err := resource.ParsePropertyPath(pattern)
	contract.AssertNoError(err)
	return pp
}

func readContainerResourceStatus(
	p *googleCloudProvider,
	uri string,
) (map[string]interface{}, error) {
	// Read the current state of the resource from the API.
	newState, err := retryRequest(p.client, "GET", uri, "", nil)
	if err != nil {
		return nil, fmt.Errorf("error fetching nodepool status: %w", err)
	}
	return newState, nil
}

func waitForNodepoolRestingState(
	p *googleCloudProvider,
	urn resource.URN,
	res *resources.CloudAPIResource,
	inputs,
	oldState resource.PropertyMap,
) (map[string]interface{}, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancelFunc()
	retryPolicy := backoff.Backoff{
		Min:    1 * time.Second,
		Max:    15 * time.Second,
		Factor: 1.5,
		Jitter: true,
	}

	uri, err := res.Read.Endpoint.URI(inputs.Mappable(), oldState.Mappable())
	if err != nil {
		return nil, err
	}

	for {
		resp, err := readContainerResourceStatus(p, uri)
		if err != nil {
			return nil, err
		}

		status, hasStatus := resp["status"].(string)
		if hasStatus && isNodepoolInRestingStatus(status) {
			return resp, nil
		} else {
			_ = p.host.LogStatus(context.Background(), diag.Info, urn,
				fmt.Sprintf("Waiting for nodepool to reach resting state. "+
					"Current status: %q", status))
		}
		select {
		case <-time.After(retryPolicy.Duration()):
		case <-ctx.Done():
			return nil, fmt.Errorf("timeout polling for nodepool resting state: %q", uri)
		}
	}

}

func isNodepoolInRestingStatus(status string) bool {
	switch status {
	case "RUNNING", "RUNNING_WITH_ERROR", "ERROR":
		return true
	}
	return false
}

func updateNodePoolMapping(
	diffPropPath resource.PropertyPath,
	apiFieldName string,
	defaultIfMissing propValueIfMissingFunc,
	operationSuffix string,
	httpMethod string,
) containerCustomUpdateHandlerFunc {
	return func(p *googleCloudProvider,
		urn resource.URN,
		label string,
		res *resources.CloudAPIResource,
		newInputs,
		oldState resource.PropertyMap) error {
		tok := urn.Type()
		version := strings.Split(tok.Module().String(), "/")[1]
		endpoint := resources.CloudAPIEndpoint{
			Template: fmt.Sprintf(`https://container.googleapis.com/%s/`+
				`projects/{projectsId}/locations/{locationsId}/clusters/{clustersId}/nodePools/{nodePoolsId}%s`,
				version, operationSuffix),
			Values: append(res.Create.Endpoint.Values,
				resources.CloudAPIResourceParam{
					Name:    "nodePoolsId",
					SdkName: "name",
					Kind:    "path",
				}),
		}
		uri, err := endpoint.URI(newInputs.Mappable(), oldState.Mappable())
		if err != nil {
			return err
		}

		// If key is omitted, use default.
		_, ok := diffPropPath.Get(resource.NewObjectProperty(newInputs))
		if !ok {
			updated, _ := diffPropPath.Add(resource.NewObjectProperty(newInputs), defaultIfMissing(oldState, diffPropPath))
			newInputs = updated.ObjectValue()
		}

		body := p.prepareAPIInputs(newInputs, oldState, map[string]resources.CloudAPIProperty{
			apiFieldName: {SdkName: diffPropPath.String()},
		})
		op, err := retryRequest(p.client, httpMethod, uri, "", body)
		if err != nil {
			return fmt.Errorf("error sending request: %s: %q %+v", err, uri, body)
		}

		_, err = p.waitForResourceOpCompletion(urn, res.Update.CloudAPIOperation, op, nil)
		if err != nil {
			return errors.Wrapf(err, "waiting for completion")
		}
		return nil
	}
}

func updateNodePoolConfig(diffPropPath, targetPropPath resource.PropertyPath,
	defaultIfMissing propValueIfMissingFunc) containerCustomUpdateHandlerFunc {
	return func(p *googleCloudProvider,
		urn resource.URN,
		label string,
		res *resources.CloudAPIResource,
		newInputs,
		oldState resource.PropertyMap) error {
		logging.V(9).Infof("[%s] updateNodePoolConfig called for property path: %q", label, diffPropPath.String())
		tok := urn.Type()
		version := strings.Split(tok.Module().String(), "/")[1]
		endpoint := resources.CloudAPIEndpoint{
			Template: fmt.Sprintf(`https://container.googleapis.com/%s/`+
				`projects/{projectsId}/locations/{locationsId}/clusters/{clustersId}/nodePools/{nodePoolsId}`, version),
			Values: append(res.Create.Endpoint.Values,
				resources.CloudAPIResourceParam{
					Name:    "nodePoolsId",
					SdkName: "name",
					Kind:    "path",
				}),
		}
		uri, err := endpoint.URI(newInputs.Mappable(), oldState.Mappable())
		if err != nil {
			return err
		}
		// Get the new config object to copy the value for fieldName from
		newConfigField, ok := diffPropPath.Get(resource.NewObjectProperty(newInputs))
		if !ok {
			logging.V(9).Infof("[%s] no update found at path: %q. Assuming the field is deleted.", label,
				diffPropPath.String())
			newConfigField = defaultIfMissing(oldState, diffPropPath)
		}
		// Overwrite the config object where the one value for fieldName has been updated.
		newConfig := resource.NewObjectProperty(resource.NewPropertyMapFromMap(nil))
		if targetPropPath == nil {
			targetPropPath = diffPropPath
		}
		// Now updatedConfigField essentially is a config object with just one field - `fieldName`
		updatedConfigField, ok := targetPropPath.Add(newConfig, newConfigField)
		if !ok {
			return fmt.Errorf("failed to inject node config at property path: %q", targetPropPath.String())
		}
		logging.V(9).Infof("[%s] staged property value: %+v", label, updatedConfigField.ObjectValue().Mappable())
		body := p.prepareAPIInputs(updatedConfigField.ObjectValue(), oldState, map[string]resources.CloudAPIProperty{
			"config": {Flatten: true},
		})
		logging.V(9).Infof("[%s] nodepool config update request body: %v", label, body)
		op, err := retryRequest(p.client, "PUT", uri, "", body)
		if err != nil {
			return fmt.Errorf("error sending request: %s: %q %+v", err, uri, body)
		}

		_, err = p.waitForResourceOpCompletion(urn, res.Update.CloudAPIOperation, op, nil)
		if err != nil {
			return errors.Wrapf(err, "waiting for completion")
		}
		return nil
	}
}

// *** Support for method invocations on container resources ***
// Most of this was adapted from a pattern established in pulumi-yaml:
// https://github.com/pulumi/pulumi-yaml/blob/a36000e5bf6b64c050ec8777578aea3630dc1b7f/pkg/pulumiyaml/run.go

// callableResource models the receiver resource for a method call in generic terms.
// While currently we only have the getKubeconfig method on the cluster resource,
// this can be reused for others in the future and should move to a separate package/file.
// This bag-based output shape defined here helps us avoid adding an unholy dependency on the
// generated Go SDK here (albeit which might have given us strongly typed outputs instead).
type callableResource interface {
	GetOutput(k string) pulumi.Output
	CustomResource() *pulumi.CustomResourceState
	ProviderResource() *pulumi.ProviderResourceState
}

type customResourceState struct {
	pulumi.CustomResourceState
	name    string
	Outputs pulumi.MapOutput `pulumi:""`
}

// GetOutput returns the named output of the resource.
func (st *customResourceState) GetOutput(k string) pulumi.Output {
	return st.Outputs.ApplyT(func(outputs map[string]interface{}) (interface{}, error) {
		out, ok := outputs[k]
		if !ok {
			return nil, fmt.Errorf("no output %q on resource %q", k, st.name)
		}
		return out, nil
	})
}

func (st *customResourceState) CustomResource() *pulumi.CustomResourceState {
	return &st.CustomResourceState
}

func (st *customResourceState) ProviderResource() *pulumi.ProviderResourceState {
	return nil
}

func (*customResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*callableResource)(nil)).Elem()
}

// clusterGetKubeconfigResult is the representation of the result from the getKubeconfig method call.
// Note the use of `liftSingleValueMethodReturns` in the schema is a codegen construct and applies
// in the consuming SDK. The schema still only supports object types in the output.
type clusterGetKubeconfigResult struct {
	Kubeconfig pulumi.StringOutput `pulumi:"kubeconfig"`
}

const kubeconfigFmt = `apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: %[3]s
    server: https://%[2]s
  name: %[1]s
contexts:
- context:
    cluster: %[1]s
    user: %[1]s
  name: %[1]s
current-context: %[1]s
kind: Config
preferences: {}
users:
- name: %[1]s
  user:
    exec:
      apiVersion: client.authentication.k8s.io/v1beta1
      command: gke-gcloud-auth-plugin
      installHint: Install gke-gcloud-auth-plugin for use with kubectl by following
        https://cloud.google.com/blog/products/containers-kubernetes/kubectl-auth-changes-in-gke
      provideClusterInfo: true
`

// getKubeConfigCallHandler returns a function which converts the getKubeconfig Call GRPC invocation to Pulumi's Go
// programming model.
func getKubeConfigCallHandler(label, tok string, callArgs resource.PropertyMap) func(*pulumi.Context, string,
	pulumiprov.CallArgs) (*pulumiprov.CallResult, error) {
	return func(ctx *pulumi.Context, tok string,
		args pulumiprov.CallArgs) (*pulumiprov.CallResult, error) {
		var res pulumi.Resource
		var state callableResource
		// If we get here, callArgs will already have a `__self__` and will be known resource reference.
		self := callArgs["__self__"]
		if !self.IsResourceReference() {
			return nil, fmt.Errorf("call(%s) __self__ was expected to be a resource reference", tok)
		}
		ref := self.ResourceReferenceValue()
		r := customResourceState{name: ""}
		res = &r
		state = &r
		// This hydrates the resource state from the resource reference.
		err := ctx.RegisterResource("_", "_", nil, res, pulumi.URN_(string(ref.URN)))
		if err != nil {
			return nil, fmt.Errorf("failed to get resource: %q: %w", ref.URN, err)
		}
		logging.V(9).Infof("[%s] resource: %#v", label, state)

		kubeconfig := pulumi.Sprintf(
			kubeconfigFmt,
			state.GetOutput("name"),
			state.GetOutput("endpoint"),
			state.GetOutput("masterAuth").ApplyT(func(ma interface{}) string {
				return ma.(map[string]interface{})["clusterCaCertificate"].(string)
			}))
		result := clusterGetKubeconfigResult{
			Kubeconfig: kubeconfig,
		}
		return pulumiprov.NewCallResult(&result)
	}
}
