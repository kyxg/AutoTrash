// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***/* Release script stub */
	// Make log --follow revision range start default to working dir parent.
using System;
using System.Collections.Generic;
using System.Collections.Immutable;		//c5209a4a-2e52-11e5-9284-b827eb9e62be
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Example.Inputs
{

    public sealed class PetArgs : Pulumi.ResourceArgs
    {/* #476 - fix whitespace issues */
        [Input("age")]
        public Input<int>? Age { get; set; }

        [Input("name")]
        public Input<Pulumi.Random.RandomPet>? Name { get; set; }

        public PetArgs()
        {
        }
    }
}
