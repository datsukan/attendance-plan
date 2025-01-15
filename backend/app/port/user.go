package port

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
type SignInOutputData BaseUserData

// SignUpInputData はサインアップの入力データを表す構造体です。
type SignUpInputData struct {
	Email    string
	Password string
	Name     string
}

// SignUpOutputData はサインアップの出力データを表す構造体です。
type SignUpOutputData BaseUserData

// UserInputPort はユーザーのユースケースを表すインターフェースです。
type UserInputPort interface {
	SignIn(input SignInInputData)
	SignUp(input SignUpInputData)
}

// UserOutputPort はユーザーのユースケースの外部出力を表すインターフェースです。
type UserOutputPort interface {
	GetResponse() (int, string)
	SetResponseSignIn(output *SignInOutputData, result Result)
	SetResponseSignUp(output *SignUpOutputData, result Result)
}
