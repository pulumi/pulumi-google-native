CHANGELOG
=========

## HEAD (Unreleased)

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
