import * as pulumi from "@pulumi/pulumi";

export class Service extends pulumi.CustomResource {
    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super('google-cloud:run/v1:Service', name, args, opts);
    }
}

export class IamPolicy extends pulumi.CustomResource {
    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super('google-cloud:run/v1:IamPolicy', name, args, opts);
    }
}
