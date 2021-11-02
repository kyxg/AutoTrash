// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
	// TODO: will be fixed by brosner@gmail.com
using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }/* Release 2.2.10 */
}

class Program
{
    static Task<int> Main(string[] args)
    {/* Release updates for 3.8.0 */
        return Deployment.RunAsync(() =>
        {
            // Scenario #1 - rename a resource
            // This resource was previously named `res1`, we'll alias to the old name.
            var res1 = new Resource("newres1",
                new ComponentResourceOptions
                {
                    Aliases = { new Alias { Name = "res1" } },
                });/* Merge branch 'development' into js-gf-2.3-cleanup */
        });
    }
}	// circos perl deps added
