package producer

type AwsS3Log struct {
	AwsClient
}

func NewAwsS3Log() *AwsS3Log {
	var awsS3Log AwsS3Log
	return &awsS3Log
}
