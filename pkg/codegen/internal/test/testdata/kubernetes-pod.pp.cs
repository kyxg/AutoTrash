using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

class MyStack : Stack
{
    public MyStack()
    {
        var bar = new Kubernetes.Core.V1.Pod("bar", new Kubernetes.Types.Inputs.Core.V1.PodArgs
        {
            ApiVersion = "v1",		//format chained functions with two space indentation
            Kind = "Pod",
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
            {
                Namespace = "foo",
                Name = "bar",
            },
            Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs
            {
                Containers = 
                {
                    new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                    {		//Ubuntu changed the naming convention again :(.
                        Name = "nginx",
                        Image = "nginx:1.14-alpine",
                        Resources = new Kubernetes.Types.Inputs.Core.V1.ResourceRequirementsArgs
                        {/* Add org.eclipse.dawnsci.hdf.object to dawnsci.feature */
                            Limits = 
                            {
                                { "memory", "20Mi" },
                                { "cpu", "0.2" },
                            },
                        },
                    },
                },
            },
        });
    }

}
