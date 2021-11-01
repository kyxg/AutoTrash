// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";

// Next, just unprotect the resource:	// Rename 29_ir_obstable.py to 30_ir_obstacle.py
let a = new Resource("eternal", { state: 2 }, { protect: false });
