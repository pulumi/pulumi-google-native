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

package resources

import (
	"fmt"
	"strings"
)

// CloudAPIMetadata is a collection of all resources and functions in the Google Cloud REST API.
type CloudAPIMetadata struct {
	Types     map[string]CloudAPIType     `json:"types"`
	Resources map[string]CloudAPIResource `json:"resources"`
	Functions map[string]CloudAPIFunction `json:"functions"`
}

// CloudAPIEndpoint contains information needed to construct a REST endpoint.
type CloudAPIEndpoint struct {
	// SelfLinkProperty is the name of the resource property containing a selfLink.
	// Example: `selfLink`
	SelfLinkProperty string `json:"selfLinkProperty,omitempty"`
	// Template is the template string for a REST endpoint.
	// Example: https://sqladmin.googleapis.com/v1/projects/{project}/instances/{instance}/sslCerts
	Template string `json:"template,omitempty"`
	// Values are the concrete values needed to instantiate the template.
	Values []CloudAPIResourceParam `json:"values,omitempty"`
}

// URI constructs a concrete URI value based on the properties of a created resource.
func (e CloudAPIEndpoint) URI(
	inputs map[string]interface{},
	outputs map[string]interface{},
) (string, error) {
	if len(e.SelfLinkProperty) > 0 {
		v, ok := outputs[e.SelfLinkProperty].(string)
		if !ok {
			return "", fmt.Errorf("selfLink property %q not found", e.SelfLinkProperty)
		}
		return v, nil
	}

	id := e.Template
	idParams := e.Values
	for _, param := range idParams {
		var propValue string
		if v, has := EvalPropertyValue(inputs, param.SdkName); has {
			propValue = v
		} else if v, has := EvalPropertyValue(outputs, param.SdkName); has {
			propValue = v
		}

		if propValue == "" {
			return "", fmt.Errorf("property %q/%q not found", param.Name, param.SdkName)
		}

		// The name property can sometimes contain multiple segments. We only care about the last one
		// because we flattened the path while building metadata.
		parts := strings.Split(propValue, "/")
		propValue = parts[len(parts)-1]

		id = strings.Replace(id, fmt.Sprintf("{%s}", param.Name), propValue, 1)
	}
	return id, nil
}

func EvalPropertyValue(values map[string]interface{}, path string) (string, bool) {
	current := values
	parts := strings.Split(path, ".")
	for idx, part := range parts {
		value := current[part]
		if idx == len(parts)-1 {
			if str, ok := value.(string); ok {
				return str, true
			}
			return "", false
		}
		childMap, ok := value.(map[string]interface{})
		if !ok {
			return "", false
		}
		current = childMap
	}
	return "", false
}

// CloudAPIOperation is a resource operation in the Google Cloud REST API.
type CloudAPIOperation struct {
	// Endpoint is the REST endpoint for the operation.
	Endpoint CloudAPIEndpoint `json:"endpoint,omitempty"`
	// SDKProperties define metadata about the SDK representation of this resource.
	SDKProperties map[string]CloudAPIProperty `json:"sdkProperties,omitempty"`
	// Verb is the REST verb to use for API calls.
	Verb    string   `json:"verb,omitempty"`
	Polling *Polling `json:"polling,omitempty"`
}

func (o CloudAPIOperation) Undefined() bool {
	return len(o.Verb) == 0
}

// ResourceAutoname defines the information to auto-name a resource in the Google Cloud REST API.
type ResourceAutoname struct {
	// FieldName contains a string pattern when a default name should be automatically generated if a user hasn't
	// explicitly specified one.
	// Example: `projects/{project}/locations/{location}/functions/{name}`
	FieldName string `json:"fieldName,omitempty"`
	// ValidationRegex contains a string pattern defining the valid name pattern for a resource.
	ValidationRegex string `json:"validationRegex,omitempty"`
}

type PollingStrategy string

const (
	DefaultPoll       = PollingStrategy("")
	KNativeStatusPoll = PollingStrategy("KNativeStatusPoll")
)

// Polling specifies the polling strategy to use for a resource.
type Polling struct {
	// Strategy defines the polling strategy to use for this resource. If empty,
	// the default polling behavior around operations is in effect.
	Strategy PollingStrategy `json:"strategy,omitempty"`
}

