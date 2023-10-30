[![Slack](http://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Fgoogle-native.svg)](https://npmjs.com/package/@pulumi/google-native)
[![Python version](https://badge.fury.io/py/pulumi-google-native.svg)](https://pypi.org/project/pulumi-google-native)
[![NuGet version](https://badge.fury.io/nu/pulumi.googlenative.svg)](https://badge.fury.io/nu/pulumi.googlenative)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-google-native/sdk/go)](https://pkg.go.dev/github.com/pulumi/pulumi-google-native/sdk/go)
[![License](https://img.shields.io/npm/l/%40pulumi%2Fgoogle-native.svg)](https://github.com/pulumi/pulumi-google-native/blob/master/LICENSE)

# Native Google Cloud Pulumi Provider (developer preview)

The native Google Cloud Provider for Pulumi lets you provision Google Cloud resources in your cloud programs.

This provider uses the Google Cloud REST API directly and therefore provides full access to Google Cloud.

> **Warning:** Google Cloud Native is available in developer preview. While upstream changes continue to be delivered, active development is paused. Breaking changes may be introduced in minor version releases.

To use this package, [install the Pulumi CLI](https://www.pulumi.com/docs/get-started/install/).

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @pulumi/google-native
```

or `yarn`:

```bash
yarn add @pulumi/google-native
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_google_native
```

### Go

To use from Go, use `go get` to grab the latest version of the library

```bash
go get github.com/pulumi/pulumi-google-native/sdk
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.GoogleNative
```

## Concepts

The native Google Cloud package provides a strongly-typed means to build cloud applications that create
and interact closely with Google Cloud resources.  Resources are exposed for the entire Google Cloud surface area,
including (but not limited to) 'compute', 'container', 'run', 'storage', and more.

The native Google Cloud provider works directly with the Google Cloud API instead of depending on a
handwritten layer as with the [Google Cloud Classic Provider](https://github.com/pulumi/pulumi-gcp). This approach ensures higher
quality and higher fidelity with Google Cloud.

## Configuring credentials

To learn how to configure credentials refer to the [Google Cloud configuration options](https://www.pulumi.com/registry/packages/google-native/installation-configuration/#configuration-options).

When developing locally, we recommend that you install the Google Cloud SDK and then authorize access with a user
account. Other configuration settings should be set either via `pulumi config set google-native:<KEY> <VALUE>` or
pass options to the constructor of a new [google-native `Provider`](https://www.pulumi.com/registry/packages/google-native/api-docs/provider/).

## Building

### Dependencies

- Go 1.15
- NodeJS 10.X.X or later
- Python 3.6 or later
- .NET Core 3.1

Please refer to [Contributing to Pulumi](https://github.com/pulumi/pulumi/blob/master/CONTRIBUTING.md) for installation
guidance.

### Building locally

Run the following commands to install Go modules, generate all SDKs, and build the provider: 

```bash
make ensure
make build
```

Add the `bin` folder to your `$PATH` or copy the `bin/pulumi-resource-google-native` file to another location in your `$PATH`.

### Running an example

Navigate to one of the `examples` and run Pulumi:

```bash
cd ./exampes/simple
yarn link @pulumi/google-native
pulumi up
```
