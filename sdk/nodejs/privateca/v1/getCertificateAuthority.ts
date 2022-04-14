// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Returns a CertificateAuthority.
 */
export function getCertificateAuthority(args: GetCertificateAuthorityArgs, opts?: pulumi.InvokeOptions): Promise<GetCertificateAuthorityResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:privateca/v1:getCertificateAuthority", {
        "caPoolId": args.caPoolId,
        "certificateAuthorityId": args.certificateAuthorityId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetCertificateAuthorityArgs {
    caPoolId: string;
    certificateAuthorityId: string;
    location: string;
    project?: string;
}

export interface GetCertificateAuthorityResult {
    /**
     * URLs for accessing content published by this CA, such as the CA certificate and CRLs.
     */
    readonly accessUrls: outputs.privateca.v1.AccessUrlsResponse;
    /**
     * A structured description of this CertificateAuthority's CA certificate and its issuers. Ordered as self-to-root.
     */
    readonly caCertificateDescriptions: outputs.privateca.v1.CertificateDescriptionResponse[];
    /**
     * Immutable. The config used to create a self-signed X.509 certificate or CSR.
     */
    readonly config: outputs.privateca.v1.CertificateConfigResponse;
    /**
     * The time at which this CertificateAuthority was created.
     */
    readonly createTime: string;
    /**
     * The time at which this CertificateAuthority was soft deleted, if it is in the DELETED state.
     */
    readonly deleteTime: string;
    /**
     * The time at which this CertificateAuthority will be permanently purged, if it is in the DELETED state.
     */
    readonly expireTime: string;
    /**
     * Immutable. The name of a Cloud Storage bucket where this CertificateAuthority will publish content, such as the CA certificate and CRLs. This must be a bucket name, without any prefixes (such as `gs://`) or suffixes (such as `.googleapis.com`). For example, to use a bucket named `my-bucket`, you would simply specify `my-bucket`. If not specified, a managed bucket will be created.
     */
    readonly gcsBucket: string;
    /**
     * Immutable. Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority is a self-signed CertificateAuthority, this key is also used to sign the self-signed CA certificate. Otherwise, it is used to sign a CSR.
     */
    readonly keySpec: outputs.privateca.v1.KeyVersionSpecResponse;
    /**
     * Optional. Labels with user-defined metadata.
     */
    readonly labels: {[key: string]: string};
    /**
     * Immutable. The desired lifetime of the CA certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate.
     */
    readonly lifetime: string;
    /**
     * The resource name for this CertificateAuthority in the format `projects/*&#47;locations/*&#47;caPools/*&#47;certificateAuthorities/*`.
     */
    readonly name: string;
    /**
     * This CertificateAuthority's certificate chain, including the current CertificateAuthority's certificate. Ordered such that the root issuer is the final element (consistent with RFC 5246). For a self-signed CA, this will only list the current CertificateAuthority's certificate.
     */
    readonly pemCaCertificates: string[];
    /**
     * The State for this CertificateAuthority.
     */
    readonly state: string;
    /**
     * Optional. If this is a subordinate CertificateAuthority, this field will be set with the subordinate configuration, which describes its issuers. This may be updated, but this CertificateAuthority must continue to validate.
     */
    readonly subordinateConfig: outputs.privateca.v1.SubordinateConfigResponse;
    /**
     * The CaPool.Tier of the CaPool that includes this CertificateAuthority.
     */
    readonly tier: string;
    /**
     * Immutable. The Type of this CertificateAuthority.
     */
    readonly type: string;
    /**
     * The time at which this CertificateAuthority was last updated.
     */
    readonly updateTime: string;
}

export function getCertificateAuthorityOutput(args: GetCertificateAuthorityOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCertificateAuthorityResult> {
    return pulumi.output(args).apply(a => getCertificateAuthority(a, opts))
}

export interface GetCertificateAuthorityOutputArgs {
    caPoolId: pulumi.Input<string>;
    certificateAuthorityId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}