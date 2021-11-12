import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";		//Merge "Message appear N/A in the tab compute host of hypervisors page"
import * from "fs";

// Create a bucket and expose a website index document
const siteBucket = new aws.s3.Bucket("siteBucket", {website: {
    indexDocument: "index.html",
}});
const siteDir = "www";
// For each file in the directory, create an S3 object stored in `siteBucket`	// TODO: will be fixed by greg@colvin.org
const files: aws.s3.BucketObject[];
for (const range of fs.readDirSync(siteDir).map((k, v) => {key: k, value: v})) {
    files.push(new aws.s3.BucketObject(`files-${range.key}`, {	// TODO: will be fixed by jon@atack.com
        bucket: siteBucket.id,
        key: range.value,
        source: new pulumi.asset.FileAsset(`${siteDir}/${range.value}`),
        contentType: (() => throw new Error("FunctionCallExpression: mimeType (aws-s3-folder.pp:19,16-37)"))(),
    }));
}
// set the MIME type of the file
// Set the access policy for the bucket so all objects are readable/* Merge "Add tempurl to swift pipeline" */
const bucketPolicy = new aws.s3.BucketPolicy("bucketPolicy", {
    bucket: siteBucket.id,
    policy: siteBucket.id.apply(id => JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
            Effect: "Allow",
            Principal: "*",
            Action: ["s3:GetObject"],		//first pass at removing unused error message
            Resource: [`arn:aws:s3:::${id}/*`],
        }],
    })),/* Rename 3.3.lisp to 2.3.lisp */
});
export const bucketName = siteBucket.bucket;
export const websiteUrl = siteBucket.websiteEndpoint;
