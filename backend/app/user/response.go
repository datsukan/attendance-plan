package user

// UserResponse はユーザーのレスポンスを表す構造体です。
type UserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// SignInResponse はサインインのレスポンスを表す構造体です。
type SignInResponse UserResponse

// SignUpResponse はサインアップのレスポンスを表す構造体です。
type SignUpResponse UserResponse

// ToSignInResponse はサインインのレスポンスに変換します。
func ToSignInResponse(output *SignInOutputData) SignInResponse {
	if output == nil {
		return SignInResponse{}
	}

	return SignInResponse(*output)
}

// ToSignUpResponse はサインアップのレスポンスに変換します。
func ToSignUpResponse(output *SignUpOutputData) SignUpResponse {
	if output == nil {
		return SignUpResponse{}
	}

	return SignUpResponse(*output)
}
