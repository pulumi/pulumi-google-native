// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha
{
    public static class GetInstanceGroupManager
    {
        /// <summary>
        /// Returns all of the details about the specified managed instance group. Gets a list of available managed instance groups by making a list() request.
        /// </summary>
        public static Task<GetInstanceGroupManagerResult> InvokeAsync(GetInstanceGroupManagerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceGroupManagerResult>("google-native:compute/alpha:getInstanceGroupManager", args ?? new GetInstanceGroupManagerArgs(), options.WithDefaults());

        /// <summary>
        /// Returns all of the details about the specified managed instance group. Gets a list of available managed instance groups by making a list() request.
        /// </summary>
        public static Output<GetInstanceGroupManagerResult> Invoke(GetInstanceGroupManagerInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstanceGroupManagerResult>("google-native:compute/alpha:getInstanceGroupManager", args ?? new GetInstanceGroupManagerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceGroupManagerArgs : Pulumi.InvokeArgs
    {
        [Input("instanceGroupManager", required: true)]
        public string InstanceGroupManager { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetInstanceGroupManagerArgs()
        {
        }
    }

    public sealed class GetInstanceGroupManagerInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("instanceGroupManager", required: true)]
        public Input<string> InstanceGroupManager { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetInstanceGroupManagerInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstanceGroupManagerResult
    {
        /// <summary>
        /// Specifies the instances configs overrides that should be applied for all instances in the MIG.
        /// </summary>
        public readonly Outputs.InstanceGroupManagerAllInstancesConfigResponse AllInstancesConfig;
        /// <summary>
        /// The autohealing policy for this managed instance group. You can specify only one value.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstanceGroupManagerAutoHealingPolicyResponse> AutoHealingPolicies;
        /// <summary>
        /// The base instance name to use for instances in this group. The value must be 1-58 characters long. Instances are named by appending a hyphen and a random four-character string to the base instance name. The base instance name must comply with RFC1035.
        /// </summary>
        public readonly string BaseInstanceName;
        /// <summary>
        /// The creation timestamp for this managed instance group in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// The list of instance actions and the number of instances in this managed instance group that are scheduled for each of those actions.
        /// </summary>
        public readonly Outputs.InstanceGroupManagerActionsSummaryResponse CurrentActions;
        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Policy specifying the intended distribution of managed instances across zones in a regional managed instance group.
        /// </summary>
        public readonly Outputs.DistributionPolicyResponse DistributionPolicy;
        /// <summary>
        /// The action to perform in case of zone failure. Only one value is supported, NO_FAILOVER. The default is NO_FAILOVER.
        /// </summary>
        public readonly string FailoverAction;
        /// <summary>
        /// Fingerprint of this resource. This field may be used in optimistic locking. It will be ignored when inserting an InstanceGroupManager. An up-to-date fingerprint must be provided in order to update the InstanceGroupManager, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an InstanceGroupManager.
        /// </summary>
        public readonly string Fingerprint;
        /// <summary>
        /// The URL of the Instance Group resource.
        /// </summary>
        public readonly string InstanceGroup;
        /// <summary>
        /// Instance lifecycle policy for this Instance Group Manager.
        /// </summary>
        public readonly Outputs.InstanceGroupManagerInstanceLifecyclePolicyResponse InstanceLifecyclePolicy;
        /// <summary>
        /// The URL of the instance template that is specified for this managed instance group. The group uses this template to create all new instances in the managed instance group. The templates for existing instances in the group do not change unless you run recreateInstances, run applyUpdatesToInstances, or set the group's updatePolicy.type to PROACTIVE.
        /// </summary>
        public readonly string InstanceTemplate;
        /// <summary>
        /// The resource type, which is always compute#instanceGroupManager for managed instance groups.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Pagination behavior of listManagedInstances API method for this Managed Instance Group.
        /// </summary>
        public readonly string ListManagedInstancesResults;
        /// <summary>
        /// The name of the managed instance group. The name must be 1-63 characters long, and comply with RFC1035.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Named ports configured for the Instance Groups complementary to this Instance Group Manager.
        /// </summary>
        public readonly ImmutableArray<Outputs.NamedPortResponse> NamedPorts;
        /// <summary>
        /// The URL of the region where the managed instance group resides (for regional resources).
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The URL for this managed instance group. The server defines this URL.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Server-defined URL for this resource with the resource id.
        /// </summary>
        public readonly string SelfLinkWithId;
        /// <summary>
        /// The service account to be used as credentials for all operations performed by the managed instance group on instances. The service accounts needs all permissions required to create and delete instances. By default, the service account {projectNumber}@cloudservices.gserviceaccount.com is used.
        /// </summary>
        public readonly string ServiceAccount;
        /// <summary>
        /// Standby policy for stopped and suspended instances.
        /// </summary>
        public readonly Outputs.InstanceGroupManagerStandbyPolicyResponse StandbyPolicy;
        /// <summary>
        /// Stateful configuration for this Instanced Group Manager
        /// </summary>
        public readonly Outputs.StatefulPolicyResponse StatefulPolicy;
        /// <summary>
        /// The status of this managed instance group.
        /// </summary>
        public readonly Outputs.InstanceGroupManagerStatusResponse Status;
        /// <summary>
        /// The URLs for all TargetPool resources to which instances in the instanceGroup field are added. The target pools automatically apply to all of the instances in the managed instance group.
        /// </summary>
        public readonly ImmutableArray<string> TargetPools;
        /// <summary>
        /// The target number of running instances for this managed instance group. You can reduce this number by using the instanceGroupManager deleteInstances or abandonInstances methods. Resizing the group also changes this number.
        /// </summary>
        public readonly int TargetSize;
        /// <summary>
        /// The target number of stopped instances for this managed instance group. This number changes when you: - Stop instance using the stopInstances method or start instances using the startInstances method. - Manually change the targetStoppedSize using the update method. 
        /// </summary>
        public readonly int TargetStoppedSize;
        /// <summary>
        /// The target number of suspended instances for this managed instance group. This number changes when you: - Suspend instance using the suspendInstances method or resume instances using the resumeInstances method. - Manually change the targetSuspendedSize using the update method. 
        /// </summary>
        public readonly int TargetSuspendedSize;
        /// <summary>
        /// The update policy for this managed instance group.
        /// </summary>
        public readonly Outputs.InstanceGroupManagerUpdatePolicyResponse UpdatePolicy;
        /// <summary>
        /// Specifies the instance templates used by this managed instance group to create instances. Each version is defined by an instanceTemplate and a name. Every version can appear at most once per instance group. This field overrides the top-level instanceTemplate field. Read more about the relationships between these fields. Exactly one version must leave the targetSize field unset. That version will be applied to all remaining instances. For more information, read about canary updates.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstanceGroupManagerVersionResponse> Versions;
        /// <summary>
        /// The URL of a zone where the managed instance group is located (for zonal resources).
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetInstanceGroupManagerResult(
            Outputs.InstanceGroupManagerAllInstancesConfigResponse allInstancesConfig,

            ImmutableArray<Outputs.InstanceGroupManagerAutoHealingPolicyResponse> autoHealingPolicies,

            string baseInstanceName,

            string creationTimestamp,

            Outputs.InstanceGroupManagerActionsSummaryResponse currentActions,

            string description,

            Outputs.DistributionPolicyResponse distributionPolicy,

            string failoverAction,

            string fingerprint,

            string instanceGroup,

            Outputs.InstanceGroupManagerInstanceLifecyclePolicyResponse instanceLifecyclePolicy,

            string instanceTemplate,

            string kind,

            string listManagedInstancesResults,

            string name,

            ImmutableArray<Outputs.NamedPortResponse> namedPorts,

            string region,

            string selfLink,

            string selfLinkWithId,

            string serviceAccount,

            Outputs.InstanceGroupManagerStandbyPolicyResponse standbyPolicy,

            Outputs.StatefulPolicyResponse statefulPolicy,

            Outputs.InstanceGroupManagerStatusResponse status,

            ImmutableArray<string> targetPools,

            int targetSize,

            int targetStoppedSize,

            int targetSuspendedSize,

            Outputs.InstanceGroupManagerUpdatePolicyResponse updatePolicy,

            ImmutableArray<Outputs.InstanceGroupManagerVersionResponse> versions,

            string zone)
        {
            AllInstancesConfig = allInstancesConfig;
            AutoHealingPolicies = autoHealingPolicies;
            BaseInstanceName = baseInstanceName;
            CreationTimestamp = creationTimestamp;
            CurrentActions = currentActions;
            Description = description;
            DistributionPolicy = distributionPolicy;
            FailoverAction = failoverAction;
            Fingerprint = fingerprint;
            InstanceGroup = instanceGroup;
            InstanceLifecyclePolicy = instanceLifecyclePolicy;
            InstanceTemplate = instanceTemplate;
            Kind = kind;
            ListManagedInstancesResults = listManagedInstancesResults;
            Name = name;
            NamedPorts = namedPorts;
            Region = region;
            SelfLink = selfLink;
            SelfLinkWithId = selfLinkWithId;
            ServiceAccount = serviceAccount;
            StandbyPolicy = standbyPolicy;
            StatefulPolicy = statefulPolicy;
            Status = status;
            TargetPools = targetPools;
            TargetSize = targetSize;
            TargetStoppedSize = targetStoppedSize;
            TargetSuspendedSize = targetSuspendedSize;
            UpdatePolicy = updatePolicy;
            Versions = versions;
            Zone = zone;
        }
    }
}