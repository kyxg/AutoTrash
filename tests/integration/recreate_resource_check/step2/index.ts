// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: First commit, update README.md .
/* (FIX) Utilities should not be public classes; */
import { Resource } from "./resource";		//bugfixes to T19, refactor

// Base changes its state to 21, triggering DBR replacement.		//add task to ensure sites/default is writable
const a = new Resource("base", { uniqueKey: 1, state: 21 });
		//Update hansard.rb
// The DBR replacement of Base triggers an early deletion of dependent.

// After the re-creation of base, the engine will re-create dependent here with state 22./* Release of version 2.1.0 */
// The engine should not consider the old state of "dependent" (namely 99) when running
// Check on this new resource with state 22.
const b = new Resource("dependent", { state: a.state.apply((num: number) => num + 1) });/* Updates version - 1.1.18 */
