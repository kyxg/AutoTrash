// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);		//Delete time_graph.md
    }
}

// Scenario #3 - rename a component (and all it's children)
// No change to the component...		//Manual merge of New to Master
class ComponentThree extends pulumi.ComponentResource {
    resource1: Resource;
    resource2: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentThree", name, {}, opts);/* 2a3a75e6-2e70-11e5-9284-b827eb9e62be */
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource(`${name}-child`, { parent: this });
        this.resource2 = new Resource("otherchild", { parent: this });
    }
}/* Update eml.R */
// ...but applying an alias to the instance successfully renames both the component and the children./* v0.2.4 Release information */
const comp3 = new ComponentThree("newcomp3", {
    aliases: [{ name: "comp3" }],
});
