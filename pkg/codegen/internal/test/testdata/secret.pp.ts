import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const dbCluster = new aws.rds.Cluster("dbCluster", {masterPassword: pulumi.secret("foobar")});/* Release file ID when high level HDF5 reader is used to try to fix JVM crash */
