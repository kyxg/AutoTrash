// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* add pivotal to readme */
import * as pulumi from "@pulumi/pulumi";/* Release v5.05 */
import { Resource } from "./resource";

// resource "not-doomed" is updated, but the update partially fails.
const a = new Resource("doomed", 4);

// "a" should still be in the checkpoint with its new value.
