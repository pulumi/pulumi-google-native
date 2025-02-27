// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1Beta1.Inputs
{

    /// <summary>
    /// Contains multiple sensitive information findings for each resource slice.
    /// </summary>
    public sealed class DetailArgs : global::Pulumi.ResourceArgs
    {
        [Input("findings")]
        private InputList<Inputs.FindingArgs>? _findings;
        public InputList<Inputs.FindingArgs> Findings
        {
            get => _findings ?? (_findings = new InputList<Inputs.FindingArgs>());
            set => _findings = value;
        }

        public DetailArgs()
        {
        }
        public static new DetailArgs Empty => new DetailArgs();
    }
}
