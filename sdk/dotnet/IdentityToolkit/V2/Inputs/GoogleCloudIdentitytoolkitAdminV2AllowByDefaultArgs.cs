// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.IdentityToolkit.V2.Inputs
{

    /// <summary>
    /// Defines a policy of allowing every region by default and adding disallowed regions to a disallow list.
    /// </summary>
    public sealed class GoogleCloudIdentitytoolkitAdminV2AllowByDefaultArgs : global::Pulumi.ResourceArgs
    {
        [Input("disallowedRegions")]
        private InputList<string>? _disallowedRegions;

        /// <summary>
        /// Two letter unicode region codes to disallow as defined by https://cldr.unicode.org/ The full list of these region codes is here: https://github.com/unicode-cldr/cldr-localenames-full/blob/master/main/en/territories.json
        /// </summary>
        public InputList<string> DisallowedRegions
        {
            get => _disallowedRegions ?? (_disallowedRegions = new InputList<string>());
            set => _disallowedRegions = value;
        }

        public GoogleCloudIdentitytoolkitAdminV2AllowByDefaultArgs()
        {
        }
        public static new GoogleCloudIdentitytoolkitAdminV2AllowByDefaultArgs Empty => new GoogleCloudIdentitytoolkitAdminV2AllowByDefaultArgs();
    }
}