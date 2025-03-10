package infrastructure

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

func NewDB() *dynamo.DB {
	dynamoDbRegion := os.Getenv("AWS_REGION")

	// デフォルトでは東京リージョンを指定する。
	if len(dynamoDbRegion) == 0 {
		dynamoDbRegion = "ap-northeast-1"
	}

	dynamoDbEndpoint := os.Getenv("DYNAMO_ENDPOINT")

	config := &aws.Config{
		Region:   aws.String(dynamoDbRegion),
		Endpoint: aws.String(dynamoDbEndpoint),
	}

	// DynamoDB Local を利用する場合
	if os.Getenv("ENV") != "prd" {
		config.DisableSSL = aws.Bool(true)
		config.CredentialsChainVerboseErrors = aws.Bool(true)
		config.Credentials = credentials.NewStaticCredentials("dummy", "dummy", "dummy")
	}

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}

	return dynamo.New(sess)
}
