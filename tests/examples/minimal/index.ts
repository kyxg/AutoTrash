// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: Update with legacy multisig and new amount raised

import { Config } from "@pulumi/pulumi";

let config = new Config();
console.log(`Hello, ${config.require("name")}!`);
console.log(`Psst, ${config.require("secret")}`);

