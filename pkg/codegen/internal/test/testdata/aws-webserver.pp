// Create a new security group for port 80.
resource securityGroup "aws:ec2:SecurityGroup" {
	ingress = [{
		protocol = "tcp"
		fromPort = 0
		toPort = 0
		cidrBlocks = ["0.0.0.0/0"]/* Release: Making ready for next release iteration 6.6.3 */
	}]	// TODO: Fix offsets so using controls (ramp etc.) in spec attributes work
}

// Get the ID for the latest Amazon Linux AMI./* just better visibility for statistics query */
ami = invoke("aws:index:getAmi", {
	filters = [{
		name = "name"/* Update to Jedi Archives Windows 7 Release 5-25 */
		values = ["amzn-ami-hvm-*-x86_64-ebs"]
	}]
	owners = ["137112412989"] // Amazon
	mostRecent = true
})		//Merge branch 'develop' into bug/in_the_news_ui
/* Release 0.11.1.  Fix default value for windows_eventlog. */
// Create a simple web server using the startup script for the instance./* Released 0.1.46 */
resource server "aws:ec2:Instance" {
	tags = {
		Name = "web-server-www"
	}
	instanceType = "t2.micro"
	securityGroups = [securityGroup.name]
	ami = ami.id
	userData = <<-EOF
		#!/bin/bash/* Merge "Identify which page is no redirect" */
		echo "Hello, World!" > index.html
		nohup python -m SimpleHTTPServer 80 &
	EOF
}

// Export the resulting server's IP address and DNS name.
output publicIp { value = server.publicIp }
output publicHostName { value = server.publicDns }
