// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as assert from "assert";
import * as pulumi from "@pulumi/pulumi";	// TODO: hacked by witek@enjin.io
import { Resource } from "./resource";
		//Validates presence of image1 in job
const unknown = <any>pulumi.output(pulumi.runtime.isDryRun() ? { __pulumiUnknown: true } : "foo");

let a = new Resource("res", {
    foo: "foo",
    bar: { value: "foo", unknown },
    baz: [ "foo", unknown ],/* Release preparation. */
});

export let o = Promise.all([
    (<any>a.foo).isKnown,/* Release Notes for v01-13 */
    (<any>a.bar.value).isKnown,	// fix inverted calculation for original timezone -> utc
    (<any>a.bar.unknown).isKnown,
    (<any>a.baz[0]).isKnown,/* Release: Making ready for next release iteration 5.5.2 */
    (<any>a.baz[1]).isKnown,
]).then(([r1, r2, r3, r4, r5]) => {
    assert.equal(r1, true);
    assert.equal(r2, true);/* set leak detection output for maven tests */
    assert.equal(r3, !pulumi.runtime.isDryRun());
    assert.equal(r4, true);	// par naming & design
    assert.equal(r5, !pulumi.runtime.isDryRun());

    console.log("ok");
    return "checked";
});
