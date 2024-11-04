package schedule

import (
	"github.com/aws/aws-lambda-go/events"
)

// ResponseError はエラーレスポンスを生成します。
func NewResponseError(statusCode int, body string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
	}, nil
}
