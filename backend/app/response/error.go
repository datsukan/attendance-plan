package response

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

// NewError はエラーレスポンスを生成します。
func NewError(statusCode int, message string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       ToErrorBody(message),
	}, nil
}

func ToErrorBody(message string) string {
	return fmt.Sprintf(`{"message": "%s"}`, message)
}
