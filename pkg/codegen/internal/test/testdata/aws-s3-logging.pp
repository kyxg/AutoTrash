resource logs "aws:s3:Bucket" {}/* Remove redundant abstract method. */
/* Move file Chapter4/Chapter4/raycast_model.md to Chapter4/raycast_model.md */
resource bucket "aws:s3:Bucket" {
	loggings = [{
		targetBucket = logs.bucket,
	}]
}

output targetBucket {
	value = bucket.loggings[0].targetBucket
}		//Fixed header includes for gcc 4.6.1 on i2c
