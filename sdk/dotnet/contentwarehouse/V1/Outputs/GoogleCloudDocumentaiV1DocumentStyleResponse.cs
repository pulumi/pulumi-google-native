// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.contentwarehouse.V1.Outputs
{

    /// <summary>
    /// Annotation for common text style attributes. This adheres to CSS conventions as much as possible.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDocumentaiV1DocumentStyleResponse
    {
        /// <summary>
        /// Text background color.
        /// </summary>
        public readonly Outputs.GoogleTypeColorResponse BackgroundColor;
        /// <summary>
        /// Text color.
        /// </summary>
        public readonly Outputs.GoogleTypeColorResponse Color;
        /// <summary>
        /// Font family such as `Arial`, `Times New Roman`. https://www.w3schools.com/cssref/pr_font_font-family.asp
        /// </summary>
        public readonly string FontFamily;
        /// <summary>
        /// Font size.
        /// </summary>
        public readonly Outputs.GoogleCloudDocumentaiV1DocumentStyleFontSizeResponse FontSize;
        /// <summary>
        /// Font weight. Possible values are normal, bold, bolder, and lighter. https://www.w3schools.com/cssref/pr_font_weight.asp
        /// </summary>
        public readonly string FontWeight;
        /// <summary>
        /// Text anchor indexing into the Document.text.
        /// </summary>
        public readonly Outputs.GoogleCloudDocumentaiV1DocumentTextAnchorResponse TextAnchor;
        /// <summary>
        /// Text decoration. Follows CSS standard. https://www.w3schools.com/cssref/pr_text_text-decoration.asp
        /// </summary>
        public readonly string TextDecoration;
        /// <summary>
        /// Text style. Possible values are normal, italic, and oblique. https://www.w3schools.com/cssref/pr_font_font-style.asp
        /// </summary>
        public readonly string TextStyle;

        [OutputConstructor]
        private GoogleCloudDocumentaiV1DocumentStyleResponse(
            Outputs.GoogleTypeColorResponse backgroundColor,

            Outputs.GoogleTypeColorResponse color,

            string fontFamily,

            Outputs.GoogleCloudDocumentaiV1DocumentStyleFontSizeResponse fontSize,

            string fontWeight,

            Outputs.GoogleCloudDocumentaiV1DocumentTextAnchorResponse textAnchor,

            string textDecoration,

            string textStyle)
        {
            BackgroundColor = backgroundColor;
            Color = color;
            FontFamily = fontFamily;
            FontSize = fontSize;
            FontWeight = fontWeight;
            TextAnchor = textAnchor;
            TextDecoration = textDecoration;
            TextStyle = textStyle;
        }
    }
}