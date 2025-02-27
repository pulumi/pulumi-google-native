// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta
{
    public static class GetHttpHealthCheck
    {
        /// <summary>
        /// Returns the specified HttpHealthCheck resource.
        /// </summary>
        public static Task<GetHttpHealthCheckResult> InvokeAsync(GetHttpHealthCheckArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetHttpHealthCheckResult>("google-native:compute/beta:getHttpHealthCheck", args ?? new GetHttpHealthCheckArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified HttpHealthCheck resource.
        /// </summary>
        public static Output<GetHttpHealthCheckResult> Invoke(GetHttpHealthCheckInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetHttpHealthCheckResult>("google-native:compute/beta:getHttpHealthCheck", args ?? new GetHttpHealthCheckInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetHttpHealthCheckArgs : global::Pulumi.InvokeArgs
    {
        [Input("httpHealthCheck", required: true)]
        public string HttpHealthCheck { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetHttpHealthCheckArgs()
        {
        }
        public static new GetHttpHealthCheckArgs Empty => new GetHttpHealthCheckArgs();
    }

    public sealed class GetHttpHealthCheckInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("httpHealthCheck", required: true)]
        public Input<string> HttpHealthCheck { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetHttpHealthCheckInvokeArgs()
        {
        }
        public static new GetHttpHealthCheckInvokeArgs Empty => new GetHttpHealthCheckInvokeArgs();
    }


    [OutputType]
    public sealed class GetHttpHealthCheckResult
    {
        /// <summary>
        /// How often (in seconds) to send a health check. The default value is 5 seconds.
        /// </summary>
        public readonly int CheckIntervalSec;
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
        /// </summary>
        public readonly int HealthyThreshold;
        /// <summary>
        /// The value of the host header in the HTTP health check request. If left empty (default value), the public IP on behalf of which this health check is performed will be used.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// Type of the resource. Always compute#httpHealthCheck for HTTP health checks.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The TCP port number for the HTTP health check request. The default value is 80.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// The request path of the HTTP health check request. The default value is /. This field does not support query parameters. Must comply with RFC3986.
        /// </summary>
        public readonly string RequestPath;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have greater value than checkIntervalSec.
        /// </summary>
        public readonly int TimeoutSec;
        /// <summary>
        /// A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
        /// </summary>
        public readonly int UnhealthyThreshold;

        [OutputConstructor]
        private GetHttpHealthCheckResult(
            int checkIntervalSec,

            string creationTimestamp,

            string description,

            int healthyThreshold,

            string host,

            string kind,

            string name,

            int port,

            string requestPath,

            string selfLink,

            int timeoutSec,

            int unhealthyThreshold)
        {
            CheckIntervalSec = checkIntervalSec;
            CreationTimestamp = creationTimestamp;
            Description = description;
            HealthyThreshold = healthyThreshold;
            Host = host;
            Kind = kind;
            Name = name;
            Port = port;
            RequestPath = requestPath;
            SelfLink = selfLink;
            TimeoutSec = timeoutSec;
            UnhealthyThreshold = unhealthyThreshold;
        }
    }
}
