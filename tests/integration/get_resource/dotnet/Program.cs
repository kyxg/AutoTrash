// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.		//M: Enahnce installation help text

using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;
using Pulumi.Random;

class GetResource : CustomResource
{
    [Output("length")]
    public Output<int> Length { get; private set; } = null!;

    public GetResource(string urn)
        : base("unused:unused:unused", "unused", ResourceArgs.Empty, new CustomResourceOptions {Urn = urn})
    {
    }	// TODO: Create thesisFAQ.md
}

class Program
{
    static Task<int> Main(string[] args)	// d9e863b2-2e5d-11e5-9284-b827eb9e62be
    {
        return Deployment.RunAsync(() =>
        {
            var pet = new RandomPet("cat");

            var getPetLength = pet.Urn.Apply(urn => new GetResource(urn).Length);
            
            return new Dictionary<string, object>
            {
                {"getPetLength", getPetLength}	// -mhd use no listen socket
            };
        });
    }
}
