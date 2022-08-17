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

	"github.com/pulumi/pulumi-google-native/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"google.golang.org/api/cloudkms/v1"
	"google.golang.org/api/option"
)

type customMutationsHandler interface {
	mutablePropertyPaths() []string
	update(providerInstance *googleCloudProvider,
		urn resource.URN,
		label string,
		res *resources.CloudAPIResource,
		inputs resource.PropertyMap,
		oldState resource.PropertyMap,
	) (appliedInputs resource.PropertyMap, resp map[string]interface{}, err error)
}

var clusterMutationHandlers = clusterMutations{
	updateHandlers: map[string]containerCustomUpdateHandlerFunc{
		"masterAuthorizedNetworksConfig": updateClusterNestedField(
			mustParsePropertyPath("masterAuthorizedNetworksConfig"),
			mustParsePropertyPath("update.desiredMasterAuthorizedNetworksConfig"),
			defaultValue(emptyObjVal),
			"",
			"PUT"),
		"addonsConfig": updateClusterMapping(
			mustParsePropertyPath("addonsConfig"),
			"addonsConfig",
			defaultValue(emptyObjVal),
			":setAddons",
			"POST",
			nil),
		"legacyAbac.enabled": updateClusterNestedField(
			mustParsePropertyPath("legacyAbac.enabled"),
			mustParsePropertyPath("enabled"),
			extractFromDefaults(resource.NewBoolProperty(false)),
			":setLegacyAbac",
			"POST"),
		"locations": updateClusterMapping(
			mustParsePropertyPath("locations"),
			"locations",
			defaultValue(resource.NewArrayProperty([]resource.PropertyValue{})),
			":setLocations",
			"POST",
			nil),
		"maintenancePolicy": updateClusterMapping(
			mustParsePropertyPath("maintenancePolicy"),
			"maintenancePolicy",
			defaultValue(emptyObjVal),
			":setMaintenancePolicy",
			"POST",
			nil),
		// How to set `action` field?
		//"masterAuth.password": updateClusterNestedField(
		//	mustParsePropertyPath("masterAuth.password"),
		//	nil,
		//	defaultValue(emptyObjVal),
		//	":setMasterAuth",
		//	"POST"),
		// The following endpoints don't actually work. The API requires both monitoring and logging service
		// to be modified simultaneously which is possible only through the main update API.
		//"monitoringService": updateClusterMapping(
		//	mustParsePropertyPath("monitoringService"),
		//	"monitoringService",
		//	extractFromDefaults(resource.NewStringProperty("")),
		//	":setMonitoring",
		//	"POST",
		//	nil),
		//"loggingService": updateClusterMapping(
		//	mustParsePropertyPath("loggingService"),
		//	"loggingService",
		//	extractFromDefaults(resource.NewStringProperty("")),
		//	":setLogging",
		//	"POST",
		//	nil),
		"networkPolicy": updateClusterMapping(
			mustParsePropertyPath("networkPolicy"),
			"networkPolicy",
			defaultValue(emptyObjVal),
			":setNetworkPolicy",
			"POST",
			nil),
		"resourceLabels": updateClusterMapping(
			mustParsePropertyPath("resourceLabels"),
			"resourceLabels",
			defaultValue(emptyObjVal),
			":setResourceLabels",
			"POST",
			map[string]resources.CloudAPIProperty{"labelFingerprint": {CopyFromOutputs: true}}),
		"initialClusterVersion": updateClusterMapping(
			mustParsePropertyPath("initialClusterVersion"),
			"masterVersion",
			extractFromDefaults(resource.NewStringProperty("")),
			":updateMaster",
			"POST",
			nil),
		"authenticatorGroupsConfig": updateClusterNestedField(
			mustParsePropertyPath("authenticatorGroupsConfig"),
			mustParsePropertyPath("update.desiredAuthenticatorGroupsConfig"),
			defaultValue(emptyObjVal),
			"",
			"PUT"),
		"binaryAuthorization": updateClusterNestedField(
			mustParsePropertyPath("binaryAuthorization"),
			mustParsePropertyPath("update.desiredBinaryAuthorization"),
			defaultValue(emptyObjVal),
			"",
			"PUT"),
		"autoscaling": updateClusterNestedField(
			mustParsePropertyPath("autoscaling"),
			mustParsePropertyPath("update.desiredClusterAutoscaling"),
			defaultValue(emptyObjVal),
			"",
			"PUT"),
		"shieldedNodes": updateClusterNestedField(
			mustParsePropertyPath("shieldedNodes"),
			mustParsePropertyPath("update.desiredShieldedNodes"),
			defaultValue(emptyObjVal),
			"",
			"PUT"),
		"releaseChannel": updateClusterNestedField(
			mustParsePropertyPath("releaseChannel"),
			mustParsePropertyPath("update.desiredReleaseChannel"),
			defaultValue(emptyObjVal),
			"",
			"PUT"),
		"resourceUsageExportConfig": updateClusterNestedField(
			mustParsePropertyPath("resourceUsageExportConfig"),
			mustParsePropertyPath("update.desiredResourceUsageExportConfig"),
			defaultValue(emptyObjVal),
			"",
			"PUT"),
		"networkConfig.serviceExternalIpsConfig": updateClusterNestedField(
			mustParsePropertyPath("networkConfig.serviceExternalIpsConfig"),
			mustParsePropertyPath("update.desiredServiceExternalIpsConfig"),
			defaultValue(emptyObjVal),
			"",
			"PUT"),
		"networkConfig.enableIntraNodeVisibility": updateClusterNestedField(
			mustParsePropertyPath("networkConfig.enableIntraNodeVisibility"),
			mustParsePropertyPath("update.desiredIntraNodeVisibilityConfig.enabled"),
			defaultValue(resource.NewBoolProperty(false)),
			"",
			"PUT"),
		"networkConfig.privateIpv6GoogleAccess": updateClusterNestedField(
			mustParsePropertyPath("networkConfig.privateIpv6GoogleAccess"),
			mustParsePropertyPath("update.desiredPrivateIpv6GoogleAccess"),
			extractFromDefaults(resource.NewStringProperty("")),
			"",
			"PUT"),
		"networkConfig.enableL4ilbSubsetting": updateClusterNestedField(
			mustParsePropertyPath("networkConfig.enableL4ilbSubsetting"),
			mustParsePropertyPath("update.desiredL4ilbSubsettingConfig"),
			defaultValue(emptyObjVal),
			"",
			"PUT"),
		"nodeConfig.imageType": updateClusterNestedField(
			mustParsePropertyPath("nodeConfig.imageType"),
			mustParsePropertyPath("update.desiredImageType"),
			extractFromDefaults(resource.NewStringProperty("")),
			"",
			"PUT"),
		"notificationConfig": updateClusterNestedField(
			mustParsePropertyPath("notificationConfig"),
			mustParsePropertyPath("update.desiredNotificationConfig"),
			defaultValue(emptyObjVal),
			"",
			"PUT"),
		"verticalPodAutoscaling": updateClusterNestedField(
			mustParsePropertyPath("verticalPodAutoscaling"),
			mustParsePropertyPath("update.desiredVerticalPodAutoscaling"),
			defaultValue(emptyObjVal),
			"",
			"PUT"),
		"meshCertificates": updateClusterNestedField(
			mustParsePropertyPath("meshCertificates"),
			mustParsePropertyPath("update.desiredMeshCertificates"),
			defaultValue(emptyObjVal),
			"",
			"PUT"),
		"databaseEncryption": updateClusterNestedField(
			mustParsePropertyPath("databaseEncryption"),
			mustParsePropertyPath("update.desiredDatabaseEncryption"),
			defaultValue(emptyObjVal),
			"",
			"PUT"),
		"identityServiceConfig": updateClusterNestedField(
			mustParsePropertyPath("identityServiceConfig"),
			mustParsePropertyPath("update.desiredIdentityServiceConfig"),
			defaultValue(emptyObjVal),
			"",
			"PUT"),
		"workloadIdentityConfig": updateClusterNestedField(
			mustParsePropertyPath("workloadIdentityConfig"),
			mustParsePropertyPath("update.desiredWorkloadIdentityConfig"),
			defaultValue(emptyObjVal),
			"",
			"PUT"),
		"loggingConfig": updateClusterNestedField(
			mustParsePropertyPath("loggingConfig"),
			mustParsePropertyPath("update.desiredLoggingConfig"),
			defaultValue(emptyObjVal),
			"",
			"PUT"),
		"monitoringConfig": updateClusterNestedField(
			mustParsePropertyPath("monitoringConfig"),
			mustParsePropertyPath("update.desiredMonitoringConfig"),
			defaultValue(emptyObjVal),
			"",
			"PUT"),
		"networkConfig.datapathProvider": updateClusterNestedField(
			mustParsePropertyPath("networkConfig.datapathProvider"),
			mustParsePropertyPath("update.desiredDatapathProvider"),
			defaultValue(resource.NewStringProperty("DATAPATH_PROVIDER_UNSPECIFIED")),
			"",
			"PUT"),
		"networkConfig.dnsConfig": updateClusterNestedField(
			mustParsePropertyPath("networkConfig.dnsConfig"),
			mustParsePropertyPath("update.desiredDnsConfig"),
			defaultValue(emptyObjVal),
			"",
			"PUT"),
		"networkConfig.gcfsConfig": updateClusterNestedField(
			mustParsePropertyPath("networkConfig.gcfsConfig"),
			mustParsePropertyPath("update.desiredGcfsConfig"),
			defaultValue(emptyObjVal),
			"",
			"PUT"),
		"privateClusterConfig": updateClusterNestedField(
			mustParsePropertyPath("privateClusterConfig"),
			mustParsePropertyPath("update.desiredPrivateClusterConfig"),
			defaultValue(emptyObjVal),
			"",
			"PUT"),
		// V1 Beta1 only
		"podSecurityPolicyConfig": updateClusterNestedField(
			mustParsePropertyPath("podSecurityPolicyConfig"),
			mustParsePropertyPath("update.desiredPodSecurityPolicyConfig"),
			defaultValue(emptyObjVal),
			"",
			"PUT"),
		"clusterTelemetry": updateClusterNestedField(
			mustParsePropertyPath("clusterTelemetry"),
			mustParsePropertyPath("update.desiredClusterTelemetry"),
			defaultValue(emptyObjVal),
			"",
			"PUT"),
	}}

