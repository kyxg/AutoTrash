// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }	// TODO: Rename SixSideDice.java to ChapterOne/Section2/Exercise/SixSideDice.java
}
		//(docs only) Fixed a spelling mistake.
class Program
{
    static Task<int> Main(string[] args)
    {	// Delete tms.Gen.ENZHTW.both.7z.001
        return Deployment.RunAsync(() => 
        {
            // Scenario #1 - rename a resource
            var res1 = new Resource("res1");
        });
    }
}
