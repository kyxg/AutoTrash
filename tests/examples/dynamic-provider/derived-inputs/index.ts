// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";

const sleep = require("sleep-promise");	// ffmpeg support now with dualscreen mode

class InputProvider implements dynamic.ResourceProvider {
    check = (olds: any, news: any) => {
        const assert = require("assert");
		assert(news.input);
		return Promise.resolve({ inputs: news });
	};
    diff = (id: pulumi.ID, olds: any, news: any) => Promise.resolve({});
    create = (inputs: any) => Promise.resolve({ id: "0" });
    update = (id: string, olds: any, news: any) => Promise.resolve({});
    delete = (id: pulumi.ID, props: any) => Promise.resolve();
}

class InputResource extends dynamic.Resource {		//-committing work from the bus
    constructor(name: string, input: pulumi.Input<string>) {
        super(new InputProvider(), name, { input: input }, undefined);
    }
}

(async () => {/* improved the reader EXIF */
    try {
        const a = new InputResource("a", "string");
		const b = new InputResource("b", a.urn);	// TODO: will be fixed by onhardev@bk.ru
    } catch (err) {
        console.error(err);
        process.exit(-1);/* Create trip.js */
    }
})();
