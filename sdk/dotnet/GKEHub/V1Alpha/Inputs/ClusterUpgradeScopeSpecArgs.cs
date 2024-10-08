// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1Alpha.Inputs
{

    /// <summary>
    /// **ClusterUpgrade**: The configuration for the scope-level ClusterUpgrade feature.
    /// </summary>
    public sealed class ClusterUpgradeScopeSpecArgs : global::Pulumi.ResourceArgs
    {
        [Input("gkeUpgradeOverrides")]
        private InputList<Inputs.ClusterUpgradeGKEUpgradeOverrideArgs>? _gkeUpgradeOverrides;

        /// <summary>
        /// Allow users to override some properties of each GKE upgrade.
        /// </summary>
        public InputList<Inputs.ClusterUpgradeGKEUpgradeOverrideArgs> GkeUpgradeOverrides
        {
            get => _gkeUpgradeOverrides ?? (_gkeUpgradeOverrides = new InputList<Inputs.ClusterUpgradeGKEUpgradeOverrideArgs>());
            set => _gkeUpgradeOverrides = value;
        }

        /// <summary>
        /// Post conditions to evaluate to mark an upgrade COMPLETE. Required.
        /// </summary>
        [Input("postConditions", required: true)]
        public Input<Inputs.ClusterUpgradePostConditionsArgs> PostConditions { get; set; } = null!;

        [Input("upstreamScopes")]
        private InputList<string>? _upstreamScopes;

        /// <summary>
        /// This scope consumes upgrades that have COMPLETE status code in the upstream scopes. See UpgradeStatus.Code for code definitions. The scope name should be in the form: `projects/{p}/locations/global/scopes/{s}` Where {p} is the project, {s} is a valid Scope in this project. {p} WILL match the Feature's project. This is defined as repeated for future proof reasons. Initial implementation will enforce at most one upstream scope.
        /// </summary>
        public InputList<string> UpstreamScopes
        {
            get => _upstreamScopes ?? (_upstreamScopes = new InputList<string>());
            set => _upstreamScopes = value;
        }

        public ClusterUpgradeScopeSpecArgs()
        {
        }
        public static new ClusterUpgradeScopeSpecArgs Empty => new ClusterUpgradeScopeSpecArgs();
    }
}
