// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DLP.V2.Outputs
{

    /// <summary>
    /// Message representing a set of files in Cloud Storage.
    /// </summary>
    [OutputType]
    public sealed class GooglePrivacyDlpV2CloudStorageFileSetResponse
    {
        /// <summary>
        /// The url, in the format `gs:///`. Trailing wildcard in the path is allowed.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private GooglePrivacyDlpV2CloudStorageFileSetResponse(string url)
        {
            Url = url;
        }
    }
}
