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
    /// matchRule criteria for request header matches.
    /// </summary>
    public sealed class HttpHeaderMatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The value should exactly match contents of exactMatch. Only one of exactMatch, prefixMatch, suffixMatch, regexMatch, presentMatch or rangeMatch must be set.
        /// </summary>
        [Input("exactMatch")]
        public Input<string>? ExactMatch { get; set; }

        /// <summary>
        /// The name of the HTTP header to match. For matching against the HTTP request's authority, use a headerMatch with the header name ":authority". For matching a request's method, use the headerName ":method". When the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true, only non-binary user-specified custom metadata and the `content-type` header are supported. The following transport-level headers cannot be used in header matching rules: `:authority`, `:method`, `:path`, `:scheme`, `user-agent`, `accept-encoding`, `content-encoding`, `grpc-accept-encoding`, `grpc-encoding`, `grpc-previous-rpc-attempts`, `grpc-tags-bin`, `grpc-timeout` and `grpc-trace-bin`.
        /// </summary>
        [Input("headerName")]
        public Input<string>? HeaderName { get; set; }

        /// <summary>
        /// If set to false, the headerMatch is considered a match if the preceding match criteria are met. If set to true, the headerMatch is considered a match if the preceding match criteria are NOT met. The default setting is false. 
        /// </summary>
        [Input("invertMatch")]
        public Input<bool>? InvertMatch { get; set; }

        /// <summary>
        /// The value of the header must start with the contents of prefixMatch. Only one of exactMatch, prefixMatch, suffixMatch, regexMatch, presentMatch or rangeMatch must be set.
        /// </summary>
        [Input("prefixMatch")]
        public Input<string>? PrefixMatch { get; set; }

        /// <summary>
        /// A header with the contents of headerName must exist. The match takes place whether or not the request's header has a value. Only one of exactMatch, prefixMatch, suffixMatch, regexMatch, presentMatch or rangeMatch must be set.
        /// </summary>
        [Input("presentMatch")]
        public Input<bool>? PresentMatch { get; set; }

        /// <summary>
        /// The header value must be an integer and its value must be in the range specified in rangeMatch. If the header does not contain an integer, number or is empty, the match fails. For example for a range [-5, 0] - -3 will match. - 0 will not match. - 0.25 will not match. - -3someString will not match. Only one of exactMatch, prefixMatch, suffixMatch, regexMatch, presentMatch or rangeMatch must be set. rangeMatch is not supported for load balancers that have loadBalancingScheme set to EXTERNAL.
        /// </summary>
        [Input("rangeMatch")]
        public Input<Inputs.Int64RangeMatchArgs>? RangeMatch { get; set; }

        /// <summary>
        /// The value of the header must match the regular expression specified in regexMatch. For more information about regular expression syntax, see Syntax. For matching against a port specified in the HTTP request, use a headerMatch with headerName set to PORT and a regular expression that satisfies the RFC2616 Host header's port specifier. Only one of exactMatch, prefixMatch, suffixMatch, regexMatch, presentMatch or rangeMatch must be set. regexMatch only applies to load balancers that have loadBalancingScheme set to INTERNAL_SELF_MANAGED.
        /// </summary>
        [Input("regexMatch")]
        public Input<string>? RegexMatch { get; set; }

        /// <summary>
        /// The value of the header must end with the contents of suffixMatch. Only one of exactMatch, prefixMatch, suffixMatch, regexMatch, presentMatch or rangeMatch must be set.
        /// </summary>
        [Input("suffixMatch")]
        public Input<string>? SuffixMatch { get; set; }

        public HttpHeaderMatchArgs()
        {
        }
    }
}