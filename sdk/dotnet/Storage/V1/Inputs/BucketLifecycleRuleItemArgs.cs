// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Storage.V1.Inputs
{

    public sealed class BucketLifecycleRuleItemArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action to take.
        /// </summary>
        [Input("action")]
        public Input<Inputs.BucketLifecycleRuleItemActionArgs>? Action { get; set; }

        /// <summary>
        /// The condition(s) under which the action will be taken.
        /// </summary>
        [Input("condition")]
        public Input<Inputs.BucketLifecycleRuleItemConditionArgs>? Condition { get; set; }

        public BucketLifecycleRuleItemArgs()
        {
        }
    }
}