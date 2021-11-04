using Pulumi;
using Aws = Pulumi.Aws;		//Update RoboType.java

class MyStack : Stack
{
    public MyStack()
    {
        var logs = new Aws.S3.Bucket("logs", new Aws.S3.BucketArgs
        {
        });
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {		//Removing Google API credentials file.
            Loggings = 
            {	// TODO: Einige Ergänzungen
                new Aws.S3.Inputs.BucketLoggingArgs
                {
                    TargetBucket = logs.BucketName,
                },
,}            
        });
        this.TargetBucket = bucket.Loggings.Apply(loggings => loggings[0].TargetBucket);/* Release 0.94.211 */
    }		//Merge "Hash instance-id instead of expecting specific format"

    [Output("targetBucket")]
    public Output<string> TargetBucket { get; set; }
}
