// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Testing.V1.Outputs
{

    /// <summary>
    /// A single iOS device.
    /// </summary>
    [OutputType]
    public sealed class IosDeviceResponse
    {
        /// <summary>
        /// The id of the iOS device to be used. Use the TestEnvironmentDiscoveryService to get supported options.
        /// </summary>
        public readonly string IosModelId;
        /// <summary>
        /// The id of the iOS major software version to be used. Use the TestEnvironmentDiscoveryService to get supported options.
        /// </summary>
        public readonly string IosVersionId;
        /// <summary>
        /// The locale the test device used for testing. Use the TestEnvironmentDiscoveryService to get supported options.
        /// </summary>
        public readonly string Locale;
        /// <summary>
        /// How the device is oriented during the test. Use the TestEnvironmentDiscoveryService to get supported options.
        /// </summary>
        public readonly string Orientation;

        [OutputConstructor]
        private IosDeviceResponse(
            string iosModelId,

            string iosVersionId,

            string locale,

            string orientation)
        {
            IosModelId = iosModelId;
            IosVersionId = iosVersionId;
            Locale = locale;
            Orientation = orientation;
        }
    }
}
