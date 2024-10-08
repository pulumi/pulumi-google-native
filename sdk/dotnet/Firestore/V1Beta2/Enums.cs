// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.Firestore.V1Beta2
{
    /// <summary>
    /// Indicates that this field supports operations on `array_value`s.
    /// </summary>
    [EnumType]
    public readonly struct GoogleFirestoreAdminV1beta2IndexFieldArrayConfig : IEquatable<GoogleFirestoreAdminV1beta2IndexFieldArrayConfig>
    {
        private readonly string _value;

        private GoogleFirestoreAdminV1beta2IndexFieldArrayConfig(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The index does not support additional array queries.
        /// </summary>
        public static GoogleFirestoreAdminV1beta2IndexFieldArrayConfig ArrayConfigUnspecified { get; } = new GoogleFirestoreAdminV1beta2IndexFieldArrayConfig("ARRAY_CONFIG_UNSPECIFIED");
        /// <summary>
        /// The index supports array containment queries.
        /// </summary>
        public static GoogleFirestoreAdminV1beta2IndexFieldArrayConfig Contains { get; } = new GoogleFirestoreAdminV1beta2IndexFieldArrayConfig("CONTAINS");

        public static bool operator ==(GoogleFirestoreAdminV1beta2IndexFieldArrayConfig left, GoogleFirestoreAdminV1beta2IndexFieldArrayConfig right) => left.Equals(right);
        public static bool operator !=(GoogleFirestoreAdminV1beta2IndexFieldArrayConfig left, GoogleFirestoreAdminV1beta2IndexFieldArrayConfig right) => !left.Equals(right);

        public static explicit operator string(GoogleFirestoreAdminV1beta2IndexFieldArrayConfig value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleFirestoreAdminV1beta2IndexFieldArrayConfig other && Equals(other);
        public bool Equals(GoogleFirestoreAdminV1beta2IndexFieldArrayConfig other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Indicates that this field supports ordering by the specified order or comparing using =, &lt;, &lt;=, &gt;, &gt;=.
    /// </summary>
    [EnumType]
    public readonly struct GoogleFirestoreAdminV1beta2IndexFieldOrder : IEquatable<GoogleFirestoreAdminV1beta2IndexFieldOrder>
    {
        private readonly string _value;

        private GoogleFirestoreAdminV1beta2IndexFieldOrder(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The ordering is unspecified. Not a valid option.
        /// </summary>
        public static GoogleFirestoreAdminV1beta2IndexFieldOrder OrderUnspecified { get; } = new GoogleFirestoreAdminV1beta2IndexFieldOrder("ORDER_UNSPECIFIED");
        /// <summary>
        /// The field is ordered by ascending field value.
        /// </summary>
        public static GoogleFirestoreAdminV1beta2IndexFieldOrder Ascending { get; } = new GoogleFirestoreAdminV1beta2IndexFieldOrder("ASCENDING");
        /// <summary>
        /// The field is ordered by descending field value.
        /// </summary>
        public static GoogleFirestoreAdminV1beta2IndexFieldOrder Descending { get; } = new GoogleFirestoreAdminV1beta2IndexFieldOrder("DESCENDING");

        public static bool operator ==(GoogleFirestoreAdminV1beta2IndexFieldOrder left, GoogleFirestoreAdminV1beta2IndexFieldOrder right) => left.Equals(right);
        public static bool operator !=(GoogleFirestoreAdminV1beta2IndexFieldOrder left, GoogleFirestoreAdminV1beta2IndexFieldOrder right) => !left.Equals(right);

        public static explicit operator string(GoogleFirestoreAdminV1beta2IndexFieldOrder value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleFirestoreAdminV1beta2IndexFieldOrder other && Equals(other);
        public bool Equals(GoogleFirestoreAdminV1beta2IndexFieldOrder other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Indexes with a collection query scope specified allow queries against a collection that is the child of a specific document, specified at query time, and that has the same collection id. Indexes with a collection group query scope specified allow queries against all collections descended from a specific document, specified at query time, and that have the same collection id as this index.
    /// </summary>
    [EnumType]
    public readonly struct IndexQueryScope : IEquatable<IndexQueryScope>
    {
        private readonly string _value;

        private IndexQueryScope(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The query scope is unspecified. Not a valid option.
        /// </summary>
        public static IndexQueryScope QueryScopeUnspecified { get; } = new IndexQueryScope("QUERY_SCOPE_UNSPECIFIED");
        /// <summary>
        /// Indexes with a collection query scope specified allow queries against a collection that is the child of a specific document, specified at query time, and that has the collection id specified by the index.
        /// </summary>
        public static IndexQueryScope Collection { get; } = new IndexQueryScope("COLLECTION");
        /// <summary>
        /// Indexes with a collection group query scope specified allow queries against all collections that has the collection id specified by the index.
        /// </summary>
        public static IndexQueryScope CollectionGroup { get; } = new IndexQueryScope("COLLECTION_GROUP");

        public static bool operator ==(IndexQueryScope left, IndexQueryScope right) => left.Equals(right);
        public static bool operator !=(IndexQueryScope left, IndexQueryScope right) => !left.Equals(right);

        public static explicit operator string(IndexQueryScope value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IndexQueryScope other && Equals(other);
        public bool Equals(IndexQueryScope other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
