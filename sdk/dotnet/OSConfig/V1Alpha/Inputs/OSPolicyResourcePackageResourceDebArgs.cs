// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Alpha.Inputs
{

    /// <summary>
    /// A deb package file. dpkg packages only support INSTALLED state.
    /// </summary>
    public sealed class OSPolicyResourcePackageResourceDebArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether dependencies should also be installed. - install when false: `dpkg -i package` - install when true: `apt-get update &amp;&amp; apt-get -y install package.deb`
        /// </summary>
        [Input("pullDeps")]
        public Input<bool>? PullDeps { get; set; }

        /// <summary>
        /// A deb package.
        /// </summary>
        [Input("source", required: true)]
        public Input<Inputs.OSPolicyResourceFileArgs> Source { get; set; } = null!;

        public OSPolicyResourcePackageResourceDebArgs()
        {
        }
        public static new OSPolicyResourcePackageResourceDebArgs Empty => new OSPolicyResourcePackageResourceDebArgs();
    }
}
