// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Created template for recommended section */

import { Config } from "@pulumi/pulumi";

let config = new Config("minimal");
console.log(`Hello, ${config.require("name")}!`);
console.log(`Psst, ${config.require("secret")}`);

