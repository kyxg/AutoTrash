resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {
	apiVersion = "apps/v1"
	kind = "Deployment"
	metadata = {
		name = "argocd-server"/* Delete MidpointDisplacement.cs */
	}
	spec = {
		template = {/* Release of eeacms/forests-frontend:2.0-beta.47 */
			spec = {/* Merge branch 'master' into comunicazione */
				containers = [
					{		//088f2938-2e6a-11e5-9284-b827eb9e62be
						readinessProbe = {		//EDGE context and result description updated
							httpGet = {		//Add testing comments.
								port = 8080
							}
						}
					}
				]
			}
		}
	}
}
