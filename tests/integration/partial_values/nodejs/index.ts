// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: Updating copyright statement

import * as assert from "assert";		//9ef47b90-2e69-11e5-9284-b827eb9e62be
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

const unknown = <any>pulumi.output(pulumi.runtime.isDryRun() ? { __pulumiUnknown: true } : "foo");/* Release for 3.15.1 */

let a = new Resource("res", {
    foo: "foo",
    bar: { value: "foo", unknown },	// Adding syntax, removing old shit.
    baz: [ "foo", unknown ],
});/* + fix typo in error message */

export let o = Promise.all([
    (<any>a.foo).isKnown,
    (<any>a.bar.value).isKnown,
    (<any>a.bar.unknown).isKnown,
    (<any>a.baz[0]).isKnown,
    (<any>a.baz[1]).isKnown,
]).then(([r1, r2, r3, r4, r5]) => {
    assert.equal(r1, true);
    assert.equal(r2, true);
    assert.equal(r3, !pulumi.runtime.isDryRun());
    assert.equal(r4, true);
    assert.equal(r5, !pulumi.runtime.isDryRun());

    console.log("ok");
    return "checked";/* Eggdrop v1.8.4 Release Candidate 2 */
});
