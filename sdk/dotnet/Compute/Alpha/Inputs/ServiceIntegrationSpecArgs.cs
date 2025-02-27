// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    /// <summary>
    /// Specifies the parameters to configure an integration with instances.
    /// </summary>
    public sealed class ServiceIntegrationSpecArgs : global::Pulumi.ResourceArgs
    {
        [Input("backupDr")]
        public Input<Inputs.ServiceIntegrationSpecBackupDRSpecArgs>? BackupDr { get; set; }

        public ServiceIntegrationSpecArgs()
        {
        }
        public static new ServiceIntegrationSpecArgs Empty => new ServiceIntegrationSpecArgs();
    }
}
