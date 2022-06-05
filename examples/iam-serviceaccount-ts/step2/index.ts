import * as iam from "@pulumi/google-native/iam/v1";

const serviceAccount = new iam.ServiceAccount("service-account", {
    accountId: "testaccount1",
    displayName: "testaccount-updated", // Update display name
});

export const serviceAccountName = serviceAccount.name;
