package component

// ResponseResult は処理結果を表す構造体です。
type ResponseResult struct {
	StatusCode int
	HasError   bool
	Message    string
}

// NewResponseResult は ResponseResult を生成します。
func NewResponseResult(statusCode int, hasError bool, message string) ResponseResult {
	return ResponseResult{
		StatusCode: statusCode,
		HasError:   hasError,
		Message:    message,
	}
}
