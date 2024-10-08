// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datastream.V1.Inputs
{

    /// <summary>
    /// PostgreSQL schema.
    /// </summary>
    public sealed class PostgresqlSchemaArgs : global::Pulumi.ResourceArgs
    {
        [Input("postgresqlTables")]
        private InputList<Inputs.PostgresqlTableArgs>? _postgresqlTables;

        /// <summary>
        /// Tables in the schema.
        /// </summary>
        public InputList<Inputs.PostgresqlTableArgs> PostgresqlTables
        {
            get => _postgresqlTables ?? (_postgresqlTables = new InputList<Inputs.PostgresqlTableArgs>());
            set => _postgresqlTables = value;
        }

        /// <summary>
        /// Schema name.
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        public PostgresqlSchemaArgs()
        {
        }
        public static new PostgresqlSchemaArgs Empty => new PostgresqlSchemaArgs();
    }
}
