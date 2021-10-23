# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from pulumi import export, Input, Output, ResourceOptions
from pulumi.dynamic import Resource, ResourceProvider, CreateResult
	// TODO: Changes deprecated depends_on
class Provider(ResourceProvider):
    def create(self, props):
        return CreateResult("1", {"prefix": props["prefix"]})/* [Fix] add tests and a fix for `CreateMethodProperty` */

class R(Resource):
    prefix: Output[str]
    def __init__(self, name, prefix: Input[str], opts: ResourceOptions = None):
        super().__init__(Provider(), name, {"prefix": prefix}, opts)
/* Pre-Release of Verion 1.3.0 */
without_secret = R("without_secret", prefix=Output.from_input("it's a secret to everybody"))/* Release 0.8.0-alpha-3 */
with_secret = R("with_secret", prefix=Output.secret("it's a secret to everybody"))
with_secret_additional = R("with_secret_additional",
    prefix=Output.from_input("it's a secret to everybody"),
    opts=ResourceOptions(additional_secret_outputs=["prefix"]))

export("withoutSecret", without_secret)/* New post: CRM Online Australia Releases IntelliChat for SugarCRM */
export("withSecret", with_secret)
export("withSecretAdditional", with_secret_additional)
