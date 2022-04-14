// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./appProfile";
export * from "./backup";
export * from "./cluster";
export * from "./getAppProfile";
export * from "./getBackup";
export * from "./getCluster";
export * from "./getInstance";
export * from "./getInstanceClusterBackupIamPolicy";
export * from "./getInstanceIamPolicy";
export * from "./getInstanceTableIamPolicy";
export * from "./getTable";
export * from "./instance";
export * from "./instanceClusterBackupIamPolicy";
export * from "./instanceIamPolicy";
export * from "./instanceTableIamPolicy";
export * from "./table";

// Export enums:
export * from "../../types/enums/bigtableadmin/v2";

// Import resources to register:
import { AppProfile } from "./appProfile";
import { Backup } from "./backup";
import { Cluster } from "./cluster";
import { Instance } from "./instance";
import { InstanceClusterBackupIamPolicy } from "./instanceClusterBackupIamPolicy";
import { InstanceIamPolicy } from "./instanceIamPolicy";
import { InstanceTableIamPolicy } from "./instanceTableIamPolicy";
import { Table } from "./table";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:bigtableadmin/v2:AppProfile":
                return new AppProfile(name, <any>undefined, { urn })
            case "google-native:bigtableadmin/v2:Backup":
                return new Backup(name, <any>undefined, { urn })
            case "google-native:bigtableadmin/v2:Cluster":
                return new Cluster(name, <any>undefined, { urn })
            case "google-native:bigtableadmin/v2:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "google-native:bigtableadmin/v2:InstanceClusterBackupIamPolicy":
                return new InstanceClusterBackupIamPolicy(name, <any>undefined, { urn })
            case "google-native:bigtableadmin/v2:InstanceIamPolicy":
                return new InstanceIamPolicy(name, <any>undefined, { urn })
            case "google-native:bigtableadmin/v2:InstanceTableIamPolicy":
                return new InstanceTableIamPolicy(name, <any>undefined, { urn })
            case "google-native:bigtableadmin/v2:Table":
                return new Table(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "bigtableadmin/v2", _module)