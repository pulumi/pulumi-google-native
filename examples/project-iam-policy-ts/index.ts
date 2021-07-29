// Copyright 2016-2021, Pulumi Corporation.

import * as pulumi from "@pulumi/pulumi";
import * as google from "@pulumi/google-native";

const config = new pulumi.Config("google-native");
const project = config.require("project");

const serviceAccount = new google.iam.v1.ServiceAccount("serviceAccount", {
    accountId: "gke-cloud-connector-1"
});

const existingPolicy = google.cloudresourcemanager.v1.getProjectIamPolicy({resource: project});
existingPolicy.then(p => {
    const bindings = [{
        members: pulumi.all([pulumi.interpolate `serviceAccount:${serviceAccount.email}`]),
        role: "roles/owner",
    }];
    p.bindings.map(b => bindings.push({members: pulumi.all(b.members), role: b.role}));
    // Update the bindings to include the service account
    new google.cloudresourcemanager.v1.ProjectIamPolicy("owner", {
        bindings: bindings,
        resource: project,
    });
});
