// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V2.Inputs
{

    /// <summary>
    /// Ephemeral storage which can be backed by real disks (HD, SSD), network storage or memory (i.e. tmpfs). For now only in memory (tmpfs) is supported. It is ephemeral in the sense that when the sandbox is taken down, the data is destroyed with it (it does not persist across sandbox runs).
    /// </summary>
    public sealed class GoogleCloudRunV2EmptyDirVolumeSourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The medium on which the data is stored. Acceptable values today is only MEMORY or none. When none, the default will currently be backed by memory but could change over time. +optional
        /// </summary>
        [Input("medium")]
        public Input<Pulumi.GoogleNative.Run.V2.GoogleCloudRunV2EmptyDirVolumeSourceMedium>? Medium { get; set; }

        /// <summary>
        /// Limit on the storage usable by this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. This field's values are of the 'Quantity' k8s type: https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir +optional
        /// </summary>
        [Input("sizeLimit")]
        public Input<string>? SizeLimit { get; set; }

        public GoogleCloudRunV2EmptyDirVolumeSourceArgs()
        {
        }
        public static new GoogleCloudRunV2EmptyDirVolumeSourceArgs Empty => new GoogleCloudRunV2EmptyDirVolumeSourceArgs();
    }
}