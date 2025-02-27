// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Aiplatform.V1
{
    public static class GetDeploymentResourcePool
    {
        /// <summary>
        /// Get a DeploymentResourcePool.
        /// </summary>
        public static Task<GetDeploymentResourcePoolResult> InvokeAsync(GetDeploymentResourcePoolArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDeploymentResourcePoolResult>("google-native:aiplatform/v1:getDeploymentResourcePool", args ?? new GetDeploymentResourcePoolArgs(), options.WithDefaults());

        /// <summary>
        /// Get a DeploymentResourcePool.
        /// </summary>
        public static Output<GetDeploymentResourcePoolResult> Invoke(GetDeploymentResourcePoolInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDeploymentResourcePoolResult>("google-native:aiplatform/v1:getDeploymentResourcePool", args ?? new GetDeploymentResourcePoolInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDeploymentResourcePoolArgs : global::Pulumi.InvokeArgs
    {
        [Input("deploymentResourcePoolId", required: true)]
        public string DeploymentResourcePoolId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetDeploymentResourcePoolArgs()
        {
        }
        public static new GetDeploymentResourcePoolArgs Empty => new GetDeploymentResourcePoolArgs();
    }

    public sealed class GetDeploymentResourcePoolInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("deploymentResourcePoolId", required: true)]
        public Input<string> DeploymentResourcePoolId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetDeploymentResourcePoolInvokeArgs()
        {
        }
        public static new GetDeploymentResourcePoolInvokeArgs Empty => new GetDeploymentResourcePoolInvokeArgs();
    }


    [OutputType]
    public sealed class GetDeploymentResourcePoolResult
    {
        /// <summary>
        /// Timestamp when this DeploymentResourcePool was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The underlying DedicatedResources that the DeploymentResourcePool uses.
        /// </summary>
        public readonly Outputs.GoogleCloudAiplatformV1DedicatedResourcesResponse DedicatedResources;
        /// <summary>
        /// Immutable. The resource name of the DeploymentResourcePool. Format: `projects/{project}/locations/{location}/deploymentResourcePools/{deployment_resource_pool}`
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetDeploymentResourcePoolResult(
            string createTime,

            Outputs.GoogleCloudAiplatformV1DedicatedResourcesResponse dedicatedResources,

            string name)
        {
            CreateTime = createTime;
            DedicatedResources = dedicatedResources;
            Name = name;
        }
    }
}
