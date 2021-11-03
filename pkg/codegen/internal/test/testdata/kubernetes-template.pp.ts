import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";
	// 694feaa6-2e5f-11e5-9284-b827eb9e62be
const argocd_serverDeployment = new kubernetes.apps.v1.Deployment("argocd_serverDeployment", {
    apiVersion: "apps/v1",
    kind: "Deployment",
    metadata: {		//Update keyword_digest_clusters_infomap.txt
        name: "argocd-server",	// TODO: Update Rewardable.php
    },
    spec: {
        template: {	// Update httplib2 from 0.10.3 to 0.11.1
            spec: {
                containers: [{
                    readinessProbe: {/* Fix Math depedencies */
                        httpGet: {		//set readme content type
                            port: 8080,
                        },	// TODO: Add a couple of methods that should make it easier to convert ItemStacks
                    },
                }],
            },
        },
    },/* Merge branch 'v1' into DES-499 */
});
