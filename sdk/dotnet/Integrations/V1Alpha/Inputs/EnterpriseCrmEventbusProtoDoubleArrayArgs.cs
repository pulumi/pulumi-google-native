// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha.Inputs
{

    public sealed class EnterpriseCrmEventbusProtoDoubleArrayArgs : global::Pulumi.ResourceArgs
    {
        [Input("values")]
        private InputList<double>? _values;
        public InputList<double> Values
        {
            get => _values ?? (_values = new InputList<double>());
            set => _values = value;
        }

        public EnterpriseCrmEventbusProtoDoubleArrayArgs()
        {
        }
        public static new EnterpriseCrmEventbusProtoDoubleArrayArgs Empty => new EnterpriseCrmEventbusProtoDoubleArrayArgs();
    }
}
