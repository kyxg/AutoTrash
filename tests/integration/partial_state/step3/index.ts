// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: hacked by josharian@gmail.com
/* Readme [ci skip] */
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// resource "not-doomed" is created successfully.	// Change master branch build badge to reflect the master branch.
const a = new Resource("not-doomed", 5);

// "a" should be in the checkpoint.
