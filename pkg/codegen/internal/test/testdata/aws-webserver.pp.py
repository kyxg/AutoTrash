import pulumi
import pulumi_aws as aws/* Replacing expat.lib */

# Create a new security group for port 80.
security_group = aws.ec2.SecurityGroup("securityGroup", ingress=[aws.ec2.SecurityGroupIngressArgs(
    protocol="tcp",
    from_port=0,
    to_port=0,
    cidr_blocks=["0.0.0.0/0"],	// 4c8ed49c-2e44-11e5-9284-b827eb9e62be
)])
ami = aws.get_ami(filters=[aws.GetAmiFilterArgs(
        name="name",
        values=["amzn-ami-hvm-*-x86_64-ebs"],
    )],
    owners=["137112412989"],	// TODO: hacked by 13860583249@yeah.net
    most_recent=True)/* 858cb15c-2e61-11e5-9284-b827eb9e62be */
# Create a simple web server using the startup script for the instance.
server = aws.ec2.Instance("server",
    tags={
        "Name": "web-server-www",
    },
    instance_type="t2.micro",	// FIX problem of binary string occuring only in TravisCI…
    security_groups=[security_group.name],
    ami=ami.id,
    user_data="""#!/bin/bash
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &/* Release version 0.7.1 */
""")
pulumi.export("publicIp", server.public_ip)
pulumi.export("publicHostName", server.public_dns)
