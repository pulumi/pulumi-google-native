// Copyright 2016-2021, Pulumi Corporation.

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

// CloudAPIResource is a resource in Google Cloud REST API.
type CloudAPIResource struct {
	BaseUrl          string                      `json:"baseUrl"`
	CreatePath       string                      `json:"createPath"`
	CreateParams     []CloudAPIResourceParam     `json:"createParams"`
	CreateVerb       string                      `json:"createVerb,omitempty"`
	CreateProperties map[string]CloudAPIProperty `json:"createProperties,omitempty"`
	UpdateVerb       string                      `json:"updateVerb,omitempty"`
	UpdateProperties map[string]CloudAPIProperty `json:"updateProperties,omitempty"`
	NoDelete         bool                        `json:"noDelete,omitempty"`
	// IdProperty contains the name of the output property that represents resource ID (a self link).
	IdProperty string `json:"idProperty,omitempty"`
	// IdPath is the template for building resource ID with ID parameter values. It should only
	// be defined if IdProperty is missing.
	IdPath   string            `json:"idPath,omitempty"`
	IdParams map[string]string `json:"idParams,omitempty"`
	// AutoNamePattern contains a string pattern when a default name should be automatically generated if a user hasn't
	// explicitly specified one. Example: projects/{project}/locations/{location}/functions/{name}.
	AutoNamePattern string `json:"autoNamePattern,omitempty"`
}

// CloudAPIFunction is a function in Google Cloud REST API.
type CloudAPIFunction struct {
	Url    string                  `json:"url"`
	Params []CloudAPIResourceParam `json:"params"`
}

// CloudAPIResourceParam is a URL parameter of an API resource.
type CloudAPIResourceParam struct {
	Name     string `json:"name"`
	SdkName  string `json:"sdkName,omitempty"`
	Location string `json:"location"`
}

// ResourceUrl returns the resource API URL by joining the base URL with the resource path.
func (r *CloudAPIResource) ResourceUrl(path string) string {
	return CombineUrl(r.BaseUrl, path)
}

// CombineUrl concats a base URL of the service with a path of an endpoint.
func CombineUrl(baseUrl, path string) string {
	if strings.HasPrefix(path, "https://") {
		return path
	}
	return fmt.Sprintf("%s/%s", strings.TrimRight(baseUrl, "/"), strings.TrimLeft(path, "/"))
}

// CloudAPIProperty is a property of a body of an API call payload.
type CloudAPIProperty struct {
	Ref                  string            `json:"$ref,omitempty"`
	Items                *CloudAPIProperty `json:"items,omitempty"`
	AdditionalProperties *CloudAPIProperty `json:"additionalProperties,omitempty"`
	// The name of the container property that was "flattened" during SDK generation, i.e. extra layer that exists
	// in the API payload but does not exist in the SDK.
	Container string `json:"container,omitempty"`
	SdkName   string `json:"sdkName,omitempty"`
	// CopyFromOutputs equal to true means that the value for this property during an update should be taken from
	// the previous state of the resource, not user inputs.
	CopyFromOutputs bool `json:"copyFromOutputs,omitempty"`
}

// CloudAPIType represents the shape of an auxiliary type in the API.
type CloudAPIType struct {
	Properties map[string]CloudAPIProperty `json:"properties,omitempty"`
}
