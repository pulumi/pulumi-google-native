// Copyright 2016-2021, Pulumi Corporation.

package gen

import "github.com/pulumi/pulumi-google-native/provider/pkg/resources"

// resourceNameByTypeOverrides is a map of Pulumi resource names by the type name
// that is present in the discovery document.
var resourceNameByTypeOverrides = map[string]string{
	// Apigee.
	"GoogleCloudApigeeV1ApiProduct":       "ApiProduct",
	"GoogleCloudApigeeV1CanaryEvaluation": "CanaryEvaluation",
	"GoogleCloudApigeeV1DataCollector":    "DataCollector",
	"GoogleCloudApigeeV1DebugSession":     "DebugSession",
	"GoogleCloudApigeeV1RatePlan":         "RatePlan",
	"GoogleCloudApigeeV1TargetServer":     "TargetServer",

	// Cloud Search.
	"DataSource":        "DataSource",
	"SearchApplication": "SearchApplication",

	// Diag Flow.
	"GoogleCloudDialogflowV2SessionEntityType":        "SessionEntityType",
	"GoogleCloudDialogflowV2beta1SessionEntityType":   "SessionEntityType",
	"GoogleCloudDialogflowCxV3SessionEntityType":      "SessionEntityType",
	"GoogleCloudDialogflowCxV3beta1SessionEntityType": "SessionEntityType",

	// Remote Build Execution.
	"GoogleDevtoolsRemotebuildexecutionAdminV1alphaCreateWorkerPoolRequest": "WorkerPool",

	// Run.
	"DomainMapping": "DomainMapping",
}

// resourceNameByPathOverrides is a map of Pulumi resource names by the API flat path
// that is present in the discovery document.
var resourceNameByPathOverrides = map[string]string{
	// Cloud Run: remove domain mapping on the KNative path.
	"apis/domains.cloudrun.com/v1/namespaces/{namespacesId}/domainmappings": "",

	// Apigee.
	"v1/organizations/{organizationsId}/envgroups/{envgroupsId}/attachments": "EnvgroupAttachment",
	"v1/organizations/{organizationsId}/instances/{instancesId}/attachments": "InstanceAttachment",

	// DLP.
	"v2/organizations/{organizationsId}/deidentifyTemplates":                         "",
	"v2/organizations/{organizationsId}/locations/{locationsId}/deidentifyTemplates": "OrganizationsDeidentifyTemplate",
	"v2/organizations/{organizationsId}/inspectTemplates":                            "",
	"v2/organizations/{organizationsId}/locations/{locationsId}/inspectTemplates":    "OrganizationInspectTemplate",
	"v2/organizations/{organizationsId}/locations/{locationsId}/jobTriggers":         "OrganizationJobTrigger",

	// Essential Contracts.
	"v1/folders/{foldersId}/contacts":             "FolderContact",
	"v1/organizations/{organizationsId}/contacts": "OrganizationContact",

	// IAM.
	"v1/organizations/{organizationsId}/roles": "OrganizationRole",

	// Logging.
	"v2/billingAccounts/{billingAccountsId}/exclusions":                                        "BillingAccountExclusion",
	"v2/folders/{foldersId}/exclusions":                                                        "FolderExclusion",
	"v2/organizations/{organizationsId}/exclusions":                                            "OrganizationExclusion",
	"v2/{v2Id}/{v2Id1}/exclusions":                                                             "",
	"v2/billingAccounts/{billingAccountsId}/sinks":                                             "BillingAccountSink",
	"v2/folders/{foldersId}/sinks":                                                             "FolderSink",
	"v2/organizations/{organizationsId}/sinks":                                                 "OrganizationSink",
	"v2/billingAccounts/{billingAccountsId}/locations/{locationsId}/buckets":                   "BillingAccountBucket",
	"v2/folders/{foldersId}/locations/{locationsId}/buckets":                                   "FolderBucket",
	"v2/organizations/{organizationsId}/locations/{locationsId}/buckets":                       "OrganizationBucket",
	"v2/billingAccounts/{billingAccountsId}/locations/{locationsId}/buckets/{bucketsId}/views": "BillingAccountBucketView",
	"v2/folders/{foldersId}/locations/{locationsId}/buckets/{bucketsId}/views":                 "FolderBucketView",
	"v2/organizations/{organizationsId}/locations/{locationsId}/buckets/{bucketsId}/views":     "OrganizationBucketView",
	"v2/projects/{projectsId}/locations/{locationsId}/buckets/{bucketsId}/views":               "BucketView",

	// Org Policy.
	"v2/folders/{foldersId}/policies":             "FolderPolicy",
	"v2/organizations/{organizationsId}/policies": "OrganizationPolicy",

	// Policy Simulator.
	"v1/folders/{foldersId}/locations/{locationsId}/replays":                  "FolderReplay",
	"v1/organizations/{organizationsId}/locations/{locationsId}/replays":      "OrganizationReplay",
	"v1beta1/folders/{foldersId}/locations/{locationsId}/replays":             "FolderReplay",
	"v1beta1/organizations/{organizationsId}/locations/{locationsId}/replays": "OrganizationReplay",

	// Storage.
	"b/{bucket}/o": "BucketObject",
}

