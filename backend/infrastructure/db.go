package infrastructure

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

func NewDB() *dynamo.DB {
	dynamoDbRegion := os.Getenv("AWS_REGION")
	disableSsl := false

	// DynamoDB Local を利用する場合は Endpoint の URL を設定する。
	dynamoDbEndpoint := os.Getenv("DYNAMO_ENDPOINT")
	if len(dynamoDbEndpoint) != 0 {
		disableSsl = true
	}

	// デフォルトでは東京リージョンを指定する。
	if len(dynamoDbRegion) == 0 {
		dynamoDbRegion = "ap-northeast-1"
	}

	sess, err := session.NewSession(&aws.Config{
		Region:     aws.String(dynamoDbRegion),
		Endpoint:   aws.String(dynamoDbEndpoint),
		DisableSSL: aws.Bool(disableSsl),
	})
	if err != nil {
		panic(err)
	}

	return dynamo.New(sess)
}
