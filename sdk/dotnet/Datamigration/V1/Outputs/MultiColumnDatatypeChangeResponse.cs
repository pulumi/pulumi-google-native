// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datamigration.V1.Outputs
{

    /// <summary>
    /// Options to configure rule type MultiColumnDatatypeChange. The rule is used to change the data type and associated properties of multiple columns at once. The rule filter field can refer to one or more entities. The rule scope can be one of:Column. This rule requires additional filters to be specified beyond the basic rule filter field, which is the source data type, but the rule supports additional filtering capabilities such as the minimum and maximum field length. All additional filters which are specified are required to be met in order for the rule to be applied (logical AND between the fields).
    /// </summary>
    [OutputType]
    public sealed class MultiColumnDatatypeChangeResponse
    {
        /// <summary>
        /// Optional. Custom engine specific features.
        /// </summary>
        public readonly ImmutableDictionary<string, object> CustomFeatures;
        /// <summary>
        /// New data type.
        /// </summary>
        public readonly string NewDataType;
        /// <summary>
        /// Optional. Column fractional seconds precision - used only for timestamp based datatypes - if not specified and relevant uses the source column fractional seconds precision.
        /// </summary>
        public readonly int OverrideFractionalSecondsPrecision;
        /// <summary>
        /// Optional. Column length - e.g. varchar (50) - if not specified and relevant uses the source column length.
        /// </summary>
        public readonly string OverrideLength;
        /// <summary>
        /// Optional. Column precision - when relevant - if not specified and relevant uses the source column precision.
        /// </summary>
        public readonly int OverridePrecision;
        /// <summary>
        /// Optional. Column scale - when relevant - if not specified and relevant uses the source column scale.
        /// </summary>
        public readonly int OverrideScale;
        /// <summary>
        /// Filter on source data type.
        /// </summary>
        public readonly string SourceDataTypeFilter;
        /// <summary>
        /// Optional. Filter for fixed point number data types such as NUMERIC/NUMBER.
        /// </summary>
        public readonly Outputs.SourceNumericFilterResponse SourceNumericFilter;
        /// <summary>
        /// Optional. Filter for text-based data types like varchar.
        /// </summary>
        public readonly Outputs.SourceTextFilterResponse SourceTextFilter;

        [OutputConstructor]
        private MultiColumnDatatypeChangeResponse(
            ImmutableDictionary<string, object> customFeatures,

            string newDataType,

            int overrideFractionalSecondsPrecision,

            string overrideLength,

            int overridePrecision,

            int overrideScale,

            string sourceDataTypeFilter,

            Outputs.SourceNumericFilterResponse sourceNumericFilter,

            Outputs.SourceTextFilterResponse sourceTextFilter)
        {
            CustomFeatures = customFeatures;
            NewDataType = newDataType;
            OverrideFractionalSecondsPrecision = overrideFractionalSecondsPrecision;
            OverrideLength = overrideLength;
            OverridePrecision = overridePrecision;
            OverrideScale = overrideScale;
            SourceDataTypeFilter = sourceDataTypeFilter;
            SourceNumericFilter = sourceNumericFilter;
            SourceTextFilter = sourceTextFilter;
        }
    }
}
