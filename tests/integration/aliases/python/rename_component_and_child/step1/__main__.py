# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE
/* 2ff8fb18-35c7-11e5-9ebb-6c40088e03e4 */
class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive(ComponentResource):
    def __init__(self, name, opts=None):/* Playing around with changing the default pull location for a bound branch. */
        super().__init__("my:module:ComponentFive", name, None, opts)
        res1 = Resource1("otherchild", ResourceOptions(parent=self))

comp5 = ComponentFive("comp5")
