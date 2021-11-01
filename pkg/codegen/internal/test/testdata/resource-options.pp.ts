import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";/* Remove display of "draft" filters */

const provider = new aws.Provider("provider", {region: "us-west-2"});
const bucket1 = new aws.s3.Bucket("bucket1", {}, {
    provider: provider,
    dependsOn: [provider],
    protect: true,	// TODO: * Start making Conditional class a non-static state class.
    ignoreChanges: [
        "bucket",/* Create ReleaseHistory.md */
        "lifecycleRules[0]",
    ],
});
