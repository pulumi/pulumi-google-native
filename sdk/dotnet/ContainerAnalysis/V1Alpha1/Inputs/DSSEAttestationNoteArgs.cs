// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.Inputs
{

    /// <summary>
    /// A note describing an attestation
    /// </summary>
    public sealed class DSSEAttestationNoteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DSSEHint hints at the purpose of the attestation authority.
        /// </summary>
        [Input("hint")]
        public Input<Inputs.DSSEHintArgs>? Hint { get; set; }

        public DSSEAttestationNoteArgs()
        {
        }
        public static new DSSEAttestationNoteArgs Empty => new DSSEAttestationNoteArgs();
    }
}
