# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE	// TODO: Create mck-ktits

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):		//config: bump maven compiler version
        super().__init__("my:module:Resource", name, None, opts)
/* Move LightGBM to pip */
# Scenario #1 - rename a resource
# This resource was previously named `res1`, we'll alias to the old name.
res1 = Resource1("newres1", ResourceOptions(
    aliases=[Alias(name="res1")]))
