// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Aiplatform.V1.Outputs
{

    /// <summary>
    /// Contains the schemata used in Model's predictions and explanations via PredictionService.Predict, PredictionService.Explain and BatchPredictionJob.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudAiplatformV1PredictSchemataResponse
    {
        /// <summary>
        /// Immutable. Points to a YAML file stored on Google Cloud Storage describing the format of a single instance, which are used in PredictRequest.instances, ExplainRequest.instances and BatchPredictionJob.input_config. The schema is defined as an OpenAPI 3.0.2 [Schema Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.2.md#schemaObject). AutoML Models always have this field populated by Vertex AI. Note: The URI given on output will be immutable and probably different, including the URI scheme, than the one given on input. The output URI will point to a location where the user only has a read access.
        /// </summary>
        public readonly string InstanceSchemaUri;
        /// <summary>
        /// Immutable. Points to a YAML file stored on Google Cloud Storage describing the parameters of prediction and explanation via PredictRequest.parameters, ExplainRequest.parameters and BatchPredictionJob.model_parameters. The schema is defined as an OpenAPI 3.0.2 [Schema Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.2.md#schemaObject). AutoML Models always have this field populated by Vertex AI, if no parameters are supported, then it is set to an empty string. Note: The URI given on output will be immutable and probably different, including the URI scheme, than the one given on input. The output URI will point to a location where the user only has a read access.
        /// </summary>
        public readonly string ParametersSchemaUri;
        /// <summary>
        /// Immutable. Points to a YAML file stored on Google Cloud Storage describing the format of a single prediction produced by this Model, which are returned via PredictResponse.predictions, ExplainResponse.explanations, and BatchPredictionJob.output_config. The schema is defined as an OpenAPI 3.0.2 [Schema Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.2.md#schemaObject). AutoML Models always have this field populated by Vertex AI. Note: The URI given on output will be immutable and probably different, including the URI scheme, than the one given on input. The output URI will point to a location where the user only has a read access.
        /// </summary>
        public readonly string PredictionSchemaUri;

        [OutputConstructor]
        private GoogleCloudAiplatformV1PredictSchemataResponse(
            string instanceSchemaUri,

            string parametersSchemaUri,

            string predictionSchemaUri)
        {
            InstanceSchemaUri = instanceSchemaUri;
            ParametersSchemaUri = parametersSchemaUri;
            PredictionSchemaUri = predictionSchemaUri;
        }
    }
}