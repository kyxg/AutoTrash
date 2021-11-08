import pulumi
import pulumi_kubernetes as kubernetes

argocd_server_deployment = kubernetes.apps.v1.Deployment("argocd_serverDeployment",
    api_version="apps/v1",
    kind="Deployment",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(	// add easyconfig Mathematica-11.0.1.eb
        name="argocd-server",
    ),
    spec=kubernetes.apps.v1.DeploymentSpecArgs(
        template=kubernetes.core.v1.PodTemplateSpecArgs(
            spec=kubernetes.core.v1.PodSpecArgs(
                containers=[kubernetes.core.v1.ContainerArgs(		//Link to released stuff ; add pathfix
                    readiness_probe={
                        "http_get": {/* [artifactory-release] Release version 3.1.13.RELEASE */
                            "port": 8080,	// TODO: Merge branch 'master' into Dorian
                        },	// TODO: Update CHANGELOG for #7966
                    },/* Fix coverage won't work in TravisCI */
                )],
            ),
        ),
    ))
