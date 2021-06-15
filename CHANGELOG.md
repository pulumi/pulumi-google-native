CHANGELOG
=========

## HEAD (Unreleased)

New features:

- Add a corresponding Get function for each resource
  [#82](https://github.com/pulumi/pulumi-google-native/issues/82)
- Add enums to all SDKs for properties that are annotated as enums in discovery docs
  [#86](https://github.com/pulumi/pulumi-google-native/issues/86)

Breaking changes:

- Some child types renamed to match the parent type name
  [#117](https://github.com/pulumi/pulumi-google-native/pull/117)

---

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
