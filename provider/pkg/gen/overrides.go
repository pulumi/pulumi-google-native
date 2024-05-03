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

import (
	"github.com/pulumi/pulumi-google-native/provider/pkg/resources"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
)

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
	"run_v1.json:apis/domains.cloudrun.com/v1/namespaces/{namespacesId}/domainmappings": "",
	"run_v1.json:apis/serving.knative.dev/v1/namespaces/{namespacesId}/services":        "",

	// AI Platform.
	"aiplatform_v1beta1.json:v1beta1/projects/{projectsId}/locations/{locationsId}/featureGroups/{featureGroupsId}/features":                             "FeatureGroupFeature",
	"aiplatform_v1beta1.json:v1beta1/projects/{projectsId}/locations/{locationsId}/featurestores/{featurestoresId}/entityTypes/{entityTypesId}/features": "FeatureStoreFeature",
	"aiplatform_v1.json:v1/projects/{projectsId}/locations/{locationsId}/featureGroups/{featureGroupsId}/features":                                       "FeatureGroupFeature",
	"aiplatform_v1.json:v1/projects/{projectsId}/locations/{locationsId}/featurestores/{featurestoresId}/entityTypes/{entityTypesId}/features":           "FeatureStoreFeature",

	// Apigee.
	"apigee_v1.json:v1/organizations/{organizationsId}/appgroups/{appgroupsId}/apps":                                        "AppGroupApp",
	"apigee_v1.json:v1/organizations/{organizationsId}/developers/{developersId}/apps":                                      "DeveloperApp",
	"apigee_v1.json:v1/organizations/{organizationsId}/appgroups/{appgroupsId}/apps/{appsId}/keys":                          "AppGroupAppKey",
	"apigee_v1.json:v1/organizations/{organizationsId}/developers/{developersId}/apps/{appsId}/keys":                        "DeveloperAppKey",
	"apigee_v1.json:v1/organizations/{organizationsId}/envgroups/{envgroupsId}/attachments":                                 "EnvgroupAttachment",
	"apigee_v1.json:v1/organizations/{organizationsId}/instances/{instancesId}/attachments":                                 "InstanceAttachment",
	"apigee_v1.json:v1/organizations/{organizationsId}/environments/{environmentsId}/keyvaluemaps/{keyvaluemapsId}/entries": "EnvironmentEntry",

	// App Engine Alpha v1. Get rid of locations and operations nested under locations.
	"appengine_v1alpha.json:v1alpha/projects/{projectsId}/locations/{locationsId}":                           "",
	"appengine_v1alpha.json:v1alpha/projects/{projectsId}/locations/{locationsId}/operations/{operationsId}": "",
	"appengine_v1beta.json:v1beta/projects/{projectsId}/locations/{locationsId}":                             "",
	"appengine_v1beta.json:v1beta/projects/{projectsId}/locations/{locationsId}/operations/{operationsId}":   "",

	// ApigeeRegistry
	"apigeeregistry_v1.json:v1/projects/{projectsId}/locations/{locationsId}/apis/{apisId}/deployments/{deploymentsId}/artifacts": "DeploymentArtifact",
	"apigeeregistry_v1.json:v1/projects/{projectsId}/locations/{locationsId}/apis/{apisId}/versions/{versionsId}/artifacts":       "VersionArtifact",

	// DLP.
	"dlp_v2.json:v2/organizations/{organizationsId}/deidentifyTemplates":                         "",
	"dlp_v2.json:v2/organizations/{organizationsId}/locations/{locationsId}/deidentifyTemplates": "OrganizationsDeidentifyTemplate",
	"dlp_v2.json:v2/organizations/{organizationsId}/inspectTemplates":                            "",
	"dlp_v2.json:v2/organizations/{organizationsId}/locations/{locationsId}/discoveryConfigs":    "OrganizationDiscoveryConfig",
	"dlp_v2.json:v2/organizations/{organizationsId}/locations/{locationsId}/inspectTemplates":    "OrganizationInspectTemplate",
	"dlp_v2.json:v2/organizations/{organizationsId}/locations/{locationsId}/jobTriggers":         "OrganizationJobTrigger",
	"dlp_v2.json:v2/organizations/{organizationsId}/locations/{locationsId}/storedInfoTypes":     "",
	"dlp_v2.json:v2/organizations/{organizationsId}/storedInfoTypes":                             "",

	// Essential Contacts.
	"essentialcontacts_v1.json:v1/folders/{foldersId}/contacts":             "FolderContact",
	"essentialcontacts_v1.json:v1/organizations/{organizationsId}/contacts": "OrganizationContact",

	// GKE Hub.
	"gkehub_v1alpha.json:v1alpha/projects/{projectsId}/locations/{locationsId}/memberships/{membershipsId}/rbacrolebindings": "MembershipRbacRoleBinding",
	"gkehub_v1alpha.json:v1alpha/projects/{projectsId}/locations/{locationsId}/scopes/{scopesId}/rbacrolebindings":           "ScopeRbacRoleBinding",
	"gkehub_v1beta.json:v1beta/projects/{projectsId}/locations/{locationsId}/memberships/{membershipsId}/rbacrolebindings":   "MembershipRbacRoleBinding",
	"gkehub_v1beta.json:v1beta/projects/{projectsId}/locations/{locationsId}/scopes/{scopesId}/rbacrolebindings":             "ScopeRbacRoleBinding",

	// IAM.
	"iam_v1.json:v1/organizations/{organizationsId}/roles": "OrganizationRole",
	// Original: "iam_v1.json:v1/projects/{projectsId}/locations/{locationsId}/workloadIdentityPools/{workloadIdentityPoolsId}/providers"
	"iam_v1.json:v1/locations/{locationsId}/workforcePools/{workforcePoolsId}/providers":                                                        "WorkforcePoolProvider",
	"iam_v1.json:v1/locations/{locationsId}/workforcePools/{workforcePoolsId}/providers/{providersId}/keys":                                     "WorkforcePoolKey",
	"iam_v1.json:v1/projects/{projectsId}/locations/{locationsId}/workloadIdentityPools/{workloadIdentityPoolsId}/providers/{providersId}/keys": "WorkloadIdentityPoolKey",

	// Integrations.
	"integrations_v1alpha.json:v1alpha/projects/{projectsId}/locations/{locationsId}/products/{productsId}/integrations/{integrationsId}/versions":                 "Version",
	"integrations_v1alpha.json:v1alpha/projects/{projectsId}/locations/{locationsId}/products/{productsId}/integrationtemplates/{integrationtemplatesId}/versions": "TemplatesVersion",

	// Logging.
	"logging_v2.json:v2/billingAccounts/{billingAccountsId}/exclusions":                                        "BillingAccountExclusion",
	"logging_v2.json:v2/folders/{foldersId}/exclusions":                                                        "FolderExclusion",
	"logging_v2.json:v2/organizations/{organizationsId}/exclusions":                                            "OrganizationExclusion",
	"logging_v2.json:v2/{v2Id}/{v2Id1}/exclusions":                                                             "",
	"logging_v2.json:v2/billingAccounts/{billingAccountsId}/sinks":                                             "BillingAccountSink",
	"logging_v2.json:v2/folders/{foldersId}/sinks":                                                             "FolderSink",
	"logging_v2.json:v2/organizations/{organizationsId}/sinks":                                                 "OrganizationSink",
	"logging_v2.json:v2/billingAccounts/{billingAccountsId}/locations/{locationsId}/buckets":                   "BillingAccountBucket",
	"logging_v2.json:v2/folders/{foldersId}/locations/{locationsId}/buckets":                                   "FolderBucket",
	"logging_v2.json:v2/organizations/{organizationsId}/locations/{locationsId}/buckets":                       "OrganizationBucket",
	"logging_v2.json:v2/billingAccounts/{billingAccountsId}/locations/{locationsId}/buckets/{bucketsId}/views": "BillingAccountBucketView",
	"logging_v2.json:v2/folders/{foldersId}/locations/{locationsId}/buckets/{bucketsId}/views":                 "FolderBucketView",
	"logging_v2.json:v2/organizations/{organizationsId}/locations/{locationsId}/buckets/{bucketsId}/views":     "OrganizationBucketView",
	"logging_v2.json:v2/projects/{projectsId}/locations/{locationsId}/buckets/{bucketsId}/views":               "BucketView",
	"logging_v2.json:v2/{v2Id}/{v2Id1}/locations/{locationsId}/buckets":                                        "",
	"logging_v2.json:v2/{v2Id}/{v2Id1}/sinks":                                                                  "",
	"logging_v2.json:v2/billingAccounts/{billingAccountsId}/locations/{locationsId}/buckets/{bucketsId}/links": "BillingAccountBucketLink",
	"logging_v2.json:v2/folders/{foldersId}/locations/{locationsId}/buckets/{bucketsId}/links":                 "FolderBucketLink",
	"logging_v2.json:v2/organizations/{organizationsId}/locations/{locationsId}/buckets/{bucketsId}/links":     "OrganizationBucketLink",

	// Network Security.
	"networksecurity_v1beta1.json:v1beta1/organizations/{organizationsId}/locations/{locationsId}/addressGroups": "OrganizationAddressGroup",
	"networksecurity_v1beta1.json:v1beta1/projects/{projectsId}/locations/{locationsId}/addressGroups":           "AddressGroup",
	"networksecurity_v1.json:v1/organizations/{organizationsId}/locations/{locationsId}/addressGroups":           "OrganizationAddressGroup",
	"networksecurity_v1.json:v1/projects/{projectsId}/locations/{locationsId}/addressGroups":                     "AddressGroup",

	// Org Policy.
	"orgpolicy_v2.json:v2/folders/{foldersId}/policies":             "FolderPolicy",
	"orgpolicy_v2.json:v2/organizations/{organizationsId}/policies": "OrganizationPolicy",

	// Policy Simulator.
	"policysimulator_v1.json:v1/folders/{foldersId}/locations/{locationsId}/replays":                       "FolderReplay",
	"policysimulator_v1.json:v1/organizations/{organizationsId}/locations/{locationsId}/replays":           "OrganizationReplay",
	"policysimulator_v1alpha.json:v1alpha/folders/{foldersId}/locations/{locationsId}/replays":             "FolderReplay",
	"policysimulator_v1alpha.json:v1alpha/organizations/{organizationsId}/locations/{locationsId}/replays": "OrganizationReplay",
	"policysimulator_v1beta.json:v1beta/folders/{foldersId}/locations/{locationsId}/replays":               "FolderReplay",
	"policysimulator_v1beta.json:v1beta/organizations/{organizationsId}/locations/{locationsId}/replays":   "OrganizationReplay",
	"policysimulator_v1beta1.json:v1beta1/folders/{foldersId}/locations/{locationsId}/replays":             "FolderReplay",
	"policysimulator_v1beta1.json:v1beta1/organizations/{organizationsId}/locations/{locationsId}/replays": "OrganizationReplay",

	// Security Center.
	"securitycenter_v1.json:v1/folders/{foldersId}/muteConfigs":                                               "",
	"securitycenter_v1.json:v1/folders/{foldersId}/notificationConfigs":                                       "FolderNotificationConfig",
	"securitycenter_v1.json:v1/folders/{foldersId}/bigQueryExports":                                           "FolderBigQueryExport",
	"securitycenter_v1.json:v1/organizations/{organizationsId}/bigQueryExports":                               "OrganizationBigQueryExport",
	"securitycenter_v1.json:v1/organizations/{organizationsId}/muteConfigs":                                   "OrganizationMuteConfig",
	"securitycenter_v1.json:v1/organizations/{organizationsId}/notificationConfigs":                           "OrganizationNotificationConfig",
	"securitycenter_v1.json:v1/projects/{projectsId}/bigQueryExports":                                         "ProjectBigQueryExport",
	"securitycenter_v1.json:v1/folders/{foldersId}/securityHealthAnalyticsSettings/customModules":             "FolderSecurityHealthAnalyticsSettingCustomModule",
	"securitycenter_v1.json:v1/organizations/{organizationsId}/securityHealthAnalyticsSettings/customModules": "OrganizationSecurityHealthAnalyticsSettingCustomModule",
	"securitycenter_v1.json:v1/organizations/{organizationsId}/eventThreatDetectionSettings/customModules":    "OrganizationEventThreatDetectionSettingCustomModule",
	"securitycenter_v1.json:v1/projects/{projectsId}/securityHealthAnalyticsSettings/customModules":           "ProjectSecurityHealthAnalyticsSettingCustomModule",

	// Storage.
	"storage_v1.json:b/{bucket}/o":              "BucketObject",
	"storage_v1.json:b/{bucket}/anywhereCaches": "",
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

