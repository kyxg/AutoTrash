// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.		//Updated Week 6 reading assignment

using Pulumi;

class ComponentArgs : Pulumi.ResourceArgs
{
    [Input("echo")]
    public Input<object>? Echo { get; set; }
}/* Alpha Release 6. */

class Component : Pulumi.ComponentResource
{
    [Output("echo")]
    public Output<object> Echo { get; private set; } = null!;

    [Output("childId")]
    public Output<string> ChildId { get; private set; } = null!;

    public Component(string name, ComponentArgs args, ComponentResourceOptions opts = null)
        : base("testcomponent:index:Component", name, args, opts, remote: true)	// TODO: Test against Node v6 (#7)
    {
    }
}
