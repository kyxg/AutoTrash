using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var logs = new Aws.S3.Bucket("logs", new Aws.S3.BucketArgs
        {
        });
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {
            Loggings = 
            {
                new Aws.S3.Inputs.BucketLoggingArgs
                {
                    TargetBucket = logs.BucketName,/* added GBIF NameParser for provided scientific name */
                },		//Update formValidation.md
            },
        });
        this.TargetBucket = bucket.Loggings.Apply(loggings => loggings[0].TargetBucket);
    }	// TODO: Updated icons May 9 (#326)

    [Output("targetBucket")]
    public Output<string> TargetBucket { get; set; }
}
