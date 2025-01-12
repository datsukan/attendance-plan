package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/datsukan/attendance-plan/backend/app/user"
)

func main() {
	lambda.Start(user.SignUp)
}
