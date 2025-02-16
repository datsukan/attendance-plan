package response

import "github.com/datsukan/attendance-plan/backend/app/port"

// UserResponse はユーザーのレスポンスを表す構造体です。
type UserResponse struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	SessionToken string `json:"session_token"`
}

// SignInResponse はサインインのレスポンスを表す構造体です。
type SignInResponse UserResponse

// SignUpResponse はサインアップのレスポンスを表す構造体です。
type SignUpResponse UserResponse

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
