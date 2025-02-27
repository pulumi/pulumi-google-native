// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataCatalog.V1Beta1.Outputs
{

    /// <summary>
    /// The template for an individual field within a tag template.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDatacatalogV1beta1TagTemplateFieldResponse
    {
        /// <summary>
        /// The description for this field. Defaults to an empty string.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The display name for this field. Defaults to an empty string.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Whether this is a required field. Defaults to false.
        /// </summary>
        public readonly bool IsRequired;
        /// <summary>
        /// The resource name of the tag template field in URL format. Example: * projects/{project_id}/locations/{location}/tagTemplates/{tag_template}/fields/{field} Note that this TagTemplateField may not actually be stored in the location in this name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The order of this field with respect to other fields in this tag template. A higher value indicates a more important field. The value can be negative. Multiple fields can have the same order, and field orders within a tag do not have to be sequential.
        /// </summary>
        public readonly int Order;
        /// <summary>
        /// The type of value this tag field can contain.
        /// </summary>
        public readonly Outputs.GoogleCloudDatacatalogV1beta1FieldTypeResponse Type;

        [OutputConstructor]
        private GoogleCloudDatacatalogV1beta1TagTemplateFieldResponse(
            string description,

            string displayName,

            bool isRequired,

            string name,

            int order,

            Outputs.GoogleCloudDatacatalogV1beta1FieldTypeResponse type)
        {
            Description = description;
            DisplayName = displayName;
            IsRequired = isRequired;
            Name = name;
            Order = order;
            Type = type;
        }
    }
}