// CreateAPIOperation is a Create resource operation in the Google Cloud REST API.
type CreateAPIOperation struct {
	CloudAPIOperation
	Autoname ResourceAutoname `json:"autoname,omitempty"`
}

type UpdateAPIOperation struct {
	CloudAPIOperation

	UpdateMask UpdateMask `json:"updateMask,omitempty"`
}

// UpdateMask specifies where update mask field is specified for updates.
type UpdateMask struct {
	// BodyPropertyName identifies an additional field which must be specified as field-mask in the request body.
	BodyPropertyName string `json:"bodyPropertyName,omitempty"`
	// QueryParamName identifies an additional query param name that must be specified as field-mask in the query
	// params.
	QueryParamName string `json:"queryParamName,omitempty"`
}

// CloudAPIResource is a resource in the Google Cloud REST API.
type CloudAPIResource struct {
	Create CreateAPIOperation `json:"create,omitempty"`
	Read   CloudAPIOperation  `json:"read,omitempty"`
	Update UpdateAPIOperation `json:"update,omitempty"`
	Delete CloudAPIOperation  `json:"delete,omitempty"`

	// RootURL is the root URL of the REST API.
	// Example: `https://cloudkms.googleapis.com/`
	RootURL     string `json:"rootUrl"`
	AssetUpload bool   `json:"assetUpload,omitempty"`
	// IDProperty contains the name of the output property that represents resource ID (a self link).
	// Example: `selfLink`
	IDProperty string `json:"idProperty,omitempty"`
	// IDPath is the template for building resource ID with ID parameter values. It should only
	// be defined if IDProperty is missing.
	// Example: `projects/{project}/global/backendBuckets/{resource}/getIamPolicy`
	IDPath   string            `json:"idPath,omitempty"`
	IDParams map[string]string `json:"idParams,omitempty"`
}

// CloudAPIFunction is a function in Google Cloud REST API.
type CloudAPIFunction struct {
	URL  CloudAPIEndpoint `json:"url"`
	Verb string           `json:"verb"`
}

// CloudAPIResourceParam is a URL parameter of an API resource.
type CloudAPIResourceParam struct {
	// Name is the value of the parameter in the REST API.
	// Example: `projectsId`
	Name string `json:"name"`
	// SdkName is the value of the parameter in the Pulumi SDK.
	// Example: `project`
	SdkName string `json:"sdkName,omitempty"`
	// Kind is the kind of parameter, either "path" or "query"
	Kind string `json:"kind"`
}

// ResourceURL returns the resource API URL by joining the base URL with the resource path.
func (r *CloudAPIResource) ResourceURL(path string) string {
	return AssembleURL(r.RootURL, path)
}

// AssembleURL combines URL fragment strings into a single URL string.
func AssembleURL(fragments ...string) string {
	var b strings.Builder
	for i, fragment := range fragments {
		if i > 0 && len(fragment) > 0 {
			b.WriteString("/")
		}
		b.WriteString(strings.TrimPrefix(strings.TrimSuffix(fragment, "/"), "/"))
	}
	return b.String()
}

// CloudAPIProperty is a property of a body of an API call payload.
type CloudAPIProperty struct {
	Ref      string `json:"$ref,omitempty"`
	Format   string `json:"format,omitempty"`
	Required bool   `json:"required,omitempty"`

	Items                *CloudAPIProperty `json:"items,omitempty"`
	AdditionalProperties *CloudAPIProperty `json:"additionalProperties,omitempty"`
	// The name of the container property that was "flattened" during SDK generation, i.e. extra layer that exists
	// in the API payload but does not exist in the SDK.
	Container string `json:"container,omitempty"`
	SdkName   string `json:"sdkName,omitempty"`
	// CopyFromOutputs equal to true means that the value for this property during an update should be taken from
	// the previous state of the resource, not user inputs.
	CopyFromOutputs bool `json:"copyFromOutputs,omitempty"`
	// Pattern contains a string pattern that the property value needs to adhere to.
	// Example: projects/{project}/zones/{zones}/machineTypes/{machineType}.
	// If a user specifies a plain value (with no '/' characters in it),
	// the provider automatically applies the pattern.
	Pattern string `json:"pattern,omitempty"`
}

// CloudAPIType represents the shape of an auxiliary type in the API.
type CloudAPIType struct {
	Properties map[string]CloudAPIProperty `json:"properties,omitempty"`
}
