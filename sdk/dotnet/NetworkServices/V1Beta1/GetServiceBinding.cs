// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkServices.V1Beta1
{
    public static class GetServiceBinding
    {
        /// <summary>
        /// Gets details of a single ServiceBinding.
        /// </summary>
        public static Task<GetServiceBindingResult> InvokeAsync(GetServiceBindingArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServiceBindingResult>("google-native:networkservices/v1beta1:getServiceBinding", args ?? new GetServiceBindingArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single ServiceBinding.
        /// </summary>
        public static Output<GetServiceBindingResult> Invoke(GetServiceBindingInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetServiceBindingResult>("google-native:networkservices/v1beta1:getServiceBinding", args ?? new GetServiceBindingInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServiceBindingArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("serviceBindingId", required: true)]
        public string ServiceBindingId { get; set; } = null!;

        public GetServiceBindingArgs()
        {
        }
    }

    public sealed class GetServiceBindingInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("serviceBindingId", required: true)]
        public Input<string> ServiceBindingId { get; set; } = null!;

        public GetServiceBindingInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServiceBindingResult
    {
        /// <summary>
        /// The timestamp when the resource was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. A free-text description of the resource. Max length 1024 characters.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Optional. Set of label tags associated with the ServiceBinding resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Name of the ServiceBinding resource. It matches pattern `projects/*/locations/global/serviceBindings/service_binding_name&gt;`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The full service directory service name of the format /projects/*/locations/*/namespaces/*/services/*
        /// </summary>
        public readonly string Service;
        /// <summary>
        /// The timestamp when the resource was updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetServiceBindingResult(
            string createTime,

            string description,

            ImmutableDictionary<string, string> labels,

            string name,

            string service,

            string updateTime)
        {
            CreateTime = createTime;
            Description = description;
            Labels = labels;
            Name = name;
            Service = service;
            UpdateTime = updateTime;
        }
    }
}