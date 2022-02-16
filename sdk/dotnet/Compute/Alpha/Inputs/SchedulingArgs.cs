// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    /// <summary>
    /// Sets the scheduling options for an Instance. NextID: 21
    /// </summary>
    public sealed class SchedulingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether the instance should be automatically restarted if it is terminated by Compute Engine (not terminated by a user). You can only set the automatic restart option for standard instances. Preemptible instances cannot be automatically restarted. By default, this is set to true so an instance is automatically restarted if it is terminated by Compute Engine.
        /// </summary>
        [Input("automaticRestart")]
        public Input<bool>? AutomaticRestart { get; set; }

        /// <summary>
        /// Specifies the availability domain (AD), which this instance should be scheduled on. The AD belongs to the spread GroupPlacementPolicy resource policy that has been assigned to the instance. Specify a value between 1-max count of availability domains in your GroupPlacementPolicy. See go/placement-policy-extension for more details.
        /// </summary>
        [Input("availabilityDomain")]
        public Input<int>? AvailabilityDomain { get; set; }

        /// <summary>
        /// Current number of vCPUs available for VM. 0 or unset means default vCPUs of the current machine type.
        /// </summary>
        [Input("currentCpus")]
        public Input<int>? CurrentCpus { get; set; }

        /// <summary>
        /// Current amount of memory (in MB) available for VM. 0 or unset means default amount of memory of the current machine type.
        /// </summary>
        [Input("currentMemoryMb")]
        public Input<string>? CurrentMemoryMb { get; set; }

        /// <summary>
        /// Specify the time in seconds for host error detection, the value must be within the range of [90, 330] with the increment of 30, if unset, the default behavior of host error recovery will be used.
        /// </summary>
        [Input("hostErrorTimeoutSeconds")]
        public Input<int>? HostErrorTimeoutSeconds { get; set; }

        /// <summary>
        /// Specifies the termination action for the instance.
        /// </summary>
        [Input("instanceTerminationAction")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.SchedulingInstanceTerminationAction>? InstanceTerminationAction { get; set; }

        /// <summary>
        /// Defines whether the instance is tolerant of higher cpu latency. This can only be set during instance creation, or when the instance is not currently running. It must not be set if the preemptible option is also set.
        /// </summary>
        [Input("latencyTolerant")]
        public Input<bool>? LatencyTolerant { get; set; }

        /// <summary>
        /// An opaque location hint used to place the instance close to other resources. This field is for use by internal tools that use the public API.
        /// </summary>
        [Input("locationHint")]
        public Input<string>? LocationHint { get; set; }

        /// <summary>
        /// Specifies the number of hours after VM instance creation where the VM won't be scheduled for maintenance.
        /// </summary>
        [Input("maintenanceFreezeDurationHours")]
        public Input<int>? MaintenanceFreezeDurationHours { get; set; }

        /// <summary>
        /// For more information about maintenance intervals, see Setting maintenance intervals.
        /// </summary>
        [Input("maintenanceInterval")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.SchedulingMaintenanceInterval>? MaintenanceInterval { get; set; }

        /// <summary>
        /// Specifies the max run duration for the given instance. If specified, the instance termination action will be performed at the end of the run duration.
        /// </summary>
        [Input("maxRunDuration")]
        public Input<Inputs.DurationArgs>? MaxRunDuration { get; set; }

        /// <summary>
        /// The minimum number of virtual CPUs this instance will consume when running on a sole-tenant node.
        /// </summary>
        [Input("minNodeCpus")]
        public Input<int>? MinNodeCpus { get; set; }

        [Input("nodeAffinities")]
        private InputList<Inputs.SchedulingNodeAffinityArgs>? _nodeAffinities;

        /// <summary>
        /// A set of node affinity and anti-affinity configurations. Refer to Configuring node affinity for more information. Overrides reservationAffinity.
        /// </summary>
        public InputList<Inputs.SchedulingNodeAffinityArgs> NodeAffinities
        {
            get => _nodeAffinities ?? (_nodeAffinities = new InputList<Inputs.SchedulingNodeAffinityArgs>());
            set => _nodeAffinities = value;
        }

        /// <summary>
        /// Defines the maintenance behavior for this instance. For standard instances, the default behavior is MIGRATE. For preemptible instances, the default and only possible behavior is TERMINATE. For more information, see Set VM availability policies.
        /// </summary>
        [Input("onHostMaintenance")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.SchedulingOnHostMaintenance>? OnHostMaintenance { get; set; }

        /// <summary>
        /// Defines whether the instance is preemptible. This can only be set during instance creation or while the instance is stopped and therefore, in a `TERMINATED` state. See Instance Life Cycle for more information on the possible instance states.
        /// </summary>
        [Input("preemptible")]
        public Input<bool>? Preemptible { get; set; }

        /// <summary>
        /// Specifies the provisioning model of the instance.
        /// </summary>
        [Input("provisioningModel")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.SchedulingProvisioningModel>? ProvisioningModel { get; set; }

        /// <summary>
        /// Specifies the timestamp, when the instance will be terminated, in RFC3339 text format. If specified, the instance termination action will be performed at the termination time.
        /// </summary>
        [Input("terminationTime")]
        public Input<string>? TerminationTime { get; set; }

        public SchedulingArgs()
        {
        }
    }
}