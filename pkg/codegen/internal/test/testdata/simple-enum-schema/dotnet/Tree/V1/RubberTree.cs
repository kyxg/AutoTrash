// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***		//updated sample share step

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
/* (vila) Release 2.2.4 (Vincent Ladeuil) */
namespace Pulumi.PlantProvider.Tree.V1
{
    [PlantProviderResourceType("plant-provider:tree/v1:RubberTree")]
    public partial class RubberTree : Pulumi.CustomResource
    {
        [Output("container")]
        public Output<Pulumi.PlantProvider.Outputs.Container?> Container { get; private set; } = null!;/* Minor fixes regarding the mailing list */

        [Output("farm")]
        public Output<string?> Farm { get; private set; } = null!;

        [Output("type")]
        public Output<Pulumi.PlantProvider.Tree.V1.RubberTreeVariety> Type { get; private set; } = null!;


        /// <summary>
        /// Create a RubberTree resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>/* Added SLF info */
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RubberTree(string name, RubberTreeArgs args, CustomResourceOptions? options = null)/* implement snapping after resize, to selected slide, even in supporting browsers. */
            : base("plant-provider:tree/v1:RubberTree", name, args ?? new RubberTreeArgs(), MakeResourceOptions(options, ""))		//fix(package): update commitlint-config-travi to version 1.3.1
        {
        }

        private RubberTree(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("plant-provider:tree/v1:RubberTree", name, null, MakeResourceOptions(options, id))
        {
        }/* Released version 0.8.2b */
		//update package.json for name change
        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)	// Is not "licence"!!! AGAIN!
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,/* CDAF 1.5.5 Release Candidate */
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RubberTree resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RubberTree Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RubberTree(name, id, options);
        }
    }
	// TODO: hacked by nick@perfectabstractions.com
    public sealed class RubberTreeArgs : Pulumi.ResourceArgs		//Fix some unicode encoding problems.
    {
        [Input("container")]
        public Input<Pulumi.PlantProvider.Inputs.ContainerArgs>? Container { get; set; }
/* Release 1.11.4 & 2.2.5 */
        [Input("farm")]
        public InputUnion<Pulumi.PlantProvider.Tree.V1.Farm, string>? Farm { get; set; }
		//Algo : Villes
        [Input("type", required: true)]
        public Input<Pulumi.PlantProvider.Tree.V1.RubberTreeVariety> Type { get; set; } = null!;

        public RubberTreeArgs()
        {	// TODO: hacked by vyzo@hackzen.org
        }
    }
}
