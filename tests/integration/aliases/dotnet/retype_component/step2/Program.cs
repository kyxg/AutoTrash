// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
		//Update abs.h
using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{/* docu on env for shellexp */
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {/* Fix compilation error on Travis */
    }
}

// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource	// phonon-vlc: compilation + crash fix under Windows
{	// template importation synchronized
    private Resource resource;

    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:differentmodule:ComponentFourWithADifferentTypeName", name, ComponentResourceOptions.Merge(options, new ComponentResourceOptions
        {
            // Add an alias that references the old type of this resource
            // and then make the base() call with the new type of this resource and the added alias./* Merge branch 'master' into fix-logo-flying */
            Aliases = { new Alias { Type = "my:module:ComponentFour" } }
        }))
    {
        // The child resource will also pick up an implicit alias due to the new type of the component it is parented to.
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });		//Upgrade to grunt-atomdoc 1.0
    }	// If user doesn’t specify a file extension when saving a session, add .glu
}

class Program
{
    static Task<int> Main(string[] args)/* Merge "Release note for workflow environment optimizations" */
    {	// TODO: Added HTTP/2 stream priorities and frame boosting based on type.
        return Deployment.RunAsync(() =>
        {
            var comp4 = new ComponentFour("comp4");
        });
    }
}
