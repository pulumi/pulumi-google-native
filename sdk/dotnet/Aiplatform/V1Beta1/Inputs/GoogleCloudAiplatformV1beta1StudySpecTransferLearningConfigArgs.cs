// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Aiplatform.V1Beta1.Inputs
{

    /// <summary>
    /// This contains flag for manually disabling transfer learning for a study. The names of prior studies being used for transfer learning (if any) are also listed here.
    /// </summary>
    public sealed class GoogleCloudAiplatformV1beta1StudySpecTransferLearningConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Flag to to manually prevent vizier from using transfer learning on a new study. Otherwise, vizier will automatically determine whether or not to use transfer learning.
        /// </summary>
        [Input("disableTransferLearning")]
        public Input<bool>? DisableTransferLearning { get; set; }

        public GoogleCloudAiplatformV1beta1StudySpecTransferLearningConfigArgs()
        {
        }
        public static new GoogleCloudAiplatformV1beta1StudySpecTransferLearningConfigArgs Empty => new GoogleCloudAiplatformV1beta1StudySpecTransferLearningConfigArgs();
    }
}
