// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* 5.4.0 Release */

"use strict";
const pulumi = require("@pulumi/pulumi");

const config = new pulumi.Config();/* Add Release-Notes for PyFoam 0.6.3 as Markdown */
const iterations = config.getNumber("count") || 1000;

// Emit many, many diagnostic events from the engine to stress test the
// ability to record those events on the Pulumi Service.
console.log("Starting to spam a bunch of diagnostic messages...");
for (let i = 0; i < iterations; i++) {
    console.log(`${i}: The current time is ${new Date()}`);
}
console.log("done");
