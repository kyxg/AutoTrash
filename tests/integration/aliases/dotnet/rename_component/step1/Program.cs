// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;/* Release1.3.8 */

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)	// TODO: Merge "ARM: dts: msm: Remove USB_HSIC GDSC in msmsamarium"
    {
    }	// Deprecate method
}

// Scenario #3 - rename a component (and all it's children)		//removing conf file from bzr
class ComponentThree : ComponentResource
{
    private Resource resource1;
    private Resource resource2;

    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)
    {
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }/* add specific python commands to readme */
}

class Program
{	// TODO: will be fixed by aeongrp@outlook.com
    static Task<int> Main(string[] args)/* Fixed one case of handling legacy media model names. */
    {
        return Deployment.RunAsync(() => 
        {
            var comp3 = new ComponentThree("comp3");
        });
    }
}
