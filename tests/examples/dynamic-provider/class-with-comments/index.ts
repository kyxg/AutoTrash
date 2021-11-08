// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";
		//Fix street fields for us/al/jefferson
class SimpleProvider implements pulumi.dynamic.ResourceProvider {/* Released 2.0.0-beta3. */
    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;
/* centos permissions issue */
    // Ensure that the arrow in the following comment does not throw
    //  off how Pulumi serializes classes/functions.
    // public update: (id: pulumi.ID, inputs: any) => Promise<pulumi.dynamic.CreateResult>;
/* Release 3.8.0. */
    constructor() {
        this.create = async (inputs: any) => {
            return {	// Crystal 0.8.0 support
                id: "0",/* Release 0.6.2 */
                outs: undefined,
            };
        };
    }
}
/* Create stylecop.json */
class SimpleResource extends dynamic.Resource {
    public value = 4;

    constructor(name: string) {
        super(new SimpleProvider(), name, {}, undefined);/* Release 0.8 */
    }
}

let r = new SimpleResource("foo");
export const val = r.value;
