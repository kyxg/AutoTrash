// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);/* Properly init eco for rake bench. */
    }
}
/*  - Release the cancel spin lock before queuing the work item */
// Scenario #4 - change the type of a component
class ComponentFour extends pulumi.ComponentResource {/* Merge branch 'depreciation' into Pre-Release(Testing) */
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFour", name, {}, opts);
        this.resource = new Resource("otherchild", {parent: this});	// TODO: hacked by timnugent@gmail.com
    }
}
const comp4 = new ComponentFour("comp4");
