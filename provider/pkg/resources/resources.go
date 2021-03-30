// Copyright 2016-2020, Pulumi Corporation.

package resources

// CloudAPIMetadata is a collection of all resources and functions in the Google Cloud REST API.
type CloudAPIMetadata struct {
	Resources map[string]CloudAPIResource `json:"resources"`
}

// CloudAPIResource is a resource in Google Cloud REST API.
type CloudAPIResource struct {
	BaseUrl      string   `json:"baseUrl"`
	CreatePath   string   `json:"createPath"`
	CreateParams []string `json:"createParams"`
	DeletePath   string   `json:"deletePath,omitempty"`
	DeleteParams []string `json:"deleteParams,omitempty"`
}
