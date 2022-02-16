// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./provider";

// Export sub-modules:
import * as accesscontextmanager from "./accesscontextmanager";
import * as apigateway from "./apigateway";
import * as apigee from "./apigee";
import * as apikeys from "./apikeys";
import * as appengine from "./appengine";
import * as artifactregistry from "./artifactregistry";
import * as assuredworkloads from "./assuredworkloads";
import * as baremetalsolution from "./baremetalsolution";
import * as bigquery from "./bigquery";
import * as bigqueryconnection from "./bigqueryconnection";
import * as bigquerydatatransfer from "./bigquerydatatransfer";
import * as bigqueryreservation from "./bigqueryreservation";
import * as bigtableadmin from "./bigtableadmin";
import * as billingbudgets from "./billingbudgets";
import * as binaryauthorization from "./binaryauthorization";
import * as certificatemanager from "./certificatemanager";
import * as cloudasset from "./cloudasset";
import * as cloudbilling from "./cloudbilling";
import * as cloudbuild from "./cloudbuild";
import * as cloudchannel from "./cloudchannel";
import * as clouddeploy from "./clouddeploy";
import * as cloudfunctions from "./cloudfunctions";
import * as cloudidentity from "./cloudidentity";
import * as cloudiot from "./cloudiot";
import * as cloudkms from "./cloudkms";
import * as cloudresourcemanager from "./cloudresourcemanager";
import * as cloudscheduler from "./cloudscheduler";
import * as cloudsearch from "./cloudsearch";
import * as cloudsupport from "./cloudsupport";
import * as cloudtasks from "./cloudtasks";
import * as cloudtrace from "./cloudtrace";
import * as composer from "./composer";
import * as compute from "./compute";
import * as config from "./config";
import * as connectors from "./connectors";
import * as contactcenterinsights from "./contactcenterinsights";
import * as container from "./container";
import * as containeranalysis from "./containeranalysis";
import * as datacatalog from "./datacatalog";
import * as dataflow from "./dataflow";
import * as datafusion from "./datafusion";
import * as datalabeling from "./datalabeling";
import * as datamigration from "./datamigration";
import * as datapipelines from "./datapipelines";
import * as dataplex from "./dataplex";
import * as dataproc from "./dataproc";
import * as datastore from "./datastore";
import * as datastream from "./datastream";
import * as deploymentmanager from "./deploymentmanager";
import * as dialogflow from "./dialogflow";
import * as dlp from "./dlp";
import * as dns from "./dns";
import * as documentai from "./documentai";
import * as domains from "./domains";
import * as essentialcontacts from "./essentialcontacts";
import * as eventarc from "./eventarc";
import * as file from "./file";
import * as firebase from "./firebase";
import * as firebaseappcheck from "./firebaseappcheck";
import * as firebasedatabase from "./firebasedatabase";
import * as firebasehosting from "./firebasehosting";
import * as firebaseml from "./firebaseml";
import * as firebaserules from "./firebaserules";
import * as firestore from "./firestore";
import * as gameservices from "./gameservices";
import * as genomics from "./genomics";
import * as gkehub from "./gkehub";
import * as healthcare from "./healthcare";
import * as iam from "./iam";
import * as iap from "./iap";
import * as ids from "./ids";
import * as jobs from "./jobs";
import * as logging from "./logging";
import * as managedidentities from "./managedidentities";
import * as memcache from "./memcache";
import * as metastore from "./metastore";
import * as ml from "./ml";
import * as monitoring from "./monitoring";
import * as networkconnectivity from "./networkconnectivity";
import * as networkmanagement from "./networkmanagement";
import * as networksecurity from "./networksecurity";
import * as networkservices from "./networkservices";
import * as notebooks from "./notebooks";
import * as orgpolicy from "./orgpolicy";
import * as osconfig from "./osconfig";
import * as oslogin from "./oslogin";
import * as policysimulator from "./policysimulator";
import * as privateca from "./privateca";
import * as pubsub from "./pubsub";
import * as pubsublite from "./pubsublite";
import * as recaptchaenterprise from "./recaptchaenterprise";
import * as recommendationengine from "./recommendationengine";
import * as redis from "./redis";
import * as remotebuildexecution from "./remotebuildexecution";
import * as retail from "./retail";
import * as run from "./run";
import * as runtimeconfig from "./runtimeconfig";
import * as secretmanager from "./secretmanager";
import * as securitycenter from "./securitycenter";
import * as servicedirectory from "./servicedirectory";
import * as servicemanagement from "./servicemanagement";
import * as sourcerepo from "./sourcerepo";
import * as spanner from "./spanner";
import * as speech from "./speech";
import * as sqladmin from "./sqladmin";
import * as storage from "./storage";
import * as storagetransfer from "./storagetransfer";
import * as testing from "./testing";
import * as toolresults from "./toolresults";
import * as tpu from "./tpu";
import * as transcoder from "./transcoder";
import * as translate from "./translate";
import * as types from "./types";
import * as vision from "./vision";
import * as vmmigration from "./vmmigration";
import * as vpcaccess from "./vpcaccess";
import * as websecurityscanner from "./websecurityscanner";
import * as workflowexecutions from "./workflowexecutions";
import * as workflows from "./workflows";

export {
    accesscontextmanager,
    apigateway,
    apigee,
    apikeys,
    appengine,
    artifactregistry,
    assuredworkloads,
    baremetalsolution,
    bigquery,
    bigqueryconnection,
    bigquerydatatransfer,
    bigqueryreservation,
    bigtableadmin,
    billingbudgets,
    binaryauthorization,
    certificatemanager,
    cloudasset,
    cloudbilling,
    cloudbuild,
    cloudchannel,
    clouddeploy,
    cloudfunctions,
    cloudidentity,
    cloudiot,
    cloudkms,
    cloudresourcemanager,
    cloudscheduler,
    cloudsearch,
    cloudsupport,
    cloudtasks,
    cloudtrace,
    composer,
    compute,
    config,
    connectors,
    contactcenterinsights,
    container,
    containeranalysis,
    datacatalog,
    dataflow,
    datafusion,
    datalabeling,
    datamigration,
    datapipelines,
    dataplex,
    dataproc,
    datastore,
    datastream,
    deploymentmanager,
    dialogflow,
    dlp,
    dns,
    documentai,
    domains,
    essentialcontacts,
    eventarc,
    file,
    firebase,
    firebaseappcheck,
    firebasedatabase,
    firebasehosting,
    firebaseml,
    firebaserules,
    firestore,
    gameservices,
    genomics,
    gkehub,
    healthcare,
    iam,
    iap,
    ids,
    jobs,
    logging,
    managedidentities,
    memcache,
    metastore,
    ml,
    monitoring,
    networkconnectivity,
    networkmanagement,
    networksecurity,
    networkservices,
    notebooks,
    orgpolicy,
    osconfig,
    oslogin,
    policysimulator,
    privateca,
    pubsub,
    pubsublite,
    recaptchaenterprise,
    recommendationengine,
    redis,
    remotebuildexecution,
    retail,
    run,
    runtimeconfig,
    secretmanager,
    securitycenter,
    servicedirectory,
    servicemanagement,
    sourcerepo,
    spanner,
    speech,
    sqladmin,
    storage,
    storagetransfer,
    testing,
    toolresults,
    tpu,
    transcoder,
    translate,
    types,
    vision,
    vmmigration,
    vpcaccess,
    websecurityscanner,
    workflowexecutions,
    workflows,
};

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("google-native", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:google-native") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});