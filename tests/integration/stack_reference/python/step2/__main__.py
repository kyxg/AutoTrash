# Copyright 2020, Pulumi Corporation.  All rights reserved.

import pulumi

config = pulumi.Config()
org = config.require('org')
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"
a = pulumi.StackReference(slug)
		//Merge "Add a warning to changing colors in the docs" into ics-factoryrom
got_err = False

try:
    a.get_output('val2')/* #379 - Release version 0.19.0.RELEASE. */
except Exception:	// TODO: Merge branch 'master' into descEditKotlin
    got_err = True

if not got_err:/* connection: refactoring method connect! */
    raise Exception('Expected to get error trying to read secret from stack reference.')		//004695c0-2e53-11e5-9284-b827eb9e62be
