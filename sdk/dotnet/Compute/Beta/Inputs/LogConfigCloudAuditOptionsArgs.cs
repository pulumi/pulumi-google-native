// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    /// <summary>
    /// This is deprecated and has no effect. Do not use.
    /// </summary>
    public sealed class LogConfigCloudAuditOptionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// This is deprecated and has no effect. Do not use.
        /// </summary>
        [Input("authorizationLoggingOptions")]
        public Input<Inputs.AuthorizationLoggingOptionsArgs>? AuthorizationLoggingOptions { get; set; }

        /// <summary>
        /// This is deprecated and has no effect. Do not use.
        /// </summary>
        [Input("logName")]
        public Input<Pulumi.GoogleNative.Compute.Beta.LogConfigCloudAuditOptionsLogName>? LogName { get; set; }

        public LogConfigCloudAuditOptionsArgs()
        {
        }
    }
}