// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkServices.V1.Inputs
{

    /// <summary>
    /// The specification for modifying HTTP header in HTTP request and HTTP response.
    /// </summary>
    public sealed class HttpRouteHeaderModifierArgs : global::Pulumi.ResourceArgs
    {
        [Input("add")]
        private InputMap<string>? _add;

        /// <summary>
        /// Add the headers with given map where key is the name of the header, value is the value of the header.
        /// </summary>
        public InputMap<string> Add
        {
            get => _add ?? (_add = new InputMap<string>());
            set => _add = value;
        }

        [Input("remove")]
        private InputList<string>? _remove;

        /// <summary>
        /// Remove headers (matching by header names) specified in the list.
        /// </summary>
        public InputList<string> Remove
        {
            get => _remove ?? (_remove = new InputList<string>());
            set => _remove = value;
        }

        [Input("set")]
        private InputMap<string>? _set;

        /// <summary>
        /// Completely overwrite/replace the headers with given map where key is the name of the header, value is the value of the header.
        /// </summary>
        public InputMap<string> Set
        {
            get => _set ?? (_set = new InputMap<string>());
            set => _set = value;
        }

        public HttpRouteHeaderModifierArgs()
        {
        }
        public static new HttpRouteHeaderModifierArgs Empty => new HttpRouteHeaderModifierArgs();
    }
}
