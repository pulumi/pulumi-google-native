// Copyright 2016-2021, Pulumi Corporation.

import * as pulumi from "@pulumi/pulumi";
import * as google from "@pulumi/google-native";

const config = new pulumi.Config("google-native");
const region = config.require("region");

const bucket = new google.storage.v1.Bucket("bucket");

const bucketObject = new google.storage.v1.BucketObject("bucketObject", {
    name: "zip",
    bucket: bucket.name,
    source: new pulumi.asset.AssetArchive({
        ".": new pulumi.asset.FileArchive("./pythonfunc"),
    }),
});

const func = new google.cloudfunctions.v1.Function("function-py", {
    location: region,
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
    location: region,
    functionId: func.name.apply(name => name.split("/")[name.split("/").length-1]),
    bindings: [
        {
            members: ["allUsers"],
            role: "roles/cloudfunctions.invoker",
        }
    ],
}, { dependsOn: func});

export const functionUrl = func.httpsTrigger.url;
