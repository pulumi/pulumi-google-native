// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1.Inputs
{

    /// <summary>
    /// Security configurations to manage scoring.
    /// </summary>
    public sealed class GoogleCloudApigeeV1SecurityProfileScoringConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the config.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Path of the component config used for scoring.
        /// </summary>
        [Input("scorePath")]
        public Input<string>? ScorePath { get; set; }

        /// <summary>
        /// Title of the config.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        public GoogleCloudApigeeV1SecurityProfileScoringConfigArgs()
        {
        }
        public static new GoogleCloudApigeeV1SecurityProfileScoringConfigArgs Empty => new GoogleCloudApigeeV1SecurityProfileScoringConfigArgs();
    }
}
