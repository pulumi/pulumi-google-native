// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

import * as gcp from "@pulumi/gcp-native";
import * as pulumi from "@pulumi/pulumi";

const config = new pulumi.Config("gcp-native");
const project = config.require("project");
const zone = config.require("zone");

const networkName = "webserver-net";
const computeNetwork = new gcp.compute.v1.Network("network", {
    autoCreateSubnetworks: true,
    project: project,
    network: networkName,
    name: networkName,
});

const firewallName = "webserver-fw";
const computeFirewall = new gcp.compute.v1.Firewall("firewall", {
    network: computeNetwork.selfLink,
    firewall: firewallName,
    name: firewallName,
    project: project,
    allowed: [{
        IPProtocol: "tcp",
        ports: ["22", "80"],
    }],
});

// (optional) create a simple web server using the startup script for the instance
const startupScript = `#!/bin/bash
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &`;

const instanceName = "webserver-instance";
const computeInstance = new gcp.compute.v1.Instance("instance", {
    name: instanceName,
    instance: instanceName,
    project: project,
    zone: zone,
    machineType: `projects/${project}/zones/${zone}/machineTypes/f1-micro`,
    metadata: {
        items: [{
            "key": "startup-script",
            "value": startupScript,
        }],
    },
    disks: [
        {
            boot: true,
            initializeParams: {
                sourceImage: "projects/debian-cloud/global/images/debian-9-stretch-v20181210",
            },
        },
    ],
    networkInterfaces: [{
        network: computeNetwork.id,
        accessConfigs: [{}], // must be empty to request an ephemeral IP
    }],
    serviceAccounts: [{
        scopes: ["https://www.googleapis.com/auth/cloud-platform"],
    }],
}, { dependsOn: [computeFirewall] });

exports.instanceLink = computeInstance.selfLink;
exports.instanceIP = computeInstance.networkInterfaces[0].accessConfigs[0].natIP;
