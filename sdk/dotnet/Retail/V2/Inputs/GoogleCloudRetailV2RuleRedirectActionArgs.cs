// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Retail.V2.Inputs
{

    /// <summary>
    /// Redirects a shopper to a specific page. * Rule Condition: - Must specify Condition.query_terms. * Action Input: Request Query * Action Result: Redirects shopper to provided uri.
    /// </summary>
    public sealed class GoogleCloudRetailV2RuleRedirectActionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// URL must have length equal or less than 2000 characters.
        /// </summary>
        [Input("redirectUri")]
        public Input<string>? RedirectUri { get; set; }

        public GoogleCloudRetailV2RuleRedirectActionArgs()
        {
        }
        public static new GoogleCloudRetailV2RuleRedirectActionArgs Empty => new GoogleCloudRetailV2RuleRedirectActionArgs();
    }
}