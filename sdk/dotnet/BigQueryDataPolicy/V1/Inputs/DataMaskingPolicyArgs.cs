// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigQueryDataPolicy.V1.Inputs
{

    /// <summary>
    /// The data masking policy that is used to specify data masking rule.
    /// </summary>
    public sealed class DataMaskingPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A predefined masking expression.
        /// </summary>
        [Input("predefinedExpression")]
        public Input<Pulumi.GoogleNative.BigQueryDataPolicy.V1.DataMaskingPolicyPredefinedExpression>? PredefinedExpression { get; set; }

        /// <summary>
        /// The name of the BigQuery routine that contains the custom masking routine, in the format of `projects/{project_number}/datasets/{dataset_id}/routines/{routine_id}`.
        /// </summary>
        [Input("routine")]
        public Input<string>? Routine { get; set; }

        public DataMaskingPolicyArgs()
        {
        }
        public static new DataMaskingPolicyArgs Empty => new DataMaskingPolicyArgs();
    }
}