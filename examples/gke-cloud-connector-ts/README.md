## GKE Config Connector Example

This example is based on the [upstream installation docs](https://cloud.google.com/config-connector/docs/how-to/install-upgrade-uninstall)

Currently, the following manual steps are required:
1. The Kubernetes provider doesn't natively support patching unmanaged resources, so the update has to be run in two
parts. First, the existing CustomResource is imported into Pulumi state, and then is modified in a subsequent update.
2. IAM Policy isn't deleted when the stack is destroyed, and requires manual cleanup.
