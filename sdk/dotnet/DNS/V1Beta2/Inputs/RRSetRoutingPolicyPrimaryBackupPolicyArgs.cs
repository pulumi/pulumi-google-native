// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V1Beta2.Inputs
{

    /// <summary>
    /// Configures a RRSetRoutingPolicy such that all queries are responded with the primary_targets if they are healthy. And if all of them are unhealthy, then we fallback to a geo localized policy.
    /// </summary>
    public sealed class RRSetRoutingPolicyPrimaryBackupPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Backup targets provide a regional failover policy for the otherwise global primary targets. If serving state is set to BACKUP, this policy essentially becomes a geo routing policy.
        /// </summary>
        [Input("backupGeoTargets")]
        public Input<Inputs.RRSetRoutingPolicyGeoPolicyArgs>? BackupGeoTargets { get; set; }

        [Input("kind")]
        public Input<string>? Kind { get; set; }

        [Input("primaryTargets")]
        public Input<Inputs.RRSetRoutingPolicyHealthCheckTargetsArgs>? PrimaryTargets { get; set; }

        /// <summary>
        /// When serving state is PRIMARY, this field provides the option of sending a small percentage of the traffic to the backup targets.
        /// </summary>
        [Input("trickleTraffic")]
        public Input<double>? TrickleTraffic { get; set; }

        public RRSetRoutingPolicyPrimaryBackupPolicyArgs()
        {
        }
        public static new RRSetRoutingPolicyPrimaryBackupPolicyArgs Empty => new RRSetRoutingPolicyPrimaryBackupPolicyArgs();
    }
}