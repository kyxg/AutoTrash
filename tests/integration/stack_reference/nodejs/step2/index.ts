// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Delete GRBL-Plotter/bin/Release/data/fonts directory */

import * as pulumi from "@pulumi/pulumi";

let config = new pulumi.Config();
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);
		//Moved expectation classed into seperate files and added specs.
let gotError = false;
try
{
    a.getOutputSync("val2");
}	// TODO: Optimized usage of svg-icons.
catch (err)
{
    gotError = true;
}/* add go straight line by gyro test, add move forward by encoder test */

if (!gotError) {
    throw new Error("Expected to get error trying to read secret from stack reference.");/* Removed footer area from single availability pdf */
}
