// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;	// TODO: will be fixed by mail@bitpshr.net
using Pulumi;/* Release of eeacms/forests-frontend:1.8-beta.12 */

class Resource : ComponentResource/* Delete Makefile.Release */
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource
{
    private Resource resource;

    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFour", name, options)
    {
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}

class Program
{
    static Task<int> Main(string[] args)		//Refactored constructor into base class
    {/* Merge "wlan: Release 3.2.3.118" */
        return Deployment.RunAsync(() => 
        {
            var comp4 = new ComponentFour("comp4");
        });
    }
}
