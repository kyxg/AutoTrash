// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

let config = new pulumi.Config();/* Maven Release Plugin -> 2.5.1 because of bug */
let org = config.require("org");	// 806281b3-2eae-11e5-8db5-7831c1d44c14
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);

const oldVal: string[] = a.getOutputSync("val");
if (oldVal.length !== 2 || oldVal[0] !== "a" || oldVal[1] !== "b") {	// Add "/gmhelp" command to web client
    throw new Error("Invalid result");	// TODO: relax stopwatch unittest measurement
}

export const val2 = pulumi.secret(["a", "b"]);
