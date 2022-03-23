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

	// Security Center.
	"v1/folders/{foldersId}/muteConfigs":                 "",
	"v1/folders/{foldersId}/bigQueryExports":             "FolderBigQueryExport",
	"v1/organizations/{organizationsId}/bigQueryExports": "OrganizationBigQueryExport",
	"v1/organizations/{organizationsId}/muteConfigs":     "OrganizationMuteConfig",
	"v1/projects/{projectsId}/bigQueryExports":           "ProjectBigQueryExport",

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

// autonameOverrides is a map of exceptions to the property used for auto-naming.
// The key is the resource token, and the value is the property to use for auto-naming.
var autonameOverrides = map[string]string{
	"google-native:cloudkms/v1:CryptoKey":        "cryptoKeyId",
	"google-native:cloudkms/v1:CryptoKeyVersion": "cryptoKeyId",
	"google-native:cloudkms/v1:EkmConnection":    "ekmConnectionId",
	"google-native:cloudkms/v1:ImportJob":        "importJobId",
	"google-native:cloudkms/v1:KeyRing":          "keyRingId",
}

// metadataOverrides is a map of values static overlays to merge into the metadata for
// individual resource tokens. In case of conflict, the values in this mapping are preferred.
var metadataOverrides = map[string]resources.CloudAPIResource{
	"google-native:run/v1:Service": {
		Create: resources.CreateAPIOperation{
			CloudAPIOperation: resources.CloudAPIOperation{
				Polling: resources.Polling{Strategy: resources.KNativeStatusPoll},
			},
		},
		Update: resources.UpdateAPIOperation{
			CloudAPIOperation: resources.CloudAPIOperation{
				Polling: resources.Polling{Strategy: resources.KNativeStatusPoll},
			},
		},
	},
	"google-native:run/v1:DomainMapping": {
		Create: resources.CreateAPIOperation{
			CloudAPIOperation: resources.CloudAPIOperation{
				Polling: resources.Polling{Strategy: resources.KNativeStatusPoll},
			},
		},
	},
}

// csharpNamespaceOverrides is a map of canonical C# namespaces per lowercase module name. It only lists the ones
// that aren't successfully inferred from the discovery document.
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
