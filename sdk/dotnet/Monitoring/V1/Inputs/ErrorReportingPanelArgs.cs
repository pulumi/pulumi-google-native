// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V1.Inputs
{

    /// <summary>
    /// A widget that displays a list of error groups.
    /// </summary>
    public sealed class ErrorReportingPanelArgs : global::Pulumi.ResourceArgs
    {
        [Input("projectNames")]
        private InputList<string>? _projectNames;

        /// <summary>
        /// The resource name of the Google Cloud Platform project. Written as projects/{projectID} or projects/{projectNumber}, where {projectID} and {projectNumber} can be found in the Google Cloud console (https://support.google.com/cloud/answer/6158840).Examples: projects/my-project-123, projects/5551234.
        /// </summary>
        public InputList<string> ProjectNames
        {
            get => _projectNames ?? (_projectNames = new InputList<string>());
            set => _projectNames = value;
        }

        [Input("services")]
        private InputList<string>? _services;

        /// <summary>
        /// An identifier of the service, such as the name of the executable, job, or Google App Engine service name. This field is expected to have a low number of values that are relatively stable over time, as opposed to version, which can be changed whenever new code is deployed.Contains the service name for error reports extracted from Google App Engine logs or default if the App Engine default service is used.
        /// </summary>
        public InputList<string> Services
        {
            get => _services ?? (_services = new InputList<string>());
            set => _services = value;
        }

        [Input("versions")]
        private InputList<string>? _versions;

        /// <summary>
        /// Represents the source code version that the developer provided, which could represent a version label or a Git SHA-1 hash, for example. For App Engine standard environment, the version is set to the version of the app.
        /// </summary>
        public InputList<string> Versions
        {
            get => _versions ?? (_versions = new InputList<string>());
            set => _versions = value;
        }

        public ErrorReportingPanelArgs()
        {
        }
        public static new ErrorReportingPanelArgs Empty => new ErrorReportingPanelArgs();
    }
}