// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    /// <summary>
    /// Represents a match condition that incoming traffic is evaluated against. Exactly one field must be specified.
    /// </summary>
    public sealed class FirewallPolicyRuleMatcherArgs : Pulumi.ResourceArgs
    {
        [Input("destIpRanges")]
        private InputList<string>? _destIpRanges;

        /// <summary>
        /// CIDR IP address range. Maximum number of destination CIDR IP ranges allowed is 5000.
        /// </summary>
        public InputList<string> DestIpRanges
        {
            get => _destIpRanges ?? (_destIpRanges = new InputList<string>());
            set => _destIpRanges = value;
        }

        [Input("layer4Configs")]
        private InputList<Inputs.FirewallPolicyRuleMatcherLayer4ConfigArgs>? _layer4Configs;

        /// <summary>
        /// Pairs of IP protocols and ports that the rule should match.
        /// </summary>
        public InputList<Inputs.FirewallPolicyRuleMatcherLayer4ConfigArgs> Layer4Configs
        {
            get => _layer4Configs ?? (_layer4Configs = new InputList<Inputs.FirewallPolicyRuleMatcherLayer4ConfigArgs>());
            set => _layer4Configs = value;
        }

        [Input("srcIpRanges")]
        private InputList<string>? _srcIpRanges;

        /// <summary>
        /// CIDR IP address range. Maximum number of source CIDR IP ranges allowed is 5000.
        /// </summary>
        public InputList<string> SrcIpRanges
        {
            get => _srcIpRanges ?? (_srcIpRanges = new InputList<string>());
            set => _srcIpRanges = value;
        }

        [Input("srcSecureTags")]
        private InputList<Inputs.FirewallPolicyRuleSecureTagArgs>? _srcSecureTags;

        /// <summary>
        /// List of secure tag values, which should be matched at the source of the traffic. For INGRESS rule, if all the srcSecureTag are INEFFECTIVE, and there is no srcIpRange, this rule will be ignored. Maximum number of source tag values allowed is 256.
        /// </summary>
        public InputList<Inputs.FirewallPolicyRuleSecureTagArgs> SrcSecureTags
        {
            get => _srcSecureTags ?? (_srcSecureTags = new InputList<Inputs.FirewallPolicyRuleSecureTagArgs>());
            set => _srcSecureTags = value;
        }

        public FirewallPolicyRuleMatcherArgs()
        {
        }
    }
}