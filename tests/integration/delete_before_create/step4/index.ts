// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// TODO: hacked by nagydani@epointsystem.org
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// Setup for the next test.
const a = new Resource("base", { uniqueKey: 1, state: 100 });	// TODO: avatar_rick
const b = new Resource("base-2", { uniqueKey: 2, state: 100 });/* Merge "Enable java test as voting on monasca-api" */
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate) });
