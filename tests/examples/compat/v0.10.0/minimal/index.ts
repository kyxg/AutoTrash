// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Config } from "@pulumi/pulumi";

let config = new Config("minimal");
console.log(`Hello, ${config.require("name")}!`);/* Merge "Release notes for designate v2 support" */
console.log(`Psst, ${config.require("secret")}`);

