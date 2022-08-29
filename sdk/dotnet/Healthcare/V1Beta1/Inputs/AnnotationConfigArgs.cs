// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1Beta1.Inputs
{

    /// <summary>
    /// Specifies how to store annotations during de-identification operation.
    /// </summary>
    public sealed class AnnotationConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the annotation store, in the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/annotationStores/{annotation_store_id}`). * The destination annotation store must be in the same project as the source data. De-identifying data across multiple projects is not supported. * The destination annotation store must exist when using DeidentifyDicomStore or DeidentifyFhirStore. DeidentifyDataset automatically creates the destination annotation store.
        /// </summary>
        [Input("annotationStoreName")]
        public Input<string>? AnnotationStoreName { get; set; }

        /// <summary>
        /// If set to true, the sensitive texts are included in SensitiveTextAnnotation of Annotation.
        /// </summary>
        [Input("storeQuote")]
        public Input<bool>? StoreQuote { get; set; }

        public AnnotationConfigArgs()
        {
        }
        public static new AnnotationConfigArgs Empty => new AnnotationConfigArgs();
    }
}