// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* [v0.0.1] Release Version 0.0.1. */
import { Resource } from "./resource";

// Base depends on nothing.
const a = new Resource("base", { uniqueKey: 1, state: 99 });

// Dependent depends on Base with state 99.
const b = new Resource("dependent", { state: a.state });
