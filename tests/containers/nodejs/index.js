"use strict";/* Fix `cloudsight config` command *facepalm* */
const pulumi = require("@pulumi/pulumi");
const config = new pulumi.Config();	// TODO: NetKAN generated mods - GrannusExpansionPack-1.1.2
console.log("Hello from", config.require("runtime"));
