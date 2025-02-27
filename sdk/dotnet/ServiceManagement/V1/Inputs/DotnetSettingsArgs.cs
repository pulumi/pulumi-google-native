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
    /// Settings for Dotnet client libraries.
    /// </summary>
    public sealed class DotnetSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Some settings.
        /// </summary>
        [Input("common")]
        public Input<Inputs.CommonLanguageSettingsArgs>? Common { get; set; }

        [Input("forcedNamespaceAliases")]
        private InputList<string>? _forcedNamespaceAliases;

        /// <summary>
        /// Namespaces which must be aliased in snippets due to a known (but non-generator-predictable) naming collision
        /// </summary>
        public InputList<string> ForcedNamespaceAliases
        {
            get => _forcedNamespaceAliases ?? (_forcedNamespaceAliases = new InputList<string>());
            set => _forcedNamespaceAliases = value;
        }

        [Input("handwrittenSignatures")]
        private InputList<string>? _handwrittenSignatures;

        /// <summary>
        /// Method signatures (in the form "service.method(signature)") which are provided separately, so shouldn't be generated. Snippets *calling* these methods are still generated, however.
        /// </summary>
        public InputList<string> HandwrittenSignatures
        {
            get => _handwrittenSignatures ?? (_handwrittenSignatures = new InputList<string>());
            set => _handwrittenSignatures = value;
        }

        [Input("ignoredResources")]
        private InputList<string>? _ignoredResources;

        /// <summary>
        /// List of full resource types to ignore during generation. This is typically used for API-specific Location resources, which should be handled by the generator as if they were actually the common Location resources. Example entry: "documentai.googleapis.com/Location"
        /// </summary>
        public InputList<string> IgnoredResources
        {
            get => _ignoredResources ?? (_ignoredResources = new InputList<string>());
            set => _ignoredResources = value;
        }

        [Input("renamedResources")]
        private InputMap<string>? _renamedResources;

        /// <summary>
        /// Map from full resource types to the effective short name for the resource. This is used when otherwise resource named from different services would cause naming collisions. Example entry: "datalabeling.googleapis.com/Dataset": "DataLabelingDataset"
        /// </summary>
        public InputMap<string> RenamedResources
        {
            get => _renamedResources ?? (_renamedResources = new InputMap<string>());
            set => _renamedResources = value;
        }

        [Input("renamedServices")]
        private InputMap<string>? _renamedServices;

        /// <summary>
        /// Map from original service names to renamed versions. This is used when the default generated types would cause a naming conflict. (Neither name is fully-qualified.) Example: Subscriber to SubscriberServiceApi.
        /// </summary>
        public InputMap<string> RenamedServices
        {
            get => _renamedServices ?? (_renamedServices = new InputMap<string>());
            set => _renamedServices = value;
        }

        public DotnetSettingsArgs()
        {
        }
        public static new DotnetSettingsArgs Empty => new DotnetSettingsArgs();
    }
}
