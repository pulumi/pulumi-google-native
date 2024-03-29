// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Inputs
{

    /// <summary>
    /// Encryption settings for the encrypting customer core content. NEXT ID: 2
    /// </summary>
    public sealed class GoogleCloudDataprocV1WorkflowTemplateEncryptionConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The Cloud KMS key name to use for encrypting customer core content.
        /// </summary>
        [Input("kmsKey")]
        public Input<string>? KmsKey { get; set; }

        public GoogleCloudDataprocV1WorkflowTemplateEncryptionConfigArgs()
        {
        }
        public static new GoogleCloudDataprocV1WorkflowTemplateEncryptionConfigArgs Empty => new GoogleCloudDataprocV1WorkflowTemplateEncryptionConfigArgs();
    }
}
