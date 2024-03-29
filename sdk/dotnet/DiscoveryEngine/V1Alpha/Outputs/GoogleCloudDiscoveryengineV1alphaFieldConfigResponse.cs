// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DiscoveryEngine.V1Alpha.Outputs
{

    /// <summary>
    /// Configurations for fields of a schema. For example, configuring a field is indexable, or searchable.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDiscoveryengineV1alphaFieldConfigResponse
    {
        /// <summary>
        /// If completable_option is COMPLETABLE_ENABLED, field values are directly used and returned as suggestions for Autocomplete in CompletionService.CompleteQuery. If completable_option is unset, the server behavior defaults to COMPLETABLE_DISABLED for fields that support setting completable options, which are just `string` fields. For those fields that do not support setting completable options, the server will skip completable option setting, and setting completable_option for those fields will throw `INVALID_ARGUMENT` error.
        /// </summary>
        public readonly string CompletableOption;
        /// <summary>
        /// If dynamic_facetable_option is DYNAMIC_FACETABLE_ENABLED, field values are available for dynamic facet. Could only be DYNAMIC_FACETABLE_DISABLED if FieldConfig.indexable_option is INDEXABLE_DISABLED. Otherwise, an `INVALID_ARGUMENT` error will be returned. If dynamic_facetable_option is unset, the server behavior defaults to DYNAMIC_FACETABLE_DISABLED for fields that support setting dynamic facetable options. For those fields that do not support setting dynamic facetable options, such as `object` and `boolean`, the server will skip dynamic facetable option setting, and setting dynamic_facetable_option for those fields will throw `INVALID_ARGUMENT` error.
        /// </summary>
        public readonly string DynamicFacetableOption;
        /// <summary>
        /// Field path of the schema field. For example: `title`, `description`, `release_info.release_year`.
        /// </summary>
        public readonly string FieldPath;
        /// <summary>
        /// Raw type of the field.
        /// </summary>
        public readonly string FieldType;
        /// <summary>
        /// If indexable_option is INDEXABLE_ENABLED, field values are indexed so that it can be filtered or faceted in SearchService.Search. If indexable_option is unset, the server behavior defaults to INDEXABLE_DISABLED for fields that support setting indexable options. For those fields that do not support setting indexable options, such as `object` and `boolean` and key properties, the server will skip indexable_option setting, and setting indexable_option for those fields will throw `INVALID_ARGUMENT` error.
        /// </summary>
        public readonly string IndexableOption;
        /// <summary>
        /// Type of the key property that this field is mapped to. Empty string if this is not annotated as mapped to a key property. Example types are `title`, `description`. Full list is defined by `keyPropertyMapping` in the schema field annotation. If the schema field has a `KeyPropertyMapping` annotation, `indexable_option` and `searchable_option` of this field cannot be modified.
        /// </summary>
        public readonly string KeyPropertyType;
        /// <summary>
        /// If recs_filterable_option is FILTERABLE_ENABLED, field values are filterable by filter expression in RecommendationService.Recommend. If FILTERABLE_ENABLED but the field type is numerical, field values are not filterable by text queries in RecommendationService.Recommend. Only textual fields are supported. If recs_filterable_option is unset, the default setting is FILTERABLE_DISABLED for fields that support setting filterable options. When a field set to [FILTERABLE_DISABLED] is filtered, a warning is generated and an empty result is returned.
        /// </summary>
        public readonly string RecsFilterableOption;
        /// <summary>
        /// If retrievable_option is RETRIEVABLE_ENABLED, field values are included in the search results. If retrievable_option is unset, the server behavior defaults to RETRIEVABLE_DISABLED for fields that support setting retrievable options. For those fields that do not support setting retrievable options, such as `object` and `boolean`, the server will skip retrievable option setting, and setting retrievable_option for those fields will throw `INVALID_ARGUMENT` error.
        /// </summary>
        public readonly string RetrievableOption;
        /// <summary>
        /// If searchable_option is SEARCHABLE_ENABLED, field values are searchable by text queries in SearchService.Search. If SEARCHABLE_ENABLED but field type is numerical, field values will not be searchable by text queries in SearchService.Search, as there are no text values associated to numerical fields. If searchable_option is unset, the server behavior defaults to SEARCHABLE_DISABLED for fields that support setting searchable options. Only `string` fields that have no key property mapping support setting searchable_option. For those fields that do not support setting searchable options, the server will skip searchable option setting, and setting searchable_option for those fields will throw `INVALID_ARGUMENT` error.
        /// </summary>
        public readonly string SearchableOption;

        [OutputConstructor]
        private GoogleCloudDiscoveryengineV1alphaFieldConfigResponse(
            string completableOption,

            string dynamicFacetableOption,

            string fieldPath,

            string fieldType,

            string indexableOption,

            string keyPropertyType,

            string recsFilterableOption,

            string retrievableOption,

            string searchableOption)
        {
            CompletableOption = completableOption;
            DynamicFacetableOption = dynamicFacetableOption;
            FieldPath = fieldPath;
            FieldType = fieldType;
            IndexableOption = indexableOption;
            KeyPropertyType = keyPropertyType;
            RecsFilterableOption = recsFilterableOption;
            RetrievableOption = retrievableOption;
            SearchableOption = searchableOption;
        }
    }
}
