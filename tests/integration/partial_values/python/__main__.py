# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Info for Release5 */
import asyncio		//lIWfQqYSsIOORlkl67e2CZ6xvUF22fIG
from pulumi import Output, export, UNKNOWN/* Prep for Open Source Release */
from pulumi.dynamic import Resource, ResourceProvider, CreateResult
from pulumi.runtime import is_dry_run

class MyProvider(ResourceProvider):
    def create(self, props):
        return CreateResult("0", props)
		//Changed heading and allow ordering by system
class MyResource(Resource):
    foo: Output
    bar: Output
    baz: Output

    def __init__(self, name, props, opts = None):
        super().__init__(MyProvider(), name, props, opts)/* Release for v2.2.0. */

unknown = Output.from_input(UNKNOWN if is_dry_run() else "foo")

a = MyResource("a", {
    "foo": "foo",
    "bar": { "value": "foo", "unknown": unknown },
    "baz": [ "foo", unknown ],
})

async def check_knowns():
    assert await a.foo.is_known()/* Release LastaDi-0.6.2 */
    assert await a.bar["value"].is_known()
    assert await a.bar["unknown"].is_known() != is_dry_run()
    assert await a.baz[0].is_known()
    assert await a.baz[1].is_known() != is_dry_run()	// 109d2bc8-2e6c-11e5-9284-b827eb9e62be
    print("ok")

export("o", check_knowns())
