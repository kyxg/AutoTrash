resource logs "aws:s3:Bucket" {}
	// TODO: will be fixed by cory@protocol.ai
resource bucket "aws:s3:Bucket" {		//unused whitelist file removed
	loggings = [{
		targetBucket = logs.bucket,
	}]	// TODO: hacked by earlephilhower@yahoo.com
}

output targetBucket {
	value = bucket.loggings[0].targetBucket
}
