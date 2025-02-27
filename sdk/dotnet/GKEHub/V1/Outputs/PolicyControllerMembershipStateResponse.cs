// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1.Outputs
{

    /// <summary>
    /// **Policy Controller**: State for a single cluster.
    /// </summary>
    [OutputType]
    public sealed class PolicyControllerMembershipStateResponse
    {
        /// <summary>
        /// Currently these include (also serving as map keys): 1. "admission" 2. "audit" 3. "mutation"
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.PolicyControllerOnClusterStateResponse> ComponentStates;
        /// <summary>
        /// The overall content state observed by the Hub Feature controller.
        /// </summary>
        public readonly Outputs.PolicyControllerPolicyContentStateResponse PolicyContentState;
        /// <summary>
        /// The overall Policy Controller lifecycle state observed by the Hub Feature controller.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private PolicyControllerMembershipStateResponse(
            ImmutableDictionary<string, Outputs.PolicyControllerOnClusterStateResponse> componentStates,

            Outputs.PolicyControllerPolicyContentStateResponse policyContentState,

            string state)
        {
            ComponentStates = componentStates;
            PolicyContentState = policyContentState;
            State = state;
        }
    }
}
