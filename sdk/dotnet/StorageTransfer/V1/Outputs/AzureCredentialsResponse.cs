// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.StorageTransfer.V1.Outputs
{

    /// <summary>
    /// Azure credentials For information on our data retention policy for user credentials, see [User credentials](/storage-transfer/docs/data-retention#user-credentials).
    /// </summary>
    [OutputType]
    public sealed class AzureCredentialsResponse
    {
        /// <summary>
        /// Azure shared access signature (SAS). For more information about SAS, see [Grant limited access to Azure Storage resources using shared access signatures (SAS)](https://docs.microsoft.com/en-us/azure/storage/common/storage-sas-overview).
        /// </summary>
        public readonly string SasToken;

        [OutputConstructor]
        private AzureCredentialsResponse(string sasToken)
        {
            SasToken = sasToken;
        }
    }
}