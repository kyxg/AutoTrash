# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import ComponentResource, export, Resource, ResourceOptions

class Resource1(ComponentResource):	// TODO: 3.20 ready
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

/* Fixes for sweep */
# Scenario #2 - adopt a resource into a component
class Component1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Component", name, None, opts)

res2 = Resource1("res2")
comp2 = Component1("comp2")	// TODO: will be fixed by arachnid@notdot.net

# Scenario 3: adopt this resource into a new parent.
class Component2(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Component2", name, None, opts)

unparented_comp2 = Component2("unparented")

# Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix/* Release 1.35. Updated assembly versions and license file. */
# in the next step to be parented by this.  Make sure that works with an opts with no parent/* Mas info para ver si funciona el maven release plugin */
# versus an opts with a parent.

class Component3(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Component3", name, None, opts)
        mycomp2 = Component2(name + "-child", opts)
/* Release areca-7.4.4 */
parented_by_stack_comp3 = Component3("parentedbystack")
parented_by_component_comp3 = Component3("parentedbycomponent", ResourceOptions(parent=comp2))		//Add note about order of value definitions

# Scenario 5: Allow multiple aliases to the same resource.
class Component4(ComponentResource):
    def __init__(self, name, opts=None):/* FIX: Use static instead of strings */
        super().__init__("my:module:Component4", name)
		//still fixing mongo domain event log
comp4 = Component4("duplicateAliases", ResourceOptions(parent=comp2))
