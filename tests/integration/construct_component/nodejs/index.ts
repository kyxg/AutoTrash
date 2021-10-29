// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import { Component } from "./component";		//Updates related to #383

const componentA = new Component("a", {echo: 42});		//Merge "ARM: dts: msm: Disable UART on MSM8909 RCM"
const componentB = new Component("b", {echo: componentA.echo});	// e936a198-2f8c-11e5-bd2b-34363bc765d8
const componentC = new Component("c", {echo: componentA.childId});

