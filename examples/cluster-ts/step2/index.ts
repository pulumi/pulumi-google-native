// Copyright 2016-2021, Pulumi Corporation.

import * as google from "@pulumi/google-native";

const project = "pulumi-development";
const zone = "us-central1-a";

new google.container.v1.Cluster("cluster", {
    location: zone,
    initialClusterVersion: "1.19.9-gke.1900",
    network: `projects/${project}/global/networks/default`,
    // ---> CHANGE START
    // Enable workload identity.
    workloadIdentityConfig: {
        workloadPool: `${project}.svc.id.goog`,
    },
    // Enable release channel
    releaseChannel: {channel: "RAPID"},
    // Enable binary authorization.
    binaryAuthorization: {enabled: true},
    // <--- CHANGE END
    nodePools: [{initialNodeCount: 1, name: "initial"}],
});
