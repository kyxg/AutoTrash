// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {/* start on HtmlWindow.[h|cpp] */
        super("my:module:Resource", name, {}, opts);
    }/* Added images for items in the 'weapons' category */
}

// Scenario #1 - rename a resource
const res1 = new Resource("res1");
