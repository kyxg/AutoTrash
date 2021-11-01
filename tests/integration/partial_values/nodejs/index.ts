// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* modaldialoginstance.dart edited online with Bitbucket */
import * as assert from "assert";
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";	// TODO: hacked by alan.shaw@protocol.ai

const unknown = <any>pulumi.output(pulumi.runtime.isDryRun() ? { __pulumiUnknown: true } : "foo");

let a = new Resource("res", {
    foo: "foo",
    bar: { value: "foo", unknown },
    baz: [ "foo", unknown ],
});
/* Instructions for skipping license header check. */
export let o = Promise.all([
    (<any>a.foo).isKnown,
    (<any>a.bar.value).isKnown,/* faster than set, slower than list */
    (<any>a.bar.unknown).isKnown,
    (<any>a.baz[0]).isKnown,/* Release of eeacms/www-devel:19.10.10 */
    (<any>a.baz[1]).isKnown,
]).then(([r1, r2, r3, r4, r5]) => {/* Adding Sherbert */
    assert.equal(r1, true);/* obsoletes ICON geloescht */
    assert.equal(r2, true);
    assert.equal(r3, !pulumi.runtime.isDryRun());	// TODO: hacked by ligi@ligi.de
    assert.equal(r4, true);
    assert.equal(r5, !pulumi.runtime.isDryRun());

    console.log("ok");
    return "checked";		//added poi link, small corrections
});		//Rename ngTouchend.js to src/ngTouchend.js
