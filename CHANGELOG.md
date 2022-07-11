CHANGELOG
=========

## HEAD (Unreleased)
- Handle network timeouts as retryable [#524](https://github.com/pulumi/pulumi-google-native/pull/524)
- Add path parameters as required outputs [#539](https://github.com/pulumi/pulumi-google-native/pull/539)
- Fix handling for resource moving from named to autonamed resources [#551](https://github.com/pulumi/pulumi-google-native/pull/551)
- Replace on changes when changes to path or required query params [#551](https://github.com/pulumi/pulumi-google-native/pull/551)

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
