// Copyright 2016-2021, Pulumi Corporation.

package gen

// resourceNamePropertyOverrides is a list of exceptions populated for the buildIdParams method above.
var resourceNamePropertyOverrides = map[string]string{
	"apigee/v1:OrganizationEnvironmentKeystoreAlias.aliasesId":       "alias",
	"appengine/v1:AppFirewallIngressRule.ingressRulesId":             "priority",
	"appengine/v1beta:AppFirewallIngressRule.ingressRulesId":         "priority",
	"bigquery/v2:Routine.routinesId":                                 "routineReference.routineId",
	"dataproc/v1:RegionJob.jobId":                                    "reference.jobId",
	"dataproc/v1beta2:RegionJob.jobId":                               "reference.jobId",
	"dns/v1:Change.changeId":                                         "id",
	"dns/v1beta2:Change.changeId":                                    "id",
	"dns/v1beta2:ResponsePolicy.responsePolicy":                      "id",
	"dns/v1beta2:ResponsePolicyRule.responsePolicyRule":              "ruleName",
	"firebasehosting/v1beta1:SiteDomain.domainsId":                   "domainName",
	"recommendationengine/v1beta1:CatalogCatalogItem.catalogItemsId": "id",
	"run/v1alpha1:NamespaceJob.jobsId":                               "metadata.name",
	"run/v1:Domainmapping.domainmappingsId":                          "metadata.name",
	"run/v1:NamespaceDomainmapping.domainmappingsId":                 "metadata.name",
	"run/v1:NamespaceService.servicesId":                             "metadata.name",
	"run/v1:Service.servicesId":                                      "metadata.name",
}
