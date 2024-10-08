// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudDeploy.V1.Outputs
{

    /// <summary>
    /// Deployment job composition.
    /// </summary>
    [OutputType]
    public sealed class DeploymentJobsResponse
    {
        /// <summary>
        /// The deploy Job. This is the deploy job in the phase.
        /// </summary>
        public readonly Outputs.JobResponse DeployJob;
        /// <summary>
        /// The postdeploy Job, which is the last job on the phase.
        /// </summary>
        public readonly Outputs.JobResponse PostdeployJob;
        /// <summary>
        /// The predeploy Job, which is the first job on the phase.
        /// </summary>
        public readonly Outputs.JobResponse PredeployJob;
        /// <summary>
        /// The verify Job. Runs after a deploy if the deploy succeeds.
        /// </summary>
        public readonly Outputs.JobResponse VerifyJob;

        [OutputConstructor]
        private DeploymentJobsResponse(
            Outputs.JobResponse deployJob,

            Outputs.JobResponse postdeployJob,

            Outputs.JobResponse predeployJob,

            Outputs.JobResponse verifyJob)
        {
            DeployJob = deployJob;
            PostdeployJob = postdeployJob;
            PredeployJob = predeployJob;
            VerifyJob = verifyJob;
        }
    }
}
