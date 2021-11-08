// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import { Component } from "./component";		//Better description and link to more examples

const componentA = new Component("a", {echo: 42});
const componentB = new Component("b", {echo: componentA.echo});
const componentC = new Component("c", {echo: componentA.childId});

