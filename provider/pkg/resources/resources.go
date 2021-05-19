// Copyright 2016-2021, Pulumi Corporation.

package resources

import (
	"fmt"
	"strings"
)

// CloudAPIMetadata is a collection of all resources and functions in the Google Cloud REST API.
type CloudAPIMetadata struct {
	Resources map[string]CloudAPIResource `json:"resources"`
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
	IdPath           string                      `json:"idPath,omitempty"`
	IdParams         map[string]string           `json:"idParams,omitempty"`
	NoDelete         bool                        `json:"noDelete,omitempty"`
}

// CloudAPIResourceParam is a URL parameter of an API resource.
type CloudAPIResourceParam struct {
	Name     string `json:"name"`
	SdkName  string `json:"sdkName,omitempty"`
	Location string `json:"location"`
}

// RelativePath joins the resource base URL with the given path.
func (r *CloudAPIResource) RelativePath(rel string) string {
	return fmt.Sprintf("%s/%s", strings.TrimRight(r.BaseUrl, "/"), strings.TrimLeft(rel, "/"))
}

// CloudAPIProperty is a property of a body of an API call payload.
type CloudAPIProperty struct {
	// The name of the container property that was "flattened" during SDK generation, i.e. extra layer that exists
	// in the API payload but does not exist in the SDK.
	Container string `json:"container,omitempty"`
	SdkName   string `json:"sdkName,omitempty"`
}
