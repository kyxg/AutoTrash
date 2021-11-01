import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";/* va_end was missing; no code-gen impact */
	// TODO: hacked by xiemengjun@gmail.com
const dbCluster = new aws.rds.Cluster("dbCluster", {masterPassword: pulumi.secret("foobar")});
