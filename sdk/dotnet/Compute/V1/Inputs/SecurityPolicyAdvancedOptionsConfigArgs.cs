// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    public sealed class SecurityPolicyAdvancedOptionsConfigArgs : Pulumi.ResourceArgs
    {
        [Input("jsonParsing")]
        public Input<Pulumi.GoogleNative.Compute.V1.SecurityPolicyAdvancedOptionsConfigJsonParsing>? JsonParsing { get; set; }

        [Input("logLevel")]
        public Input<Pulumi.GoogleNative.Compute.V1.SecurityPolicyAdvancedOptionsConfigLogLevel>? LogLevel { get; set; }

        public SecurityPolicyAdvancedOptionsConfigArgs()
        {
        }
    }
}