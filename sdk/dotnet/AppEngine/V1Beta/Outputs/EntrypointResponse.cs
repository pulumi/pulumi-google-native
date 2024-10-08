// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1Beta.Outputs
{

    /// <summary>
    /// The entrypoint for the application.
    /// </summary>
    [OutputType]
    public sealed class EntrypointResponse
    {
        /// <summary>
        /// The format should be a shell command that can be fed to bash -c.
        /// </summary>
        public readonly string Shell;

        [OutputConstructor]
        private EntrypointResponse(string shell)
        {
            Shell = shell;
        }
    }
}
