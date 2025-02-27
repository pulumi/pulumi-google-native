// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Aiplatform.V1.Inputs
{

    /// <summary>
    /// Value is the value of the field.
    /// </summary>
    public sealed class GoogleCloudAiplatformV1ValueArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A double value.
        /// </summary>
        [Input("doubleValue")]
        public Input<double>? DoubleValue { get; set; }

        /// <summary>
        /// An integer value.
        /// </summary>
        [Input("intValue")]
        public Input<string>? IntValue { get; set; }

        /// <summary>
        /// A string value.
        /// </summary>
        [Input("stringValue")]
        public Input<string>? StringValue { get; set; }

        public GoogleCloudAiplatformV1ValueArgs()
        {
        }
        public static new GoogleCloudAiplatformV1ValueArgs Empty => new GoogleCloudAiplatformV1ValueArgs();
    }
}
