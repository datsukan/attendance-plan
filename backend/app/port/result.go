package port

// Result は処理結果を表す構造体です。
type Result struct {
	StatusCode   int
	HasError     bool
	ErrorMessage string
}

// NewSuccessResult は成功時の Result を生成します。
func NewSuccessResult(statusCode int) Result {
	return Result{StatusCode: statusCode}
}

// NewErrorResult はエラー時の Result を生成します。
func NewErrorResult(statusCode int, errorMessage string) Result {
	return Result{
		StatusCode:   statusCode,
		HasError:     true,
		ErrorMessage: errorMessage,
	}
}
