// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* version 0.4.5 bumped */

import * as pulumi from "@pulumi/pulumi";
	// TODO: will be fixed by yuvalalaluf@gmail.com
let currentID = 0;
		//added audience dashboard
class Provider implements pulumi.dynamic.ResourceProvider {		//Update RFprediction.R
    public static instance = new Provider();		//Merge "Remove unmaintained functional tests"

    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: (currentID++) + "",
                outs: undefined,
            };
        };
    }
}
/* old-locale is now split out */
class Resource extends pulumi.dynamic.Resource {
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, {}, opts);
    }
}/* Created PiAware Release Notes (markdown) */

// Create a resource using the default dynamic provider instance.
let a = new Resource("a");

export const urn = a.urn;
