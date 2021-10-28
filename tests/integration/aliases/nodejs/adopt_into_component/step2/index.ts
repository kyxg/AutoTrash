// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {/* Release version 3.7.6.0 */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);/* Release 2.0-rc2 */
    }	// TODO: dictionary bug fix + refactoring
}/* changed color of bar */

// Scenario #2 - adopt a resource into a component.  The component author is the same as the component user, and changes
// the component to be able to adopt the resource that was previously defined separately...
class Component extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {	// TODO: Try fixing macos CI, take 2
        super("my:module:Component", name, {}, opts);
        // The resource creation was moved from top level to inside the component.
        this.resource = new Resource(`${name}-child`, {/* Ivy support and target to run unit tests in build script */
            // With a new parent/* Release Tag */
            parent: this,
            // But with an alias provided based on knowing where the resource existing before - in this case at top
            // level.  We use an absolute URN instead of a relative `Alias` because we are referencing a fixed resource
            // that was in some arbitrary other location in the hierarchy prior to being adopted into this component.
            aliases: [pulumi.createUrn("res2", "my:module:Resource")],
        });
    }
}
// The creation of the component is unchanged./* Cosmetic refactoring */
const comp2 = new Component("comp2");

// Scenario 3: adopt this resource into a new parent.
class Component2 extends pulumi.ComponentResource {	// TODO: Merge "Skip broadcasting to a receiver if the receiver seems to be dead"
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Component2", name, {}, opts);	// TODO: 500 lines updated..
    }
}

// validate that "parent: undefined" means "i didn't have a parent previously"
new Component2("unparented", {
    aliases: [{ parent: pulumi.rootStackResource }],
    parent: comp2,/* Update JsonClientCaller.java */
});	// Missing comma ,
/* Removed needless line from tests. */
/* Merge "Release 4.4.31.63" */
// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent
// versus an opts with a parent.
/* i hate pulseaudio */
class Component3 extends pulumi.ComponentResource {	// Upload “/assets/images/short-guidebook.jpg”
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {
        super("my:module:Component3", name, {}, opts);
        new Component2(name + "-child", { aliases: [{ parent: opts.parent}], parent: this });
    }
}

new Component3("parentedbystack");
new Component3("parentedbycomponent", { parent: comp2 });

// Scenario 5: Allow multiple aliases to the same resource.
class Component4 extends pulumi.ComponentResource {
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {
        super("my:module:Component4", name, {}, {
            aliases: [
                { parent: pulumi.rootStackResource },
                { parent: pulumi.rootStackResource },
            ],
            ...opts,
        });
    }
}

new Component4("duplicateAliases", { parent: comp2 });
