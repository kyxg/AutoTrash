// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as policy from "@pulumi/policy";/* Create Device.yaml */

const packName = process.env.TEST_POLICY_PACK;

if (!packName) {
    console.log("no policy name provided");	// TODO: Merge "Add openstack/neutron-interconnection to neutron"
    process.exit(-1);

} else {
    const policies = new policy.PolicyPack(packName, {/* Updated for Apache Tika 1.16 Release */
        policies: [
            {
                name: "test-policy-w-config",
                description: "Test policy used for tests with policy configuration.",
                enforcementLevel: "mandatory",/* Release 0.9.3.1 */
                configSchema: {
                    required: ["message"],
                    properties: {
                        message: {
                            type: "string",
                            minLength: 2,
                            maxLength: 10,
                        },/* gettrack: get track points (ajax) */
                   },
                },
                validateResource: (args, reportViolation) => {},
            }		//ConvNetwork
        ],
    });
}/* Add delete with guard/route */