// unsupportedPropertiesOverrides is a map of properties in types that
// are not supported and should be included in the Pulumi schema spec.
var unsupportedPropertiesOverrides = map[string][]string{
	// API ref: https://cloud.google.com/service-infrastructure/docs/service-management/reference/rest/v1/services.configs#backendrule
	"BackendRule": {
		// The `overridesByRequestProtocol` property in `BackendRule`,
		// while absent in the API reference docs, is present in the
		// discover doc that we fetch.
		//
		// But it doesn't make sense to include
		// `overridesByRequestProtocol` in the recursion.
		// For now don't include this property which means
		// we avoid the recursion.
		//
		// Discovery doc: https://servicemanagement.googleapis.com/$discovery/rest?version=v1
		"overridesByRequestProtocol",
	},
	"BackendRuleResponse": {
		// Similar to the property in `BackendRule`. See notes above.
		"overridesByRequestProtocol",
	},
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

// autonameExcludes is a set of resource tokens which should be explicitly excluded from autonaming.
var autonameExcludes = codegen.NewStringSet(
	"google-native:bigtableadmin/v1:Cluster",
	"google-native:bigtableadmin/v2:Cluster",
	"google-native:bigtableadmin/v2:Instance",
	"google-native:cloudresourcemanager/v3:TagKey",
	"google-native:cloudresourcemanager/v3:TagValue",
	"google-native:dialogflow/v3:Agent",
	"google-native:dialogflow/v3beta1:Agent",
	"google-native:monitoring/v3:NotificationChannel",
	"google-native:monitoring/v3:AlertPolicy",
	"google-native:monitoring/v3:UptimeCheckConfig",
	"google-native:vpcaccess/v1:Connector",
	"google-native:iam/v1:Role",
	"google-native:iam/v1:OrganizationRole",
	"google-native:run/v2:Service")

// metadataOverrides is a map of values static overlays to merge into the metadata for
// individual resource tokens. In case of conflict, the values in this mapping are preferred.
var metadataOverrides = map[string]resources.CloudAPIResource{
	"google-native:run/v1:Service": {
		Create: resources.CreateAPIOperation{
			CloudAPIOperation: resources.CloudAPIOperation{
				Polling: &resources.Polling{Strategy: resources.KNativeStatusPoll},
			},
		},
		Update: resources.UpdateAPIOperation{
			CloudAPIOperation: resources.CloudAPIOperation{
				Polling: &resources.Polling{Strategy: resources.KNativeStatusPoll},
			},
		},
	},
	"google-native:run/v1:DomainMapping": {
		Create: resources.CreateAPIOperation{
			CloudAPIOperation: resources.CloudAPIOperation{
				Polling: &resources.Polling{Strategy: resources.KNativeStatusPoll},
			},
		},
	},
	"google-native:bigtableadmin/v2:Instance": {
		Create: resources.CreateAPIOperation{
			CloudAPIOperation: resources.CloudAPIOperation{
				SDKProperties: map[string]resources.CloudAPIProperty{
					"clusters": {
						Required: true,
						ForceNew: true, // clusters are a sub-resource which don't provide a simple
						// means to update. Mark clusters as immutable and require replacement on
						//change to avoid errors with updates.
					},
				},
			},
		},
	},
	"google-native:storage/v1:Bucket": {
		Create: resources.CreateAPIOperation{
			CloudAPIOperation: resources.CloudAPIOperation{
				SDKProperties: map[string]resources.CloudAPIProperty{
					"location": {
						ForceNew: true,
					},
					"project": {
						ForceNew: true,
					},
					"name": {
						ForceNew: true,
					},
				},
			},
		},
	},
	"google-native:iam/v1:ServiceAccount": {
		Update: resources.UpdateAPIOperation{
			UpdateMask: resources.UpdateMask{
				BodyPropertyName: "updateMask",
			},
		},
	},
	"google-native:apigee/v1:Api": {
		Update: resources.UpdateAPIOperation{
			CloudAPIOperation: resources.CloudAPIOperation{
				Endpoint: resources.CloudAPIEndpoint{
					Template: "https://apigee.googleapis.com/v1/organizations/{organizationsId}/apis",
					Values: []resources.CloudAPIResourceParam{
						{
							Name: "action",
							Kind: "query",
						},
						{
							Name: "name",
							Kind: "query",
						},
						{
							Name: "validate",
							Kind: "query",
						},
						{
							Name:    "organizationsId",
							SdkName: "organizationId",
							Kind:    "path",
						},
					},
				},
				SDKProperties: map[string]resources.CloudAPIProperty{
					"contentType": {},
					"data": {
						Format: "byte",
					},
					"extensions": {
						Items: &resources.CloudAPIProperty{},
					},
				},

				Verb: "POST",
			},
			UpdateMask: resources.UpdateMask{},
		},
	},
	"google-native:container/v1:Cluster": {
		Create: resources.CreateAPIOperation{
			CloudAPIOperation: resources.CloudAPIOperation{
				Polling: &resources.Polling{Strategy: resources.ClusterAwaitRestingStatePoll},
			},
			RecordDefaults: map[string]resources.CloudAPIProperty{
				"loggingService":                        {},
				"monitoringService":                     {},
				"initialClusterVersion":                 {},
				"locations":                             {},
				"networkConfig.privateIpv6GoogleAccess": {},
				"nodeConfig.imageType":                  {},
			},
		},
		Update: resources.UpdateAPIOperation{
			CloudAPIOperation: resources.CloudAPIOperation{
				Polling: &resources.Polling{Strategy: resources.ClusterAwaitRestingStatePoll},
			},
		},
	},
	"google-native:container/v1beta1:Cluster": {
		Create: resources.CreateAPIOperation{
			CloudAPIOperation: resources.CloudAPIOperation{
				Polling: &resources.Polling{Strategy: resources.ClusterAwaitRestingStatePoll},
			},
			RecordDefaults: map[string]resources.CloudAPIProperty{
				"loggingService":                        {},
				"monitoringService":                     {},
				"initialClusterVersion":                 {},
				"locations":                             {},
				"networkConfig.privateIpv6GoogleAccess": {},
				"nodeConfig.imageType":                  {},
			},
		},
		Update: resources.UpdateAPIOperation{
			CloudAPIOperation: resources.CloudAPIOperation{
				Polling: &resources.Polling{Strategy: resources.ClusterAwaitRestingStatePoll},
			},
		},
	},
	"google-native:container/v1:NodePool": {
		Create: resources.CreateAPIOperation{
			CloudAPIOperation: resources.CloudAPIOperation{
				Polling: &resources.Polling{Strategy: resources.NodepoolAwaitRestingStatePoll},
			},
			RecordDefaults: map[string]resources.CloudAPIProperty{
				"version":          {},
				"config.imageType": {},
				"locations":        {},
			},
		},
		Update: resources.UpdateAPIOperation{
			CloudAPIOperation: resources.CloudAPIOperation{
				Polling: &resources.Polling{Strategy: resources.NodepoolAwaitRestingStatePoll},
			},
		},
	},
	"google-native:container/v1beta1:NodePool": {
		Create: resources.CreateAPIOperation{
			CloudAPIOperation: resources.CloudAPIOperation{
				Polling: &resources.Polling{Strategy: resources.NodepoolAwaitRestingStatePoll},
			},
			RecordDefaults: map[string]resources.CloudAPIProperty{
				"version":          {},
				"config.imageType": {},
				"locations":        {},
			},
		},
		Update: resources.UpdateAPIOperation{
			CloudAPIOperation: resources.CloudAPIOperation{
				Polling: &resources.Polling{Strategy: resources.NodepoolAwaitRestingStatePoll},
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
