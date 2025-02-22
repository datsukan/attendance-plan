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

// GetUserRequest はユーザー取得のリクエストパラメータの構造体です。
type GetUserRequest struct {
	UserID string
}

// UpdateUser はユーザー情報更新のリクエストパラメータの構造体です。
type PutUserRequest struct {
	UserID string
	Name   string `json:"name"`
}

// DeleteUser はユーザー削除のリクエストパラメータの構造体です。
type DeleteUserRequest struct {
	UserID string
}

// ResetEmail はメールアドレスリセットのリクエストパラメータの構造体です。
type ResetEmailRequest struct {
	UserID string
	Email  string `json:"email"`
}

// SetEmail はメールアドレス設定のリクエストパラメータの構造体です。
type SetEmailRequest struct {
	UserIDToken string `json:"id_token"`
	EmailToken  string `json:"email_token"`
}

// ToSignInRequest はサインインのリクエストパラメータへ変換します。
func ToSignInRequest(r events.APIGatewayProxyRequest) (*SignInRequest, error) {
	var req SignInRequest
	if err := json.Unmarshal([]byte(r.Body), &req); err != nil {
		return nil, err
	}

	return &req, nil
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
func ToSignUpRequest(r events.APIGatewayProxyRequest) (*SignUpRequest, error) {
	var req SignUpRequest
	if err := json.Unmarshal([]byte(r.Body), &req); err != nil {
		return nil, err
	}

	return &req, nil
}

// ValidateSignUpRequest はサインアップのリクエストパラメータを検証します。
func ValidateSignUpRequest(req *SignUpRequest) error {
	if req.Email == "" {
		return fmt.Errorf("メールアドレスを入力してください")
	}

	return nil
}

// ToPasswordResetRequest はパスワードリセットのリクエストパラメータへ変換します。
func ToPasswordResetRequest(r events.APIGatewayProxyRequest) (*PasswordResetRequest, error) {
	var req PasswordResetRequest
	if err := json.Unmarshal([]byte(r.Body), &req); err != nil {
		return nil, err
	}

	return &req, nil
}

// ValidatePasswordResetRequest はパスワードリセットのリクエストパラメータを検証します。
func ValidatePasswordResetRequest(req *PasswordResetRequest) error {
	if req.Email == "" {
		return fmt.Errorf("メールアドレスを入力してください")
	}

	return nil
}

// ToPasswordSetRequest はパスワード設定のリクエストパラメータへ変換します。
func ToPasswordSetRequest(r events.APIGatewayProxyRequest) (*PasswordSetRequest, error) {
	var req PasswordSetRequest
	if err := json.Unmarshal([]byte(r.Body), &req); err != nil {
		return nil, err
	}

	return &req, nil
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

// ToGetUserRequest はユーザー取得のリクエストパラメータへ変換します。
func ToGetUserRequest(r events.APIGatewayProxyRequest) *GetUserRequest {
	return &GetUserRequest{
		UserID: r.PathParameters["user_id"],
	}
}

// ValidateGetUserRequest はユーザー取得のリクエストパラメータを検証します。
func ValidateGetUserRequest(req *GetUserRequest) error {
	if req.UserID == "" {
		return fmt.Errorf("ユーザーIDが指定されていません")
	}

	return nil
}

// ToPutUserRequest はユーザー情報更新のリクエストパラメータへ変換します。
func ToPutUserRequest(r events.APIGatewayProxyRequest) (*PutUserRequest, error) {
	var req PutUserRequest
	if err := json.Unmarshal([]byte(r.Body), &req); err != nil {
		return nil, err
	}

	req.UserID = r.PathParameters["user_id"]

	return &req, nil
}

// ValidatePutUserRequest はユーザー情報更新のリクエストパラメータを検証します。
func ValidatePutUserRequest(req *PutUserRequest) error {
	if req.UserID == "" {
		return fmt.Errorf("ユーザーIDが指定されていません")
	}

	return nil
}

// ToDeleteUserRequest はユーザー削除のリクエストパラメータへ変換します。
func ToDeleteUserRequest(r events.APIGatewayProxyRequest) *DeleteUserRequest {
	return &DeleteUserRequest{
		UserID: r.PathParameters["user_id"],
	}
}

// ValidateDeleteUserRequest はユーザー削除のリクエストパラメータを検証します。
func ValidateDeleteUserRequest(req *DeleteUserRequest) error {
	if req.UserID == "" {
		return fmt.Errorf("ユーザーIDが指定されていません")
	}

	return nil
}

// ToResetEmailRequest はメールアドレスリセットのリクエストパラメータへ変換します。
func ToResetEmailRequest(r events.APIGatewayProxyRequest) (*ResetEmailRequest, error) {
	var req ResetEmailRequest
	if err := json.Unmarshal([]byte(r.Body), &req); err != nil {
		return nil, err
	}

	req.UserID = r.PathParameters["user_id"]

	return &req, nil
}

// ValidateResetEmailRequest はメールアドレスリセットのリクエストパラメータを検証します。
func ValidateResetEmailRequest(req *ResetEmailRequest) error {
	if req.UserID == "" {
		return fmt.Errorf("ユーザーIDが指定されていません")
	}

	if req.Email == "" {
		return fmt.Errorf("メールアドレスを入力してください")
	}

	return nil
}

// ToSetEmailRequest はメールアドレス設定のリクエストパラメータへ変換します。
func ToSetEmailRequest(r events.APIGatewayProxyRequest) (*SetEmailRequest, error) {
	var req SetEmailRequest
	if err := json.Unmarshal([]byte(r.Body), &req); err != nil {
		return nil, err
	}

	return &req, nil
}

// ValidateSetEmailRequest はメールアドレス設定のリクエストパラメータを検証します。
func ValidateSetEmailRequest(req *SetEmailRequest) error {
	if req.UserIDToken == "" {
		return fmt.Errorf("ユーザーIDトークンが指定されていません")
	}

	if req.EmailToken == "" {
		return fmt.Errorf("メールアドレストークンが指定されていません")
	}

	return nil
}
