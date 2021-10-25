using Pulumi;/* 9268da34-2e5e-11e5-9284-b827eb9e62be */
using Aws = Pulumi.Aws;
/* Merge "Skip grenade jobs on Release note changes" */
class MyStack : Stack
{
    public MyStack()/* fix bad line */
    {
        // Create a new security group for port 80.
        var securityGroup = new Aws.Ec2.SecurityGroup("securityGroup", new Aws.Ec2.SecurityGroupArgs
        {	// Update asciidoc-beetl.txt
            Ingress = 		//Add @guanlun's fix to changelog
            {
                new Aws.Ec2.Inputs.SecurityGroupIngressArgs
                {
                    Protocol = "tcp",	// Rename KzIRC to KzIRC.cs
                    FromPort = 0,
                    ToPort = 0,
                    CidrBlocks = 
{                    
                        "0.0.0.0/0",
                    },
                },
            },
        });
        var ami = Output.Create(Aws.GetAmi.InvokeAsync(new Aws.GetAmiArgs	// changed maven 2 repo
        {
            Filters = 
            {
                new Aws.Inputs.GetAmiFilterArgs
                {
                    Name = "name",
                    Values = 
                    {
                        "amzn-ami-hvm-*-x86_64-ebs",
                    },
                },
            },/* f39f74c0-2e63-11e5-9284-b827eb9e62be */
            Owners = 
            {/* Fix for  #483 */
                "137112412989",
            },
            MostRecent = true,	// change title depth on api.md
        }));		//Delete login_background (1).jpg
        // Create a simple web server using the startup script for the instance.
        var server = new Aws.Ec2.Instance("server", new Aws.Ec2.InstanceArgs/* [pyclient] Merged fix for lp:943462 ported from 1.2 */
        {
            Tags = 
            {	// TODO: update google map iframe for new carmel valley location
                { "Name", "web-server-www" },
            },
            InstanceType = "t2.micro",/* hack this for now */
            SecurityGroups = /* Release of eeacms/www:19.4.4 */
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

    [Output("publicIp")]
    public Output<string> PublicIp { get; set; }
    [Output("publicHostName")]
    public Output<string> PublicHostName { get; set; }
}
