// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1Beta1.Inputs
{

    /// <summary>
    /// Inspect text and transform sensitive text. Configurable using `TextConfig`. Supported [Value Representations] (http://dicom.nema.org/medical/dicom/2018e/output/chtml/part05/sect_6.2.html#table_6.2-1): AE, LO, LT, PN, SH, ST, UC, UT, DA, DT, AS
    /// </summary>
    public sealed class CleanTextTagArgs : global::Pulumi.ResourceArgs
    {
        public CleanTextTagArgs()
        {
        }
        public static new CleanTextTagArgs Empty => new CleanTextTagArgs();
    }
}