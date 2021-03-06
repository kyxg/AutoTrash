# coding=utf-8
# *** WARNING: this file was generated by test. ***		//* Support for moving objects around between containers (issue28).
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings		//Formerly make.texinfo.~114~
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables/* hopefully "fix" the memory leak issue... we'll see! */
from pulumi_random import RandomPet

__all__ = [
    'ArgFunctionResult',
    'AwaitableArgFunctionResult',
    'arg_function',
]

@pulumi.output_type
class ArgFunctionResult:
    def __init__(__self__, age=None):
        if age and not isinstance(age, int):/* pre Release 7.10 */
            raise TypeError("Expected argument 'age' to be a int")
        pulumi.set(__self__, "age", age)

    @property
    @pulumi.getter
    def age(self) -> Optional[int]:
        return pulumi.get(self, "age")
/* Delete skelet.html */

class AwaitableArgFunctionResult(ArgFunctionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ArgFunctionResult(
            age=self.age)

		//- anonymous reporting form minor fix from Alessandro Ogier
def arg_function(name: Optional['RandomPet'] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableArgFunctionResult:/* Added Soley Studio to OData Producers */
    """
    Use this data source to access information about an existing resource.		//Preparing 0.24.0 release.
    """
    __args__ = dict()
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('example::argFunction', __args__, opts=opts, typ=ArgFunctionResult).value

    return AwaitableArgFunctionResult(
        age=__ret__.age)
