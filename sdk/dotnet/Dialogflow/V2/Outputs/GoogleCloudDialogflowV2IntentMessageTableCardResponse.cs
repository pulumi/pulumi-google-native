// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2.Outputs
{

    /// <summary>
    /// Table card for Actions on Google.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowV2IntentMessageTableCardResponse
    {
        /// <summary>
        /// Optional. List of buttons for the card.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDialogflowV2IntentMessageBasicCardButtonResponse> Buttons;
        /// <summary>
        /// Optional. Display properties for the columns in this table.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDialogflowV2IntentMessageColumnPropertiesResponse> ColumnProperties;
        /// <summary>
        /// Optional. Image which should be displayed on the card.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2IntentMessageImageResponse Image;
        /// <summary>
        /// Optional. Rows in this table of data.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDialogflowV2IntentMessageTableCardRowResponse> Rows;
        /// <summary>
        /// Optional. Subtitle to the title.
        /// </summary>
        public readonly string Subtitle;
        /// <summary>
        /// Title of the card.
        /// </summary>
        public readonly string Title;

        [OutputConstructor]
        private GoogleCloudDialogflowV2IntentMessageTableCardResponse(
            ImmutableArray<Outputs.GoogleCloudDialogflowV2IntentMessageBasicCardButtonResponse> buttons,

            ImmutableArray<Outputs.GoogleCloudDialogflowV2IntentMessageColumnPropertiesResponse> columnProperties,

            Outputs.GoogleCloudDialogflowV2IntentMessageImageResponse image,

            ImmutableArray<Outputs.GoogleCloudDialogflowV2IntentMessageTableCardRowResponse> rows,

            string subtitle,

            string title)
        {
            Buttons = buttons;
            ColumnProperties = columnProperties;
            Image = image;
            Rows = rows;
            Subtitle = subtitle;
            Title = title;
        }
    }
}
