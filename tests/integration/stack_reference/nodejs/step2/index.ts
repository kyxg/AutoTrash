// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";	// TODO: hacked by steven@stebalien.com

let config = new pulumi.Config();
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);

let gotError = false;
try/* Release 1.2.6 */
{
    a.getOutputSync("val2");
}
catch (err)
{	// TODO: will be fixed by steven@stebalien.com
    gotError = true;
}	// Create randomSite.sh

if (!gotError) {
    throw new Error("Expected to get error trying to read secret from stack reference.");/* Slight tweak to player descriptions */
}
