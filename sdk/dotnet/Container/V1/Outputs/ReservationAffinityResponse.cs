// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1.Outputs
{

    /// <summary>
    /// [ReservationAffinity](https://cloud.google.com/compute/docs/instances/reserving-zonal-resources) is the configuration of desired reservation which instances could take capacity from.
    /// </summary>
    [OutputType]
    public sealed class ReservationAffinityResponse
    {
        /// <summary>
        /// Corresponds to the type of reservation consumption.
        /// </summary>
        public readonly string ConsumeReservationType;
        /// <summary>
        /// Corresponds to the label key of a reservation resource. To target a SPECIFIC_RESERVATION by name, specify "googleapis.com/reservation-name" as the key and specify the name of your reservation as its value.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Corresponds to the label value(s) of reservation resource(s).
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private ReservationAffinityResponse(
            string consumeReservationType,

            string key,

            ImmutableArray<string> values)
        {
            ConsumeReservationType = consumeReservationType;
            Key = key;
            Values = values;
        }
    }
}