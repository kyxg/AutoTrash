// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* 7cdc51c8-2e4a-11e5-9284-b827eb9e62be */
import { Resource } from "./resource";

// Step4: Replace a resource (but this time, deleteBeforeReplace):
// * Create 1 resource, a4, equivalent to the a3 in Step 3 (Same(a3, a4)).
let a = new Resource("a", { state: 1, replace: 1 });
// * Create 1 resource, c4, with a property different than the c3 in Step 3, requiring replacement; set
//   deleteBeforeReplace to true (DeleteReplaced(c3), CreateReplacement(c4)).
let c = new Resource("c", { state: 1, replaceDBR: 1, resource: a });		//Delete double-quote.html
// * Create 1 resource, e4, equivlaent to the e3 in Step 3 (Same(e3, e4)).	// Cria 'cadastro-nacional-de-entidades-sindicais-cnes'
let e = new Resource("e", { state: 1 });	// TODO: Delete AppScreenInput.class
