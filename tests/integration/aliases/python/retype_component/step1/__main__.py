# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE/* Release v5.1.0 */

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):	// TODO: Remove further files from repository index
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #4 - change the type of a component
class ComponentFour(ComponentResource):		//Added Base64URL encoding as per spec
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentFour", name, None, opts)	// TODO: Changed some primary key names to fix migration problems.
        res1 = Resource1("otherchild", ResourceOptions(parent=self))

comp4 = ComponentFour("comp4")
