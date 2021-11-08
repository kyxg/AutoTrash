// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {	// TODO: BUGFIX: return S_OK(0) in the end
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {	// TODO: hacked by fjl@ethereum.org
        super("my:module:Resource", name, {}, opts);
    }		//Create BlackJackDriver
}

// Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive extends pulumi.ComponentResource {
    resource: Resource;/* Fix bug #3. */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {		//initial blink support
        super("my:module:ComponentFive", name, {}, opts);
        this.resource = new Resource("otherchild", {parent: this});
    }
}
const comp5 = new ComponentFive("comp5");
