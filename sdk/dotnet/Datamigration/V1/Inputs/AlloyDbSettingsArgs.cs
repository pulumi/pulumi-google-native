// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datamigration.V1.Inputs
{

    /// <summary>
    /// Settings for creating an AlloyDB cluster.
    /// </summary>
    public sealed class AlloyDbSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Input only. Initial user to setup during cluster creation. Required.
        /// </summary>
        [Input("initialUser", required: true)]
        public Input<Inputs.UserPasswordArgs> InitialUser { get; set; } = null!;

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels for the AlloyDB cluster created by DMS. An object containing a list of 'key', 'value' pairs.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("primaryInstanceSettings")]
        public Input<Inputs.PrimaryInstanceSettingsArgs>? PrimaryInstanceSettings { get; set; }

        /// <summary>
        /// The resource link for the VPC network in which cluster resources are created and from which they are accessible via Private IP. The network must belong to the same project as the cluster. It is specified in the form: "projects/{project_number}/global/networks/{network_id}". This is required to create a cluster.
        /// </summary>
        [Input("vpcNetwork", required: true)]
        public Input<string> VpcNetwork { get; set; } = null!;

        public AlloyDbSettingsArgs()
        {
        }
        public static new AlloyDbSettingsArgs Empty => new AlloyDbSettingsArgs();
    }
}