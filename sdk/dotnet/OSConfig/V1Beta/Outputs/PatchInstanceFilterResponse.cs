// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Beta.Outputs
{

    /// <summary>
    /// A filter to target VM instances for patching. The targeted VMs must meet all criteria specified. So if both labels and zones are specified, the patch job targets only VMs with those labels and in those zones.
    /// </summary>
    [OutputType]
    public sealed class PatchInstanceFilterResponse
    {
        /// <summary>
        /// Target all VM instances in the project. If true, no other criteria is permitted.
        /// </summary>
        public readonly bool All;
        /// <summary>
        /// Targets VM instances matching at least one of these label sets. This allows targeting of disparate groups, for example "env=prod or env=staging".
        /// </summary>
        public readonly ImmutableArray<Outputs.PatchInstanceFilterGroupLabelResponse> GroupLabels;
        /// <summary>
        /// Targets VMs whose name starts with one of these prefixes. Similar to labels, this is another way to group VMs when targeting configs, for example prefix="prod-".
        /// </summary>
        public readonly ImmutableArray<string> InstanceNamePrefixes;
        /// <summary>
        /// Targets any of the VM instances specified. Instances are specified by their URI in the form `zones/[ZONE]/instances/[INSTANCE_NAME]`, `projects/[PROJECT_ID]/zones/[ZONE]/instances/[INSTANCE_NAME]`, or `https://www.googleapis.com/compute/v1/projects/[PROJECT_ID]/zones/[ZONE]/instances/[INSTANCE_NAME]`
        /// </summary>
        public readonly ImmutableArray<string> Instances;
        /// <summary>
        /// Targets VM instances in ANY of these zones. Leave empty to target VM instances in any zone.
        /// </summary>
        public readonly ImmutableArray<string> Zones;

        [OutputConstructor]
        private PatchInstanceFilterResponse(
            bool all,

            ImmutableArray<Outputs.PatchInstanceFilterGroupLabelResponse> groupLabels,

            ImmutableArray<string> instanceNamePrefixes,

            ImmutableArray<string> instances,

            ImmutableArray<string> zones)
        {
            All = all;
            GroupLabels = groupLabels;
            InstanceNamePrefixes = instanceNamePrefixes;
            Instances = instances;
            Zones = zones;
        }
    }
}