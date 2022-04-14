// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1
{
    public static class GetTargetInstance
    {
        /// <summary>
        /// Returns the specified TargetInstance resource. Gets a list of available target instances by making a list() request.
        /// </summary>
        public static Task<GetTargetInstanceResult> InvokeAsync(GetTargetInstanceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTargetInstanceResult>("google-native:compute/v1:getTargetInstance", args ?? new GetTargetInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified TargetInstance resource. Gets a list of available target instances by making a list() request.
        /// </summary>
        public static Output<GetTargetInstanceResult> Invoke(GetTargetInstanceInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTargetInstanceResult>("google-native:compute/v1:getTargetInstance", args ?? new GetTargetInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTargetInstanceArgs : Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        [Input("targetInstance", required: true)]
        public string TargetInstance { get; set; } = null!;

        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetTargetInstanceArgs()
        {
        }
    }

    public sealed class GetTargetInstanceInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("targetInstance", required: true)]
        public Input<string> TargetInstance { get; set; } = null!;

        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetTargetInstanceInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTargetInstanceResult
    {
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// A URL to the virtual machine instance that handles traffic for this target instance. When creating a target instance, you can provide the fully-qualified URL or a valid partial URL to the desired virtual machine. For example, the following are all valid URLs: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance - zones/zone/instances/instance 
        /// </summary>
        public readonly string Instance;
        /// <summary>
        /// The type of the resource. Always compute#targetInstance for target instances.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// NAT option controlling how IPs are NAT'ed to the instance. Currently only NO_NAT (default value) is supported.
        /// </summary>
        public readonly string NatPolicy;
        /// <summary>
        /// The URL of the network this target instance uses to forward traffic. If not specified, the traffic will be forwarded to the network that the default network interface belongs to.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// URL of the zone where the target instance resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetTargetInstanceResult(
            string creationTimestamp,

            string description,

            string instance,

            string kind,

            string name,

            string natPolicy,

            string network,

            string selfLink,

            string zone)
        {
            CreationTimestamp = creationTimestamp;
            Description = description;
            Instance = instance;
            Kind = kind;
            Name = name;
            NatPolicy = natPolicy;
            Network = network;
            SelfLink = selfLink;
            Zone = zone;
        }
    }
}