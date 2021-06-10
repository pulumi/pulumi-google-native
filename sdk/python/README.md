[![Slack](http://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Fgoogle-native.svg)](https://npmjs.com/package/@pulumi/google-native)
[![Python version](https://badge.fury.io/py/pulumi-google-native.svg)](https://pypi.org/project/pulumi-google-native)
[![NuGet version](https://badge.fury.io/nu/pulumi.googlenative.svg)](https://badge.fury.io/nu/pulumi.googlenative)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-google-native/sdk/go)](https://pkg.go.dev/github.com/pulumi/pulumi-google-native/sdk/go)
[![License](https://img.shields.io/npm/l/%40pulumi%2Fgoogle-native.svg)](https://github.com/pulumi/pulumi-google-native/blob/master/LICENSE)

# Native Google Cloud Pulumi Provider (preview)

The native Google Cloud Provider for Pulumi lets you provision Google Cloud resources in your cloud programs.

This provider uses the Google Cloud REST API directly and therefore provides full access to Google Cloud.

The provider is currently in public preview and is not recommended for production deployments yet. Breaking changes will be introduced in minor version releases.

To use this package, please [install the Pulumi CLI first](https://pulumi.io/).

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/google-native

or `yarn`:

    $ yarn add @pulumi/google-native

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_google_native

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-google-native/sdk

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.GoogleNative

## Concepts

The native Google Cloud package provides a strongly-typed means to build cloud applications that create
and interact closely with Google Cloud resources.  Resources are exposed for the entire Google Cloud surface area,
including (but not limited to) 'compute', 'container', 'run', 'storage', and more.

The native Google Cloud provider works directly with the Google Cloud API instead of depending on a
handwritten layer as with the [classic provider](https://github.com/pulumi/pulumi-gcp). This approach ensures higher
quality and higher fidelity with Google Cloud.

## Configuring credentials

Credentials configuration is compatible with the classic GCP provider.

Please refer to [this quickstart guide](
https://www.pulumi.com/docs/intro/cloud-providers/gcp/setup/) for possible configuration options.

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

```
$ make ensure
$ make build
```

Add the `bin` folder to your `$PATH` or copy the `bin/pulumi-resource-google-native` file to another location in your `$PATH`.

### Running an example

Navigate to one of the `examples` and run Pulumi:

```
$ cd ./exampes/simple
$ yarn link @pulumi/google-native
$ pulumi up
``` 
