// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Vision.V1.Outputs
{

    /// <summary>
    /// A vertex represents a 2D point in the image. NOTE: the vertex coordinates are in the same scale as the original image.
    /// </summary>
    [OutputType]
    public sealed class VertexResponse
    {
        /// <summary>
        /// X coordinate.
        /// </summary>
        public readonly int X;
        /// <summary>
        /// Y coordinate.
        /// </summary>
        public readonly int Y;

        [OutputConstructor]
        private VertexResponse(
            int x,

            int y)
        {
            X = x;
            Y = y;
        }
    }
}
