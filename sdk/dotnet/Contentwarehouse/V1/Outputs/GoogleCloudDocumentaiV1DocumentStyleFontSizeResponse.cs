// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Contentwarehouse.V1.Outputs
{

    /// <summary>
    /// Font size with unit.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDocumentaiV1DocumentStyleFontSizeResponse
    {
        /// <summary>
        /// Font size for the text.
        /// </summary>
        public readonly double Size;
        /// <summary>
        /// Unit for the font size. Follows CSS naming (in, px, pt, etc.).
        /// </summary>
        public readonly string Unit;

        [OutputConstructor]
        private GoogleCloudDocumentaiV1DocumentStyleFontSizeResponse(
            double size,

            string unit)
        {
            Size = size;
            Unit = unit;
        }
    }
}