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
	"encoding/json"
	"fmt"
	"log"

	"github.com/pulumi/pulumi-google-native/provider/pkg/resources"
)

// knativeStatusCheck checks if the knative resource is ready. If so it returns true as the first return type.
// If the resource is not ready, returns false and nil for error. Lastly if there is an error, returns false and
// an error as the second return.
func knativeStatusCheck(res map[string]interface{}) (bool, error) {
	m, err := extractMetadata(res)
	if err != nil {
		return false, err
	}

	gen, ok := m["generation"]
	if !ok {
		return false, fmt.Errorf("unable to find generation in knative metadata")
	}
	generation := int(gen.(float64))

	// 2. Get status from response.
	s := KnativeStatus{}
	if err := convert(res, &s); err != nil {
		return false, fmt.Errorf("failed to parse knative status: %w", err)
	}

	// 3. Check if observed and actual generation match. If not, continue to poll
	if int(s.Status.ObservedGeneration) != generation {
		return false, nil
	}

	// 4. Check conditions to ensure resource is ready.
	for _, condition := range s.Status.Conditions {
		if condition.Type == readyStatusType {
			log.Printf("[DEBUG] checking KnativeStatus Ready condition %s: %s", condition.Status, condition.Message)
			switch condition.Status {
			case "True":
				// 4. a. Resource is ready if condition type is `Ready` status is `True`
				return true, nil
			case "Unknown":
				// 4. b. DomainMapping can enter a 'terminal' state where "Ready" status is "Unknown"
				// but the resource is waiting for external verification of DNS records.
				if condition.Reason == pendingCertificateReason {
					return true, nil
				}
				// 4. c. Otherwise, not ready as yet.
				return false, nil
			case "False":
				// 5. If the condition status is False, the resource has failed.
				return false, fmt.Errorf(`resource is in failed state "Ready:False", message: %s`,
					condition.Message)
			}
		}
	}
	return false, nil
}

func extractMetadata(res map[string]interface{}) (map[string]interface{}, error) {
	// 1. Get generation from metadata
	metadata, ok := res["metadata"]
	if !ok {
		return nil, fmt.Errorf("unable to find knative metadata")
	}

	m, ok := metadata.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unable to find generation in knative metadata")
	}
	return m, nil
}

func getKNativeSelfLinkURL(res map[string]interface{}) (string, error) {
	m, err := extractMetadata(res)
	if err != nil {
		return "", err
	}

	sl, ok := m["selfLink"]
	if !ok {
		return "", fmt.Errorf("missing selfLink metadata field in response")
	}
	asString, ok := sl.(string)
	if !ok {
		return "", fmt.Errorf("malformed selfLink metadata field in response")
	}

	labels, ok := m["labels"]
	if !ok {
		return "", fmt.Errorf("labels not found in response metadata")
	}

	labelsMap, ok := labels.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("malformed labels in response metadata")
	}
	location, ok := labelsMap["cloud.googleapis.com/location"]
	if !ok {
		return "", fmt.Errorf("missing location label in response metadata")
	}

	return resources.AssembleURL(fmt.Sprintf("https://%s-run.googleapis.com", location), asString), nil
}

func convert(item, out interface{}) error {
	bytes, err := json.Marshal(item)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, out)
	if err != nil {
		return err
	}

	return nil
}

const readyStatusType string = "Ready"
const pendingCertificateReason string = "CertificatePending"

type Condition struct {
	Type    string
	Status  string
	Reason  string
	Message string
}

// KnativeStatus is a struct that can contain a Knative style resource's Status block. It is not
// intended to be used for anything other than polling for the success of the given resource.
type KnativeStatus struct {
	Metadata struct {
		Name      string
		Namespace string
		SelfLink  string
	}
	Status struct {
		Conditions         []Condition
		ObservedGeneration float64
	}
}
