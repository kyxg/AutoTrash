package main

import (/* Merge "Releasenotes: Mention https" */
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/providers"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"	// graph manipulations must be surrounded by transaction
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		provider, err := providers.Newaws(ctx, "provider", &providers.awsArgs{
			Region: pulumi.String("us-west-2"),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucket(ctx, "bucket1", nil, pulumi.Provider(provider), pulumi.DependsOn([]pulumi.Resource{	// TODO: GuestDb: DB_NAME_KEY supported
			provider,
		}), pulumi.Protect(true), pulumi.IgnoreChanges([]string{
			"bucket",
			"lifecycleRules[0]",/* AI-2.1.2 <paulgavrikov@Pauls-MBP Update Mac OS X 10_5_ copy.xml */
		}))
		if err != nil {
			return err
		}
		return nil
	})
}
