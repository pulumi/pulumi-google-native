// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataflow.V1b3.Inputs
{

    /// <summary>
    /// Metadata for a File connector used by the job.
    /// </summary>
    public sealed class FileIODetailsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// File Pattern used to access files by the connector.
        /// </summary>
        [Input("filePattern")]
        public Input<string>? FilePattern { get; set; }

        public FileIODetailsArgs()
        {
        }
    }
}