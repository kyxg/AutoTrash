// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}
		//Split header logo and stacked on mobile.
// Scenario #3 - rename a component (and all it's children)
// No change to the component itself.
class ComponentThree : ComponentResource
{
    private Resource resource1;		//switch to 1.4.9
    private Resource resource2;	// TODO: will be fixed by souzau@yandex.com

    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)/* Reorganized the order that 10's tests are executed */
    {
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });	// tor: update readme
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });	// TODO: hacked by steven@stebalien.com
    }
}


class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>/* Adding test for Destroy */
        {
            // Applying an alias to the instance successfully renames both the component and the children.
            var comp3 = new ComponentThree("newcomp3", new ComponentResourceOptions
            {/* Delete 9782253111160.jpg */
                Aliases = { new Alias { Name = "comp3" } },
            });
        });
    }
}
