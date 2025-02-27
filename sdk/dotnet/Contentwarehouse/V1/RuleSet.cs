// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Contentwarehouse.V1
{
    /// <summary>
    /// Creates a ruleset.
    /// </summary>
    [GoogleNativeResourceType("google-native:contentwarehouse/v1:RuleSet")]
    public partial class RuleSet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Short description of the rule-set.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The resource name of the rule set. Managed internally. Format: projects/{project_number}/locations/{location}/ruleSet/{rule_set_id}. The name is ignored when creating a rule set.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// List of rules given by the customer.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.GoogleCloudContentwarehouseV1RuleResponse>> Rules { get; private set; } = null!;

        /// <summary>
        /// Source of the rules i.e., customer name.
        /// </summary>
        [Output("source")]
        public Output<string> Source { get; private set; } = null!;


        /// <summary>
        /// Create a RuleSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RuleSet(string name, RuleSetArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:contentwarehouse/v1:RuleSet", name, args ?? new RuleSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RuleSet(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:contentwarehouse/v1:RuleSet", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "location",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RuleSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RuleSet Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RuleSet(name, id, options);
        }
    }

    public sealed class RuleSetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Short description of the rule-set.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The resource name of the rule set. Managed internally. Format: projects/{project_number}/locations/{location}/ruleSet/{rule_set_id}. The name is ignored when creating a rule set.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("rules")]
        private InputList<Inputs.GoogleCloudContentwarehouseV1RuleArgs>? _rules;

        /// <summary>
        /// List of rules given by the customer.
        /// </summary>
        public InputList<Inputs.GoogleCloudContentwarehouseV1RuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.GoogleCloudContentwarehouseV1RuleArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// Source of the rules i.e., customer name.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        public RuleSetArgs()
        {
        }
        public static new RuleSetArgs Empty => new RuleSetArgs();
    }
}
