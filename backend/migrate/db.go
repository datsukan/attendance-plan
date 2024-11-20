package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

func NewDB() *dynamo.DB {
	config := &aws.Config{
		Region:                        aws.String("ap-northeast-1"),
		Endpoint:                      aws.String("http://localhost:8000"),
		DisableSSL:                    aws.Bool(true),
		CredentialsChainVerboseErrors: aws.Bool(true),
		Credentials:                   credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
	}

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}

	return dynamo.New(sess)
}
