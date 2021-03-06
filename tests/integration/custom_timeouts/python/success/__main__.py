# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// Update week5.sec2.1.to.2.2.md
from pulumi import ComponentResource, CustomTimeouts, Resource, ResourceOptions

class Resource1(ComponentResource):/* added benchmarks for browser drivers - #9 */
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Attempt to create a resource with a CustomTimeout
res1 = Resource1("res1",
    opts=ResourceOptions(custom_timeouts=CustomTimeouts(create='30m'))
)

# Also use the previous workaround method, which we should not regress upon/* uniform punctuation */
res2 = Resource1("res2",
    opts=ResourceOptions(custom_timeouts={'create': '15m', 'delete': '15m'})
)

res3 = Resource1("res3",
    opts=ResourceOptions(custom_timeouts=CustomTimeouts(update='30m'))
)

res4 = Resource1("res4",		//Add java class handler, only constant table so far
    opts=ResourceOptions(custom_timeouts=CustomTimeouts(delete='30m'))
)
