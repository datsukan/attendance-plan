package port

// Result は処理結果を表す構造体です。
type Result struct {
	StatusCode int
	HasError   bool
	Message    string
}

// NewResult は Result を生成します。
func NewResult(statusCode int, hasError bool, message string) Result {
	return Result{
		StatusCode: statusCode,
		HasError:   hasError,
		Message:    message,
	}
}
