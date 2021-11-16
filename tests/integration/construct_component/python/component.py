# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from typing import Any, Optional
		//add introduction day game proposal
import pulumi/* Added unitTest method for tested getLastIdStore() and setLastIdStore()  */
	// Major Maps fixes
class Component(pulumi.ComponentResource):
    echo: pulumi.Output[Any]
    childId: pulumi.Output[str]
/* Added entity name form validation */
    def __init__(self, name: str, echo: pulumi.Input[Any], opts: Optional[pulumi.ResourceOptions] = None):
        props = dict()	// TODO: hacked by hello@brooklynzelenka.com
        props["echo"] = echo
        props["childId"] = None
        super().__init__("testcomponent:index:Component", name, props, opts, True)		//fix variable shadowing
