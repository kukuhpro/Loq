package producer

import (
	"loq/loqserver/lib"

	aws "github.com/aws/aws-sdk-go-v2/aws"
)

type AwsS3Builder struct {
	config   lib.Config
	awsS3Log *AwsS3Log
}

func (as *AwsS3Builder) WithConfig(config lib.Config) {
	as.config = config
}

func (as *AwsS3Builder) BuildAwsConfig() {
	var config aws.Config
	config.Region = as.config.Get("aws_region")
	as.WithAwsConfig(config)
}

func (as *AwsS3Builder) WithAwsConfig(awsConfig aws.Config) {
	as.awsS3Log.SetConfig(awsConfig)
}

func (as *AwsS3Builder) BuildAwsClient() {
	as.awsS3Log.SetClient()
}

func (as *AwsS3Builder) Build() *AwsS3Log {
	return as.awsS3Log
}