// resourceNamePropertyOverrides is a list of exceptions populated for the buildIdParams method above.
var resourceNamePropertyOverrides = map[string]string{
	"apigee/v1:Alias.aliasesId":                               "alias",
	"appengine/v1:IngressRule.ingressRulesId":                 "priority",
	"appengine/v1beta:IngressRule.ingressRulesId":             "priority",
	"bigquery/v2:Routine.routinesId":                          "routineReference.routineId",
	"dataproc/v1:Job.jobId":                                   "reference.jobId",
	"dataproc/v1beta2:Job.jobId":                              "reference.jobId",
	"dns/v1:Change.changeId":                                  "id",
	"dns/v1beta2:Change.changeId":                             "id",
	"dns/v1beta2:ResponsePolicy.responsePolicy":               "id",
	"dns/v1beta2:ResponsePolicyRule.responsePolicyRule":       "ruleName",
	"firebasehosting/v1beta1:Domain.domainsId":                "domainName",
	"recommendationengine/v1beta1:CatalogItem.catalogItemsId": "id",
	"run/v1alpha1:Job.jobsId":                                 "metadata.name",
	"run/v1:DomainMapping.domainmappingsId":                   "metadata.name",
	"run/v1:Service.servicesId":                               "metadata.name",
}

// csharpNamespaceOverrides is a map of canonical C# namespaces per lowercase module name. It only lists the ones
// that aren't successfully infered from the discovery document.
var csharpNamespaceOverrides = map[string]string{
	"billingbudgets":       "BillingBudgets",
	"cloudkms":             "CloudKMS",
	"datamigration":        "DataMigration",
	"firebasedatabase":     "FirebaseDatabase",
	"iam":                  "Iam",
	"managedidentities":    "ManagedIdentities",
	"ml":                   "ML",
	"orgpolicy":            "OrgPolicy",
	"privateca":            "PrivateCA",
	"pubsub":               "PubSub",
	"pubsublite":           "PubSubLite",
	"recommendationengine": "RecommendationEngine",
	"securitycenter":       "SecurityCenter",
}

