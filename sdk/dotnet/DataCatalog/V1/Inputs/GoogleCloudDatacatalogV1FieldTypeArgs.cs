// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataCatalog.V1.Inputs
{

    public sealed class GoogleCloudDatacatalogV1FieldTypeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An enum type.
        /// </summary>
        [Input("enumType")]
        public Input<Inputs.GoogleCloudDatacatalogV1FieldTypeEnumTypeArgs>? EnumType { get; set; }

        /// <summary>
        /// Primitive types, such as string, boolean, etc.
        /// </summary>
        [Input("primitiveType")]
        public Input<Pulumi.GoogleNative.DataCatalog.V1.GoogleCloudDatacatalogV1FieldTypePrimitiveType>? PrimitiveType { get; set; }

        public GoogleCloudDatacatalogV1FieldTypeArgs()
        {
        }
        public static new GoogleCloudDatacatalogV1FieldTypeArgs Empty => new GoogleCloudDatacatalogV1FieldTypeArgs();
    }
}
