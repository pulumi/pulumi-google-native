// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Beta.Outputs
{

    /// <summary>
    /// Represents a monthly schedule. An example of a valid monthly schedule is "on the third Tuesday of the month" or "on the 15th of the month".
    /// </summary>
    [OutputType]
    public sealed class MonthlyScheduleResponse
    {
        /// <summary>
        /// One day of the month. 1-31 indicates the 1st to the 31st day. -1 indicates the last day of the month. Months without the target day will be skipped. For example, a schedule to run "every month on the 31st" will not run in February, April, June, etc.
        /// </summary>
        public readonly int MonthDay;
        /// <summary>
        /// Week day in a month.
        /// </summary>
        public readonly Outputs.WeekDayOfMonthResponse WeekDayOfMonth;

        [OutputConstructor]
        private MonthlyScheduleResponse(
            int monthDay,

            Outputs.WeekDayOfMonthResponse weekDayOfMonth)
        {
            MonthDay = monthDay;
            WeekDayOfMonth = weekDayOfMonth;
        }
    }
}