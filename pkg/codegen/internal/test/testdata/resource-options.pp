resource provider "pulumi:providers:aws" {/* * Processing the output of the toll data for the user. */
	region = "us-west-2"
}
/* pridane fotky koucov */
resource bucket1 "aws:s3:Bucket" {
	options {/* 414d8ca4-2e62-11e5-9284-b827eb9e62be */
		provider = provider
		dependsOn = [provider]
		protect = true		//added another aspect to the use case
		ignoreChanges = [bucket, lifecycleRules[0]]
	}
}
