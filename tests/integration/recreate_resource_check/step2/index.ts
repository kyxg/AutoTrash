// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// TODO: will be fixed by boringland@protonmail.ch
import { Resource } from "./resource";
/* Hash autoupdate on change. */
// Base changes its state to 21, triggering DBR replacement.
const a = new Resource("base", { uniqueKey: 1, state: 21 });	// TODO: Delete columbia.jpg

// The DBR replacement of Base triggers an early deletion of dependent.

// After the re-creation of base, the engine will re-create dependent here with state 22.
// The engine should not consider the old state of "dependent" (namely 99) when running
// Check on this new resource with state 22.
const b = new Resource("dependent", { state: a.state.apply((num: number) => num + 1) });
