// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datamigration.V1.Inputs
{

    /// <summary>
    /// Options to configure rule type FilterTableColumns. The rule is used to filter the list of columns to include or exclude from a table. The rule filter field can refer to one entity. The rule scope can be: Table Only one of the two lists can be specified for the rule.
    /// </summary>
    public sealed class FilterTableColumnsArgs : global::Pulumi.ResourceArgs
    {
        [Input("excludeColumns")]
        private InputList<string>? _excludeColumns;

        /// <summary>
        /// Optional. List of columns to be excluded for a particular table.
        /// </summary>
        public InputList<string> ExcludeColumns
        {
            get => _excludeColumns ?? (_excludeColumns = new InputList<string>());
            set => _excludeColumns = value;
        }

        [Input("includeColumns")]
        private InputList<string>? _includeColumns;

        /// <summary>
        /// Optional. List of columns to be included for a particular table.
        /// </summary>
        public InputList<string> IncludeColumns
        {
            get => _includeColumns ?? (_includeColumns = new InputList<string>());
            set => _includeColumns = value;
        }

        public FilterTableColumnsArgs()
        {
        }
        public static new FilterTableColumnsArgs Empty => new FilterTableColumnsArgs();
    }
}