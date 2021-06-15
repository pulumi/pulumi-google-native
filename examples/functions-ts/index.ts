// Copyright 2016-2021, Pulumi Corporation.

import * as pulumi from "@pulumi/pulumi";
import * as google from "@pulumi/google-native";
import * as random from "@pulumi/random"

const config = new pulumi.Config("google-native");
const project = config.require("project");
const region = config.require("region");

const randomString = new random.RandomString("name", {
    upper: false,
    number: false,
    special: false,
    length: 8,
});

const bucket = new google.storage.v1.Bucket("bucket", {
    project,
    name: pulumi.interpolate`bucket-${randomString.result}`,
});

const bucketObject = new google.storage.v1.BucketObject("bucketObject", {
    name: "zip",
    bucket: bucket.name,
    source: new pulumi.asset.AssetArchive({
        ".": new pulumi.asset.FileArchive("./pythonfunc"),
    }),
});

const functionName = pulumi.interpolate`func-${randomString.result}`;
const func = new google.cloudfunctions.v1.Function("function-py", {
    project,
    location: region,
    name: pulumi.interpolate`projects/${project}/locations/${region}/functions/${functionName}`,
    sourceArchiveUrl: pulumi.interpolate`gs://${bucket.name}/${bucketObject.name}`,
    httpsTrigger: {
        securityLevel: google.cloudfunctions.v1.HttpsTriggerSecurityLevel.SecureAlways,
    },
    entryPoint: "handler",
    timeout: "60s",
    availableMemoryMb: 128,
    runtime: "python37",
    ingressSettings: "ALLOW_ALL",
});

const invoker = new google.cloudfunctions.v1.FunctionIamPolicy("function-py-iam", {
    project,
    location: region,
    functionId: functionName, // func.name returns the long `projects/foo/locations/bat/functions/buzz` name which doesn't suit here
    bindings: [
        {
            members: ["allUsers"],
            role: "roles/cloudfunctions.invoker",
        }
    ],
}, { dependsOn: func});

export const functionUrl = func.httpsTrigger.url;
