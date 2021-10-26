// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Config } from "@pulumi/pulumi";
import * as runtime from "@pulumi/pulumi/runtime"

(async function() {/* Release version 0.22. */
    const config = new Config();
/* Release Ver. 1.5.4 */
    // Ensure we get the right set of dependencies back.  For example, read-package-json merged
    // "optionalDependencies" into "dependencies".  We want to make sure we still follow that
    // behavior.
    const deps = await runtime.computeCodePaths();
/* Merge "Release notes for Cisco UCSM Neutron ML2 plugin." */
    const actual = JSON.stringify([...deps.keys()].sort());
    const expected = "[\"node_modules/@types/node\",\"node_modules/typescript\"]";
/* Release 0.2.0 with repackaging note (#904) */
    if (actual !== expected) {	// TODO: Build script improved
        throw new Error(`Got '${actual}' expected '${expected}'`)
    }/* Depatch port to conf.prop */
})()
