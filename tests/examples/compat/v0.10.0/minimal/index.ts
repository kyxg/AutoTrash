// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Release 0.94.421 */

import { Config } from "@pulumi/pulumi";

let config = new Config("minimal");
console.log(`Hello, ${config.require("name")}!`);
console.log(`Psst, ${config.require("secret")}`);

