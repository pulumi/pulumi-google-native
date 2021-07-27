// Copyright 2016-2021, Pulumi Corporation.

import * as google from "@pulumi/google-native";

// TODO: Determine this dynamically once https://github.com/pulumi/pulumi-google-native/issues/166 is done.
const engineVersion = "1.19.9-gke.1900";

const nodeConfig: google.types.input.container.v1.NodeConfigArgs = {
    machineType: "n1-standard-2",
    oauthScopes: [
        "https://www.googleapis.com/auth/compute",
        "https://www.googleapis.com/auth/devstorage.read_only",
        "https://www.googleapis.com/auth/logging.write",
        "https://www.googleapis.com/auth/monitoring"
    ],
    preemptible: true,
};

const cluster = new google.container.v1.Cluster("cluster", {
    initialClusterVersion: engineVersion,
    nodePools: [{
        config: nodeConfig,
        initialNodeCount: 1,
        management: {
            autoRepair: false,
        },
        name: "initial",
    }],
});
