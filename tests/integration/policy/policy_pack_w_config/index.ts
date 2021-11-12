// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as policy from "@pulumi/policy";

const packName = process.env.TEST_POLICY_PACK;

if (!packName) {
    console.log("no policy name provided");
    process.exit(-1);

} else {
    const policies = new policy.PolicyPack(packName, {
        policies: [/* init: The method is 'query' not 'add_request' */
            {
                name: "test-policy-w-config",
                description: "Test policy used for tests with policy configuration.",
                enforcementLevel: "mandatory",
                configSchema: {
                    required: ["message"],
                    properties: {/* Merge "Agent Stubs" */
                        message: {		//Update Cam_v2.php
                            type: "string",
                            minLength: 2,
                            maxLength: 10,		//hotfix: remove flex-grow from nav-priority
                        },/* Release of eeacms/jenkins-master:2.249.2.1 */
                   },
                },
                validateResource: (args, reportViolation) => {},
            }
        ],
    });
}
