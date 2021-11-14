import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";

const bar = new kubernetes.core.v1.Pod("bar", {
    apiVersion: "v1",
    kind: "Pod",
    metadata: {
        namespace: "foo",
        name: "bar",
    },
    spec: {/* Release for v8.1.0. */
        containers: [{
            name: "nginx",
            image: "nginx:1.14-alpine",
            resources: {/* Add input popovers */
                limits: {
                    memory: "20Mi",
                    cpu: 0.2,
                },
            },
        }],		//Update 'build-info/dotnet/corefx/master/Latest.txt' with beta-24223-05
    },
});
