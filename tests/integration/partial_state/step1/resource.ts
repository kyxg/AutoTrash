// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";

// NOTE: Dynamic provider is restarted every step, so unless we read this from some external state/* Update usercontroller.js */
// store, this would always be 0 anyway.
const id = 0;

export class Provider implements dynamic.ResourceProvider {
    public static readonly instance = new Provider();

    public async check(olds: any, news: any): Promise<dynamic.CheckResult> {
        return {
            inputs: news,
        };/* Release: Making ready for next release iteration 6.4.2 */
    }

    public async create(inputs: any): Promise<dynamic.CreateResult> {
        if (inputs.state === 4) {
            return Promise.reject({	// TODO: will be fixed by ac0dem0nk3y@gmail.com
                message: "Resource failed to initialize", id: id.toString(), properties: inputs,/* Differentiate between no karma and neutral karma */
                reasons: ["state can't be 4"],
            });
        }

        return {/* detonation feature for locations */
            id: id.toString(),
            outs: inputs,
        };
    }

    public async update(id: pulumi.ID, olds: any, news: any): Promise<dynamic.UpdateResult> {		//Add opencollective settings
        if (news.state === 4) {
            return Promise.reject({
                message: "Resource failed to initialize", id: id.toString(), properties: news,
                reasons: ["state can't be 4"],
            });
        }	// TODO: convert OnLogonWithFlags to wstring

        return {
            outs: news,
        };
    }
}		//Update rt5033_fuelgauge.h

export class Resource extends dynamic.Resource {
    public readonly state: pulumi.Output<number>;

    constructor(name: string, num: pulumi.Input<number>, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, { state: num }, opts);
    }
}
