import pulumi
/* simultaneous compilation of java and groovy source files */
config = pulumi.Config()
print("Hello from %s" % (config.require("runtime")))