var clusterAPIResource = resources.CloudAPIResource{
	BaseUrl:    "https://container.googleapis.com/",
	CreatePath: "v1/projects/{projectsId}/locations/{locationsId}/clusters",
	CreateParams: []resources.CloudAPIResourceParam{
		{
			Name:     "projectsId",
			SdkName:  "project",
			Location: "path",
		},
		{
			Name:     "locationsId",
			SdkName:  "location",
			Location: "path",
		},
	},
	CreateVerb: "POST",
	CreateProperties: map[string]resources.CloudAPIProperty{
		"addonsConfig": {
			Container: "cluster",
		},
		"authenticatorGroupsConfig": {
			Container: "cluster",
		},
		"autopilot": {
			Container: "cluster",
		},
		"autoscaling": {
			Container: "cluster",
		},
		"binaryAuthorization": {
			Container: "cluster",
		},
		"clusterIpv4Cidr": {
			Container: "cluster",
		},
		"conditions": {
			Container: "cluster",
		},
		"confidentialNodes": {
			Container: "cluster",
		},
		"databaseEncryption": {
			Container: "cluster",
		},
		"defaultMaxPodsConstraint": {
			Container: "cluster",
		},
		"description": {
			Container: "cluster",
		},
		"enableKubernetesAlpha": {
			Container: "cluster",
		},
		"enableTpu": {
			Container: "cluster",
		},
		"initialClusterVersion": {
			Container: "cluster",
		},
		"ipAllocationPolicy": {
			Container: "cluster",
		},
		"labelFingerprint": {
			Container: "cluster",
		},
		"legacyAbac": {
			Container: "cluster",
		},
		"locations": {
			Container: "cluster",
		},
		"loggingService": {
			Container: "cluster",
		},
		"maintenancePolicy": {
			Container: "cluster",
		},
		"masterAuth": {
			Container: "cluster",
		},
		"masterAuthorizedNetworksConfig": {
			Container: "cluster",
		},
		"monitoringService": {
			Container: "cluster",
		},
		"name": {
			Container: "cluster",
		},
		"network": {
			Container: "cluster",
		},
		"networkConfig": {
			Container: "cluster",
		},
		"networkPolicy": {
			Container: "cluster",
		},
		"nodePools": {
			Container: "cluster",
		},
		"notificationConfig": {
			Container: "cluster",
		},
		"parent": {},
		"privateClusterConfig": {
			Container: "cluster",
		},
		"releaseChannel": {
			Container: "cluster",
		},
		"resourceLabels": {
			Container: "cluster",
		},
		"resourceUsageExportConfig": {
			Container: "cluster",
		},
		"shieldedNodes": {
			Container: "cluster",
		},
		"subnetwork": {
			Container: "cluster",
		},
		"verticalPodAutoscaling": {
			Container: "cluster",
		},
		"workloadIdentityConfig": {
			Container: "cluster",
		},
	},
	UpdateVerb: "PUT",
	UpdateProperties: map[string]resources.CloudAPIProperty{
		"addonsConfig": {}, //
		//"authenticatorGroupsConfig":      {}, // recreate
		//"autopilot":                      {}, // recreate
		"binaryAuthorization":            {Container: "cluster"}, //
		"autoscaling":                    {Container: "cluster"}, //
		"databaseEncryption":             {Container: "cluster"}, //
		"locations":                      {Container: "cluster"}, //
		"loggingService":                 {Container: "cluster"}, //
		"masterAuthorizedNetworksConfig": {Container: "cluster"}, //
		"monitoringService":              {Container: "cluster"}, // monitoringservice and logging service can be updated simultaneously //
		//"notificationConfig":             {}, // recreate
		//"privateClusterConfig":           {}, // recreate
		//"releaseChannel":                 {}, // recreate
		//"resourceUsageExportConfig":      {}, // recreate
		"shieldedNodes":          {Container: "cluster"}, //
		"verticalPodAutoscaling": {Container: "cluster"}, //
		"workloadIdentityConfig": {Container: "cluster"}, //
		// network config things that can change:
		// enableIntraNodeVisibility
		// enableL4ilbSubsetting
		// privateIpv6GoogleAccess
		// defaultSnatStatus
		// datapathProvider
		//"networkConfig": {}, // recreate
		"masterVersion": {Container: "cluster"}, // need to reuse initialMasterVersion for this //
		// "nodeVersion": {}, // Not sure what to do about the default nodepool.

		// non-update
		//"resourceLabels":    {}, // recreate
		"maintenancePolicy": {Container: "cluster"}, //
		//"masterAuth":        {}, // recreate
		//"networkPolicy":     {}, // recreate
		"legacyAbac": {Container: "cluster"}, //
	},
	IdProperty:      "selfLink",
	AutoNamePattern: "{name}",
}

