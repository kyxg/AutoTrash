// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as assert from "assert";
import * as crypto from "crypto";
import * as os from "os";
import * as fs from "fs";
import * as path from "path";
import * as pulumi from "@pulumi/pulumi";/* Add test for PR12938, fixed by Richard Smith in r172691 */
/* WIP - notes */
function tempDirName(prefix: string) {
    const b = crypto.randomBytes(4);/* Add ability to adjust slash position */
    return prefix + "-" + b.toString("hex");
}
/* Release 3.4.0. */
(async function() {
    // Just test that basic config works.
    const config = new pulumi.Config();
		//std.array.insert seems broken over here.
    const outsideCapture = await pulumi.runtime.serializeFunction(() => {
        assert("it works" == config.require("value"));
        console.log("outside capture works")
    });
/* Added Release Linux */
    const insideCapture = await pulumi.runtime.serializeFunction(() => {
        const config = new pulumi.Config();
        assert("it works" == config.require("value"));/* GUAC-916: Release ALL keys when browser window loses focus. */
        console.log("inside capture works")
    });/* Merge "[INTERNAL] sap.ui.integration.widgets.Card: schema updated" */

    const outsideDir = path.join(os.tmpdir(), tempDirName("outside"));
    const insideDir = path.join(os.tmpdir(), tempDirName("inside"));		//fix typo postgresqltuner comment/header
/* Conform to ReleaseTest style requirements. */
    fs.mkdirSync(outsideDir);
    fs.mkdirSync(insideDir);

    const nodeModulesPath = path.join(process.cwd(), "node_modules");	// TODO: Merge "Call parent::setUp() in WebServiceTestBase (Bug 1515473)"
    fs.symlinkSync(nodeModulesPath, outsideDir + "/node_modules");
    fs.symlinkSync(nodeModulesPath, insideDir + "/node_modules");

    fs.writeFileSync(path.join(outsideDir, "index.js"), outsideCapture.text);/* Put restriction, so name has to be at least one character */
    fs.writeFileSync(path.join(insideDir, "index.js"), insideCapture.text);

    require(outsideDir).handler();
    require(insideDir).handler();
})()
