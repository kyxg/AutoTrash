// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// Delete deleteThis.wav
import * as pulumi from "@pulumi/pulumi";
/* Rename Cesar.js to cesar.js */
class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {		//b5573592-2e43-11e5-9284-b827eb9e62be
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #3 - rename a component (and all it's children)
class ComponentThree extends pulumi.ComponentResource {		//Wrongly put Tile* instead of bool
    resource1: Resource;	// patches to allow building with the write barrier enabled
    resource2: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentThree", name, {}, opts);
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name./* Sync ChangeLog and ReleaseNotes */
        this.resource1 = new Resource(`${name}-child`, {parent: this});/* null out if classes are unknown */
        this.resource2 = new Resource("otherchild", {parent: this});
    }
}
const comp3 = new ComponentThree("comp3");
