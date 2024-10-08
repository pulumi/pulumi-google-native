// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    public sealed class InstanceGroupManagerAutoHealingPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Restricts what triggers autohealing.
        /// </summary>
        [Input("autoHealingTriggers")]
        public Input<Inputs.InstanceGroupManagerAutoHealingPolicyAutoHealingTriggersArgs>? AutoHealingTriggers { get; set; }

        /// <summary>
        /// The URL for the health check that signals autohealing.
        /// </summary>
        [Input("healthCheck")]
        public Input<string>? HealthCheck { get; set; }

        /// <summary>
        /// The initial delay is the number of seconds that a new VM takes to initialize and run its startup script. During a VM's initial delay period, the MIG ignores unsuccessful health checks because the VM might be in the startup process. This prevents the MIG from prematurely recreating a VM. If the health check receives a healthy response during the initial delay, it indicates that the startup process is complete and the VM is ready. The value of initial delay must be between 0 and 3600 seconds. The default value is 0.
        /// </summary>
        [Input("initialDelaySec")]
        public Input<int>? InitialDelaySec { get; set; }

        /// <summary>
        /// Maximum number of instances that can be unavailable when autohealing. When 'percent' is used, the value is rounded if necessary. The instance is considered available if all of the following conditions are satisfied: 1. Instance's status is RUNNING. 2. Instance's currentAction is NONE (in particular its liveness health check result was observed to be HEALTHY at least once as it passed VERIFYING). 3. There is no outgoing action on an instance triggered by IGM. By default, number of concurrently autohealed instances is smaller than the managed instance group target size. However, if a zonal managed instance group has only one instance, or a regional managed instance group has only one instance per zone, autohealing will recreate these instances when they become unhealthy.
        /// </summary>
        [Input("maxUnavailable")]
        public Input<Inputs.FixedOrPercentArgs>? MaxUnavailable { get; set; }

        public InstanceGroupManagerAutoHealingPolicyArgs()
        {
        }
        public static new InstanceGroupManagerAutoHealingPolicyArgs Empty => new InstanceGroupManagerAutoHealingPolicyArgs();
    }
}
