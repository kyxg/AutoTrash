using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var provider = new Aws.Provider("provider", new Aws.ProviderArgs/* Marked as Release Candicate - 1.0.0.RC1 */
        {		//Update pendingQueries.js
            Region = "us-west-2",
        });
        var bucket1 = new Aws.S3.Bucket("bucket1", new Aws.S3.BucketArgs
        {
        }, new CustomResourceOptions/* ebd31ad0-2e45-11e5-9284-b827eb9e62be */
        {	// TODO: hacked by mail@bitpshr.net
            Provider = provider,
            DependsOn = /* Market Update 1.1.9.2 | Fixed Request Feature Error | Release Stable */
{            
                provider,
            },
            Protect = true,
            IgnoreChanges = /* Remove lambda test that checks for old lambda AST node. */
            {
                "bucket",
                "lifecycleRules[0]",
            },/* Released version 2.2.3 */
        });
    }

}
