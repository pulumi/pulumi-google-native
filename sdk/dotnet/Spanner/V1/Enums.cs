// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.Spanner.V1
{
    /// <summary>
    /// Optional. The dialect of the Cloud Spanner Database.
    /// </summary>
    [EnumType]
    public readonly struct DatabaseDatabaseDialect : IEquatable<DatabaseDatabaseDialect>
    {
        private readonly string _value;

        private DatabaseDatabaseDialect(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default value. This value will create a database with the GOOGLE_STANDARD_SQL dialect.
        /// </summary>
        public static DatabaseDatabaseDialect DatabaseDialectUnspecified { get; } = new DatabaseDatabaseDialect("DATABASE_DIALECT_UNSPECIFIED");
        /// <summary>
        /// Google standard SQL.
        /// </summary>
        public static DatabaseDatabaseDialect GoogleStandardSql { get; } = new DatabaseDatabaseDialect("GOOGLE_STANDARD_SQL");
        /// <summary>
        /// PostgreSQL supported SQL.
        /// </summary>
        public static DatabaseDatabaseDialect Postgresql { get; } = new DatabaseDatabaseDialect("POSTGRESQL");

        public static bool operator ==(DatabaseDatabaseDialect left, DatabaseDatabaseDialect right) => left.Equals(right);
        public static bool operator !=(DatabaseDatabaseDialect left, DatabaseDatabaseDialect right) => !left.Equals(right);

        public static explicit operator string(DatabaseDatabaseDialect value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DatabaseDatabaseDialect other && Equals(other);
        public bool Equals(DatabaseDatabaseDialect other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}