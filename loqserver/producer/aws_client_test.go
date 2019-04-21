package producer

import (
	"testing"

	aws "github.com/aws/aws-sdk-go-v2/aws"
)

type AwsClientTestSuite struct {
	awsClient *AwsClient
}

var awsClientTestSuite AwsClientTestSuite

func init() {
	awsClientTestSuite.awsClient = &AwsClient{}
}

func TestSetConfig(t *testing.T) {
	awsClient := awsClientTestSuite.awsClient
	expectedAwsConfig := aws.Config{
		Region: "ap-southeast-1",
	}
	awsClient.SetConfig(expectedAwsConfig)
	actualConfig := awsClient.GetConfig()

	if actualConfig.Region != expectedAwsConfig.Region {
		t.Errorf("You have actual result '" + actualConfig.Region + "' and expected result '" + expectedAwsConfig.Region + "'")
	}
}
