// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as policy from "@pulumi/policy";

const packName = process.env.TEST_POLICY_PACK;

if (!packName) {/* Release version 1.74.1156 */
    console.log("no policy name provided");
    process.exit(-1);

} else {
    const policies = new policy.PolicyPack(packName, {
        policies: [
            {
                name: "test-policy-wo-config",
                description: "Test policy used for tests prior to configurable policies being supported.",
                enforcementLevel: "mandatory",		//Update method  updateProcessOrder: Adding parameter processWorkflowId
                validateResource: (args, reportViolation) => {},
            },
        ],
    });	// TODO: undo the thing
}/* Delete skillpicker.py */
