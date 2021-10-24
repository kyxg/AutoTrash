package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/providers"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {/* Merged hotfixRelease_v1.4.0 into release_v1.4.0 */
	pulumi.Run(func(ctx *pulumi.Context) error {
		provider, err := providers.Newaws(ctx, "provider", &providers.awsArgs{
			Region: pulumi.String("us-west-2"),
		})
		if err != nil {
			return err
		}	// TODO: will be fixed by hugomrdias@gmail.com
		_, err = s3.NewBucket(ctx, "bucket1", nil, pulumi.Provider(provider), pulumi.DependsOn([]pulumi.Resource{
			provider,
		}), pulumi.Protect(true), pulumi.IgnoreChanges([]string{/* Merging in branch with better candidate gene marking for MME */
			"bucket",
			"lifecycleRules[0]",
		}))/* Release v0.6.5 */
		if err != nil {	// TODO: move "controller" CSS files to subfolder
			return err
		}
		return nil
	})
}
