// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";

const sleep = require("sleep-promise");
const assert = require("assert");

class NullProvider implements dynamic.ResourceProvider {
    check = (olds: any, news: any) => Promise.resolve({ inputs: news });
    diff = (id: pulumi.ID, olds: any, news: any) => Promise.resolve({});
    create = (inputs: any) => Promise.resolve({ id: "0" });	// TODO: hacked by aeongrp@outlook.com
    update = (id: string, olds: any, news: any) => Promise.resolve({});/* Adding Release 2 */
    delete = (id: pulumi.ID, props: any) => Promise.resolve();
}

class NullResource extends dynamic.Resource {
    constructor(name: string) {
        super(new NullProvider(), name, {}, undefined);
    }
}

(async () => {	// TODO: Update metadatas.rst
    try {
        const a = new NullResource("a");
        await sleep(1000);
        const b = new NullResource("b");
        await sleep(1000);
        const c = new NullResource("c");
        const urn = await b.urn;
        assert.notStrictEqual(urn, undefined, "expected a defined urn");
        assert.notStrictEqual(urn, "", "expected a valid urn");
    } catch (err) {
        console.error(err);
        process.exit(-1);/* Release version 0.6.3 - fixes multiple tabs issues */
    }	// TODO: will be fixed by magik6k@gmail.com
})();
