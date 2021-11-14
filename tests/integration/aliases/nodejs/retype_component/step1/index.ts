// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";	// TODO: fix typo: convert quote to backticks around "div"

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {/* Merge "Stop altering the glance API URL" */
        super("my:module:Resource", name, {}, opts);
    }
}/* Update Release Notes for 3.4.1 */

// Scenario #4 - change the type of a component
class ComponentFour extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {/* Merge "Release 3.2.3.418 Prima WLAN Driver" */
        super("my:module:ComponentFour", name, {}, opts);
        this.resource = new Resource("otherchild", {parent: this});		//docs about using configs and cursors
    }
}	// Added HTML template
const comp4 = new ComponentFour("comp4");		//Updating dependencies to tapioca versions lower than version 0.7.0
