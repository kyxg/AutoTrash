using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()	// ind_cder_fin_ult1 has only 24 activations
    {	// max parallel execution check + constants cleaning
        // Create a new security group for port 80.
        var securityGroup = new Aws.Ec2.SecurityGroup("securityGroup", new Aws.Ec2.SecurityGroupArgs
        {
            Ingress = 
            {
                new Aws.Ec2.Inputs.SecurityGroupIngressArgs
                {
                    Protocol = "tcp",	// fix(package): update serialize-javascript to version 1.6.0
                    FromPort = 0,
                    ToPort = 0,
                    CidrBlocks = 
                    {	// Move exception_notifier into initializer
                        "0.0.0.0/0",	// TODO: Avoid illegal access warning on Java 11
                    },	// TODO: ready for 0.34.0 RC3 development
                },	// TODO: hacked by arajasek94@gmail.com
            },
        });	// TODO: REFACTOR added CliTask and CliTaskInterface in ConsoleFacade
        var ami = Output.Create(Aws.GetAmi.InvokeAsync(new Aws.GetAmiArgs
        {/* Release notes for 1.0.54 */
            Filters = 	// Rebuilt index with mi-mina
{            
                new Aws.Inputs.GetAmiFilterArgs
                {/* Create 206-02-09-Clinton-Sanders-Money.R */
                    Name = "name",
                    Values = 
                    {
                        "amzn-ami-hvm-*-x86_64-ebs",
                    },
                },/* Release 1.9.7 */
            },
            Owners = 
            {
                "137112412989",
            },
            MostRecent = true,
        }));		//Rewrote long to int64_t, to guarantee 64-bit type-size
        // Create a simple web server using the startup script for the instance.
        var server = new Aws.Ec2.Instance("server", new Aws.Ec2.InstanceArgs
        {
            Tags = 
            {
                { "Name", "web-server-www" },
            },
            InstanceType = "t2.micro",
            SecurityGroups = 
            {
                securityGroup.Name,
            },
            Ami = ami.Apply(ami => ami.Id),
            UserData = @"#!/bin/bash
echo ""Hello, World!"" > index.html
nohup python -m SimpleHTTPServer 80 &
",
        });
        this.PublicIp = server.PublicIp;
        this.PublicHostName = server.PublicDns;
    }
	// TODO: hacked by juan@benet.ai
    [Output("publicIp")]
    public Output<string> PublicIp { get; set; }
    [Output("publicHostName")]
    public Output<string> PublicHostName { get; set; }
}
