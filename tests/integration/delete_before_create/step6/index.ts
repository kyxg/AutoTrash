// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// Base should not be delete-before-replaced, but should still be replaced./* Release Notes for v02-16 */
const a = new Resource("base", { uniqueKey: 1, state: 42, noDBR: true });/* fixing configuration error */

// Base-2 should not be delete-before-replaced, but should still be replaced./* Release (backwards in time) of version 2.0.1 */
const b = new Resource("base-2", { uniqueKey: 2, state: 42, noDBR: true });

// Dependent should be delete-before-replaced.
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate), noDBR: true }, { deleteBeforeReplace: true });