var nodepoolMutationHanders = nodepoolMutations{
	updateHandlers: map[string]containerCustomUpdateHandlerFunc{
		"autoscaling": updateNodePoolMapping(
			mustParsePropertyPath("autoscaling"),
			"autoscaling",
			defaultValue(resource.NewObjectProperty(
				resource.PropertyMap{
					"enabled": resource.NewBoolProperty(false),
				})),
			":setAutoscaling",
			"POST"),
		"management": updateNodePoolMapping(
			mustParsePropertyPath("management"),
			"management",
			defaultValue(emptyObjVal),
			":setManagement",
			"POST"),
		"upgradeSettings": updateNodePoolMapping(
			mustParsePropertyPath("upgradeSettings"),
			"upgradeSettings",
			defaultValue(emptyObjVal),
			"",
			"PUT"),
		"locations": updateNodePoolMapping(
			mustParsePropertyPath("locations"),
			"locations",
			defaultValue(resource.NewArrayProperty([]resource.PropertyValue{})),
			"",
			"PUT"),
		"version": updateNodePoolMapping(
			mustParsePropertyPath("version"),
			"nodeVersion",
			extractFromDefaults(resource.NewStringProperty("")),
			"",
			"PUT"),
		"networkConfig": updateNodePoolMapping(
			mustParsePropertyPath("networkConfig"),
			"nodeNetworkConfig",
			defaultValue(emptyObjVal),
			"",
			"PUT"),
		"initialNodeCount": updateNodePoolMapping(
			mustParsePropertyPath("initialNodeCount"),
			"nodeCount",
			defaultValue(resource.NewNumberProperty(0)),
			":setSize",
			"POST"),
		"config.confidentialNodes": updateNodePoolConfig(mustParsePropertyPath("config.confidentialNodes"), nil,
			defaultValue(emptyObjVal)),
		"config.gcfsConfig": updateNodePoolConfig(mustParsePropertyPath("config.gcfsConfig"), nil,
			defaultValue(emptyObjVal)),
		"config.gvnic": updateNodePoolConfig(mustParsePropertyPath("config.gvnic"), nil,
			defaultValue(emptyObjVal)),
		"config.imageType": updateNodePoolConfig(mustParsePropertyPath("config.imageType"), nil,
			extractFromDefaults(resource.NewStringProperty(""))),
		"config.kubeletConfig": updateNodePoolConfig(mustParsePropertyPath("config.kubeletConfig"), nil,
			defaultValue(emptyObjVal)),
		"config.labels": updateNodePoolConfig(mustParsePropertyPath("config.labels"),
			mustParsePropertyPath("config.labels.labels"),
			defaultValue(emptyObjVal)),
		"config.linuxNodeConfig": updateNodePoolConfig(mustParsePropertyPath("config.linuxNodeConfig"), nil,
			defaultValue(emptyObjVal)),
		"config.tags": updateNodePoolConfig(mustParsePropertyPath("config.tags"),
			mustParsePropertyPath("config.tags.tags"),
			defaultValue(emptyArrayVal)),
		"config.taints": updateNodePoolConfig(mustParsePropertyPath("config.taints"),
			mustParsePropertyPath("config.taints.taints"),
			defaultValue(emptyArrayVal)),
		"config.workloadMetadataConfig": updateNodePoolConfig(mustParsePropertyPath("config.workloadMetadataConfig"), nil,
			defaultValue(emptyObjVal)),
	},
}

