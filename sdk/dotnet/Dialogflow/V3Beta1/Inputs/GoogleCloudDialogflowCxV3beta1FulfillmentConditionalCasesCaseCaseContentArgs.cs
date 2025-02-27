// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3Beta1.Inputs
{

    /// <summary>
    /// The list of messages or conditional cases to activate for this case.
    /// </summary>
    public sealed class GoogleCloudDialogflowCxV3beta1FulfillmentConditionalCasesCaseCaseContentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Additional cases to be evaluated.
        /// </summary>
        [Input("additionalCases")]
        public Input<Inputs.GoogleCloudDialogflowCxV3beta1FulfillmentConditionalCasesArgs>? AdditionalCases { get; set; }

        /// <summary>
        /// Returned message.
        /// </summary>
        [Input("message")]
        public Input<Inputs.GoogleCloudDialogflowCxV3beta1ResponseMessageArgs>? Message { get; set; }

        public GoogleCloudDialogflowCxV3beta1FulfillmentConditionalCasesCaseCaseContentArgs()
        {
        }
        public static new GoogleCloudDialogflowCxV3beta1FulfillmentConditionalCasesCaseCaseContentArgs Empty => new GoogleCloudDialogflowCxV3beta1FulfillmentConditionalCasesCaseCaseContentArgs();
    }
}
