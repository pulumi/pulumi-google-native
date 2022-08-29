// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigQuery.V2.Inputs
{

    /// <summary>
    /// Options for a user-defined Spark routine.
    /// </summary>
    public sealed class SparkOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("archiveUris")]
        private InputList<string>? _archiveUris;

        /// <summary>
        /// Archive files to be extracted into the working directory of each executor. For more information about Apache Spark, see [Apache Spark](https://spark.apache.org/docs/latest/index.html).
        /// </summary>
        public InputList<string> ArchiveUris
        {
            get => _archiveUris ?? (_archiveUris = new InputList<string>());
            set => _archiveUris = value;
        }

        /// <summary>
        /// Fully qualified name of the user-provided Spark connection object. Format: ```"projects/{project_id}/locations/{location_id}/connections/{connection_id}"```
        /// </summary>
        [Input("connection")]
        public Input<string>? Connection { get; set; }

        /// <summary>
        /// Custom container image for the runtime environment.
        /// </summary>
        [Input("containerImage")]
        public Input<string>? ContainerImage { get; set; }

        [Input("fileUris")]
        private InputList<string>? _fileUris;

        /// <summary>
        /// Files to be placed in the working directory of each executor. For more information about Apache Spark, see [Apache Spark](https://spark.apache.org/docs/latest/index.html).
        /// </summary>
        public InputList<string> FileUris
        {
            get => _fileUris ?? (_fileUris = new InputList<string>());
            set => _fileUris = value;
        }

        [Input("jarUris")]
        private InputList<string>? _jarUris;

        /// <summary>
        /// JARs to include on the driver and executor CLASSPATH. For more information about Apache Spark, see [Apache Spark](https://spark.apache.org/docs/latest/index.html).
        /// </summary>
        public InputList<string> JarUris
        {
            get => _jarUris ?? (_jarUris = new InputList<string>());
            set => _jarUris = value;
        }

        /// <summary>
        /// The main file URI of the Spark application. Exactly one of the definition_body field and the main_file_uri field must be set.
        /// </summary>
        [Input("mainFileUri")]
        public Input<string>? MainFileUri { get; set; }

        [Input("properties")]
        private InputMap<string>? _properties;

        /// <summary>
        /// Configuration properties as a set of key/value pairs, which will be passed on to the Spark application. For more information, see [Apache Spark](https://spark.apache.org/docs/latest/index.html).
        /// </summary>
        public InputMap<string> Properties
        {
            get => _properties ?? (_properties = new InputMap<string>());
            set => _properties = value;
        }

        [Input("pyFileUris")]
        private InputList<string>? _pyFileUris;

        /// <summary>
        /// Python files to be placed on the PYTHONPATH for PySpark application. Supported file types: `.py`, `.egg`, and `.zip`. For more information about Apache Spark, see [Apache Spark](https://spark.apache.org/docs/latest/index.html).
        /// </summary>
        public InputList<string> PyFileUris
        {
            get => _pyFileUris ?? (_pyFileUris = new InputList<string>());
            set => _pyFileUris = value;
        }

        /// <summary>
        /// Runtime version. If not specified, the default runtime version is used.
        /// </summary>
        [Input("runtimeVersion")]
        public Input<string>? RuntimeVersion { get; set; }

        public SparkOptionsArgs()
        {
        }
        public static new SparkOptionsArgs Empty => new SparkOptionsArgs();
    }
}