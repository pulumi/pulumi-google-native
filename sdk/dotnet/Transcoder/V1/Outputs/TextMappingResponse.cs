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
    /// The mapping for the JobConfig.edit_list atoms with text EditAtom.inputs.
    /// </summary>
    [OutputType]
    public sealed class TextMappingResponse
    {
        /// <summary>
        /// The EditAtom.key that references atom with text inputs in the JobConfig.edit_list.
        /// </summary>
        public readonly string AtomKey;
        /// <summary>
        /// The Input.key that identifies the input file.
        /// </summary>
        public readonly string InputKey;
        /// <summary>
        /// The zero-based index of the track in the input file.
        /// </summary>
        public readonly int InputTrack;

        [OutputConstructor]
        private TextMappingResponse(
            string atomKey,

            string inputKey,

            int inputTrack)
        {
            AtomKey = atomKey;
            InputKey = inputKey;
            InputTrack = inputTrack;
        }
    }
}
