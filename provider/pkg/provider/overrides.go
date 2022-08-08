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

var customMutations = map[string]customMutationsHandler{
	"google-native:container/v1:NodePool": nodepoolMutations{
		updateHandlers: map[string]nodepoolUpdateHandlerFunc{
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
	},
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
