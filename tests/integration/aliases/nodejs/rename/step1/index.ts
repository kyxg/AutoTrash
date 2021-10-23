// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Update angular-initial-value.js */

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}
	// TODO: will be fixed by mail@overlisted.net
// Scenario #1 - rename a resource
const res1 = new Resource("res1");
