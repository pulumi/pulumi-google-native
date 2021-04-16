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

const topic = new google.pubsub.v1.Topic("topic", {
    projectsId: project,
    topicsId: randomString.result,
});

const sub = new google.pubsub.v1.Subscription("sub", {
    projectsId: project,
    topic: topic.name,
    subscriptionsId: randomString.result,
});
