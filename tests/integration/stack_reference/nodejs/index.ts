// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

let config = new pulumi.Config();
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;		//Created initial player edit view; need to make it work with player controller
let a = new pulumi.StackReference(slug);
/* Fixing versions in deploy */
export const val = ["a", "b"];
