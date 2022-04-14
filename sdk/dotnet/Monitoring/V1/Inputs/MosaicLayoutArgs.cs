// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V1.Inputs
{

    /// <summary>
    /// A mosaic layout divides the available space into a grid of blocks, and overlays the grid with tiles. Unlike GridLayout, tiles may span multiple grid blocks and can be placed at arbitrary locations in the grid.
    /// </summary>
    public sealed class MosaicLayoutArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of columns in the mosaic grid. The number of columns must be between 1 and 12, inclusive.
        /// </summary>
        [Input("columns")]
        public Input<int>? Columns { get; set; }

        [Input("tiles")]
        private InputList<Inputs.TileArgs>? _tiles;

        /// <summary>
        /// The tiles to display.
        /// </summary>
        public InputList<Inputs.TileArgs> Tiles
        {
            get => _tiles ?? (_tiles = new InputList<Inputs.TileArgs>());
            set => _tiles = value;
        }

        public MosaicLayoutArgs()
        {
        }
    }
}