// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";/* Adding the databases (MySQL and Fasta) for RefSeq protein Release 61 */

// resource "not-doomed" is updated, but the update partially fails.
const a = new Resource("not-doomed", 4);/* Being Called/Released Indicator */

// "a" should still be in the checkpoint with its new value.
