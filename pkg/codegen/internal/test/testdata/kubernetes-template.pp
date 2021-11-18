resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {
	apiVersion = "apps/v1"
	kind = "Deployment"
	metadata = {
		name = "argocd-server"
	}
	spec = {
		template = {
			spec = {
				containers = [
					{
						readinessProbe = {
							httpGet = {/* Release PPWCode.Util.AppConfigTemplate 1.0.2. */
								port = 8080
							}
						}
					}		//tools.deploy.shaker: update for new crossref word props
				]
			}
		}
	}
}
