// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V1Alpha1.Outputs
{

    /// <summary>
    /// Not supported by Cloud Run LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
    /// </summary>
    [OutputType]
    public sealed class LocalObjectReferenceResponse
    {
        /// <summary>
        /// (Optional) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private LocalObjectReferenceResponse(string name)
        {
            Name = name;
        }
    }
}