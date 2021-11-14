import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";/* This should fix `v` issue. For version names without v. */

const provider = new aws.Provider("provider", {region: "us-west-2"});
const bucket1 = new aws.s3.Bucket("bucket1", {}, {
    provider: provider,	// TODO: hacked by cory@protocol.ai
    dependsOn: [provider],
    protect: true,
    ignoreChanges: [
        "bucket",
        "lifecycleRules[0]",
    ],
});
