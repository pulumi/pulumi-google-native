// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContactCenterAIPlatform.V1Alpha1.Inputs
{

    /// <summary>
    /// Message storing the instance configuration.
    /// </summary>
    public sealed class InstanceConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The instance size of this the instance configuration.
        /// </summary>
        [Input("instanceSize")]
        public Input<Pulumi.GoogleNative.ContactCenterAIPlatform.V1Alpha1.InstanceConfigInstanceSize>? InstanceSize { get; set; }

        public InstanceConfigArgs()
        {
        }
        public static new InstanceConfigArgs Empty => new InstanceConfigArgs();
    }
}
