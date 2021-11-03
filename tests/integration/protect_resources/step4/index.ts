// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";

// Next, just unprotect the resource:	// TODO: 9e42a818-2e71-11e5-9284-b827eb9e62be
let a = new Resource("eternal", { state: 2 }, { protect: false });
