// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BareMetalSolution.V2.Outputs
{

    /// <summary>
    /// Configuration parameters for a new volume.
    /// </summary>
    [OutputType]
    public sealed class VolumeConfigResponse
    {
        /// <summary>
        /// The GCP service of the storage volume. Available gcp_service are in https://cloud.google.com/bare-metal/docs/bms-planning.
        /// </summary>
        public readonly string GcpService;
        /// <summary>
        /// LUN ranges to be configured. Set only when protocol is PROTOCOL_FC.
        /// </summary>
        public readonly ImmutableArray<Outputs.LunRangeResponse> LunRanges;
        /// <summary>
        /// Machine ids connected to this volume. Set only when protocol is PROTOCOL_FC.
        /// </summary>
        public readonly ImmutableArray<string> MachineIds;
        /// <summary>
        /// The name of the volume config.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// NFS exports. Set only when protocol is PROTOCOL_NFS.
        /// </summary>
        public readonly ImmutableArray<Outputs.NfsExportResponse> NfsExports;
        /// <summary>
        /// Volume protocol.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// The requested size of this volume, in GB.
        /// </summary>
        public readonly int SizeGb;
        /// <summary>
        /// Whether snapshots should be enabled.
        /// </summary>
        public readonly bool SnapshotsEnabled;
        /// <summary>
        /// The type of this Volume.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// User note field, it can be used by customers to add additional information for the BMS Ops team .
        /// </summary>
        public readonly string UserNote;

        [OutputConstructor]
        private VolumeConfigResponse(
            string gcpService,

            ImmutableArray<Outputs.LunRangeResponse> lunRanges,

            ImmutableArray<string> machineIds,

            string name,

            ImmutableArray<Outputs.NfsExportResponse> nfsExports,

            string protocol,

            int sizeGb,

            bool snapshotsEnabled,

            string type,

            string userNote)
        {
            GcpService = gcpService;
            LunRanges = lunRanges;
            MachineIds = machineIds;
            Name = name;
            NfsExports = nfsExports;
            Protocol = protocol;
            SizeGb = sizeGb;
            SnapshotsEnabled = snapshotsEnabled;
            Type = type;
            UserNote = userNote;
        }
    }
}