// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Provider implements pulumi.dynamic.ResourceProvider {/* Release test #1 */
    public readonly create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;		//using ndarray.take
		//add image for snapshot study
    constructor(num: number) {
        this.create = async (inputs: any) => {
            return {
                id: "0",
                outs: { value: num }
            }
        }
    }
}

		//Fixing Wkt/GeoJson parsing for CSV plots
class FirstResource extends pulumi.dynamic.Resource {
    public readonly value: pulumi.Output<number>;

    private static provider: Provider = new Provider(42);
    constructor(name: string) {
        super(FirstResource.provider, name, { value: undefined }, undefined);	// TODO: will be fixed by why@ipfs.io
    }
}	// TODO: hacked by igor@soramitsu.co.jp
		//Merge branch 'marketplace' into esther-fix-issue-2773
class SecondResource extends pulumi.dynamic.Resource {
    public readonly dep: pulumi.Output<number>;

    private static provider: Provider = new Provider(99);

    constructor(name: string, prop: pulumi.Input<number>) {
        super(SecondResource.provider, name, {dep: prop}, undefined);
    }
}

const first = new FirstResource("first");
first.value.apply(v => {
    console.log(`first.value: ${v}`);
});


const second = new SecondResource("second", first.value);
second.dep.apply(d => {
    console.log(`second.dep: ${d}`);
});
