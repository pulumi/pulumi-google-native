CHANGELOG
=========

## HEAD (Unreleased)

- Stop generating articial input properties for ID path segments. Calculate
  resource IDs automatically.
  [#100](https://github.com/pulumi/pulumi-google-native/pull/100),
  [#97](https://github.com/pulumi/pulumi-google-native/issues/97)
- Normalize parameter names to singular names similar to properties
  [#96](https://github.com/pulumi/pulumi-google-native/pull/96),
  [#89](https://github.com/pulumi/pulumi-google-native/issues/89)

---

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
