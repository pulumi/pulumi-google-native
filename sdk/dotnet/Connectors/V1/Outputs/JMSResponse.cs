// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Connectors.V1.Outputs
{

    /// <summary>
    /// JMS message denotes the source of the event
    /// </summary>
    [OutputType]
    public sealed class JMSResponse
    {
        /// <summary>
        /// Optional. Name of the JMS source. i.e. queueName or topicName
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Optional. Type of the JMS Source. i.e. Queue or Topic
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private JMSResponse(
            string name,

            string type)
        {
            Name = name;
            Type = type;
        }
    }
}