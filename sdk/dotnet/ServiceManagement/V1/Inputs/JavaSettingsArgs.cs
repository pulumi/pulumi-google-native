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
    /// Settings for Java client libraries.
    /// </summary>
    public sealed class JavaSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Some settings.
        /// </summary>
        [Input("common")]
        public Input<Inputs.CommonLanguageSettingsArgs>? Common { get; set; }

        /// <summary>
        /// The package name to use in Java. Clobbers the java_package option set in the protobuf. This should be used **only** by APIs who have already set the language_settings.java.package_name" field in gapic.yaml. API teams should use the protobuf java_package option where possible. Example of a YAML configuration:: publishing: java_settings: library_package: com.google.cloud.pubsub.v1
        /// </summary>
        [Input("libraryPackage")]
        public Input<string>? LibraryPackage { get; set; }

        [Input("serviceClassNames")]
        private InputMap<string>? _serviceClassNames;

        /// <summary>
        /// Configure the Java class name to use instead of the service's for its corresponding generated GAPIC client. Keys are fully-qualified service names as they appear in the protobuf (including the full the language_settings.java.interface_names" field in gapic.yaml. API teams should otherwise use the service name as it appears in the protobuf. Example of a YAML configuration:: publishing: java_settings: service_class_names: - google.pubsub.v1.Publisher: TopicAdmin - google.pubsub.v1.Subscriber: SubscriptionAdmin
        /// </summary>
        public InputMap<string> ServiceClassNames
        {
            get => _serviceClassNames ?? (_serviceClassNames = new InputMap<string>());
            set => _serviceClassNames = value;
        }

        public JavaSettingsArgs()
        {
        }
        public static new JavaSettingsArgs Empty => new JavaSettingsArgs();
    }
}