// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Pubsub.V1.Outputs
{

    /// <summary>
    /// Configuration for a Cloud Storage subscription.
    /// </summary>
    [OutputType]
    public sealed class CloudStorageConfigResponse
    {
        /// <summary>
        /// If set, message data will be written to Cloud Storage in Avro format.
        /// </summary>
        public readonly Outputs.AvroConfigResponse AvroConfig;
        /// <summary>
        /// User-provided name for the Cloud Storage bucket. The bucket must be created by the user. The bucket name must be without any prefix like "gs://". See the [bucket naming requirements] (https://cloud.google.com/storage/docs/buckets#naming).
        /// </summary>
        public readonly string Bucket;
        /// <summary>
        /// User-provided prefix for Cloud Storage filename. See the [object naming requirements](https://cloud.google.com/storage/docs/objects#naming).
        /// </summary>
        public readonly string FilenamePrefix;
        /// <summary>
        /// User-provided suffix for Cloud Storage filename. See the [object naming requirements](https://cloud.google.com/storage/docs/objects#naming).
        /// </summary>
        public readonly string FilenameSuffix;
        /// <summary>
        /// The maximum bytes that can be written to a Cloud Storage file before a new file is created. Min 1 KB, max 10 GiB. The max_bytes limit may be exceeded in cases where messages are larger than the limit.
        /// </summary>
        public readonly string MaxBytes;
        /// <summary>
        /// The maximum duration that can elapse before a new Cloud Storage file is created. Min 1 minute, max 10 minutes, default 5 minutes. May not exceed the subscription's acknowledgement deadline.
        /// </summary>
        public readonly string MaxDuration;
        /// <summary>
        /// An output-only field that indicates whether or not the subscription can receive messages.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// If set, message data will be written to Cloud Storage in text format.
        /// </summary>
        public readonly Outputs.TextConfigResponse TextConfig;

        [OutputConstructor]
        private CloudStorageConfigResponse(
            Outputs.AvroConfigResponse avroConfig,

            string bucket,

            string filenamePrefix,

            string filenameSuffix,

            string maxBytes,

            string maxDuration,

            string state,

            Outputs.TextConfigResponse textConfig)
        {
            AvroConfig = avroConfig;
            Bucket = bucket;
            FilenamePrefix = filenamePrefix;
            FilenameSuffix = filenameSuffix;
            MaxBytes = maxBytes;
            MaxDuration = maxDuration;
            State = state;
            TextConfig = textConfig;
        }
    }
}