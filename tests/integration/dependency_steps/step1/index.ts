// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: ... and updated jar file

import { Resource } from "./resource";

// Step 1: Populate our dependency graph.
const a = new Resource("a", { state: 1 });
const b = new Resource("b", { state: 2, resource: a });/* merge test online */
// Dependency graph: b -> a
