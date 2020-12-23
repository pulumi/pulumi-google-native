// Copyright 2016-2020, Pulumi Corporation.

package resources

// CloudAPIResource is a resource in Google Cloud REST API.
type CloudAPIResource struct {
	BaseUrl      string   `json:"baseUrl"`
	CreatePath   string   `json:"createPath"`
	CreateParams []string `json:"createParams"`
	DeletePath   string   `json:"deletePath"`
	DeleteParams []string `json:"deleteParams"`
}

// TODO: Generate the metadata with the schema.
var CloudAPIResources = map[string]CloudAPIResource{
	"google-cloud:run/v1:Service": {
		BaseUrl:      "https://run.googleapis.com",
		CreatePath:   "v1/{parent}/services",
		CreateParams: []string{"parent"},
		DeletePath:   "v1/{parent}/services/{metadata.name}",
		DeleteParams: []string{"parent", "metadata.name"},
	},
	"google-cloud:run/v1:Policy": {
		BaseUrl:      "https://run.googleapis.com",
		CreatePath:   "v1/{resource}:setIamPolicy",
		CreateParams: []string{"resource"},
	},
	"google-cloud:cloudfunctions/v1:CloudFunction": {
		BaseUrl:      "https://cloudfunctions.googleapis.com",
		CreatePath:   "v1/{location}/functions",
		CreateParams: []string{"location"},
		DeletePath:   "v1/{name}",
		DeleteParams: []string{"name"},
	},
	"google-cloud:cloudfunctions/v1:Policy": {
		BaseUrl:      "https://cloudfunctions.googleapis.com",
		CreatePath:   "v1/{resource}:setIamPolicy",
		CreateParams: []string{"resource"},
	},
	"google-cloud:container/v1:Cluster": {
		BaseUrl:      "https://container.googleapis.com",
		CreatePath:   "v1/{parent}/clusters",
		CreateParams: []string{"parent"},
		DeletePath:   "v1/{parent}/clusters/{cluster.name}",
		DeleteParams: []string{"parent", "cluster.name"},
	},
	"google-cloud:container/v1:NodePool": {
		BaseUrl:      "https://container.googleapis.com",
		CreatePath:   "v1/{parent}/nodePools",
		CreateParams: []string{"parent"},
		DeletePath:   "v1/{parent}/nodePools/{nodePool.name}",
		DeleteParams: []string{"parent", "nodePool.name"},
	},
	"google-cloud:storage/v1:Bucket": {
		BaseUrl:      "https://storage.googleapis.com/storage/v1",
		CreatePath:   "b?project={project}",
		CreateParams: []string{"project"},
		DeletePath:   "b/{name}",
		DeleteParams: []string{"name"},
	},
	"google-cloud:storage/v1:BucketObject": {
		BaseUrl:      "https://storage.googleapis.com/storage/v1",
		CreatePath:   "b/{bucket}/o",
		CreateParams: []string{"bucket"},
		DeletePath:   "b/{bucket}/o/{name}",
		DeleteParams: []string{"bucket", "name"},
	},
}
