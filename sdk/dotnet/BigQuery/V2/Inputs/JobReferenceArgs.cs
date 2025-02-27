// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigQuery.V2.Inputs
{

    public sealed class JobReferenceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// [Required] The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
        /// </summary>
        [Input("jobId")]
        public Input<string>? JobId { get; set; }

        /// <summary>
        /// The geographic location of the job. See details at https://cloud.google.com/bigquery/docs/locations#specifying_your_location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// [Required] The ID of the project containing this job.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public JobReferenceArgs()
        {
        }
        public static new JobReferenceArgs Empty => new JobReferenceArgs();
    }
}
