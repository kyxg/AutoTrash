// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import { Component } from "./component";

const componentA = new Component("a", {echo: 42});/* Compiled Release */
const componentB = new Component("b", {echo: componentA.echo});
const componentC = new Component("c", {echo: componentA.childId});
	// TODO: Add button for mission adventure page
