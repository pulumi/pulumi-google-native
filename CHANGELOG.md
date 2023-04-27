CHANGELOG
=========

## HEAD (Unreleased)

Bug fixes:

- Remove `validateOnly` query parameter from SDK properties
  [#865](https://github.com/pulumi/pulumi-google-native/issues/865)

## v0.30.0 (2023-04-14)
Upstream breaking changes:
- Resource "google-native:sqladmin/v1beta4:Instance" missing input "availableMaintenanceVersions"
  - This input has been marked as "Output only" by GCP, and is now deprecated.
- Resource "google-native:sqladmin/v1:Instance" missing input "availableMaintenanceVersions"
  - This input has been marked as "Output only" by GCP, and is now deprecated.

## v0.29.0 (2023-04-07)
### Does the PR have any schema changes?

Upstream breaking changes:
- Resource "google-native:artifactregistry/v1beta2:Repository" missing input "createTime"
- Resource "google-native:artifactregistry/v1beta2:Repository" missing input "updateTime"
- Resource "google-native:privateca/v1beta1:Certificate" missing
- Resource "google-native:artifactregistry/v1:Repository" missing input "createTime"
- Resource "google-native:artifactregistry/v1:Repository" missing input "updateTime"
- Resource "google-native:privateca/v1beta1:CertificateAuthority" missing
- Resource "google-native:datafusion/v1:Instance" missing input "accelerators"
- Resource "google-native:datafusion/v1:Instance" missing input "availableVersion"
- Resource "google-native:batch/v1:JobIamPolicy" missing
- Resource "google-native:datafusion/v1beta1:Instance" missing input "accelerators"
- Resource "google-native:datafusion/v1beta1:Instance" missing input "availableVersion"
- Resource "google-native:recaptchaenterprise/v1:Key" missing input "createTime"
- Resource "google-native:appengine/v1:App" missing input "gcrDomain"
- Resource "google-native:batch/v1:JobIamBinding" missing
- Resource "google-native:artifactregistry/v1beta1:Repository" missing input "createTime"
- Resource "google-native:artifactregistry/v1beta1:Repository" missing input "updateTime"
- Resource "google-native:batch/v1:JobIamMember" missing
- Function "google-native:batch/v1:getJobIamPolicy" missing
- Function "google-native:privateca/v1beta1:getCertificateAuthority" missing
- Function "google-native:privateca/v1beta1:getCertificate" missing
- Type "google-native:privateca/v1beta1:ReusableConfigValues" missing
- Type "google-native:privateca/v1beta1:PublicKeyType" missing
- Type "google-native:privateca/v1beta1:KeyUsageResponse" missing
- Type "google-native:batch/v1:AuditLogConfig" missing
- Type "google-native:privateca/v1beta1:SubjectAltNames" missing
- Type "google-native:privateca/v1beta1:SubordinateConfigChain" missing
- Type "google-native:privateca/v1beta1:CertificateAuthorityPolicyResponse" missing
- Type "google-native:analyticshub/v1:RestrictedExportConfig" missing property "restrictDirectTableAccess"
- Type "google-native:datafusion/v1:AcceleratorState" missing
- Type "google-native:privateca/v1beta1:X509Extension" missing
- Type "google-native:privateca/v1beta1:KeyVersionSpecAlgorithm" missing
- Type "google-native:privateca/v1beta1:SubjectConfig" missing
- Type "google-native:privateca/v1beta1:CertificateConfig" missing
- Type "google-native:privateca/v1beta1:CaOptions" missing
- Type "google-native:batch/v1:AuditConfigResponse" missing
- Type "google-native:batch/v1:AuditLogConfigResponse" missing
- Type "google-native:privateca/v1beta1:IssuanceModesResponse" missing
- Type "google-native:analyticshub/v1beta1:RestrictedExportConfigResponse" missing property "restrictDirectTableAccess"
- Type "google-native:datafusion/v1:Version" missing
- Type "google-native:privateca/v1beta1:KeyVersionSpec" missing
- Type "google-native:privateca/v1beta1:KeyUsageOptions" missing
- Type "google-native:privateca/v1beta1:CaOptionsResponse" missing
- Type "google-native:privateca/v1beta1:AccessUrlsResponse" missing
- Type "google-native:privateca/v1beta1:CertificateConfigResponse" missing
- Type "google-native:privateca/v1beta1:KeyUsage" missing
- Type "google-native:privateca/v1beta1:ExtendedKeyUsageOptionsResponse" missing
- Type "google-native:privateca/v1beta1:AllowedConfigListResponse" missing
- Type "google-native:privateca/v1beta1:SubjectResponse" missing
- Type "google-native:privateca/v1beta1:SubjectDescriptionResponse" missing
- Type "google-native:datafusion/v1:AcceleratorAcceleratorType" missing
- Type "google-native:privateca/v1beta1:ReusableConfigWrapperResponse" missing
- Type "google-native:privateca/v1beta1:CertificateAuthorityPolicy" missing
- Type "google-native:datafusion/v1:VersionType" missing
- Type "google-native:privateca/v1beta1:ObjectIdResponse" missing
- Type "google-native:privateca/v1beta1:KeyUsageOptionsResponse" missing
- Type "google-native:batch/v1:BindingResponse" missing
- Type "google-native:privateca/v1beta1:PublicKey" missing
- Type "google-native:batch/v1:AuditConfig" missing
- Type "google-native:privateca/v1beta1:ObjectId" missing
- Type "google-native:privateca/v1beta1:Subject" missing
- Type "google-native:privateca/v1beta1:ReusableConfigValuesResponse" missing
- Type "google-native:privateca/v1beta1:X509ExtensionResponse" missing
- Type "google-native:privateca/v1beta1:SubordinateConfig" missing
- Type "google-native:privateca/v1beta1:SubjectConfigResponse" missing
- Type "google-native:privateca/v1beta1:KeyVersionSpecResponse" missing
- Type "google-native:privateca/v1beta1:KeyIdResponse" missing
- Type "google-native:privateca/v1beta1:AllowedConfigList" missing
- Type "google-native:privateca/v1beta1:IssuanceModes" missing
- Type "google-native:privateca/v1beta1:ReusableConfigWrapper" missing
- Type "google-native:datafusion/v1beta1:AcceleratorAcceleratorType" missing
- Type "google-native:datafusion/v1:Accelerator" missing
- Type "google-native:privateca/v1beta1:SubjectAltNamesResponse" missing
- Type "google-native:privateca/v1beta1:IssuingOptions" missing
- Type "google-native:privateca/v1beta1:SubordinateConfigChainResponse" missing
- Type "google-native:privateca/v1beta1:CertificateDescriptionResponse" missing
- Type "google-native:privateca/v1beta1:IssuingOptionsResponse" missing
- Type "google-native:privateca/v1beta1:CertificateFingerprintResponse" missing
- Type "google-native:privateca/v1beta1:CertificateAuthorityTier" missing
- Type "google-native:privateca/v1beta1:RevocationDetailsResponse" missing
- Type "google-native:privateca/v1beta1:CertificateAuthorityType" missing
- Type "google-native:analyticshub/v1beta1:RestrictedExportConfig" missing property "restrictDirectTableAccess"
- Type "google-native:datafusion/v1beta1:Accelerator" missing
- Type "google-native:batch/v1:AuditLogConfigLogType" missing
- Type "google-native:datafusion/v1beta1:AcceleratorState" missing
- Type "google-native:privateca/v1beta1:PublicKeyResponse" missing
- Type "google-native:batch/v1:Binding" missing
- Type "google-native:datafusion/v1beta1:Version" missing
- Type "google-native:privateca/v1beta1:AllowedSubjectAltNamesResponse" missing
- Type "google-native:privateca/v1beta1:AllowedSubjectAltNames" missing
- Type "google-native:privateca/v1beta1:SubordinateConfigResponse" missing
- Type "google-native:batch/v1:ExprResponse" missing
- Type "google-native:batch/v1:Expr" missing
- Type "google-native:privateca/v1beta1:ExtendedKeyUsageOptions" missing
- Type "google-native:datafusion/v1beta1:VersionType" missing
- Type "google-native:analyticshub/v1:RestrictedExportConfigResponse" missing property "restrictDirectTableAccess"

#### New resources:

- `appengine/v1.Application`
- `compute/alpha.StoragePool`
- `compute/alpha.StoragePoolIamBinding`
- `compute/alpha.StoragePoolIamMember`
- `compute/alpha.StoragePoolIamPolicy`
- `compute/beta.RegionInstanceTemplate`
- `dataform/v1beta1.RepositoryIamBinding`
- `dataform/v1beta1.RepositoryIamMember`
- `dataform/v1beta1.RepositoryIamPolicy`
- `dataform/v1beta1.RepositoryWorkspaceIamBinding`
- `dataform/v1beta1.RepositoryWorkspaceIamMember`
- `dataform/v1beta1.RepositoryWorkspaceIamPolicy`
- `datalineage/v1.LineageEvent`
- `datalineage/v1.Process`
- `datalineage/v1.Run`
- `firebaseappdistribution/v1.Group`
- `gkehub/v1.Binding`
- `gkehub/v1.Scope`
- `gkehub/v1alpha.Binding`
- `gkehub/v1alpha.Namespace`
- `gkehub/v1alpha.Rbacrolebinding`
- `gkehub/v1alpha.Scope`
- `gkehub/v1beta.Binding`
- `gkehub/v1beta.Namespace`
- `gkehub/v1beta.Rbacrolebinding`
- `gkehub/v1beta.Scope`
- `iam/v1.WorkforcePool`
- `iam/v1.WorkforcePoolIamBinding`
- `iam/v1.WorkforcePoolIamMember`
- `iam/v1.WorkforcePoolIamPolicy`
- `iam/v1.WorkforcePoolKey`
- `iam/v1.WorkforcePoolProvider`
- `iam/v1.WorkloadIdentityPoolKey`
- `logging/v2.BillingAccountBucketLink`
- `logging/v2.FolderBucketLink`
- `logging/v2.Link`
- `logging/v2.OrganizationBucketLink`
- `migrationcenter/v1alpha1.Group`
- `migrationcenter/v1alpha1.ImportDataFile`
- `migrationcenter/v1alpha1.ImportJob`
- `migrationcenter/v1alpha1.PreferenceSet`
- `migrationcenter/v1alpha1.Report`
- `migrationcenter/v1alpha1.ReportConfig`
- `migrationcenter/v1alpha1.Source`
- `networksecurity/v1beta1.AddressGroup`
- `networksecurity/v1beta1.GatewaySecurityPolicy`
- `networksecurity/v1beta1.OrganizationAddressGroup`
- `networksecurity/v1beta1.Rule`
- `networksecurity/v1beta1.TlsInspectionPolicy`
- `networksecurity/v1beta1.UrlList`
- `networkservices/v1.MulticastConsumerAssociationIamBinding`
- `networkservices/v1.MulticastConsumerAssociationIamMember`
- `networkservices/v1.MulticastConsumerAssociationIamPolicy`
- `networkservices/v1.MulticastDomainActivationIamBinding`
- `networkservices/v1.MulticastDomainActivationIamMember`
- `networkservices/v1.MulticastDomainActivationIamPolicy`
- `networkservices/v1.MulticastDomainIamBinding`
- `networkservices/v1.MulticastDomainIamMember`
- `networkservices/v1.MulticastDomainIamPolicy`
- `networkservices/v1.MulticastGroupDefinitionIamBinding`
- `networkservices/v1.MulticastGroupDefinitionIamMember`
- `networkservices/v1.MulticastGroupDefinitionIamPolicy`
- `networkservices/v1.MulticastGroupIamBinding`
- `networkservices/v1.MulticastGroupIamMember`
- `networkservices/v1.MulticastGroupIamPolicy`
- `notebooks/v2.InstanceIamBinding`
- `notebooks/v2.InstanceIamMember`
- `notebooks/v2.InstanceIamPolicy`
- `recaptchaenterprise/v1.Firewallpolicy`
- `retail/v2.Model`
- `retail/v2alpha.Model`
- `retail/v2beta.Model`
- `servicedirectory/v1beta1.NamespaceWorkloadIamBinding`
- `servicedirectory/v1beta1.NamespaceWorkloadIamMember`
- `servicedirectory/v1beta1.NamespaceWorkloadIamPolicy`
- `vpcaccess/v1beta1.Connector`

#### New functions:

- `appengine/v1.getApplication`
- `compute/alpha.getStoragePool`
- `compute/alpha.getStoragePoolIamPolicy`
- `compute/beta.getRegionInstanceTemplate`
- `dataform/v1beta1.getRepositoryIamPolicy`
- `dataform/v1beta1.getRepositoryWorkspaceIamPolicy`
- `datalineage/v1.getLineageEvent`
- `datalineage/v1.getProcess`
- `datalineage/v1.getRun`
- `firebaseappdistribution/v1.getGroup`
- `gkehub/v1.getBinding`
- `gkehub/v1.getScope`
- `gkehub/v1alpha.getBinding`
- `gkehub/v1alpha.getNamespace`
- `gkehub/v1alpha.getRbacrolebinding`
- `gkehub/v1alpha.getScope`
- `gkehub/v1beta.getBinding`
- `gkehub/v1beta.getNamespace`
- `gkehub/v1beta.getRbacrolebinding`
- `gkehub/v1beta.getScope`
- `iam/v1.getWorkforcePool`
- `iam/v1.getWorkforcePoolIamPolicy`
- `iam/v1.getWorkforcePoolKey`
- `iam/v1.getWorkforcePoolProvider`
- `iam/v1.getWorkloadIdentityPoolKey`
- `logging/v2.getBillingAccountBucketLink`
- `logging/v2.getFolderBucketLink`
- `logging/v2.getLink`
- `logging/v2.getOrganizationBucketLink`
- `migrationcenter/v1alpha1.getGroup`
- `migrationcenter/v1alpha1.getImportDataFile`
- `migrationcenter/v1alpha1.getImportJob`
- `migrationcenter/v1alpha1.getPreferenceSet`
- `migrationcenter/v1alpha1.getReport`
- `migrationcenter/v1alpha1.getReportConfig`
- `migrationcenter/v1alpha1.getSource`
- `networksecurity/v1beta1.getAddressGroup`
- `networksecurity/v1beta1.getGatewaySecurityPolicy`
- `networksecurity/v1beta1.getOrganizationAddressGroup`
- `networksecurity/v1beta1.getRule`
- `networksecurity/v1beta1.getTlsInspectionPolicy`
- `networksecurity/v1beta1.getUrlList`
- `networkservices/v1.getMulticastConsumerAssociationIamPolicy`
- `networkservices/v1.getMulticastDomainActivationIamPolicy`
- `networkservices/v1.getMulticastDomainIamPolicy`
- `networkservices/v1.getMulticastGroupDefinitionIamPolicy`
- `networkservices/v1.getMulticastGroupIamPolicy`
- `notebooks/v2.getInstanceIamPolicy`
- `recaptchaenterprise/v1.getFirewallpolicy`
- `retail/v2.getModel`
- `retail/v2alpha.getModel`
- `retail/v2beta.getModel`
- `servicedirectory/v1beta1.getNamespaceWorkloadIamPolicy`
- `vpcaccess/v1beta1.getConnector`

## 0.28.0 (2023-02-01)

### BREAKING CHANGE
Resource "google-native:batch/v1:NodeIamMember" missing
Resource "google-native:networkconnectivity/v1alpha1:ServiceConnectionPolicyIamBinding" missing
Resource "google-native:batch/v1:TaskIamPolicy" missing
Resource "google-native:networkconnectivity/v1alpha1:ServiceConnectionMapIamBinding" missing
Resource "google-native:networkconnectivity/v1alpha1:ServiceConnectionPolicyIamMember" missing
Resource "google-native:batch/v1:TaskIamBinding" missing
Resource "google-native:networkconnectivity/v1alpha1:ServiceClassIamMember" missing
Resource "google-native:batch/v1:NodeIamBinding" missing
Resource "google-native:contentwarehouse/v1:Document" missing input "asyncEnabled"
Resource "google-native:contentwarehouse/v1:Document" missing input "structuredContentUri"
Resource "google-native:contentwarehouse/v1:Document" missing output "structuredContentUri"
Resource "google-native:contentwarehouse/v1:Document" missing output "asyncEnabled"
Resource "google-native:networkconnectivity/v1alpha1:ServiceConnectionPolicyIamPolicy" missing
Resource "google-native:networkconnectivity/v1alpha1:ServiceClassIamBinding" missing
Resource "google-native:batch/v1:TaskIamMember" missing
Resource "google-native:networkconnectivity/v1alpha1:ServiceConnectionMapIamPolicy" missing
Resource "google-native:batch/v1:NodeIamPolicy" missing
Resource "google-native:networkconnectivity/v1alpha1:ServiceConnectionMapIamMember" missing
Resource "google-native:networkconnectivity/v1alpha1:ServiceClassIamPolicy" missing
Function "google-native:networkconnectivity/v1alpha1:getServiceConnectionPolicyIamPolicy" missing
Function "google-native:networkconnectivity/v1alpha1:getServiceConnectionMapIamPolicy" missing
Function "google-native:batch/v1:getTaskIamPolicy" missing
Function "google-native:batch/v1:getNodeIamPolicy" missing
Function "google-native:contentwarehouse/v1:getDocument" missing output "structuredContentUri"
Function "google-native:contentwarehouse/v1:getDocument" missing output "asyncEnabled"
Function "google-native:networkconnectivity/v1alpha1:getServiceClassIamPolicy" missing
Type "google-native:cloudtasks/v2beta2:UriOverride" missing property "query"
Type "google-native:cloudtasks/v2beta2:UriOverride" missing property "path"
Type "google-native:cloudtasks/v2beta2:UriOverrideResponse" missing property "path"
Type "google-native:cloudtasks/v2beta2:UriOverrideResponse" missing property "query"
Type "google-native:compute/alpha:InterconnectAttachmentConfigurationConstraintsResponse" missing property "networkConnectivityCenter"

### New resources:

- `accesscontextmanager/v1.AuthorizedOrgsDesc`
- `cloudidentity/v1.InboundSamlSsoProfile`
- `cloudidentity/v1.InboundSsoAssignment`
- `compute/beta.NetworkAttachment`
- `compute/beta.NetworkAttachmentIamBinding`
- `compute/beta.NetworkAttachmentIamMember`
- `compute/beta.NetworkAttachmentIamPolicy`
- `compute/v1.NetworkAttachment`
- `compute/v1.NetworkAttachmentIamBinding`
- `compute/v1.NetworkAttachmentIamMember`
- `compute/v1.NetworkAttachmentIamPolicy`
- `contactcenteraiplatform/v1alpha1.ContactCenter`
- `dataform/v1beta1.CompilationResult`
- `dataform/v1beta1.ReleaseConfig`
- `dataform/v1beta1.Repository`
- `dataform/v1beta1.WorkflowConfig`
- `dataform/v1beta1.WorkflowInvocation`
- `dataform/v1beta1.Workspace`
- `datamigration/v1.ConversionWorkspace`
- `datamigration/v1.PrivateConnection`
- `dataplex/v1.Attribute`
- `dataplex/v1.DataAttributeBinding`
- `dataplex/v1.DataAttributeBindingIamBinding`
- `dataplex/v1.DataAttributeBindingIamMember`
- `dataplex/v1.DataAttributeBindingIamPolicy`
- `dataplex/v1.DataScan`
- `dataplex/v1.DataScanIamBinding`
- `dataplex/v1.DataScanIamMember`
- `dataplex/v1.DataScanIamPolicy`
- `dataplex/v1.DataTaxonomy`
- `dataplex/v1.DataTaxonomyAttributeIamBinding`
- `dataplex/v1.DataTaxonomyAttributeIamMember`
- `dataplex/v1.DataTaxonomyAttributeIamPolicy`
- `dataplex/v1.DataTaxonomyIamBinding`
- `dataplex/v1.DataTaxonomyIamMember`
- `dataplex/v1.DataTaxonomyIamPolicy`
- `dataproc/v1.NodeGroup`
- `discoveryengine/v1alpha.Document`
- `discoveryengine/v1beta.Document`
- `metastore/v1.Backup`
- `metastore/v1.Federation`
- `metastore/v1.FederationIamBinding`
- `metastore/v1.FederationIamMember`
- `metastore/v1.FederationIamPolicy`
- `metastore/v1.MetadataImport`
- `metastore/v1.Service`
- `metastore/v1.ServiceBackupIamBinding`
- `metastore/v1.ServiceBackupIamMember`
- `metastore/v1.ServiceBackupIamPolicy`
- `metastore/v1.ServiceIamBinding`
- `metastore/v1.ServiceIamMember`
- `metastore/v1.ServiceIamPolicy`
- `monitoring/v3.Snooze`
- `networkconnectivity/v1alpha1.InternalRange`
- `networksecurity/v1beta1.AddressGroupIamBinding`
- `networksecurity/v1beta1.AddressGroupIamMember`
- `networksecurity/v1beta1.AddressGroupIamPolicy`
- `tpu/v2alpha1.QueuedResource`
- `translate/v3.Dataset`
- `translate/v3.Model`
- `workloadmanager/v1.Evaluation`
- `workstations/v1beta.Workstation`
- `workstations/v1beta.WorkstationCluster`
- `workstations/v1beta.WorkstationClusterWorkstationConfigIamBinding`
- `workstations/v1beta.WorkstationClusterWorkstationConfigIamMember`
- `workstations/v1beta.WorkstationClusterWorkstationConfigIamPolicy`
- `workstations/v1beta.WorkstationClusterWorkstationConfigWorkstationIamBinding`
- `workstations/v1beta.WorkstationClusterWorkstationConfigWorkstationIamMember`
- `workstations/v1beta.WorkstationClusterWorkstationConfigWorkstationIamPolicy`
- `workstations/v1beta.WorkstationConfig`

### New functions:

- `accesscontextmanager/v1.getAuthorizedOrgsDesc`
- `cloudidentity/v1.getInboundSamlSsoProfile`
- `cloudidentity/v1.getInboundSsoAssignment`
- `compute/beta.getNetworkAttachment`
- `compute/beta.getNetworkAttachmentIamPolicy`
- `compute/v1.getNetworkAttachment`
- `compute/v1.getNetworkAttachmentIamPolicy`
- `contactcenteraiplatform/v1alpha1.getContactCenter`
- `dataform/v1beta1.getCompilationResult`
- `dataform/v1beta1.getReleaseConfig`
- `dataform/v1beta1.getRepository`
- `dataform/v1beta1.getWorkflowConfig`
- `dataform/v1beta1.getWorkflowInvocation`
- `dataform/v1beta1.getWorkspace`
- `datamigration/v1.getConversionWorkspace`
- `datamigration/v1.getPrivateConnection`
- `dataplex/v1.getAttribute`
- `dataplex/v1.getDataAttributeBinding`
- `dataplex/v1.getDataAttributeBindingIamPolicy`
- `dataplex/v1.getDataScan`
- `dataplex/v1.getDataScanIamPolicy`
- `dataplex/v1.getDataTaxonomy`
- `dataplex/v1.getDataTaxonomyAttributeIamPolicy`
- `dataplex/v1.getDataTaxonomyIamPolicy`
- `dataproc/v1.getNodeGroup`
- `discoveryengine/v1alpha.getDocument`
- `discoveryengine/v1beta.getDocument`
- `metastore/v1.getBackup`
- `metastore/v1.getFederation`
- `metastore/v1.getFederationIamPolicy`
- `metastore/v1.getMetadataImport`
- `metastore/v1.getService`
- `metastore/v1.getServiceBackupIamPolicy`
- `metastore/v1.getServiceIamPolicy`
- `monitoring/v3.getSnooze`
- `networkconnectivity/v1alpha1.getInternalRange`
- `networksecurity/v1beta1.getAddressGroupIamPolicy`
- `tpu/v2alpha1.getQueuedResource`
- `translate/v3.getDataset`
- `translate/v3.getModel`
- `workloadmanager/v1.getEvaluation`
- `workstations/v1beta.getWorkstation`
- `workstations/v1beta.getWorkstationCluster`
- `workstations/v1beta.getWorkstationClusterWorkstationConfigIamPolicy`
- `workstations/v1beta.getWorkstationClusterWorkstationConfigWorkstationIamPolicy`
- `workstations/v1beta.getWorkstationConfig`

## 0.27.0 (2022-10-21)
### BREAKING CHANGE
Upstream API changes affected the following:
- `Resource "google-native:securitycenter/v1:NotificationConfig" missing input "organizationId"`
- `Resource "google-native:securitycenter/v1:NotificationConfig" missing output "organizationId"`
- `Function "google-native:securitycenter/v1:getNotificationConfig" missing input "organizationId"`
- `Type "google-native:datacatalog/v1:GoogleCloudDatacatalogV1ColumnSchemaLookerColumnSpecResponse" missing`
- `Type "google-native:datacatalog/v1:GoogleCloudDatacatalogV1ColumnSchema" missing property "lookerColumnSpec"`
- `Type "google-native:datacatalog/v1:GoogleCloudDatacatalogV1UsageSignalResponse" missing property "favoriteCount"`
- `Type "google-native:datacatalog/v1:GoogleCloudDatacatalogV1ColumnSchemaResponse" missing property "lookerColumnSpec"`
- `Type "google-native:compute/alpha:Scheduling" missing property "dynamicResizeProperties"`
- `Type "google-native:datacatalog/v1:GoogleCloudDatacatalogV1ColumnSchemaLookerColumnSpecType" missing`
- `Type "google-native:compute/alpha:SchedulingDynamicResizePropertiesHotStandbyState" missing`
- `Type "google-native:compute/alpha:SchedulingDynamicResizeProperties" missing`
- `Type "google-native:datacatalog/v1:GoogleCloudDatacatalogV1ColumnSchemaLookerColumnSpec" missing`
- `Type "google-native:compute/alpha:SchedulingResponse" missing property "dynamicResizeProperties"`
- `Type "google-native:datacatalog/v1:GoogleCloudDatacatalogV1UsageSignal" missing property "favoriteCount"`
- `Type "google-native:compute/alpha:SchedulingDynamicResizePropertiesResponse" missing`

### New resources:

- `beyondcorp/v1alpha.NetConnection`
- `beyondcorp/v1alpha.NetConnectionIamBinding`
- `beyondcorp/v1alpha.NetConnectionIamMember`
- `beyondcorp/v1alpha.NetConnectionIamPolicy`
- `certificatemanager/v1.CertificateIssuanceConfig`
- `cloudidentity/v1beta1.InboundSamlSsoProfile`
- `cloudidentity/v1beta1.InboundSsoAssignment`
- `contentwarehouse/v1.Document`
- `contentwarehouse/v1.DocumentSchema`
- `contentwarehouse/v1.RuleSet`
- `contentwarehouse/v1.SynonymSet`
- `securitycenter/v1.FolderNotificationConfig`
- `securitycenter/v1.OrganizationNotificationConfig`

### New functions:

- `beyondcorp/v1alpha.getNetConnection`
- `beyondcorp/v1alpha.getNetConnectionIamPolicy`
- `certificatemanager/v1.getCertificateIssuanceConfig`
- `cloudidentity/v1beta1.getInboundSamlSsoProfile`
- `cloudidentity/v1beta1.getInboundSsoAssignment`
- `contentwarehouse/v1.getDocument`
- `contentwarehouse/v1.getDocumentSchema`
- `contentwarehouse/v1.getRuleSet`
- `contentwarehouse/v1.getSynonymSet`
- `securitycenter/v1.getFolderNotificationConfig`
- `securitycenter/v1.getOrganizationNotificationConfig`

## 0.26.1 (2022-10-07)

### BUG FIXES/ENHANCEMENTS
- Fix Conditions handling for IAM policy updates [#696](https://github.com/pulumi/pulumi-google-native/pull/696)
- Improve error handling for getClientToken method [#705](https://github.com/pulumi/pulumi-google-native/pull/705)

## 0.26.0 (2022-09-16)

### BUG FIXES/ENHANCEMENTS
- Add IamBinding and IamMember resources to make it easier to manage IAM policies [#653](https://github.com/pulumi/pulumi-google-native/pull/653)

## 0.25.0 (2022-09-01)
### BREAKING CHANGE
- Remove DNS v2 API which is not actively advertised by Google [#662](https://github.com/pulumi/pulumi-google-native/pull/662)

### BUG FIXES/ENHANCEMENTS
- Add resource method for retrieving kubeconfig from a GKE cluster [#655](https://github.com/pulumi/pulumi-google-native/pull/655)
- Stack level project configuration marked optional [#668](https://github.com/pulumi/pulumi-google-native/pull/668/)
- Retrieve configuration options for default provider from environment variables [#668](https://github.com/pulumi/pulumi-google-native/pull/668/)

## 0.24.0 (2022-08-25)
### BREAKING CHANGE
Upstream API changes affected the following:
- Resource "google-native:compute/alpha:RegionNetwork" removed
- Resource "google-native:compute/alpha:RegionNetworkIamPolicy" removed
- Function "google-native:compute/alpha:getRegionNetworkIamPolicy" removed
- Function "google-native:compute/alpha:getRegionNetwork" removed
- Resource "google-native:firebase/v1beta1:AndroidApp" removed input "sha256Hashes"
- Resource "google-native:firebase/v1beta1:AndroidApp" removed input "sha1Hashes"
- Resource "google-native:firebase/v1beta1:AndroidApp" removed output "sha1Hashes"
- Resource "google-native:firebase/v1beta1:AndroidApp" removed output "sha256Hashes"
- Type "google-native:datacatalog/v1:GoogleCloudDatacatalogV1CloudSqlBigQueryConnectionSpecType" removed
- Type "google-native:networkmanagement/v1beta1:AppEngineVersionInfoResponse" removed
- Type "google-native:datacatalog/v1:GoogleCloudDatacatalogV1BigQueryConnectionSpecConnectionType" removed
- Type "google-native:certificatemanager/v1:ProvisioningIssue" removed
- Type "google-native:datacatalog/v1:GoogleCloudDatacatalogV1DataSourceConnectionSpec" removed property "bigqueryConnectionSpec"
- Type "google-native:compute/alpha:RegionNetworkNetworkFirewallPolicyEnforcementOrder" removed
- Type "google-native:certificatemanager/v1:ManagedCertificate" removed property "provisioningIssue"
- Type "google-native:datacatalog/v1:GoogleCloudDatacatalogV1CloudSqlBigQueryConnectionSpec" removed
- Type "google-native:certificatemanager/v1:ProvisioningIssueReason" removed
- Type "google-native:datacatalog/v1:GoogleCloudDatacatalogV1BigQueryConnectionSpec" removed
- Type "google-native:networkmanagement/v1beta1:StepResponse" removed property "appEngineVersion"

### BUG FIXES/ENHANCEMENTS
- Custom mutations for v1/v1beta1 cluster and nodepool [#632](https://github.com/pulumi/pulumi-google-native/pull/632)
- Disable autonaming for `google-native:iam:v1/Role` [#646](https://github.com/pulumi/pulumi-google-native/pull/646)
- Add support for container service getServerConfig [#648](https://github.com/pulumi/pulumi-google-native/pull/648)

## 0.23.0 (2022-07-29)
- Ensure inputs in partial error checkpoints for nodepools are correctly resumable [#602](https://github.com/pulumi/pulumi-google-native/pull/602)
- Fix operation to target resolution [#615](https://github.com/pulumi/pulumi-google-native/pull/615)
- Fix schema types for query params [#616](https://github.com/pulumi/pulumi-google-native/pull/616)
- Disable autonaming for vpcaccess/v1:Connector [#618](https://github.com/pulumi/pulumi-google-native/pull/618)

### BREAKING CHANGE
- A number of properties were previously modeled as strings in the SDK incorrectly previously. See [here](https://github.com/pulumi/pulumi-google-native/pull/616#issuecomment-1210150665) for affected resources. [#616](https://github.com/pulumi/pulumi-google-native/pull/616)
- Multiple breaking changes were made to the gameservices v1 and v1beta API spec by Google. See [622](https://github.com/pulumi/pulumi-google-native/pull/622) for details

## 0.22.0 (2022-07-29)
- Add support for apigee resources that use multipart/form-data content-type [#590](https://github.com/pulumi/pulumi-google-native/pull/590)
- Add support for nodepool mutations [#588](https://github.com/pulumi/pulumi-google-native/pull/588)
- Mark immutable fields for nodepools as replacements in diff [#598](https://github.com/pulumi/pulumi-google-native/pull/598)
- Update discovery docs [#600](https://github.com/pulumi/pulumi-google-native/pull/600)

### BREAKING CHANGE
- run/v1alpha1 API version of Cloud Run was removed [#600](https://github.com/pulumi/pulumi-google-native/pull/600)

## 0.21.0 (2022-07-14)
- Handle network timeouts as retryable [#524](https://github.com/pulumi/pulumi-google-native/pull/524)
- Add path parameters as required outputs [#539](https://github.com/pulumi/pulumi-google-native/pull/539)
- Fix handling for resource moving from named to autonamed resources [#551](https://github.com/pulumi/pulumi-google-native/pull/551)
- Replace on changes when changes to path or required query params [#551](https://github.com/pulumi/pulumi-google-native/pull/551)
- All Java packages changed to be more intuitive, unfortunately this is a breaking change for Java version [#554](https://github.com/pulumi/pulumi-google-native/pull/554)

## 0.20.0 (2022-06-06)
- Multiple bug fixes in [#453](https://github.com/pulumi/pulumi-google-native/pull/453)
  - Store partial state more consistently in the presence of partial failure
  - Fixes for bigtable instance and cluster recreation
  - Add support for replace-on-changes based on annotations and manual override for replace-on-change behavior
  - Make method discovery deterministic
  - Improvements to discovering operation endpoints
- Set force-new for storage bucket resource [#516](https://github.com/pulumi/pulumi-google-native/pull/516)
- Override field-mask for iam:v1/ServiceAccount [#518](https://github.com/pulumi/pulumi-google-native/pull/518)

## 0.19.1 (2022-05-24)

- fix: Retry GKE cluster operations [#496](https://github.com/pulumi/pulumi-google-native/issues/496)

## 0.19.0 (2022-05-19)

- Update discovery docs

### BUG FIXES
- Disable autonaming for certain resources [#291](https://github.com/pulumi/pulumi-google-native/issues/291)
- Revert "use sequence numbers to generate deterministic autonames" [#435](https://github.com/pulumi/pulumi-google-native/pull/435)

### BREAKING CHANGES
- Resource `google-native:bigquery/v2:Dataset` - output `satisfiesPZS` was renamed
- Function `google-native:bigquery/v2:getDataset` output `satisfiesPZS` was renamed
- Resource `google-native:networkservices/v1beta1:TcpRoute` missing input `gateways`
- Resource `google-native:networkservices/v1beta1:TcpRoute` missing output `gateways`
- Resource `google-native:baremetalsolution/v2:Snapshot` was removed
- Resource `google-native:metastore/v1alpha:Federation` was removed
- Resource `google-native:baremetalsolution/v2:SnapshotSchedulePolicy` was removed
- Function `google-native:metastore/v1alpha:getFederation` was removed
- Function `google-native:baremetalsolution/v2:getSnapshotSchedulePolicy` was removed
- Function `google-native:baremetalsolution/v2:getSnapshot` was removed
- Type `google-native:bigquery/v2:TableFieldSchemaResponse`  `collationSpec` renamed to `collation`
- Type `google-native:run/v2:GoogleCloudRunV2RevisionTemplate` property `confidential` removed
- Type `google-native:run/v2:GoogleCloudRunV2RevisionTemplate` property `containerConcurrency` renamed to `maxInstanceRequestConcurrency`
- Type `google-native:run/v2:GoogleCloudRunV2ConditionResponse` property `domainMappingReason` removed
- Type `google-native:run/v2:GoogleCloudRunV2ConditionResponse` property `internalReason` removed
- Type `google-native:run/v2:GoogleCloudRunV2BinaryAuthorization` property `policy` removed
- Type `google-native:run/v2:GoogleCloudRunV2BinaryAuthorizationResponse` property `policy` removed
- Type `google-native:bigquery/v2:TableFieldSchema` property `collationSpec` renamed to `collation`
- Type `google-native:run/v2:GoogleCloudRunV2RevisionTemplateResponse` property `confidential` removed
- Type `google-native:run/v2:GoogleCloudRunV2RevisionTemplateResponse`  property `containerConcurrency` renamed to `maxInstanceRequestConcurrency`
- Type `google-native:networkmanagement/v1beta1:StepResponse` missing property `appEngineVersionInfo`
- Type `google-native:baremetalsolution/v2:Schedule` was removed
- Type `google-native:baremetalsolution/v2:ScheduleResponse` was removed
- Type `google-native:compute/alpha:InstanceGroupManagerAutoHealingPolicyUpdateInstances` was removed
- Type `google-native:baremetalsolution/v2:SnapshotSchedulePolicyState` was removed
- Type `google-native:compute/alpha:InstanceGroupManagerAutoHealingPolicy` property `updateInstances`
- Type `google-native:compute/alpha:InstanceGroupManagerAutoHealingPolicyResponse` missing property `updateInstances`


## 0.18.2 (2022-05-02)

Improvements:

- Support attach
  [#460](https://github.com/pulumi/pulumi-google-native/pull/460)

Bug fixes:

- Convert token expiry to string on assignment [#470](https://github.com/pulumi/pulumi-google-native/pull/470)

Improvements:
- Initial support for handling immutable resources [453](https://github.com/pulumi/pulumi-google-native/pull/453)

Bug fixes:
- https://github.com/pulumi/pulumi-google-native/issues/426
- https://github.com/pulumi/pulumi-google-native/issues/291

## 0.18.1 (2022-04-20)

Improvements:

- Add getClientConfig and getClientToken functions [#451](https://github.com/pulumi/pulumi-google-native/pull/451)

Bug fixes:

- Add custom retry logic for cloud run [#411](https://github.com/pulumi/pulumi-google-native/pull/411)

## 0.18.0 (2022-03-30)
- Reflect renamed types as defined in cloud run v2 API specs [#419](https://github.com/pulumi/pulumi-google-native/pull/419)
- Following resources have been dropped [#419](https://github.com/pulumi/pulumi-google-native/pull/419):
  * `google-native:gkehub/v2alpha:FeatureConfig`
  * `google-native:gkehub/v2alpha:Feature`
- Following resources had some of their inputs removed (marked as output fields) [#419](https://github.com/pulumi/pulumi-google-native/pull/419):
  * `google-native:compute/alpha:FutureReservation`
  * `google-native:compute/alpha:NetworkEdgeSecurityService`

## 0.17.1 (2022-03-17)
Bug fixes:

- Fix regression in schema decompressor [#397](https://github.com/pulumi/pulumi-google-native/pull/397)

## 0.17.0 (2022-03-16)
Improvements:

- Re-introduce deprecated fields [#368](https://github.com/pulumi/pulumi-google-native/pull/368)
- Check for override functions on resource delete [#381](https://github.com/pulumi/pulumi-google-native/pull/381)
- Add missing descriptions for many resource properties [#384](https://github.com/pulumi/pulumi-google-native/pull/384)
- Handle updateMask for Update operations [#386](https://github.com/pulumi/pulumi-google-native/pull/386)

## 0.16.0 (2022-03-11)
Improvements:

- Use schema information for all REST operations [#374](https://github.com/pulumi/pulumi-google-native/pull/374)

## 0.15.0 (2022-02-28)
Improvements:

- Update to support the latest resource definitions and the latest Pulumi SDK
- Implement auto-naming for KMS resources [#351](https://github.com/pulumi/pulumi-google-native/pull/351)

#### Breaking Changes:
Resource "google-native:vmmigration/v1:CloneJob" missing input "name"
Resource "google-native:vmmigration/v1:Group" missing input "name"
Resource "google-native:vmmigration/v1alpha1:CloneJob" missing input "name"
Resource "google-native:vmmigration/v1alpha1:Group" missing input "name"
Resource "google-native:vmmigration/v1:TargetProject" missing input "name"
Resource "google-native:vmmigration/v1alpha1:TargetProject" missing input "name"
Type "google-native:vmmigration/v1alpha1:ReplicatingStepResponse" missing
Type "google-native:vmmigration/v1alpha1:PostProcessingStepResponse" missing
Type "google-native:gameservices/v1:AuditConfigResponse" missing property "exemptedMembers"
Type "google-native:vmmigration/v1alpha1:InitializingReplicationStepResponse" missing
Type "google-native:gameservices/v1beta:AuditConfig" missing property "exemptedMembers"
Type "google-native:vmmigration/v1alpha1:CycleStepResponse" missing
Type "google-native:vmmigration/v1alpha1:ReplicationCycleResponse" missing property "steps"
Type "google-native:vmmigration/v1alpha1:ReplicationCycleResponse" missing property "endTime"
Type "google-native:gameservices/v1beta:AuditConfigResponse" missing property "exemptedMembers"
Type "google-native:gameservices/v1:AuditConfig" missing property "exemptedMembers"

#### New resources:

- `apigee/v1.EndpointAttachment`
- `certificatemanager/v1.Certificate`
- `certificatemanager/v1.CertificateMap`
- `certificatemanager/v1.CertificateMapEntry`
- `certificatemanager/v1.DnsAuthorization`
- `cloudfunctions/v2alpha.Function`
- `cloudfunctions/v2alpha.FunctionIamPolicy`
- `cloudfunctions/v2beta.Function`
- `cloudfunctions/v2beta.FunctionIamPolicy`
- `dataplex/v1.Asset`
- `dataplex/v1.Contentitem`
- `dataplex/v1.Entity`
- `dataplex/v1.Environment`
- `dataplex/v1.Lake`
- `dataplex/v1.LakeAssetIamPolicy`
- `dataplex/v1.LakeContentIamPolicy`
- `dataplex/v1.LakeEnvironmentIamPolicy`
- `dataplex/v1.LakeIamPolicy`
- `dataplex/v1.LakeTaskIamPolicy`
- `dataplex/v1.LakeZoneIamPolicy`
- `dataplex/v1.Partition`
- `dataplex/v1.Task`
- `dataplex/v1.Zone`
- `gkehub/v2alpha.Feature`
- `gkehub/v2alpha.FeatureConfig`
- `metastore/v1beta.ServiceDatabaseIamPolicy`
- `metastore/v1beta.ServiceDatabaseTableIamPolicy`

#### New functions:

- `apigee/v1.getEndpointAttachment`
- `certificatemanager/v1.getCertificate`
- `certificatemanager/v1.getCertificateMap`
- `certificatemanager/v1.getCertificateMapEntry`
- `certificatemanager/v1.getDnsAuthorization`
- `cloudfunctions/v2alpha.getFunction`
- `cloudfunctions/v2alpha.getFunctionIamPolicy`
- `cloudfunctions/v2beta.getFunction`
- `cloudfunctions/v2beta.getFunctionIamPolicy`
- `dataplex/v1.getAsset`
- `dataplex/v1.getContentitem`
- `dataplex/v1.getEntity`
- `dataplex/v1.getEnvironment`
- `dataplex/v1.getLake`
- `dataplex/v1.getLakeAssetIamPolicy`
- `dataplex/v1.getLakeContentIamPolicy`
- `dataplex/v1.getLakeEnvironmentIamPolicy`
- `dataplex/v1.getLakeIamPolicy`
- `dataplex/v1.getLakeTaskIamPolicy`
- `dataplex/v1.getLakeZoneIamPolicy`
- `dataplex/v1.getPartition`
- `dataplex/v1.getTask`
- `dataplex/v1.getZone`
- `gkehub/v2alpha.getFeature`
- `gkehub/v2alpha.getFeatureConfig`
- `metastore/v1beta.getServiceDatabaseIamPolicy`
- `metastore/v1beta.getServiceDatabaseTableIamPolicy`

## 0.14.0 (2022-02-02)

- Update minimum supported Go version to 1.16

Improvements:

- Update to support the latest resource definitions and the latest Pulumi SDK
- Upgrade Google API client to v0.66.0

---

## 0.13.0 (2022-01-27)

Improvements:

- Update to support the latest resource definitions and the latest Pulumi SDK
- Expose project, region, and zone properties on Provider [#310](https://github.com/pulumi/pulumi-google-native/pull/310)

## 0.12.0 (2022-01-26)

Improvements:

- Update to support the latest resource definitions and the latest Pulumi SDK

## 0.11.0 (2022-01-10)

Improvements:

- Update to support the latest resource definitions and the latest Pulumi SDK

Bug fixes:

- Fix Project resource projectId property in SDKs.

## 0.10.1 (2021-12-22)

Bug fixes:

- Serve valid (decompressed schema) from GetSchema RPC calls

## 0.10.0 (2021-12-06)

Improvements:

- Update to support the latest resource definitions and the latest Pulumi SDK

Bug fixes:

- Disable unconditional replaces. No replaces will be triggered by the provider
  at this time. Use the `replaceOnChanges` resource option to trigger a
  replacement operation manually.
  [#185](https://github.com/pulumi/pulumi-google-native/issues/185)

## 0.9.0 (2021-11-24)

- Update to support the latest resource definitions and the latest Pulumi SDK

## 0.8.0 (2021-10-08)

- Update to support the latest resource definitions and the latest Pulumi SDK

## 0.7.0 (2021-08-06)

- Use the appropriate HTTP verbs for invokes [#175](https://github.com/pulumi/pulumi-google-native/pull/175)

- Avoid resetting project IAM policy on deletion [#175](https://github.com/pulumi/pulumi-google-native/pull/175)

- Improve replacement detection [#181](https://github.com/pulumi/pulumi-google-native/pull/181)

## 0.6.0 (2021-07-24)

- Update dependencies to pulumi/pulumi 3.7.0

- Refresh credentials for client if backing credentials are expired. [#156](https://github.com/pulumi/pulumi-google-native/pull/156)

- Allow for shorthand property values instead of duplicating project, zone, etc.
  [#56](https://github.com/pulumi/pulumi-google-native/issues/56)

- Populate `project`, `location`, `zone` properties automatically from config values
  [#94](https://github.com/pulumi/pulumi-google-native/issues/94)

- Fix a panic when an array property is updated to a larger number of elements
  [#160](https://github.com/pulumi/pulumi-google-native/issues/160)

- Support initialization failures by checkpointing partially created resources into the state
  [#149](https://github.com/pulumi/pulumi-google-native/issues/149)

## 0.5.0 (2021-07-12)

Enhancements:

- Populate fingerprint properties automatically
  [#126](https://github.com/pulumi/pulumi-google-native/issues/126)

- Auto-naming. Resource names (as seen in Google Cloud) are now automatically
  generated unless they are specified in user's program. A dash plus a random
  suffix of length 7 are appended to the logical name of Pulumi resources.
  Auto-naming is not yet available in some resource types; those resources
  have a note "Auto-naming is currently not supported for this resource"
  in their description.
  [#93](https://github.com/pulumi/pulumi-google-native/issues/93)

- Fix the representation of properties marked as `[Output only]`
  [#134](https://github.com/pulumi/pulumi-google-native/issues/134),
  [#135](https://github.com/pulumi/pulumi-google-native/issues/135)

- Normalize property naming to lowerCamelCase
  [#146](https://github.com/pulumi/pulumi-google-native/pull/146)

- Mark a property as required if description starts with `Required`
  [#145](https://github.com/pulumi/pulumi-google-native/pull/145)

- `storage/*.getObject` functions renamed to `storage/*.getBucketObject` to be
  consistent with the resource name

- Respect all metadata when creating `storage/*.BucketObject`
  [#74](https://github.com/pulumi/pulumi-google-native/issues/74)

## 0.4.0 (2021-06-16)

New features:

- Add a corresponding Get function for each resource
  [#82](https://github.com/pulumi/pulumi-google-native/issues/82)
- Add enums to all SDKs for properties that are annotated as enums in discovery docs
  [#86](https://github.com/pulumi/pulumi-google-native/issues/86)

Bug fixes:

- Fix idpath for storagetransfer/v1:TransferJob [#116](https://github.com/pulumi/pulumi-google-native/pull/116)

Breaking changes:

- Some child types renamed to match the parent type name
  [#117](https://github.com/pulumi/pulumi-google-native/pull/117)

## 0.3.0 (2021-06-09)

All the following changes are extensively breaking. We do our best to stabilize
the resource shapes as soon as possible.

- Refine resource naming to avoid excidingly long names and unobserved conflicts.
  [#92](https://github.com/pulumi/pulumi-google-native/issues/92).

- Stop generating artificial input properties for ID path segments. Calculate
  resource IDs automatically.
  [#100](https://github.com/pulumi/pulumi-google-native/pull/100),
  [#97](https://github.com/pulumi/pulumi-google-native/issues/97)

- Normalize parameter names to singular names similar to properties
  [#96](https://github.com/pulumi/pulumi-google-native/pull/96),
  [#89](https://github.com/pulumi/pulumi-google-native/issues/89)

- Extend the pattern to catch more deprecation messages for properties and parameters
  [#99](https://github.com/pulumi/pulumi-google-native/issues/99)

## 0.2.0 (2021-05-08)

- Fix refresh to re-populate inputs and provide accurate diff and drift detection
  [#65](https://github.com/pulumi/pulumi-google-native/issues/65)

- Generate properties for optional query parameters. Fix pub/sub schema creation
  [#67](https://github.com/pulumi/pulumi-google-native/issues/67)

## 0.1.1 (2021-04-26)

- Fix ManagedZoneRrset creation
  [#61](https://github.com/pulumi/pulumi-google-native/issues/61)

## 0.1.0 (2021-04-19)

The first beta release of the native Google Cloud Provider is out!
