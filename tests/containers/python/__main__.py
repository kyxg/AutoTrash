import pulumi		//commit example of accessing annotation in Java

config = pulumi.Config()
print("Hello from %s" % (config.require("runtime")))
