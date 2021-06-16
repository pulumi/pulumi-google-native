// Copyright 2016-2021, Pulumi Corporation.

import * as gcp from "@pulumi/gcp";
import * as gnative from "@pulumi/google-native";
import * as storage from "@pulumi/google-native/storage/v1";
import * as pulumi from "@pulumi/pulumi";
import * as random from "@pulumi/random";
import * as fs from "fs";
import * as mime from "mime";

const randomString = new random.RandomString("name", {
    upper: false,
    number: false,
    special: false,
    length: 5,
});

const config = new pulumi.Config("google-native");
const project = config.require("project");
const bucketName = pulumi.interpolate`pulumi-transfer-test-dst-${randomString.result}`;
const transferJobName = "transfer-test";

// Create a Google Cloud resource (Storage Bucket)
const bucket = new storage.Bucket("my-bucket", {
    name: bucketName,
    project: project,
});

// TODO: Add support for this in Native provider: https://github.com/pulumi/pulumi-google-native/issues/119
const storageTransferServiceAccount = gcp.storage.getTransferProjectServieAccount(
    {
        project: project,
    }
)

const iam = new gcp.storage.BucketIAMMember("iam-dest", {
    bucket: bucket.name,
    role: "roles/storage.admin",
    member: pulumi.interpolate`serviceAccount:${storageTransferServiceAccount.then(sa => sa.email)}`,
});

const srcBucket = new gnative.storage.v1.Bucket("src-bucket", {
    name: pulumi.interpolate`pulumi-transfer-test-src-${randomString.result}`,
    project: project
});

const iamSrc = new gcp.storage.BucketIAMMember("iam-src", {
    bucket: srcBucket.name,
    role: "roles/storage.admin",
    member: pulumi.interpolate`serviceAccount:${storageTransferServiceAccount.then(sa => sa.email)}`,
});

const contents = [] as pulumi.Resource[];
crawlDirectory("bucket-content", (filepath: string) => {
    const contentFile = new gnative.storage.v1.BucketObject(filepath, {
        name: filepath.replace("/", "_"),
        bucket: srcBucket.name,
        contentType: mime.getType(filepath) || undefined,
        source: new pulumi.asset.FileAsset(filepath),
    }, {
        parent: srcBucket,
    });
});

new gnative.storagetransfer.v1.TransferJob(
    transferJobName,
    {
        name: pulumi.interpolate`transferJobs/transferjob-${randomString.result}`,
        description: "s3-to-gcs-sync",
        status: "ENABLED",
        schedule: {
            repeatInterval: "3600s", // every hour
            scheduleStartDate: {
                day: 1,
                month: 5,
                year: 2021,
            },
        },
        project: project,
        transferSpec: {
            gcsDataSource: {
                bucketName: srcBucket.name,
            },
            gcsDataSink: {
                bucketName: bucket.name,
            },
            objectConditions: {
                includePrefixes: ["/bucket-content"],
            },
            transferOptions: {
                deleteObjectsFromSourceAfterTransfer: false,
            },
        },
    },
    {dependsOn: contents.concat([iam, iamSrc])}
);


// crawlDirectory recursive crawls the provided directory, applying the provided function
// to every file it contains. Doesn't handle cycles from symlinks.
function crawlDirectory(dir: string, f: (_: string) => void) {
    const files = fs.readdirSync(dir);
    for (const file of files) {
        const filePath = `${dir}/${file}`;
        const stat = fs.statSync(filePath);
        if (stat.isDirectory()) {
            crawlDirectory(filePath, f);
        }
        if (stat.isFile()) {
            f(filePath);
        }
    }
}

// Export the bucket self-link
export const bucketSelfLink = bucket.selfLink;

