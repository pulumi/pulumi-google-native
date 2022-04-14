// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Transcoder.V1.Outputs
{

    /// <summary>
    /// Multiplexing settings for output stream.
    /// </summary>
    [OutputType]
    public sealed class MuxStreamResponse
    {
        /// <summary>
        /// The container format. The default is `mp4` Supported container formats: - `ts` - `fmp4`- the corresponding file extension is `.m4s` - `mp4` - `vtt` See also: [Supported input and output formats](https://cloud.google.com/transcoder/docs/concepts/supported-input-and-output-formats)
        /// </summary>
        public readonly string Container;
        /// <summary>
        /// List of `ElementaryStream.key`s multiplexed in this stream.
        /// </summary>
        public readonly ImmutableArray<string> ElementaryStreams;
        /// <summary>
        /// The name of the generated file. The default is `MuxStream.key` with the extension suffix corresponding to the `MuxStream.container`. Individual segments also have an incremental 10-digit zero-padded suffix starting from 0 before the extension, such as `mux_stream0000000123.ts`.
        /// </summary>
        public readonly string FileName;
        /// <summary>
        /// A unique key for this multiplexed stream. HLS media manifests will be named `MuxStream.key` with the `.m3u8` extension suffix.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Segment settings for `ts`, `fmp4` and `vtt`.
        /// </summary>
        public readonly Outputs.SegmentSettingsResponse SegmentSettings;

        [OutputConstructor]
        private MuxStreamResponse(
            string container,

            ImmutableArray<string> elementaryStreams,

            string fileName,

            string key,

            Outputs.SegmentSettingsResponse segmentSettings)
        {
            Container = container;
            ElementaryStreams = elementaryStreams;
            FileName = fileName;
            Key = key;
            SegmentSettings = segmentSettings;
        }
    }
}