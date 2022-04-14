// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataLabeling.V1Beta1
{
    /// <summary>
    /// Creates an instruction for how data should be labeled.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:datalabeling/v1beta1:Instruction")]
    public partial class Instruction : Pulumi.CustomResource
    {
        /// <summary>
        /// The names of any related resources that are blocking changes to the instruction.
        /// </summary>
        [Output("blockingResources")]
        public Output<ImmutableArray<string>> BlockingResources { get; private set; } = null!;

        /// <summary>
        /// Creation time of instruction.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Deprecated: this instruction format is not supported any more. Instruction from a CSV file, such as for classification task. The CSV file should have exact two columns, in the following format: * The first column is labeled data, such as an image reference, text. * The second column is comma separated labels associated with data.
        /// </summary>
        [Output("csvInstruction")]
        public Output<Outputs.GoogleCloudDatalabelingV1beta1CsvInstructionResponse> CsvInstruction { get; private set; } = null!;

        /// <summary>
        /// The data type of this instruction.
        /// </summary>
        [Output("dataType")]
        public Output<string> DataType { get; private set; } = null!;

        /// <summary>
        /// Optional. User-provided description of the instruction. The description can be up to 10000 characters long.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The display name of the instruction. Maximum of 64 characters.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Instruction resource name, format: projects/{project_id}/instructions/{instruction_id}
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Instruction from a PDF document. The PDF should be in a Cloud Storage bucket.
        /// </summary>
        [Output("pdfInstruction")]
        public Output<Outputs.GoogleCloudDatalabelingV1beta1PdfInstructionResponse> PdfInstruction { get; private set; } = null!;

        /// <summary>
        /// Last update time of instruction.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Instruction resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instruction(string name, InstructionArgs args, CustomResourceOptions? options = null)
            : base("google-native:datalabeling/v1beta1:Instruction", name, args ?? new InstructionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instruction(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:datalabeling/v1beta1:Instruction", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Instruction resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instruction Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Instruction(name, id, options);
        }
    }

    public sealed class InstructionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Deprecated: this instruction format is not supported any more. Instruction from a CSV file, such as for classification task. The CSV file should have exact two columns, in the following format: * The first column is labeled data, such as an image reference, text. * The second column is comma separated labels associated with data.
        /// </summary>
        [Input("csvInstruction")]
        public Input<Inputs.GoogleCloudDatalabelingV1beta1CsvInstructionArgs>? CsvInstruction { get; set; }

        /// <summary>
        /// The data type of this instruction.
        /// </summary>
        [Input("dataType", required: true)]
        public Input<Pulumi.GoogleNative.DataLabeling.V1Beta1.InstructionDataType> DataType { get; set; } = null!;

        /// <summary>
        /// Optional. User-provided description of the instruction. The description can be up to 10000 characters long.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the instruction. Maximum of 64 characters.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Instruction from a PDF document. The PDF should be in a Cloud Storage bucket.
        /// </summary>
        [Input("pdfInstruction")]
        public Input<Inputs.GoogleCloudDatalabelingV1beta1PdfInstructionArgs>? PdfInstruction { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        public InstructionArgs()
        {
        }
    }
}