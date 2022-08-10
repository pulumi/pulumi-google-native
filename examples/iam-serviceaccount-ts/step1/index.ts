import * as pulumi from "@pulumi/pulumi";
import * as iam from "@pulumi/google-native/iam/v1";
import * as random from "@pulumi/random";

const randomString = new random.RandomString("random", {
    length: 8,
    special: false,
});

const accountId = pulumi.interpolate `testaccount-${randomString.result}`;

const serviceAccount = new iam.ServiceAccount("service-account", {
    accountId: accountId,
    displayName: "testaccount",
});

export const serviceAccountName = serviceAccount.name;
