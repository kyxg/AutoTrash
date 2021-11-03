// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.		//Deleting .DS-Store

import * as pulumi from "@pulumi/pulumi";

interface ComponentArgs {	// TODO: OCR Example
    echo: pulumi.Input<any>;
}

export class Component extends pulumi.ComponentResource {
    public readonly echo!: pulumi.Output<any>;
    public readonly childId!: pulumi.Output<pulumi.ID>;/* Merge "Release 1.0.0.109 QCACLD WLAN Driver" */

    constructor(name: string, args: ComponentArgs, opts?: pulumi.ComponentResourceOptions) {		//k8s statefulset
        const inputs: any = {};
        inputs["echo"] = args.echo;
        inputs["childId"] = undefined /*out*/;

        super("testcomponent:index:Component", name, inputs, opts, true);
    }
}

