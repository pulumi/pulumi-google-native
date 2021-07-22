// Copyright 2016-2021, Pulumi Corporation.

import * as google from "@pulumi/google-native";
import * as random from "@pulumi/random"

const randomString = new random.RandomString("name", {
    upper: false,
    number: false,
    special: false,
    length: 8,
});

const topic = new google.pubsub.v1.Topic("topic", {
    topicId: randomString.result,
    name: randomString.result,
});

const sub = new google.pubsub.v1.Subscription("sub", {
    topic: topic.name,
    subscriptionId: randomString.result,
    name: randomString.result,
});

const schema = new google.pubsub.v1.Schema("schema", {
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
    type: "AVRO",
});

// This is to test that an invoke works.
export const topicReadName = topic.name.apply(async name => {
    const parts = name.split('/');
    const topicId = parts[parts.length-1];
    const resp = await google.pubsub.v1.getTopic({ topicId });
    return resp.name;
});
