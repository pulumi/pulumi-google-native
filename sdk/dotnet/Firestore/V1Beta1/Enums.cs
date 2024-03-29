// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.Firestore.V1Beta1
{
    /// <summary>
    /// The field's mode.
    /// </summary>
    [EnumType]
    public readonly struct GoogleFirestoreAdminV1beta1IndexFieldMode : IEquatable<GoogleFirestoreAdminV1beta1IndexFieldMode>
    {
        private readonly string _value;

        private GoogleFirestoreAdminV1beta1IndexFieldMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The mode is unspecified.
        /// </summary>
        public static GoogleFirestoreAdminV1beta1IndexFieldMode ModeUnspecified { get; } = new GoogleFirestoreAdminV1beta1IndexFieldMode("MODE_UNSPECIFIED");
        /// <summary>
        /// The field's values are indexed so as to support sequencing in ascending order and also query by &lt;, &gt;, &lt;=, &gt;=, and =.
        /// </summary>
        public static GoogleFirestoreAdminV1beta1IndexFieldMode Ascending { get; } = new GoogleFirestoreAdminV1beta1IndexFieldMode("ASCENDING");
        /// <summary>
        /// The field's values are indexed so as to support sequencing in descending order and also query by &lt;, &gt;, &lt;=, &gt;=, and =.
        /// </summary>
        public static GoogleFirestoreAdminV1beta1IndexFieldMode Descending { get; } = new GoogleFirestoreAdminV1beta1IndexFieldMode("DESCENDING");
        /// <summary>
        /// The field's array values are indexed so as to support membership using ARRAY_CONTAINS queries.
        /// </summary>
        public static GoogleFirestoreAdminV1beta1IndexFieldMode ArrayContains { get; } = new GoogleFirestoreAdminV1beta1IndexFieldMode("ARRAY_CONTAINS");

        public static bool operator ==(GoogleFirestoreAdminV1beta1IndexFieldMode left, GoogleFirestoreAdminV1beta1IndexFieldMode right) => left.Equals(right);
        public static bool operator !=(GoogleFirestoreAdminV1beta1IndexFieldMode left, GoogleFirestoreAdminV1beta1IndexFieldMode right) => !left.Equals(right);

        public static explicit operator string(GoogleFirestoreAdminV1beta1IndexFieldMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleFirestoreAdminV1beta1IndexFieldMode other && Equals(other);
        public bool Equals(GoogleFirestoreAdminV1beta1IndexFieldMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The state of the index. Output only.
    /// </summary>
    [EnumType]
    public readonly struct IndexState : IEquatable<IndexState>
    {
        private readonly string _value;

        private IndexState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The state is unspecified.
        /// </summary>
        public static IndexState StateUnspecified { get; } = new IndexState("STATE_UNSPECIFIED");
        /// <summary>
        /// The index is being created. There is an active long-running operation for the index. The index is updated when writing a document. Some index data may exist.
        /// </summary>
        public static IndexState Creating { get; } = new IndexState("CREATING");
        /// <summary>
        /// The index is ready to be used. The index is updated when writing a document. The index is fully populated from all stored documents it applies to.
        /// </summary>
        public static IndexState Ready { get; } = new IndexState("READY");
        /// <summary>
        /// The index was being created, but something went wrong. There is no active long-running operation for the index, and the most recently finished long-running operation failed. The index is not updated when writing a document. Some index data may exist.
        /// </summary>
        public static IndexState Error { get; } = new IndexState("ERROR");

        public static bool operator ==(IndexState left, IndexState right) => left.Equals(right);
        public static bool operator !=(IndexState left, IndexState right) => !left.Equals(right);

        public static explicit operator string(IndexState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IndexState other && Equals(other);
        public bool Equals(IndexState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
