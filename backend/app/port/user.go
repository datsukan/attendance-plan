package port

import "context"

// BaseUserData はユーザーの基本データを表す構造体です。
type BaseUserData struct {
	ID        string
	Email     string
	Name      string
	CreatedAt string
	UpdatedAt string
}

// SignInInputData はサインインの入力データを表す構造体です。
type SignInInputData struct {
	Email    string
	Password string
}

// SignInOutputData はサインインの出力データを表す構造体です。
type SignInOutputData struct {
	BaseUserData
	SessionToken string
}

// SignUpInputData はサインアップの入力データを表す構造体です。
type SignUpInputData struct {
	Email string
}

// SignUpOutputData はサインアップの出力データを表す構造体です。
type SignUpOutputData struct{}

// PasswordResetInputData はパスワードリセットの入力データを表す構造体です。
type PasswordResetInputData struct {
	Email string
}

// PasswordResetOutputData はパスワードリセットの出力データを表す構造体です。
type PasswordResetOutputData struct{}

// PasswordSetInputData はパスワード設定の入力データを表す構造体です。
type PasswordSetInputData struct {
	Token    string
	Password string
}

// PasswordSetOutputData はパスワード設定の出力データを表す構造体です。
type PasswordSetOutputData struct{}

// UserInputPort はユーザーのユースケースを表すインターフェースです。
type UserInputPort interface {
	SignIn(input SignInInputData)
	SignUp(ctx context.Context, input SignUpInputData)
	PasswordReset(ctx context.Context, input PasswordResetInputData)
	PasswordSet(input PasswordSetInputData)
}

// UserOutputPort はユーザーのユースケースの外部出力を表すインターフェースです。
type UserOutputPort interface {
	GetResponse() (statusCode int, body string)
	SetResponseSignIn(output *SignInOutputData, result Result)
	SetResponseSignUp(output *SignUpOutputData, result Result)
	SetResponsePasswordReset(output *PasswordResetOutputData, result Result)
	SetResponsePasswordSet(output *PasswordSetOutputData, result Result)
}
