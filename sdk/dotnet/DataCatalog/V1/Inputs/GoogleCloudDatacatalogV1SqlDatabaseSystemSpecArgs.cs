// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataCatalog.V1.Inputs
{

    /// <summary>
    /// Specification that applies to entries that are part `SQL_DATABASE` system (user_specified_type)
    /// </summary>
    public sealed class GoogleCloudDatacatalogV1SqlDatabaseSystemSpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Version of the database engine.
        /// </summary>
        [Input("databaseVersion")]
        public Input<string>? DatabaseVersion { get; set; }

        /// <summary>
        /// Host of the SQL database enum InstanceHost { UNDEFINED = 0; SELF_HOSTED = 1; CLOUD_SQL = 2; AMAZON_RDS = 3; AZURE_SQL = 4; } Host of the enclousing database instance.
        /// </summary>
        [Input("instanceHost")]
        public Input<string>? InstanceHost { get; set; }

        /// <summary>
        /// SQL Database Engine. enum SqlEngine { UNDEFINED = 0; MY_SQL = 1; POSTGRE_SQL = 2; SQL_SERVER = 3; } Engine of the enclosing database instance.
        /// </summary>
        [Input("sqlEngine")]
        public Input<string>? SqlEngine { get; set; }

        public GoogleCloudDatacatalogV1SqlDatabaseSystemSpecArgs()
        {
        }
        public static new GoogleCloudDatacatalogV1SqlDatabaseSystemSpecArgs Empty => new GoogleCloudDatacatalogV1SqlDatabaseSystemSpecArgs();
    }
}