// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3.Outputs
{

    /// <summary>
    /// Settings for Gen App Builder.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowCxV3AgentGenAppBuilderSettingsResponse
    {
        /// <summary>
        /// The full name of the Gen App Builder engine related to this agent if there is one. Format: `projects/{Project ID}/locations/{Location ID}/collections/{Collection ID}/engines/{Engine ID}`
        /// </summary>
        public readonly string Engine;

        [OutputConstructor]
        private GoogleCloudDialogflowCxV3AgentGenAppBuilderSettingsResponse(string engine)
        {
            Engine = engine;
        }
    }
}
