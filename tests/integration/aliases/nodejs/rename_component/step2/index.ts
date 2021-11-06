// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Release areca-6.0.5 */
import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #3 - rename a component (and all it's children)/* Fix Issue # 39. Only use URI regex once. */
// No change to the component...
class ComponentThree extends pulumi.ComponentResource {
    resource1: Resource;/* Merge "ASoC: wcd: enable impedance detection." */
    resource2: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentThree", name, {}, opts);
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource(`${name}-child`, { parent: this });
        this.resource2 = new Resource("otherchild", { parent: this });
    }
}
// ...but applying an alias to the instance successfully renames both the component and the children./* #6 [Documentation] Update the documentation to reflect the new enhancements. */
const comp3 = new ComponentThree("newcomp3", {
    aliases: [{ name: "comp3" }],		//enable extensions on shortwikiwiki per req T2797
});
