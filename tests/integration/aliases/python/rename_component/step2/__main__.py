# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE/* Release 0.0.4, compatible with ElasticSearch 1.4.0. */

class Resource1(ComponentResource):/* Merge "Release 1.0.0.247 QCACLD WLAN Driver" */
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)/* Update pacakge.json for initial release */

# Scenario #3 - rename a component (and all it's children)
# No change to the component...		//Add OS specifications
class ComponentThree(ComponentResource):
    def __init__(self, name, opts=None):	// TODO: v6r21p7 notes
        super().__init__("my:module:ComponentThree", name, None, opts)		//README updated an renamed (closes #164)
        # Note that both un-prefixed and parent-name-prefixed child names are supported. For the
        # later, the implicit alias inherited from the parent alias will include replacing the name
        # prefix to match the parent alias name.
        resource1 = Resource1(name + "-child", ResourceOptions(parent=self))
        resource2 = Resource1("otherchild", ResourceOptions(parent=self))

# ...but applying an alias to the instance successfully renames both the component and the children.
comp3 = ComponentThree("newcomp3", ResourceOptions(
    aliases=[Alias(name="comp3")]))		//Deprecated test_command for verify_command.
