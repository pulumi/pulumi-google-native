// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DLP.V2.Outputs
{

    /// <summary>
    /// De-id options.
    /// </summary>
    [OutputType]
    public sealed class GooglePrivacyDlpV2RequestedDeidentifyOptionsResponse
    {
        /// <summary>
        /// Snapshot of the state of the `DeidentifyTemplate` from the Deidentify action at the time this job was run.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2DeidentifyTemplateResponse SnapshotDeidentifyTemplate;
        /// <summary>
        /// Snapshot of the state of the image transformation `DeidentifyTemplate` from the `Deidentify` action at the time this job was run.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2DeidentifyTemplateResponse SnapshotImageRedactTemplate;
        /// <summary>
        /// Snapshot of the state of the structured `DeidentifyTemplate` from the `Deidentify` action at the time this job was run.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2DeidentifyTemplateResponse SnapshotStructuredDeidentifyTemplate;

        [OutputConstructor]
        private GooglePrivacyDlpV2RequestedDeidentifyOptionsResponse(
            Outputs.GooglePrivacyDlpV2DeidentifyTemplateResponse snapshotDeidentifyTemplate,

            Outputs.GooglePrivacyDlpV2DeidentifyTemplateResponse snapshotImageRedactTemplate,

            Outputs.GooglePrivacyDlpV2DeidentifyTemplateResponse snapshotStructuredDeidentifyTemplate)
        {
            SnapshotDeidentifyTemplate = snapshotDeidentifyTemplate;
            SnapshotImageRedactTemplate = snapshotImageRedactTemplate;
            SnapshotStructuredDeidentifyTemplate = snapshotStructuredDeidentifyTemplate;
        }
    }
}