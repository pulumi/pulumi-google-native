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
    /// Customize service error responses. For example, list any service specific protobuf types that can appear in error detail lists of error responses. Example: custom_error: types: - google.foo.v1.CustomError - google.foo.v1.AnotherError
    /// </summary>
    public sealed class CustomErrorArgs : Pulumi.ResourceArgs
    {
        [Input("rules")]
        private InputList<Inputs.CustomErrorRuleArgs>? _rules;

        /// <summary>
        /// The list of custom error rules that apply to individual API messages. **NOTE:** All service configuration rules follow "last one wins" order.
        /// </summary>
        public InputList<Inputs.CustomErrorRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.CustomErrorRuleArgs>());
            set => _rules = value;
        }

        [Input("types")]
        private InputList<string>? _types;

        /// <summary>
        /// The list of custom error detail types, e.g. 'google.foo.v1.CustomError'.
        /// </summary>
        public InputList<string> Types
        {
            get => _types ?? (_types = new InputList<string>());
            set => _types = value;
        }

        public CustomErrorArgs()
        {
        }
    }
}