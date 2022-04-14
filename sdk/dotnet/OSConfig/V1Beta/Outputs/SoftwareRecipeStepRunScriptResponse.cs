// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Beta.Outputs
{

    /// <summary>
    /// Runs a script through an interpreter.
    /// </summary>
    [OutputType]
    public sealed class SoftwareRecipeStepRunScriptResponse
    {
        /// <summary>
        /// Return codes that indicate that the software installed or updated successfully. Behaviour defaults to [0]
        /// </summary>
        public readonly ImmutableArray<int> AllowedExitCodes;
        /// <summary>
        /// The script interpreter to use to run the script. If no interpreter is specified the script is executed directly, which likely only succeed for scripts with [shebang lines](https://en.wikipedia.org/wiki/Shebang_\(Unix\)).
        /// </summary>
        public readonly string Interpreter;
        /// <summary>
        /// The shell script to be executed.
        /// </summary>
        public readonly string Script;

        [OutputConstructor]
        private SoftwareRecipeStepRunScriptResponse(
            ImmutableArray<int> allowedExitCodes,

            string interpreter,

            string script)
        {
            AllowedExitCodes = allowedExitCodes;
            Interpreter = interpreter;
            Script = script;
        }
    }
}