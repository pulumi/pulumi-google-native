// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha.Outputs
{

    /// <summary>
    /// Key-value pair of EventBus task parameters. Next id: 13
    /// </summary>
    [OutputType]
    public sealed class EnterpriseCrmFrontendsEventbusProtoParamSpecEntryResponse
    {
        /// <summary>
        /// The FQCN of the Java object this represents. A string, for example, would be "java.lang.String". If this is "java.lang.Object", the parameter can be of any type.
        /// </summary>
        public readonly string ClassName;
        /// <summary>
        /// If it is a collection of objects, this would be the FCQN of every individual element in the collection. If this is "java.lang.Object", the parameter is a collection of any type.
        /// </summary>
        public readonly string CollectionElementClassName;
        /// <summary>
        /// Optional fields, such as help text and other useful info.
        /// </summary>
        public readonly Outputs.EnterpriseCrmEventbusProtoParamSpecEntryConfigResponse Config;
        /// <summary>
        /// The data type of the parameter.
        /// </summary>
        public readonly string DataType;
        /// <summary>
        /// Default values for the defined keys. Each value can either be string, int, double or any proto message or a serialized object.
        /// </summary>
        public readonly Outputs.EnterpriseCrmFrontendsEventbusProtoParameterValueTypeResponse DefaultValue;
        /// <summary>
        /// If set, this entry is deprecated, so further use of this parameter should be prohibited.
        /// </summary>
        public readonly bool IsDeprecated;
        public readonly bool IsOutput;
        /// <summary>
        /// If the data_type is JSON_VALUE, then this will define its schema.
        /// </summary>
        public readonly string JsonSchema;
        /// <summary>
        /// Key is used to retrieve the corresponding parameter value. This should be unique for a given task. These parameters must be predefined in the workflow definition.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Populated if this represents a proto or proto array.
        /// </summary>
        public readonly Outputs.EnterpriseCrmEventbusProtoParamSpecEntryProtoDefinitionResponse ProtoDef;
        /// <summary>
        /// If set, the user must provide an input value for this parameter.
        /// </summary>
        public readonly bool Required;
        /// <summary>
        /// Rule used to validate inputs (individual values and collection elements) for this parameter.
        /// </summary>
        public readonly Outputs.EnterpriseCrmEventbusProtoParamSpecEntryValidationRuleResponse ValidationRule;

        [OutputConstructor]
        private EnterpriseCrmFrontendsEventbusProtoParamSpecEntryResponse(
            string className,

            string collectionElementClassName,

            Outputs.EnterpriseCrmEventbusProtoParamSpecEntryConfigResponse config,

            string dataType,

            Outputs.EnterpriseCrmFrontendsEventbusProtoParameterValueTypeResponse defaultValue,

            bool isDeprecated,

            bool isOutput,

            string jsonSchema,

            string key,

            Outputs.EnterpriseCrmEventbusProtoParamSpecEntryProtoDefinitionResponse protoDef,

            bool required,

            Outputs.EnterpriseCrmEventbusProtoParamSpecEntryValidationRuleResponse validationRule)
        {
            ClassName = className;
            CollectionElementClassName = collectionElementClassName;
            Config = config;
            DataType = dataType;
            DefaultValue = defaultValue;
            IsDeprecated = isDeprecated;
            IsOutput = isOutput;
            JsonSchema = jsonSchema;
            Key = key;
            ProtoDef = protoDef;
            Required = required;
            ValidationRule = validationRule;
        }
    }
}