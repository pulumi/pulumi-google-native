// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    /// <summary>
    /// Cloud Autoscaler policy.
    /// </summary>
    public sealed class AutoscalingPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of seconds that the autoscaler waits before it starts collecting information from a new instance. This prevents the autoscaler from collecting information when the instance is initializing, during which the collected usage would not be reliable. The default time autoscaler waits is 60 seconds. Virtual machine initialization times might vary because of numerous factors. We recommend that you test how long an instance may take to initialize. To do this, create an instance and time the startup process.
        /// </summary>
        [Input("coolDownPeriodSec")]
        public Input<int>? CoolDownPeriodSec { get; set; }

        /// <summary>
        /// Defines the CPU utilization policy that allows the autoscaler to scale based on the average CPU utilization of a managed instance group.
        /// </summary>
        [Input("cpuUtilization")]
        public Input<Inputs.AutoscalingPolicyCpuUtilizationArgs>? CpuUtilization { get; set; }

        [Input("customMetricUtilizations")]
        private InputList<Inputs.AutoscalingPolicyCustomMetricUtilizationArgs>? _customMetricUtilizations;

        /// <summary>
        /// Configuration parameters of autoscaling based on a custom metric.
        /// </summary>
        public InputList<Inputs.AutoscalingPolicyCustomMetricUtilizationArgs> CustomMetricUtilizations
        {
            get => _customMetricUtilizations ?? (_customMetricUtilizations = new InputList<Inputs.AutoscalingPolicyCustomMetricUtilizationArgs>());
            set => _customMetricUtilizations = value;
        }

        /// <summary>
        /// Configuration parameters of autoscaling based on load balancer.
        /// </summary>
        [Input("loadBalancingUtilization")]
        public Input<Inputs.AutoscalingPolicyLoadBalancingUtilizationArgs>? LoadBalancingUtilization { get; set; }

        /// <summary>
        /// The maximum number of instances that the autoscaler can scale out to. This is required when creating or updating an autoscaler. The maximum number of replicas must not be lower than minimal number of replicas.
        /// </summary>
        [Input("maxNumReplicas")]
        public Input<int>? MaxNumReplicas { get; set; }

        /// <summary>
        /// The minimum number of replicas that the autoscaler can scale in to. This cannot be less than 0. If not provided, autoscaler chooses a default value depending on maximum number of instances allowed.
        /// </summary>
        [Input("minNumReplicas")]
        public Input<int>? MinNumReplicas { get; set; }

        /// <summary>
        /// Defines operating mode for this policy.
        /// </summary>
        [Input("mode")]
        public Input<Pulumi.GoogleNative.Compute.V1.AutoscalingPolicyMode>? Mode { get; set; }

        [Input("scaleInControl")]
        public Input<Inputs.AutoscalingPolicyScaleInControlArgs>? ScaleInControl { get; set; }

        [Input("scalingSchedules")]
        private InputMap<string>? _scalingSchedules;

        /// <summary>
        /// Scaling schedules defined for an autoscaler. Multiple schedules can be set on an autoscaler, and they can overlap. During overlapping periods the greatest min_required_replicas of all scaling schedules is applied. Up to 128 scaling schedules are allowed.
        /// </summary>
        public InputMap<string> ScalingSchedules
        {
            get => _scalingSchedules ?? (_scalingSchedules = new InputMap<string>());
            set => _scalingSchedules = value;
        }

        public AutoscalingPolicyArgs()
        {
        }
    }
}