import * as pulumi from "@pulumi/pulumi";
import * as cloud from "@pulumi/google-cloud";

const project = "pulumi-development";
const region = "us-central1";

const runName = "run-native";
const service = new cloud.run.v1.Service("service", {
    projectsId: project,
    locationsId: region,
    servicesId: runName,
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

const iamHello = new cloud.run.v1.ServiceIamPolicy("allow-all", {
    projectsId: project,
    locationsId: region,
    servicesId: service.metadata.name,
    policy: {
        bindings: [{
            members: ["allUsers"],
            role: "roles/run.invoker",
        }],
    },
});

const bucketName = "bucket-nextgen";
const bucket = new cloud.storage.v1.Bucket("bucket", {
    project: project,
    bucket: bucketName,
    name: bucketName,
});

const archiveName = "zip2";
const bucketObject = new cloud.storage.v1.BucketObject(archiveName, {
    object: archiveName,
    name: archiveName,
    bucket: bucket.name,
    source: new pulumi.asset.AssetArchive({
        ".": new pulumi.asset.FileArchive("./pythonfunc"),
    }),
});

const functionName = "python-func";
const functionPython = new cloud.cloudfunctions.v1.Function(functionName, {
    projectsId: project,
    locationsId: region,
    functionsId: functionName,
    name: `projects/${project}/locations/${region}/functions/${functionName}`,
    sourceArchiveUrl: pulumi.interpolate`gs://${bucket.name}/${bucketObject.name}`,
    httpsTrigger: {},
    entryPoint: "handler",
    timeout: "60s",
    availableMemoryMb: 128,
    runtime: "python37",
    ingressSettings: "ALLOW_ALL",
});

const pyInvoker = new cloud.cloudfunctions.v1.FunctionIamPolicy("py-invoker", {
    projectsId: project,
    locationsId: region,
    functionsId: functionName,
    policy: {
        bindings: [
            {
                members: ["allUsers"],
                role: "roles/cloudfunctions.invoker",
            }
        ],
    },
}, { dependsOn: [functionPython] });

export const functionUrl = functionPython.httpsTrigger.url;

const clusterName = "gke-native";
const cluster = new cloud.container.v1.Cluster("cluster", {
    projectsId: project,
    locationsId: region,
    clustersId: clusterName,
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

const nodePoolName = "extra-node-pool";
const pool = new cloud.container.v1.ClusterNodePool(nodePoolName, {
    projectsId: project,
    locationsId: region,
    clustersId: cluster.name,
    nodePoolsId: nodePoolName,
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
        name: nodePoolName,
        version: "1.18.16-gke.500",
    },
});
