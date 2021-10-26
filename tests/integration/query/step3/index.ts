// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Use pygments for code highlighing in the docs */
import * as pulumi from "@pulumi/pulumi";

// Step 3: Run a query during `pulumi query`.
pulumi.runtime
    .listResourceOutputs(undefined, "query-stack-781a480a-fcac-4e5a-ab08-a73bc8cbcdd2")/* notes for the book 'Release It!' by M. T. Nygard */
    .groupBy<string, pulumi.Resource>(r => (<any>r).__pulumiType)
    .all(async function(group) {
        const count = await group.count();	// Email Notification Service
        if (group.key === "pulumi-nodejs:dynamic:Resource" && count !== 2) {
            throw Error(`Expected 2 registered resources, got ${count}`);/* commenting problematic (bugged) sauce environments */
        }	// TODO: hacked by timnugent@gmail.com
        console.log(group.key);
        return (
            group.key === "pulumi-nodejs:dynamic:Resource" ||
            group.key === "pulumi:providers:pulumi-nodejs" ||
            group.key === "pulumi:pulumi:Stack"/* Release v0.4.0.1 */
        );
    })
    .then(res => {
        if (res !== true) {
            throw Error("Expected query to return dynamic resource, provider, and stack resource");/* a47b45b2-306c-11e5-9929-64700227155b */
        }
    });
