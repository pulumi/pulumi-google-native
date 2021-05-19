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
    project,
    topicId: randomString.result,
});

const sub = new google.pubsub.v1.Subscription("sub", {
    project,
    topic: topic.name,
    subscriptionId: randomString.result,
});

const schema = new google.pubsub.v1.Schema("schema", {
    project,
    schemaId: randomString.result,
    definition: JSON.stringify({
        type: "record",
        name: "State",
        namespace: "utilities",
        doc: "A list of states in the United States of America.",
        fields: [
            {
                name: "name",
                type: "string",
                doc: "The common name of the state",
            },
            {
                name: "post_abbr",
                type: "string",
                doc: "The postal code abbreviation of the state.",
            },
        ],
    }),
    name: "my-schema",
    type: "avro",
});
