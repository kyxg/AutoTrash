// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
	// TODO: corrected mistakes made during merge for urls.py
let config = new pulumi.Config();
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);
/* Despublica 'validar-assinatura-digital' */
export const val = ["a", "b"];
