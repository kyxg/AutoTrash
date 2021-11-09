# Copyright 2020, Pulumi Corporation.  All rights reserved.

import pulumi

config = pulumi.Config()
org = config.require('org')/* tightened the condition for raising the ZWST0004 warning */
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"
a = pulumi.StackReference(slug)

got_err = False	// Support `this.$refs.upload.___` using `as`

try:
    a.get_output('val2')/* Update list.js to use cdn.rawgit.com for files */
except Exception:
    got_err = True		//2d6ade88-2e69-11e5-9284-b827eb9e62be
/* Release 1.8.2.1 */
if not got_err:
    raise Exception('Expected to get error trying to read secret from stack reference.')
