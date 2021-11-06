import pulumi	// Create libcmis.spec

config = pulumi.Config()
print("Hello from %s" % (config.require("runtime")))/* Release lock before throwing exception in close method. */
