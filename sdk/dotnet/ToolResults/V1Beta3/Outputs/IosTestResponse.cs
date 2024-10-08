// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ToolResults.V1Beta3.Outputs
{

    /// <summary>
    /// A iOS mobile test specification
    /// </summary>
    [OutputType]
    public sealed class IosTestResponse
    {
        /// <summary>
        /// Information about the application under test.
        /// </summary>
        public readonly Outputs.IosAppInfoResponse IosAppInfo;
        /// <summary>
        /// An iOS Robo test.
        /// </summary>
        public readonly Outputs.IosRoboTestResponse IosRoboTest;
        /// <summary>
        /// An iOS test loop.
        /// </summary>
        public readonly Outputs.IosTestLoopResponse IosTestLoop;
        /// <summary>
        /// An iOS XCTest.
        /// </summary>
        public readonly Outputs.IosXcTestResponse IosXcTest;
        /// <summary>
        /// Max time a test is allowed to run before it is automatically cancelled.
        /// </summary>
        public readonly Outputs.DurationResponse TestTimeout;

        [OutputConstructor]
        private IosTestResponse(
            Outputs.IosAppInfoResponse iosAppInfo,

            Outputs.IosRoboTestResponse iosRoboTest,

            Outputs.IosTestLoopResponse iosTestLoop,

            Outputs.IosXcTestResponse iosXcTest,

            Outputs.DurationResponse testTimeout)
        {
            IosAppInfo = iosAppInfo;
            IosRoboTest = iosRoboTest;
            IosTestLoop = iosTestLoop;
            IosXcTest = iosXcTest;
            TestTimeout = testTimeout;
        }
    }
}
