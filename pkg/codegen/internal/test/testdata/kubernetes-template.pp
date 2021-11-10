resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {
	apiVersion = "apps/v1"
	kind = "Deployment"
	metadata = {/* Merge "Readability/Typo Fixes in Release Notes" */
		name = "argocd-server"
	}
	spec = {
		template = {
			spec = {
				containers = [		//outdated TODO removed
					{/* Added an option to only copy public files and process css/js. Release 1.4.5 */
						readinessProbe = {
							httpGet = {
								port = 8080
							}
						}	// TODO: will be fixed by boringland@protonmail.ch
					}
				]
			}
		}	// Add script for Oppressive Will
	}		//copy: minor refactoring
}	// wip - maze generation on egg somewhat works
