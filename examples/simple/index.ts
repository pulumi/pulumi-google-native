import { asset } from "@pulumi/pulumi";
import * as cloud from "@pulumi/google-cloud";

const project = "pulumi-development";
const region = "us-central1";

const runName = "run-native";
const service = new cloud.run.v1.Service("service", {
    parent: `projects/${project}/locations/${region}`,
    apiVersion: "serving.knative.dev/v1",
    kind: "Service",
    metadata: {
        name: runName,
    },
    spec: {
        template: {
            spec: {
                containers: [{image: "gcr.io/cloudrun/hello"}],
            },
        },
    },
});

const iamHello = new cloud.run.v1.Policy("allow-all", {
    resource: `projects/${project}/locations/${region}/services/${runName}`,
    policy: {
        bindings: [
            {
                members: ["allUsers"],
                role: "roles/run.invoker",
            }
        ],
    },
}, { dependsOn: [service] });

const bucketName = "bucket-nextgen";
const bucket = new cloud.storage.v1.Bucket("bucket", {
    project,
    name: bucketName,
});

const archiveName = "zip2";
const bucketObject = new cloud.storage.v1.BucketObject(archiveName, {
    name: archiveName,
    bucket: bucketName,
    source: new asset.AssetArchive({
        ".": new asset.FileArchive("./pythonfunc"),
    }),
}, { dependsOn: [bucket] });

const functionName = "python-func";
const functionPython = new cloud.cloudfunctions.v1.CloudFunction(functionName, {
    location: `projects/${project}/locations/${region}`,
    name: `projects/${project}/locations/${region}/functions/${functionName}`,
    sourceArchiveUrl: `gs://${bucketName}/${archiveName}`,
    httpsTrigger: {},
    entryPoint: "handler",
    timeout: "60s",
    availableMemoryMb: 128,
    runtime: "python37",
    ingressSettings: "ALLOW_ALL",
    versionId: "1",
}, { dependsOn: [bucketObject] });

const pyInvoker = new cloud.cloudfunctions.v1.Policy("py-invoker", {
    resource: `projects/${project}/locations/${region}/functions/${functionName}`,
    policy: {
        bindings: [
            {
                members: ["allUsers"],
                role: "roles/cloudfunctions.invoker",
            }
        ],
    },
}, { dependsOn: [functionPython] });

export const functionUrl = `https://${region}-${project}.cloudfunctions.net/${functionName}`;

const clusterName = "gke-native";
const cluster = new cloud.container.v1.Cluster("cluster", {
    parent: `projects/${project}/locations/${region}`,
    cluster: {
        initialClusterVersion: "1.18.16-gke.500",
        initialNodeCount: 1,
        masterAuth: {
            password: "hDiqST+U7{t+BkQA+OD*",
            username: "admin",
        },
        name: clusterName,
        network: `projects/${project}/global/networks/default`,
        nodeConfig: {
            oauthScopes: [
                "https://www.googleapis.com/auth/devstorage.read_only",
                "https://www.googleapis.com/auth/logging.write",
                "https://www.googleapis.com/auth/monitoring",
                "https://www.googleapis.com/auth/service.management.readonly",
                "https://www.googleapis.com/auth/servicecontrol",
                "https://www.googleapis.com/auth/trace.append",
            ],            
        },
    }
});

const pool = new cloud.container.v1.NodePool("nodepool", {
    parent: `projects/${project}/locations/${region}/clusters/${clusterName}`,
    nodePool: {
        config: {
            machineType: "n1-standard-1",
            oauthScopes: [
                "https://www.googleapis.com/auth/monitoring",
                "https://www.googleapis.com/auth/devstorage.read_only",
                "https://www.googleapis.com/auth/logging.write",
                "https://www.googleapis.com/auth/compute",
            ],
            preemptible: true,
        },
        initialNodeCount: 3,
        management: {
            autoRepair: true,
        },
        name: "primary-node-pool",
        version: "1.18.16-gke.500",
    },
}, {dependsOn: [cluster]});
