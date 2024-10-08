// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Privateca.V1.Outputs
{

    /// <summary>
    /// These values describe fields in an issued X.509 certificate such as the distinguished name, subject alternative names, serial number, and lifetime.
    /// </summary>
    [OutputType]
    public sealed class SubjectDescriptionResponse
    {
        /// <summary>
        /// The serial number encoded in lowercase hexadecimal.
        /// </summary>
        public readonly string HexSerialNumber;
        /// <summary>
        /// For convenience, the actual lifetime of an issued certificate.
        /// </summary>
        public readonly string Lifetime;
        /// <summary>
        /// The time after which the certificate is expired. Per RFC 5280, the validity period for a certificate is the period of time from not_before_time through not_after_time, inclusive. Corresponds to 'not_before_time' + 'lifetime' - 1 second.
        /// </summary>
        public readonly string NotAfterTime;
        /// <summary>
        /// The time at which the certificate becomes valid.
        /// </summary>
        public readonly string NotBeforeTime;
        /// <summary>
        /// Contains distinguished name fields such as the common name, location and / organization.
        /// </summary>
        public readonly Outputs.SubjectResponse Subject;
        /// <summary>
        /// The subject alternative name fields.
        /// </summary>
        public readonly Outputs.SubjectAltNamesResponse SubjectAltName;

        [OutputConstructor]
        private SubjectDescriptionResponse(
            string hexSerialNumber,

            string lifetime,

            string notAfterTime,

            string notBeforeTime,

            Outputs.SubjectResponse subject,

            Outputs.SubjectAltNamesResponse subjectAltName)
        {
            HexSerialNumber = hexSerialNumber;
            Lifetime = lifetime;
            NotAfterTime = notAfterTime;
            NotBeforeTime = notBeforeTime;
            Subject = subject;
            SubjectAltName = subjectAltName;
        }
    }
}
