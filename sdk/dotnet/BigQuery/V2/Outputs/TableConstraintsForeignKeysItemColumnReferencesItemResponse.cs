// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigQuery.V2.Outputs
{

    [OutputType]
    public sealed class TableConstraintsForeignKeysItemColumnReferencesItemResponse
    {
        public readonly string ReferencedColumn;
        public readonly string ReferencingColumn;

        [OutputConstructor]
        private TableConstraintsForeignKeysItemColumnReferencesItemResponse(
            string referencedColumn,

            string referencingColumn)
        {
            ReferencedColumn = referencedColumn;
            ReferencingColumn = referencingColumn;
        }
    }
}