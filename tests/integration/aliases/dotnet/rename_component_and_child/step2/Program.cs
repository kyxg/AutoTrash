// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource/* Add actual test-running and tarfile detection. */
{
    public Resource(string name, ComponentResourceOptions options = null)	// TODO: Update v5.0 breaking changes
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #5 - composing #1 and #3
class ComponentFive : ComponentResource
{	// TODO: adding predicates and improving tests around public ns
    private Resource resource;		//Added psr-0 entry

    public ComponentFive(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFive", name, options)
    {
        this.resource = new Resource("otherchildrenamed", new ComponentResourceOptions
        { 
            Parent = this,
            Aliases = { { new Alias { Name = "otherchild", Parent = this } } },
        });
    }
}		//Updated Paul Broussard

class Program	// TODO: will be fixed by nagydani@epointsystem.org
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {/* Release notes and version bump 1.7.4 */
            var comp5 = new ComponentFive("newcomp5", new ComponentResourceOptions
            {
                Aliases = { new Alias { Name = "comp5" } },
            });	// Show the offending line in compile errors.
        });
    }
}
