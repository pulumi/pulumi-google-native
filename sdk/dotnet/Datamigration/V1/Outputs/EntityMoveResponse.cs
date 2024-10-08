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
    /// Options to configure rule type EntityMove. The rule is used to move an entity to a new schema. The rule filter field can refer to one or more entities. The rule scope can be one of: Table, Column, Constraint, Index, View, Function, Stored Procedure, Materialized View, Sequence, UDT
    /// </summary>
    [OutputType]
    public sealed class EntityMoveResponse
    {
        /// <summary>
        /// The new schema
        /// </summary>
        public readonly string NewSchema;

        [OutputConstructor]
        private EntityMoveResponse(string newSchema)
        {
            NewSchema = newSchema;
        }
    }
}
