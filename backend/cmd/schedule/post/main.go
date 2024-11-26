package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/datsukan/attendance-plan/backend/app/schedule"
)

func main() {
	lambda.Start(schedule.PostSchedule)
}
