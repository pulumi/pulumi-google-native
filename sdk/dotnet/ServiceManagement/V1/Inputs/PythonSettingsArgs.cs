// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ServiceManagement.V1.Inputs
{

    /// <summary>
    /// Settings for Python client libraries.
    /// </summary>
    public sealed class PythonSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Some settings.
        /// </summary>
        [Input("common")]
        public Input<Inputs.CommonLanguageSettingsArgs>? Common { get; set; }

        public PythonSettingsArgs()
        {
        }
        public static new PythonSettingsArgs Empty => new PythonSettingsArgs();
    }
}