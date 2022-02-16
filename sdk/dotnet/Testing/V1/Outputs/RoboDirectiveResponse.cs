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
    /// Directs Robo to interact with a specific UI element if it is encountered during the crawl. Currently, Robo can perform text entry or element click.
    /// </summary>
    [OutputType]
    public sealed class RoboDirectiveResponse
    {
        /// <summary>
        /// The type of action that Robo should perform on the specified element.
        /// </summary>
        public readonly string ActionType;
        /// <summary>
        /// The text that Robo is directed to set. If left empty, the directive will be treated as a CLICK on the element matching the resource_name.
        /// </summary>
        public readonly string InputText;
        /// <summary>
        /// The android resource name of the target UI element. For example, in Java: R.string.foo in xml: @string/foo Only the "foo" part is needed. Reference doc: https://developer.android.com/guide/topics/resources/accessing-resources.html
        /// </summary>
        public readonly string ResourceName;

        [OutputConstructor]
        private RoboDirectiveResponse(
            string actionType,

            string inputText,

            string resourceName)
        {
            ActionType = actionType;
            InputText = inputText;
            ResourceName = resourceName;
        }
    }
}