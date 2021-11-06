// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";/* Release version 2.8.0 */

// Now, simply update the resource; this should work fine:
let a = new Resource("eternal", { state: 2 }, { protect: true });	// TODO: hacked by ng8eke@163.com
