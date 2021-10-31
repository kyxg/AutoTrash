// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource/* uhub: link against librt when eglibc is enabled */
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {		//Switched to AESLightEngine to minimise cache timing side-channel leaks.
    }
}/* Merge "Fix Mutable default argument" */

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {
            // Scenario #1 - rename a resource
            var res1 = new Resource("res1");		//ProcessorFactory fixed.
        });	// TODO: hacked by steven@stebalien.com
    }
}/* Release 2.5.0-beta-2: update sitemap */
