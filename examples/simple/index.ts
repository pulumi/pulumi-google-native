// Copyright 2016-2021, Pulumi Corporation.

import * as pulumi from "@pulumi/pulumi";
import * as google from "@pulumi/google-native";

const project = "pulumi-development";
const region = "us-central1";

const clusterName = "gke-native";
const cluster = new google.container.v1.Cluster("cluster", {
    project,
    location: region,
    parent: `projects/${project}/locations/${region}`,
    initialClusterVersion: "1.18.17-gke.1900",
    masterAuth: {
        password: "hDiqST+U7{t+BkQA+OD*",
        username: "admin",
    },
    name: clusterName,
    network: `projects/${project}/global/networks/default`,
});

const nodePoolName = "extra-node-pool";
const pool = new google.container.v1.NodePool(nodePoolName, {
    project,
    location: region,
    clusterId: cluster.name,
    parent: `projects/${project}/locations/${region}/clusters/${clusterName}`,
    config: {
        machineType: "n1-standard-1",
        oauthScopes: [
            "https://www.googleapis.com/auth/monitoring",
            "https://www.googleapis.com/auth/devstorage.read_only",
            "https://www.googleapis.com/auth/logging.write",
            "https://www.googleapis.com/auth/compute",
        ],
        preemptible: true,
    },
    initialNodeCount: 1,
    management: {
        autoRepair: true,
    },
    name: nodePoolName,
    version: "1.18.16-gke.500",
});
