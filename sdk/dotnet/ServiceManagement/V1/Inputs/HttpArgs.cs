// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ServiceManagement.V1.Inputs
{

    /// <summary>
    /// Defines the HTTP configuration for an API service. It contains a list of HttpRule, each specifying the mapping of an RPC method to one or more HTTP REST API methods.
    /// </summary>
    public sealed class HttpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When set to true, URL path parameters will be fully URI-decoded except in cases of single segment matches in reserved expansion, where "%2F" will be left encoded. The default behavior is to not decode RFC 6570 reserved characters in multi segment matches.
        /// </summary>
        [Input("fullyDecodeReservedExpansion")]
        public Input<bool>? FullyDecodeReservedExpansion { get; set; }

        [Input("rules")]
        private InputList<Inputs.HttpRuleArgs>? _rules;

        /// <summary>
        /// A list of HTTP configuration rules that apply to individual API methods. **NOTE:** All service configuration rules follow "last one wins" order.
        /// </summary>
        public InputList<Inputs.HttpRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.HttpRuleArgs>());
            set => _rules = value;
        }

        public HttpArgs()
        {
        }
        public static new HttpArgs Empty => new HttpArgs();
    }
}
