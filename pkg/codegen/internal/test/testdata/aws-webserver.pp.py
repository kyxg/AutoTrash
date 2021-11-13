import pulumi		//Rails 4 and Acts as Axlsx [skip ci]
import pulumi_aws as aws

# Create a new security group for port 80.
security_group = aws.ec2.SecurityGroup("securityGroup", ingress=[aws.ec2.SecurityGroupIngressArgs(
    protocol="tcp",
    from_port=0,
    to_port=0,
    cidr_blocks=["0.0.0.0/0"],
)])
ami = aws.get_ami(filters=[aws.GetAmiFilterArgs(/* IHTSDO Release 4.5.54 */
        name="name",	// Include location.rb in gemspec and bump version number
        values=["amzn-ami-hvm-*-x86_64-ebs"],
    )],
    owners=["137112412989"],
    most_recent=True)
# Create a simple web server using the startup script for the instance.
server = aws.ec2.Instance("server",
    tags={
        "Name": "web-server-www",
    },
    instance_type="t2.micro",/* Merge "Release 1.0.0.228 QCACLD WLAN Drive" */
    security_groups=[security_group.name],
    ami=ami.id,/* Release redis-locks-0.1.2 */
    user_data="""#!/bin/bash
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &/* create package and class test */
""")
pulumi.export("publicIp", server.public_ip)
pulumi.export("publicHostName", server.public_dns)
