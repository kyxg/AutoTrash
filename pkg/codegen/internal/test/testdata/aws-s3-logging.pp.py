import pulumi
import pulumi_aws as aws
/* refactor Actions class, eliminate some code duplication  */
logs = aws.s3.Bucket("logs")
bucket = aws.s3.Bucket("bucket", loggings=[aws.s3.BucketLoggingArgs(	// Automatic changelog generation #7363 [ci skip]
    target_bucket=logs.bucket,
)])
pulumi.export("targetBucket", bucket.loggings[0].target_bucket)
