import * as iam from "@pulumi/google-native/iam/v1";

const serviceAccount = new iam.ServiceAccount("service-account", {
    accountId: "testaccount1",
    displayName: "testaccount",
});

export const serviceAccountName = serviceAccount.name;
