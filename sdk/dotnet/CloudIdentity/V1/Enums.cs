// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.CloudIdentity.V1
{
    /// <summary>
    /// Resource type for the Dynamic Group Query
    /// </summary>
    [EnumType]
    public readonly struct DynamicGroupQueryResourceType : IEquatable<DynamicGroupQueryResourceType>
    {
        private readonly string _value;

        private DynamicGroupQueryResourceType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default value (not valid)
        /// </summary>
        public static DynamicGroupQueryResourceType ResourceTypeUnspecified { get; } = new DynamicGroupQueryResourceType("RESOURCE_TYPE_UNSPECIFIED");
        /// <summary>
        /// For queries on User
        /// </summary>
        public static DynamicGroupQueryResourceType User { get; } = new DynamicGroupQueryResourceType("USER");

        public static bool operator ==(DynamicGroupQueryResourceType left, DynamicGroupQueryResourceType right) => left.Equals(right);
        public static bool operator !=(DynamicGroupQueryResourceType left, DynamicGroupQueryResourceType right) => !left.Equals(right);

        public static explicit operator string(DynamicGroupQueryResourceType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DynamicGroupQueryResourceType other && Equals(other);
        public bool Equals(DynamicGroupQueryResourceType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Inbound SSO behavior.
    /// </summary>
    [EnumType]
    public readonly struct InboundSsoAssignmentSsoMode : IEquatable<InboundSsoAssignmentSsoMode>
    {
        private readonly string _value;

        private InboundSsoAssignmentSsoMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Not allowed.
        /// </summary>
        public static InboundSsoAssignmentSsoMode SsoModeUnspecified { get; } = new InboundSsoAssignmentSsoMode("SSO_MODE_UNSPECIFIED");
        /// <summary>
        /// Disable SSO for the targeted users.
        /// </summary>
        public static InboundSsoAssignmentSsoMode SsoOff { get; } = new InboundSsoAssignmentSsoMode("SSO_OFF");
        /// <summary>
        /// Use an external SAML Identity Provider for SSO for the targeted users.
        /// </summary>
        public static InboundSsoAssignmentSsoMode SamlSso { get; } = new InboundSsoAssignmentSsoMode("SAML_SSO");
        /// <summary>
        /// Use the domain-wide SAML Identity Provider for the targeted users if one is configured; otherwise, this is equivalent to `SSO_OFF`. Note that this will also be equivalent to `SSO_OFF` if/when support for domain-wide SAML is removed. Google may disallow this mode at that point and existing assignments with this mode may be automatically changed to `SSO_OFF`.
        /// </summary>
        public static InboundSsoAssignmentSsoMode DomainWideSamlIfEnabled { get; } = new InboundSsoAssignmentSsoMode("DOMAIN_WIDE_SAML_IF_ENABLED");

        public static bool operator ==(InboundSsoAssignmentSsoMode left, InboundSsoAssignmentSsoMode right) => left.Equals(right);
        public static bool operator !=(InboundSsoAssignmentSsoMode left, InboundSsoAssignmentSsoMode right) => !left.Equals(right);

        public static explicit operator string(InboundSsoAssignmentSsoMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is InboundSsoAssignmentSsoMode other && Equals(other);
        public bool Equals(InboundSsoAssignmentSsoMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// When to redirect sign-ins to the IdP.
    /// </summary>
    [EnumType]
    public readonly struct SignInBehaviorRedirectCondition : IEquatable<SignInBehaviorRedirectCondition>
    {
        private readonly string _value;

        private SignInBehaviorRedirectCondition(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default and means "always"
        /// </summary>
        public static SignInBehaviorRedirectCondition RedirectConditionUnspecified { get; } = new SignInBehaviorRedirectCondition("REDIRECT_CONDITION_UNSPECIFIED");
        /// <summary>
        /// Sign-in flows where the user is prompted for their identity will not redirect to the IdP (so the user will most likely be prompted by Google for a password), but special flows like IdP-initiated SAML and sign-in following automatic redirection to the IdP by domain-specific service URLs will accept the IdP's assertion of the user's identity.
        /// </summary>
        public static SignInBehaviorRedirectCondition Never { get; } = new SignInBehaviorRedirectCondition("NEVER");

        public static bool operator ==(SignInBehaviorRedirectCondition left, SignInBehaviorRedirectCondition right) => left.Equals(right);
        public static bool operator !=(SignInBehaviorRedirectCondition left, SignInBehaviorRedirectCondition right) => !left.Equals(right);

        public static explicit operator string(SignInBehaviorRedirectCondition value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SignInBehaviorRedirectCondition other && Equals(other);
        public bool Equals(SignInBehaviorRedirectCondition other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
