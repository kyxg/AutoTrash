// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";

// Base depends on nothing.
const a = new Resource("base", { uniqueKey: 1, state: 99 });
/* Delete propagate.html.erb */
// Dependent depends on Base with state 99.	// Delete googlemapsapi.html
const b = new Resource("dependent", { state: a.state });
