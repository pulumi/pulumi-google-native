// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V1Beta2.Outputs
{

    /// <summary>
    /// A RRSetRoutingPolicy represents ResourceRecordSet data that is returned dynamically with the response varying based on configured properties such as geolocation or by weighted random selection.
    /// </summary>
    [OutputType]
    public sealed class RRSetRoutingPolicyResponse
    {
        public readonly Outputs.RRSetRoutingPolicyGeoPolicyResponse Geo;
        public readonly Outputs.RRSetRoutingPolicyGeoPolicyResponse GeoPolicy;
        public readonly string Kind;
        public readonly Outputs.RRSetRoutingPolicyPrimaryBackupPolicyResponse PrimaryBackup;
        public readonly Outputs.RRSetRoutingPolicyWrrPolicyResponse Wrr;
        public readonly Outputs.RRSetRoutingPolicyWrrPolicyResponse WrrPolicy;

        [OutputConstructor]
        private RRSetRoutingPolicyResponse(
            Outputs.RRSetRoutingPolicyGeoPolicyResponse geo,

            Outputs.RRSetRoutingPolicyGeoPolicyResponse geoPolicy,

            string kind,

            Outputs.RRSetRoutingPolicyPrimaryBackupPolicyResponse primaryBackup,

            Outputs.RRSetRoutingPolicyWrrPolicyResponse wrr,

            Outputs.RRSetRoutingPolicyWrrPolicyResponse wrrPolicy)
        {
            Geo = geo;
            GeoPolicy = geoPolicy;
            Kind = kind;
            PrimaryBackup = primaryBackup;
            Wrr = wrr;
            WrrPolicy = wrrPolicy;
        }
    }
}
