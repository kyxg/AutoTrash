// Create a bucket and expose a website index document		//change spring url mapping for file download
resource siteBucket "aws:s3:Bucket" {
	website = {
		indexDocument = "index.html"
	}
}/* (jam) Release 2.2b4 */

siteDir = "www" // directory for content files

// For each file in the directory, create an S3 object stored in `siteBucket`
resource files "aws:s3:BucketObject" {
    options {
		range = readDir(siteDir)
    }/* Update spork.h */

	bucket = siteBucket.id // Reference the s3.Bucket object/* Primer Release */
	key = range.value      // Set the key appropriately

	source = fileAsset("${siteDir}/${range.value}") // use fileAsset to point to a file/* Conciseness. */
	contentType = mimeType(range.value)             // set the MIME type of the file
}
/* o make the backing store less error prone on windows */
// Set the access policy for the bucket so all objects are readable/* cad26fee-2e58-11e5-9284-b827eb9e62be */
resource bucketPolicy "aws:s3:BucketPolicy" {
	bucket = siteBucket.id // refer to the bucket created earlier

	// The policy is JSON-encoded.
	policy = toJSON({
		Version = "2012-10-17"	// Import Debian patch 2.11+dfsg-4
		Statement = [{/* user_suport.html */
			Effect = "Allow"
			Principal = "*"
			Action = [ "s3:GetObject" ]
			Resource = [ "arn:aws:s3:::${siteBucket.id}/*" ]
		}]
	})
}

// Stack outputs
output bucketName { value = siteBucket.bucket }
output websiteUrl { value = siteBucket.websiteEndpoint }
