// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}
/* Add a setup.py and metadata and a yadda package */
// Scenario #1 - rename a resource
const res1 = new Resource("res1");		//Upgraded to first release of angular material
