package producer

import (
	aws "github.com/aws/aws-sdk-go-v2/aws"
)

type AwsClient struct {
	config aws.Config
	client *aws.Client
}

func (a *AwsClient) SetConfig(config aws.Config) {
	a.config = config
}

func (a *AwsClient) SetClient() {
	a.client = aws.NewClient(a.config, aws.Metadata{
		APIVersion: "2006-03-01",
	})
}
