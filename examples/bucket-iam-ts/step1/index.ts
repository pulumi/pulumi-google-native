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

// Create a Google Cloud resource (Storage Bucket)
const bucket = new google.storage.v1.Bucket("bucket-member-test", {
    iamConfiguration: {
        uniformBucketLevelAccess: {
            enabled: true // Required for IAM conditions
        }
    }
});

new google.storage.v1.BucketIamMember("test", {
    member: "serviceAccount:native-provider-test@pulumi-ci-gcp-provider.iam.gserviceaccount.com",
    name: bucket.name,
    role: "roles/storage.objectAdmin",
    condition: {
        title: "condition-1",
        description: "Expiring someday",
        expression: `request.time < timestamp("2030-01-01T00:00:00Z")`
    },
});

// Create a Google Cloud resource (Storage Bucket)
const bucket2 = new google.storage.v1.Bucket("bucket-binding-test", {
    iamConfiguration: {
        uniformBucketLevelAccess: {
            enabled: true // Required for IAM conditions
        }
    }
});

new google.storage.v1.BucketIamBinding("test", {
    members: ["serviceAccount:native-provider-test@pulumi-ci-gcp-provider.iam.gserviceaccount.com"],
    name: bucket2.name,
    role: "roles/storage.objectAdmin",
    condition: {
        title: "condition-1",
        description: "Expiring someday",
        expression: `request.time < timestamp("2030-01-01T00:00:00Z")`
    },
});
