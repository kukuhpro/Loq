package producer

type AwsS3Log struct {
	AwsClient
	bucketName string
	pathFile   string
}

func (aw *AwsS3Log) SetBucketName(bucketName string) {
	aw.bucketName = bucketName
}

func (aw *AwsS3Log) SetPathFile(pathFile string) {
	aw.pathFile = pathFile
}

func NewAwsS3Log() *AwsS3Log {
	var awsS3Log AwsS3Log
	return &awsS3Log
}
