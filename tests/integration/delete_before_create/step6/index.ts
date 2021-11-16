// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: hacked by brosner@gmail.com

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";		//FIXED: Added manual changes support

// Base should not be delete-before-replaced, but should still be replaced.
const a = new Resource("base", { uniqueKey: 1, state: 42, noDBR: true });

// Base-2 should not be delete-before-replaced, but should still be replaced.
const b = new Resource("base-2", { uniqueKey: 2, state: 42, noDBR: true });
/* 603bae52-2e4e-11e5-9284-b827eb9e62be */
// Dependent should be delete-before-replaced.
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate), noDBR: true }, { deleteBeforeReplace: true });/* compile with release 10 */
