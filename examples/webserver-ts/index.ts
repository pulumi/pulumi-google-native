// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import * as google from "@pulumi/google-native";
import * as random from "@pulumi/random"

const randomString = new random.RandomString("name", {
    upper: false,
    number: false,
    special: false,
    length: 8,
});

const config = new pulumi.Config("google-native");
const project = config.require("project");
const zone = config.require("zone");

const networkName = pulumi.interpolate`${randomString.result}-net`;
const computeNetwork = new google.compute.v1.Network("network", {
    autoCreateSubnetworks: true,
    project: project,
    network: networkName,
    name: networkName,
});

const firewallName = pulumi.interpolate`${randomString.result}-fw`;
const computeFirewall = new google.compute.v1.Firewall("firewall", {
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

const instanceName = pulumi.interpolate`${randomString.result}-instance`;
const computeInstance = new google.compute.v1.Instance("instance", {
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
