// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Notebooks.V1.Inputs
{

    /// <summary>
    /// Definition of a custom Compute Engine virtual machine image for starting a notebook instance with the environment installed directly on the VM.
    /// </summary>
    public sealed class VmImageArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Use this VM image family to find the image; the newest image in this family will be used.
        /// </summary>
        [Input("imageFamily")]
        public Input<string>? ImageFamily { get; set; }

        /// <summary>
        /// Use VM image name to find the image.
        /// </summary>
        [Input("imageName")]
        public Input<string>? ImageName { get; set; }

        /// <summary>
        /// The name of the Google Cloud project that this VM image belongs to. Format: `projects/{project_id}`
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public VmImageArgs()
        {
        }
    }
}