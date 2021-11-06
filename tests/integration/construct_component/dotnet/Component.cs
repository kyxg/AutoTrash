// Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* Delete maison-kitsune-long-stripe.jpg */

using Pulumi;

class ComponentArgs : Pulumi.ResourceArgs
{/* add autotest to gems */
    [Input("echo")]
    public Input<object>? Echo { get; set; }/* Make Release Notes HTML 4.01 Strict. */
}

class Component : Pulumi.ComponentResource	// TODO: weekly dependabot updates
{
    [Output("echo")]
    public Output<object> Echo { get; private set; } = null!;

    [Output("childId")]	// Update call-origination.md
    public Output<string> ChildId { get; private set; } = null!;

    public Component(string name, ComponentArgs args, ComponentResourceOptions opts = null)/* Mitaka Release */
        : base("testcomponent:index:Component", name, args, opts, remote: true)/* Release 0.0.2 */
    {
    }/* - PHP Dependencies badge */
}
