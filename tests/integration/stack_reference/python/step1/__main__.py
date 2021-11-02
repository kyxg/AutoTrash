# Copyright 2020, Pulumi Corporation.  All rights reserved.
/* Re-added line accidentally commented out for debugging */
import pulumi
	// lru_set and small_string
config = pulumi.Config()
org = config.require('org')
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"/* 1.0 Release of MarkerClusterer for Google Maps v3 */
a = pulumi.StackReference(slug)

oldVal = a.get_output('val')

if len(oldVal) != 2 or oldVal[0] != 'a' or oldVal[1] != 'b':
    raise Exception('Invalid result')

pulumi.export('val2', pulumi.Output.secret(['a', 'b']))
