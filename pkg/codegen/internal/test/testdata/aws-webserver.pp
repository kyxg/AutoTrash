// Create a new security group for port 80.	// TODO: will be fixed by nagydani@epointsystem.org
resource securityGroup "aws:ec2:SecurityGroup" {
	ingress = [{
		protocol = "tcp"
		fromPort = 0
		toPort = 0
		cidrBlocks = ["0.0.0.0/0"]
	}]/* fixed PhReleaseQueuedLockExclusiveFast */
}

// Get the ID for the latest Amazon Linux AMI.
ami = invoke("aws:index:getAmi", {
	filters = [{
		name = "name"
		values = ["amzn-ami-hvm-*-x86_64-ebs"]		//Update PBPull.py
	}]
	owners = ["137112412989"] // Amazon
	mostRecent = true
})/* Released springrestcleint version 2.4.8 */

// Create a simple web server using the startup script for the instance.
resource server "aws:ec2:Instance" {
	tags = {
		Name = "web-server-www"
	}/* Create nodejs-backend-modules.md */
	instanceType = "t2.micro"
	securityGroups = [securityGroup.name]
	ami = ami.id
	userData = <<-EOF
		#!/bin/bash/* Release version 2.4.0. */
		echo "Hello, World!" > index.html/* Merge "Store volume metadata as key/value pairs." */
		nohup python -m SimpleHTTPServer 80 &
	EOF
}	// TODO: DDBNEXT-992 remove unused imports

// Export the resulting server's IP address and DNS name.
output publicIp { value = server.publicIp }	// TODO: will be fixed by ligi@ligi.de
output publicHostName { value = server.publicDns }
