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
    /// Private Connectivity.
    /// </summary>
    [OutputType]
    public sealed class PrivateConnectivityResponse
    {
        /// <summary>
        /// The resource name (URI) of the private connection.
        /// </summary>
        public readonly string PrivateConnection;

        [OutputConstructor]
        private PrivateConnectivityResponse(string privateConnection)
        {
            PrivateConnection = privateConnection;
        }
    }
}
