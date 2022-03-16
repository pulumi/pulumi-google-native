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
	"google.golang.org/api/cloudkms/v1"
	"google.golang.org/api/option"
)

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
