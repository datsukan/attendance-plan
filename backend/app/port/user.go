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

// GetUserInputData はユーザー取得の入力データを表す構造体です。
type GetUserInputData struct {
	UserID string
}

// GetUserOutputData はユーザー取得の出力データを表す構造体です。
type GetUserOutputData struct {
	BaseUserData
}

// UpdateUserInputData はユーザー情報更新の入力データを表す構造体です。
type UpdateUserInputData struct {
	UserID string
	Name   string
}

// UpdateUserOutputData はユーザー情報更新の出力データを表す構造体です。
type UpdateUserOutputData struct {
	BaseUserData
}

// DeleteUserInputData はユーザー削除の入力データを表す構造体です。
type DeleteUserInputData struct {
	UserID string
}

// DeleteUserOutputData はユーザー削除の出力データを表す構造体です。
type DeleteUserOutputData struct{}

// ResetEmailInputData はメールアドレスリセットの入力データを表す構造体です。
type ResetEmailInputData struct {
	UserID string
	Email  string
}

// ResetEmailOutputData はメールアドレスリセットの出力データを表す構造体です。
type ResetEmailOutputData struct{}

// SetEmailInputData はメールアドレス設定の入力データを表す構造体です。
type SetEmailInputData struct {
	UserIDToken string
	EmailToken  string
}

// SetEmailOutputData はメールアドレス設定の出力データを表す構造体です。
type SetEmailOutputData struct{}

// UserInputPort はユーザーのユースケースを表すインターフェースです。
type UserInputPort interface {
	SignIn(input SignInInputData)
	SignUp(ctx context.Context, input SignUpInputData)
	PasswordReset(ctx context.Context, input PasswordResetInputData)
	PasswordSet(input PasswordSetInputData)
	GetUser(input GetUserInputData)
	UpdateUser(input UpdateUserInputData)
	DeleteUser(input DeleteUserInputData)
	ResetEmail(input ResetEmailInputData)
	SetEmail(input SetEmailInputData)
}

// UserOutputPort はユーザーのユースケースの外部出力を表すインターフェースです。
type UserOutputPort interface {
	GetResponse() (statusCode int, body string)
	SetResponseSignIn(output *SignInOutputData, result Result)
	SetResponseSignUp(output *SignUpOutputData, result Result)
	SetResponsePasswordReset(output *PasswordResetOutputData, result Result)
	SetResponsePasswordSet(output *PasswordSetOutputData, result Result)
	SetResponseGetUser(output *GetUserOutputData, result Result)
	SetResponseUpdateUser(output *UpdateUserOutputData, result Result)
	SetResponseDeleteUser(output *DeleteUserOutputData, result Result)
	SetResponseResetEmail(output *ResetEmailOutputData, result Result)
	SetResponseSetEmail(output *SetEmailOutputData, result Result)
}
