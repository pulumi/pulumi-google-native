// Copyright 2016-2020, Pulumi Corporation.

package resources

// CloudAPIMetadata is a collection of all resources and functions in the Google Cloud REST API.
type CloudAPIMetadata struct {
	Resources map[string]CloudAPIResource `json:"resources"`
}

// CloudAPIResource is a resource in Google Cloud REST API.
type CloudAPIResource struct {
	BaseUrl          string                      `json:"baseUrl"`
	CreatePath       string                      `json:"createPath"`
	CreateParams     []string                    `json:"createParams"`
	CreateProperties map[string]CloudAPIProperty `json:"createProperties,omitempty"`
	IdPath           string                      `json:"idPath,omitempty"`
	IdParams         []string                    `json:"idParams,omitempty"`
	NoGet            bool                        `json:"noGet,omitempty"`
	NoDelete         bool                        `json:"noDelete,omitempty"`
}
// CloudAPIProperty is a property of a body of an API call payload.
type CloudAPIProperty struct {
	// The name of the container property that was "flattened" during SDK generation, i.e. extra layer that exists
	// in the API payload but does not exist in the SDK.
	Container string `json:"container,omitempty"`
}
