package request

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/datsukan/attendance-plan/backend/app/model"
)

// SignInRequest はサインインのリクエストパラメータの構造体です。
type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignUpRequest はサインアップのリクエストパラメータの構造体です。
type SignUpRequest struct {
	Email string `json:"email"`
}

// PasswordResetRequest はパスワードリセットのリクエストパラメータの構造体です。
type PasswordResetRequest struct {
	Email string `json:"email"`
}

// PasswordSetRequest はパスワード設定のリクエストパラメータの構造体です。
type PasswordSetRequest struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

// ToSignInRequest はサインインのリクエストパラメータへ変換します。
func ToSignInRequest(r events.APIGatewayProxyRequest) *SignInRequest {
	req := &SignInRequest{}
	json.Unmarshal([]byte(r.Body), req)
	return req
}

// ValidateSignInRequest はサインインのリクエストパラメータを検証します。
func ValidateSignInRequest(req *SignInRequest) error {
	if req.Email == "" {
		return fmt.Errorf("メールアドレスを入力してください")
	}

	if req.Password == "" {
		return fmt.Errorf("パスワードを入力してください")
	}

	return nil
}

// ToSignUpRequest はサインアップのリクエストパラメータへ変換します。
func ToSignUpRequest(r events.APIGatewayProxyRequest) *SignUpRequest {
	req := &SignUpRequest{}
	json.Unmarshal([]byte(r.Body), req)
	return req
}

// ValidateSignUpRequest はサインアップのリクエストパラメータを検証します。
func ValidateSignUpRequest(req *SignUpRequest) error {
	if req.Email == "" {
		return fmt.Errorf("メールアドレスを入力してください")
	}

	return nil
}

// ToPasswordResetRequest はパスワードリセットのリクエストパラメータへ変換します。
func ToPasswordResetRequest(r events.APIGatewayProxyRequest) *PasswordResetRequest {
	req := &PasswordResetRequest{}
	json.Unmarshal([]byte(r.Body), req)
	return req
}

// ValidatePasswordResetRequest はパスワードリセットのリクエストパラメータを検証します。
func ValidatePasswordResetRequest(req *PasswordResetRequest) error {
	if req.Email == "" {
		return fmt.Errorf("メールアドレスを入力してください")
	}

	return nil
}

// ToPasswordSetRequest はパスワード設定のリクエストパラメータへ変換します。
func ToPasswordSetRequest(r events.APIGatewayProxyRequest) *PasswordSetRequest {
	req := &PasswordSetRequest{}
	json.Unmarshal([]byte(r.Body), req)
	return req
}

// ValidatePasswordSetRequest はパスワード設定のリクエストパラメータを検証します。
func ValidatePasswordSetRequest(req *PasswordSetRequest) error {
	if req.Token == "" {
		return fmt.Errorf("トークンが指定されていません")
	}

	if req.Password == "" {
		return fmt.Errorf("パスワードを入力してください")
	}

	password := model.Password(req.Password)
	if err := password.Validate(); err != nil {
		return err
	}

	return nil
}