var customMutations = map[string]customMutationsHandler{
	"google-native:container/v1:Cluster":       clusterMutationHandlers,
	"google-native:container/v1beta1:Cluster":  clusterMutationHandlers,
	"google-native:container/v1:NodePool":      nodepoolMutationHanders,
	"google-native:container/v1beta1:NodePool": nodepoolMutationHanders,
}

var ResourceDeleteOverrides = map[string]func(
	providerInstance *googleCloudProvider,
	res resources.CloudAPIResource,
	inputs, outputs map[string]interface{},
) error{
	"google-native:cloudkms/v1:CryptoKey": func(
		providerInstance *googleCloudProvider,
		res resources.CloudAPIResource,
		inputs, outputs map[string]interface{},
	) error {
		// TODO: may need to pass ctx for cancellation
		ctx := context.Background()
		clientkms, err := cloudkms.NewService(ctx, option.WithHTTPClient(providerInstance.client.HTTPClient()))
		if err != nil {
			return err
		}
		name := outputs["name"].(string)

		versionsClient := clientkms.Projects.Locations.KeyRings.CryptoKeys.CryptoKeyVersions
		versionsResponse, err := versionsClient.List(name).Do()
		if err != nil {
			return err
		}

		// Destroy all versions of this CryptoKey.
		for _, version := range versionsResponse.CryptoKeyVersions {
			request := &cloudkms.DestroyCryptoKeyVersionRequest{}
			_, err = versionsClient.Destroy(version.Name, request).Do()
			if err != nil {
				return err
			}
		}

		// If key rotation is set, disable to avoid creating new keys in the future.
		if _, ok := outputs["rotationPeriod"]; ok {
			keyClient := clientkms.Projects.Locations.KeyRings.CryptoKeys
			_, err = keyClient.
				Patch(
					name,
					&cloudkms.CryptoKey{
						NullFields: []string{"rotationPeriod", "nextRotationTime"},
					}).
				UpdateMask("rotationPeriod,nextRotationTime").Do()
			if err != nil {
				return err
			}
		}

		return nil
	},
}