var clusterBetaAPIResource = resources.CloudAPIResource{
	BaseUrl:    "https://container.googleapis.com/",
	CreatePath: "v1/projects/{projectsId}/locations/{locationsId}/clusters",
	CreateParams: []resources.CloudAPIResourceParam{
		{
			Name:     "projectsId",
			SdkName:  "project",
			Location: "path",
		},
		{
			Name:     "locationsId",
			SdkName:  "location",
			Location: "path",
		},
	},
	CreateVerb: "POST",
	CreateProperties: map[string]resources.CloudAPIProperty{
		"addonsConfig": {
			Container: "cluster",
		},
		"authenticatorGroupsConfig": {
			Container: "cluster",
		},
		"autopilot": {
			Container: "cluster",
		},
		"autoscaling": {
			Container: "cluster",
		},
		"binaryAuthorization": {
			Container: "cluster",
		},
		"clusterIpv4Cidr": {
			Container: "cluster",
		},
		"clusterTelemetry": {
			Container: "cluster",
		},
		"conditions": {
			Container: "cluster",
		},
		"confidentialNodes": {
			Container: "cluster",
		},
		"databaseEncryption": {
			Container: "cluster",
		},
		"defaultMaxPodsConstraint": {
			Container: "cluster",
		},
		"description": {
			Container: "cluster",
		},
		"enableKubernetesAlpha": {
			Container: "cluster",
		},
		"initialClusterVersion": {
			Container: "cluster",
			SdkName:   "clusterVersion",
		},
		"ipAllocationPolicy": {
			Container: "cluster",
		},
		"labelFingerprint": {
			Container: "cluster",
		},
		"legacyAbac": {
			Container: "cluster",
		},
		"locations": {
			Container: "cluster",
		},
		"loggingService": {
			Container: "cluster",
		},
		"maintenancePolicy": {
			Container: "cluster",
		},
		"master": {
			Container: "cluster",
		},
		"masterAuth": {
			Container: "cluster",
		},
		"masterAuthorizedNetworksConfig": {
			Container: "cluster",
		},
		"monitoringService": {
			Container: "cluster",
		},
		"name": {
			Container: "cluster",
		},
		"network": {
			Container: "cluster",
		},
		"networkConfig": {
			Container: "cluster",
		},
		"networkPolicy": {
			Container: "cluster",
		},
		"nodePoolDefaults": {
			Container: "cluster",
		},
		"nodePools": {
			Container: "cluster",
		},
		"notificationConfig": {
			Container: "cluster",
		},
		"parent": {},
		"podSecurityPolicyConfig": {
			Container: "cluster",
		},
		"privateClusterConfig": {
			Container: "cluster",
		},
		"releaseChannel": {
			Container: "cluster",
		},
		"resourceLabels": {
			Container: "cluster",
		},
		"resourceUsageExportConfig": {
			Container: "cluster",
		},
		"shieldedNodes": {
			Container: "cluster",
		},
		"subnetwork": {
			Container: "cluster",
		},
		"tpuConfig": {
			Container: "cluster",
		},
		"verticalPodAutoscaling": {
			Container: "cluster",
		},
		"workloadCertificates": {
			Container: "cluster",
		},
		"workloadIdentityConfig": {
			Container: "cluster",
		},
	},
	UpdateVerb: "PUT",
	UpdateProperties: map[string]resources.CloudAPIProperty{
		"addonsConfig": {
			Container: "cluster",
		}, //
		//"authenticatorGroupsConfig":      {}, // recreate
		//"autopilot":                      {}, // recreate
		"binaryAuthorization": {Container: "cluster"}, //
		//"autoscaling":                    {}, //
		//"clusterTelemetry":               {},
		"databaseEncryption":             {Container: "cluster"}, //
		"description":                    {Container: "cluster"}, //
		"endpoint":                       {Container: "cluster"}, //
		"locations":                      {Container: "cluster"}, //
		"loggingService":                 {Container: "cluster"}, //
		"master":                         {Container: "cluster"}, //
		"masterAuthorizedNetworksConfig": {Container: "cluster"}, //
		"monitoringService":              {Container: "cluster"}, // monitoringservice and logging service can be updated simultaneously
		// "notificationConfig":             {}, // recreate
		"podSecurityPolicyConfig": {Container: "cluster"},
		// "privateClusterConfig":           {}, // recreate
		//"releaseChannel":                 {},
		// "resourceUsageExportConfig": {}, // recreate
		"shieldedNodes": {Container: "cluster"}, //
		// "tpuConfig":                 {}, // recreate
		"verticalPodAutoscaling": {Container: "cluster"}, //
		"workloadCertificates":   {Container: "cluster"}, //
		"workloadIdentityConfig": {Container: "cluster"}, //
		// network config things that can change:
		// enableIntraNodeVisibility
		// enableL4ilbSubsetting
		// privateIpv6GoogleAccess
		// defaultSnatStatus
		// datapathProvider
		// "networkConfig": {}, // recreate
		"masterVersion": {Container: "cluster"}, // need to reuse initialMasterVersion for this //
		// "nodeVersion": {}, // Not sure what to do about the default nodepool.

		// non-update
		"resourceLabels":    {Container: "cluster"}, //
		"maintenancePolicy": {Container: "cluster"}, //
		//"masterAuth":        {}, // recreate
		// "networkPolicy": {}, // recreate
		"legacyAbac": {Container: "cluster"}, //
	},
	IdProperty:      "selfLink",
	AutoNamePattern: "{name}",
}

var metadataOverlays = map[string]resources.CloudAPIResource{
	"google-native:container/v1:Cluster":      clusterAPIResource,
	"google-native:container/v1beta1:Cluster": clusterBetaAPIResource,
}
