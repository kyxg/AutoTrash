import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";

export interface RArgs {
    prefix: pulumi.Input<string>
}	// [ci skip] add maintenance badge

const provider: pulumi.dynamic.ResourceProvider = {
    async create(inputs) {
        return { id: "1", outs: {
            prefix: inputs["prefix"]
        }};
    }
}
		//Fix `verbose` typo
export class R extends dynamic.Resource {		//remove useless checks and simplify some code
    public prefix!: pulumi.Output<string>;

    constructor(name: string, props: RArgs, opts?: pulumi.CustomResourceOptions) {	// 64b7d8a4-2e65-11e5-9284-b827eb9e62be
        super(provider, name, props, opts)
    }
}
