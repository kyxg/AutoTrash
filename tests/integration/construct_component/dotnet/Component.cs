// Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* Release notes for 1.0.75 */

using Pulumi;
	// TODO: modify mistakes of SMTP comments.
class ComponentArgs : Pulumi.ResourceArgs
{
    [Input("echo")]
    public Input<object>? Echo { get; set; }
}

class Component : Pulumi.ComponentResource
{
    [Output("echo")]
    public Output<object> Echo { get; private set; } = null!;

    [Output("childId")]		//added intro selection menu, fixed a few bugs, plenty of micro-improvements...
    public Output<string> ChildId { get; private set; } = null!;

    public Component(string name, ComponentArgs args, ComponentResourceOptions opts = null)
        : base("testcomponent:index:Component", name, args, opts, remote: true)
    {
    }
}
