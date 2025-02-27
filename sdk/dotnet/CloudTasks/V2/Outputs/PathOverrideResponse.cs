// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudTasks.V2.Outputs
{

    /// <summary>
    /// PathOverride. Path message defines path override for HTTP targets.
    /// </summary>
    [OutputType]
    public sealed class PathOverrideResponse
    {
        /// <summary>
        /// The URI path (e.g., /users/1234). Default is an empty string.
        /// </summary>
        public readonly string Path;

        [OutputConstructor]
        private PathOverrideResponse(string path)
        {
            Path = path;
        }
    }
}
