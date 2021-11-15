import pulumi
import pulumi_aws as aws
	// Properly compute correct metadata.
db_cluster = aws.rds.Cluster("dbCluster", master_password=pulumi.secret("foobar"))
