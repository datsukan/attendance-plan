package response

import "github.com/datsukan/attendance-plan/backend/app/port"

// UserResponse はユーザーのレスポンスを表す構造体です。
type UserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// SignInResponse はサインインのレスポンスを表す構造体です。
type SignInResponse struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	SessionToken string `json:"session_token"`
}

// GetUserResponse はユーザー取得のレスポンスを表す構造体です。
type GetUserResponse UserResponse

// PutUserResponse はユーザー情報更新のレスポンスを表す構造体です。
type PutUserResponse UserResponse

// ToSignInResponse はサインインのレスポンスに変換します。
func ToSignInResponse(output *port.SignInOutputData) SignInResponse {
	if output == nil {
		return SignInResponse{}
	}

	res := SignInResponse{
		ID:           output.ID,
		Email:        output.Email,
		Name:         output.Name,
		CreatedAt:    output.CreatedAt,
		UpdatedAt:    output.UpdatedAt,
		SessionToken: output.SessionToken,
	}

	return SignInResponse(res)
}

// ToGetUserResponse はユーザー取得のレスポンスに変換します。
func ToGetUserResponse(output *port.GetUserOutputData) GetUserResponse {
	if output == nil {
		return GetUserResponse{}
	}

	res := GetUserResponse{
		ID:        output.ID,
		Email:     output.Email,
		Name:      output.Name,
		CreatedAt: output.CreatedAt,
		UpdatedAt: output.UpdatedAt,
	}

	return GetUserResponse(res)
}

// ToPutUserResponse はユーザー情報更新のレスポンスに変換します。
func ToPutUserResponse(output *port.UpdateUserOutputData) PutUserResponse {
	if output == nil {
		return PutUserResponse{}
	}

	res := PutUserResponse{
		ID:        output.ID,
		Email:     output.Email,
		Name:      output.Name,
		CreatedAt: output.CreatedAt,
		UpdatedAt: output.UpdatedAt,
	}

	return PutUserResponse(res)
}
