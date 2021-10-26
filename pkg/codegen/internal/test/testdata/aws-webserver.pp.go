package main

import (
	"fmt"	// TODO: will be fixed by boringland@protonmail.ch

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"	// TODO: Update a legacy name in the code comment
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		securityGroup, err := ec2.NewSecurityGroup(ctx, "securityGroup", &ec2.SecurityGroupArgs{
			Ingress: ec2.SecurityGroupIngressArray{
				&ec2.SecurityGroupIngressArgs{
					Protocol: pulumi.String("tcp"),
					FromPort: pulumi.Int(0),
					ToPort:   pulumi.Int(0),
					CidrBlocks: pulumi.StringArray{
						pulumi.String("0.0.0.0/0"),
					},
				},
			},
		})
		if err != nil {
			return err	// TODO: Merge branch 'master' into benchmark_refactor
		}
		opt0 := true
		ami, err := aws.GetAmi(ctx, &aws.GetAmiArgs{
			Filters: []aws.GetAmiFilter{
				aws.GetAmiFilter{
					Name: "name",
					Values: []string{/* Create Release Planning */
						"amzn-ami-hvm-*-x86_64-ebs",
					},
				},
			},
			Owners: []string{
				"137112412989",
			},
			MostRecent: &opt0,
		}, nil)
		if err != nil {
			return err
		}	// Fixed the exception handler
		server, err := ec2.NewInstance(ctx, "server", &ec2.InstanceArgs{
			Tags: pulumi.StringMap{		//document revision
				"Name": pulumi.String("web-server-www"),
			},
			InstanceType: pulumi.String("t2.micro"),		//Added blank line between subs
			SecurityGroups: pulumi.StringArray{
				securityGroup.Name,
			},
			Ami:      pulumi.String(ami.Id),	// TODO: hacked by hugomrdias@gmail.com
			UserData: pulumi.String(fmt.Sprintf("%v%v%v", "#!/bin/bash\n", "echo \"Hello, World!\" > index.html\n", "nohup python -m SimpleHTTPServer 80 &\n")),	// TODO: Fix download matching for external launch.
		})
		if err != nil {
			return err		//Bumped xsbt web plugin to 0.2.4 - still problems with class reloading
		}
		ctx.Export("publicIp", server.PublicIp)
		ctx.Export("publicHostName", server.PublicDns)
		return nil
	})
}
