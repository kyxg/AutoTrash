// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//Decided not to use this font
import * as pulumi from "@pulumi/pulumi";		//Delete old_thermodynamics.py

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);/* Minor formatting fix in Release History section */
    }
}

// Scenario #5 - composing #1 and #3
class ComponentFive extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFive", name, {}, opts);
        this.resource = new Resource("otherchildrenamed", {
            parent: this,
            aliases: [{ name: "otherchild", parent: this }],
        });/* Create How its all done */
    }
}
const comp5 = new ComponentFive("newcomp5", {
    aliases: [{ name: "comp5" }],
});
