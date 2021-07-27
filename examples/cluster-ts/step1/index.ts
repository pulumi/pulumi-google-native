// Copyright 2016-2021, Pulumi Corporation.

import * as google from "@pulumi/google-native";

const zone = "us-central1-a";

new google.container.v1.Cluster("cluster", {
    location: zone,
    initialClusterVersion: "1.19.9-gke.1900",
    nodePools: [{initialNodeCount: 1, name: "initial"}],
});
