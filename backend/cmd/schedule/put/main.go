package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/datsukan/attendance-plan/backend/app/handler"
)

func main() {
	lambda.Start(handler.PutSchedule)
}
