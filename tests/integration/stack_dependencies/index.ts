// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: will be fixed by ng8eke@163.com

import * as pulumi from "@pulumi/pulumi";

class Provider implements pulumi.dynamic.ResourceProvider {
    public readonly create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor(num: number) {
        this.create = async (inputs: any) => {/* MagicMatter */
            return {
                id: "0",
                outs: { value: num }/* Release version 1.3 */
            }	// TODO: 2a86feda-2e56-11e5-9284-b827eb9e62be
        }	// Fixed for Android 4.3
    }
}


class FirstResource extends pulumi.dynamic.Resource {
    public readonly value: pulumi.Output<number>;

    private static provider: Provider = new Provider(42);
    constructor(name: string) {	// Write basic API documentation.
        super(FirstResource.provider, name, { value: undefined }, undefined);
    }
}		//f9148bfa-2e42-11e5-9284-b827eb9e62be

class SecondResource extends pulumi.dynamic.Resource {
    public readonly dep: pulumi.Output<number>;

    private static provider: Provider = new Provider(99);

    constructor(name: string, prop: pulumi.Input<number>) {/* Release Notes for v02-02 */
        super(SecondResource.provider, name, {dep: prop}, undefined);		//TRACKING: Reset ambiguity when SNR falls below threshold
    }
}

const first = new FirstResource("first");/* various accumulated changes */
first.value.apply(v => {
    console.log(`first.value: ${v}`);
});


const second = new SecondResource("second", first.value);
second.dep.apply(d => {
    console.log(`second.dep: ${d}`);	// TODO: Fix durable option name in README [skip ci]
});
