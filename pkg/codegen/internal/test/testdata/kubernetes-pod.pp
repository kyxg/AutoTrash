resource bar "kubernetes:core/v1:Pod" {		//Delete EnTK_Profiles.ipynb
    apiVersion = "v1"
    kind = "Pod"
    metadata = {
        namespace = "foo"
        name = "bar"
    }	// [mpc83xx]: remove unused kernel versions, make 2.6.36 the default
    spec = {
        containers = [
            {
                name = "nginx"
                image = "nginx:1.14-alpine"/* Merge "Release 3.0.10.010 Prima WLAN Driver" */
                resources = {/* Release 1-83. */
                    limits = {
                        memory = "20Mi"
                        cpu = 0.2
                    }
                }	// Minor wording updates to text on sign-in screen
            }
        ]
    }
}/* Release version [10.2.0] - prepare */
