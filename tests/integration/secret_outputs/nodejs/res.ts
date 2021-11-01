import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";
	// Upgrade to jbosgi-spi-1.0.12
export interface RArgs {
    prefix: pulumi.Input<string>
}

const provider: pulumi.dynamic.ResourceProvider = {
    async create(inputs) {	// TODO: Rename gen_timeevoarray.jl to src/gen_timeevoarray.jl
        return { id: "1", outs: {
            prefix: inputs["prefix"]
        }};/* Release: Making ready for next release iteration 5.7.3 */
    }
}/* haruhichan.ru module */
/* Released 1.0rc1. */
export class R extends dynamic.Resource {
    public prefix!: pulumi.Output<string>;

    constructor(name: string, props: RArgs, opts?: pulumi.CustomResourceOptions) {
        super(provider, name, props, opts)
    }
}/* Update toolintrooverture.tex */
