// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Pubsublite.V1.Outputs
{

    /// <summary>
    /// The settings for this topic's Reservation usage.
    /// </summary>
    [OutputType]
    public sealed class ReservationConfigResponse
    {
        /// <summary>
        /// The Reservation to use for this topic's throughput capacity. Structured like: projects/{project_number}/locations/{location}/reservations/{reservation_id}
        /// </summary>
        public readonly string ThroughputReservation;

        [OutputConstructor]
        private ReservationConfigResponse(string throughputReservation)
        {
            ThroughputReservation = throughputReservation;
        }
    }
}
