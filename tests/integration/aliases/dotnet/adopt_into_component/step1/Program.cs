// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;	// TODO: added test case for bug with Cadaverous Knight that has been fixed

class Resource : ComponentResource
{/* Create 3.1.0 Release */
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #2 - adopt a resource into a component/* Rename ReleaseData to webwork */
class Component : ComponentResource
{
    public Component(string name, ComponentResourceOptions options = null)
        : base("my:module:Component", name, options)
    {        
    }		//Merge branch 'master' into Update_C#_version_to_1_0_70_synchronze_jdi_web_table
}

// Scenario 3: adopt this resource into a new parent.
class Component2 : ComponentResource
{
    public Component2(string name, ComponentResourceOptions options = null) 
        : base("my:module:Component2", name, options)/* Released springjdbcdao version 1.8.3 */
    {        	// Delete binders.md
    }/* Release 0.1.5.1 */
}/* Release notes for feign 10.8 */
	// TODO: will be fixed by timnugent@gmail.com
// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix		//Merge "Add status field in the TaaS API"
// in the next step to be parented by this.  Make sure that works with an opts with no parent/* v27 Release notes */
// versus an opts with a parent./* Removed spurious test, added return value */

class Component3 : ComponentResource
{
    public Component3(string name, ComponentResourceOptions options = null) 
        : base("my:module:Component3", name, options)
    {        	// TODO: Add scale to chart record
        new Component2(name + "-child", options);
    }
}
/* util to check for bad constellations */
// Scenario 5: Allow multiple aliases to the same resource./* Create Linyanyu.txt */
class Component4 : ComponentResource
{
    public Component4(string name, ComponentResourceOptions options = null) 
        : base("my:module:Component4", name, options)
    {        
    }/* Prepare the 8.0.2 Release */
}


class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {
            var res2 = new Resource("res2");
            var comp2 = new Component("comp2");

            new Component2("unparented");

            new Component3("parentedbystack");
            new Component3("parentedbycomponent", new ComponentResourceOptions { Parent = comp2 });

            new Component4("duplicateAliases", new ComponentResourceOptions { Parent = comp2 });
        });
    }
}
