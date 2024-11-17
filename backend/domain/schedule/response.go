package schedule

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

// ResponseError はエラーレスポンスを生成します。
func NewResponseError(statusCode int, message string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       fmt.Sprintf(`{"message": "%s"}`, message),
	}, nil
}
