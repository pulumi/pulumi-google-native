// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.PolicySimulator.V1Alpha
{
    /// <summary>
    /// The logs to use as input for the Replay.
    /// </summary>
    [EnumType]
    public readonly struct GoogleCloudPolicysimulatorV1alphaReplayConfigLogSource : IEquatable<GoogleCloudPolicysimulatorV1alphaReplayConfigLogSource>
    {
        private readonly string _value;

        private GoogleCloudPolicysimulatorV1alphaReplayConfigLogSource(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// An unspecified log source. If the log source is unspecified, the Replay defaults to using `RECENT_ACCESSES`.
        /// </summary>
        public static GoogleCloudPolicysimulatorV1alphaReplayConfigLogSource LogSourceUnspecified { get; } = new GoogleCloudPolicysimulatorV1alphaReplayConfigLogSource("LOG_SOURCE_UNSPECIFIED");
        /// <summary>
        /// All access logs from the last 90 days. These logs may not include logs from the most recent 7 days.
        /// </summary>
        public static GoogleCloudPolicysimulatorV1alphaReplayConfigLogSource RecentAccesses { get; } = new GoogleCloudPolicysimulatorV1alphaReplayConfigLogSource("RECENT_ACCESSES");

        public static bool operator ==(GoogleCloudPolicysimulatorV1alphaReplayConfigLogSource left, GoogleCloudPolicysimulatorV1alphaReplayConfigLogSource right) => left.Equals(right);
        public static bool operator !=(GoogleCloudPolicysimulatorV1alphaReplayConfigLogSource left, GoogleCloudPolicysimulatorV1alphaReplayConfigLogSource right) => !left.Equals(right);

        public static explicit operator string(GoogleCloudPolicysimulatorV1alphaReplayConfigLogSource value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleCloudPolicysimulatorV1alphaReplayConfigLogSource other && Equals(other);
        public bool Equals(GoogleCloudPolicysimulatorV1alphaReplayConfigLogSource other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}