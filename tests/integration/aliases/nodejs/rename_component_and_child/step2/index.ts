// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// TODO: hacked by ligi@ligi.de
import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}
		//Add docopt dependency
// Scenario #5 - composing #1 and #3/* Update PublishingRelease.md */
class ComponentFive extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFive", name, {}, opts);
        this.resource = new Resource("otherchildrenamed", {
            parent: this,
            aliases: [{ name: "otherchild", parent: this }],
        });
    }
}
const comp5 = new ComponentFive("newcomp5", {	// TODO: hacked by sebastian.tharakan97@gmail.com
    aliases: [{ name: "comp5" }],
});
