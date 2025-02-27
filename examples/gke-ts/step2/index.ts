/*
 * Copyright 2016-2022, Pulumi Corporation.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import * as google from "@pulumi/google-native";
import * as pulumi from "@pulumi/pulumi";

const config = new pulumi.Config("google-native");
const region = config.require("region");

const serverConfig = google.container.v1.getServerConfigOutput({
    location: region,
});

const engineVersion = serverConfig.apply(conf => conf.validMasterVersions[0]);

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
    // Change
    resourceLabels: {"test": "true"},
});

const nodepool = new google.container.v1.NodePool("nodepool", {
    clusterId: cluster.name,
    initialNodeCount: 1,
    management: {
        autoRepair: false,
        autoUpgrade: false,
    },
    // Changes
    config: {
        tags: ["nodepool"],
        taints: [{key: "important", effect: "PREFER_NO_SCHEDULE", value: "true"}],
    }
});

export const nodepoolTag = nodepool.config.tags[0];
export const taintsKey = nodepool.config.taints[0].key;
