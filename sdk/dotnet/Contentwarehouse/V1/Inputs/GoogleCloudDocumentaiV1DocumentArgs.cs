// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Contentwarehouse.V1.Inputs
{

    /// <summary>
    /// Document represents the canonical document resource in Document AI. It is an interchange format that provides insights into documents and allows for collaboration between users and Document AI to iterate and optimize for quality.
    /// </summary>
    public sealed class GoogleCloudDocumentaiV1DocumentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Inline document content, represented as a stream of bytes. Note: As with all `bytes` fields, protobuffers use a pure binary representation, whereas JSON representations use base64.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        [Input("entities")]
        private InputList<Inputs.GoogleCloudDocumentaiV1DocumentEntityArgs>? _entities;

        /// <summary>
        /// A list of entities detected on Document.text. For document shards, entities in this list may cross shard boundaries.
        /// </summary>
        public InputList<Inputs.GoogleCloudDocumentaiV1DocumentEntityArgs> Entities
        {
            get => _entities ?? (_entities = new InputList<Inputs.GoogleCloudDocumentaiV1DocumentEntityArgs>());
            set => _entities = value;
        }

        [Input("entityRelations")]
        private InputList<Inputs.GoogleCloudDocumentaiV1DocumentEntityRelationArgs>? _entityRelations;

        /// <summary>
        /// Placeholder. Relationship among Document.entities.
        /// </summary>
        public InputList<Inputs.GoogleCloudDocumentaiV1DocumentEntityRelationArgs> EntityRelations
        {
            get => _entityRelations ?? (_entityRelations = new InputList<Inputs.GoogleCloudDocumentaiV1DocumentEntityRelationArgs>());
            set => _entityRelations = value;
        }

        /// <summary>
        /// Any error that occurred while processing this document.
        /// </summary>
        [Input("error")]
        public Input<Inputs.GoogleRpcStatusArgs>? Error { get; set; }

        /// <summary>
        /// An IANA published MIME type (also referred to as media type). For more information, see https://www.iana.org/assignments/media-types/media-types.xhtml.
        /// </summary>
        [Input("mimeType")]
        public Input<string>? MimeType { get; set; }

        [Input("pages")]
        private InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageArgs>? _pages;

        /// <summary>
        /// Visual page layout for the Document.
        /// </summary>
        public InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageArgs> Pages
        {
            get => _pages ?? (_pages = new InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageArgs>());
            set => _pages = value;
        }

        [Input("revisions")]
        private InputList<Inputs.GoogleCloudDocumentaiV1DocumentRevisionArgs>? _revisions;

        /// <summary>
        /// Placeholder. Revision history of this document.
        /// </summary>
        public InputList<Inputs.GoogleCloudDocumentaiV1DocumentRevisionArgs> Revisions
        {
            get => _revisions ?? (_revisions = new InputList<Inputs.GoogleCloudDocumentaiV1DocumentRevisionArgs>());
            set => _revisions = value;
        }

        /// <summary>
        /// Information about the sharding if this document is sharded part of a larger document. If the document is not sharded, this message is not specified.
        /// </summary>
        [Input("shardInfo")]
        public Input<Inputs.GoogleCloudDocumentaiV1DocumentShardInfoArgs>? ShardInfo { get; set; }

        /// <summary>
        /// Optional. UTF-8 encoded text in reading order from the document.
        /// </summary>
        [Input("text")]
        public Input<string>? Text { get; set; }

        [Input("textChanges")]
        private InputList<Inputs.GoogleCloudDocumentaiV1DocumentTextChangeArgs>? _textChanges;

        /// <summary>
        /// Placeholder. A list of text corrections made to Document.text. This is usually used for annotating corrections to OCR mistakes. Text changes for a given revision may not overlap with each other.
        /// </summary>
        public InputList<Inputs.GoogleCloudDocumentaiV1DocumentTextChangeArgs> TextChanges
        {
            get => _textChanges ?? (_textChanges = new InputList<Inputs.GoogleCloudDocumentaiV1DocumentTextChangeArgs>());
            set => _textChanges = value;
        }

        [Input("textStyles")]
        private InputList<Inputs.GoogleCloudDocumentaiV1DocumentStyleArgs>? _textStyles;

        /// <summary>
        /// Styles for the Document.text.
        /// </summary>
        public InputList<Inputs.GoogleCloudDocumentaiV1DocumentStyleArgs> TextStyles
        {
            get => _textStyles ?? (_textStyles = new InputList<Inputs.GoogleCloudDocumentaiV1DocumentStyleArgs>());
            set => _textStyles = value;
        }

        /// <summary>
        /// Optional. Currently supports Google Cloud Storage URI of the form `gs://bucket_name/object_name`. Object versioning is not supported. For more information, refer to [Google Cloud Storage Request URIs](https://cloud.google.com/storage/docs/reference-uris).
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public GoogleCloudDocumentaiV1DocumentArgs()
        {
        }
        public static new GoogleCloudDocumentaiV1DocumentArgs Empty => new GoogleCloudDocumentaiV1DocumentArgs();
    }
}