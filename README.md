# Native Google Cloud Pulumi Provider (preview)

The native Google Cloud Provider for Pulumi lets you provision Google Cloud resources in your cloud programs.

This provider uses the Google Cloud REST API directly and therefore provides full access to the Google Cloud Platform (GCP).

The provider is currently in public preview and is not recommended for production deployments yet. Breaking changes will be introduced in minor version releases.

To use this package, please [install the Pulumi CLI first](https://pulumi.io/).

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/google-native

or `yarn`:

    $ yarn add @pulumi/google-native

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_gcp_native

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-google-native/sdk

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.GcpNative

## Concepts

The native Google Cloud package provides a strongly-typed means to build cloud applications that create
and interact closely with Google Cloud resources.  Resources are exposed for the entire GCP surface area,
including (but not limited to) 'compute', 'container', 'run', 'storage', and more.

The native GCP provider works directly with the Google Cloud API instead of depending on a
handwritten layer as with the [classic provider](https://github.com/pulumi/pulumi-gcp). This approach ensures higher
quality and higher fidelity with GCP.

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
