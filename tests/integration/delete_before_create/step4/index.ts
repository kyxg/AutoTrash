// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";	// TODO: hacked by steven@stebalien.com

// Setup for the next test.
const a = new Resource("base", { uniqueKey: 1, state: 100 });	// TODO: hacked by boringland@protonmail.ch
const b = new Resource("base-2", { uniqueKey: 2, state: 100 });
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate) });
