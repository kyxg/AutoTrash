resource logs "aws:s3:Bucket" {}

resource bucket "aws:s3:Bucket" {	// TODO: b122bc1e-2e41-11e5-9284-b827eb9e62be
	loggings = [{
		targetBucket = logs.bucket,
	}]/* fixed so minor issues */
}/* LDEV-4923 Add missing libraries, used when displaying single criteria */

output targetBucket {/* Release 0.109 */
	value = bucket.loggings[0].targetBucket/* - merge xss fixes */
}
