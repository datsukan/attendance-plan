package request

import (
	"errors"
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/datsukan/attendance-plan/backend/app/model"
	"github.com/stretchr/testify/assert"
)

func TestToSignInRequest(t *testing.T) {
	r := events.APIGatewayProxyRequest{
		Body: `{"email": "test@example.com", "password": "test-password"}`,
	}
	req := ToSignInRequest(r)

	assert := assert.New(t)
	assert.Equal("test@example.com", req.Email)
	assert.Equal("test-password", req.Password)
}

func TestValidateSignInRequest(t *testing.T) {
	tests := []struct {
		name string
		req  *SignInRequest
		want error
	}{
		{
			name: "異常系: email が未指定の場合はエラー",
			req: &SignInRequest{
				Email:    "",
				Password: "test-password",
			},
			want: errors.New("メールアドレスを入力してください"),
		},
		{
			name: "異常系: password が未指定の場合はエラー",
			req: &SignInRequest{
				Email:    "test@example.com",
				Password: "",
			},
			want: errors.New("パスワードを入力してください"),
		},
		{
			name: "正常系",
			req: &SignInRequest{
				Email:    "test@example.com",
				Password: "test-password",
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ValidateSignInRequest(tt.req))
		})
	}
}

func TestToSignUpRequest(t *testing.T) {
	r := events.APIGatewayProxyRequest{
		Body: `{"email": "test@example.com"}`,
	}
	req := ToSignUpRequest(r)

	assert := assert.New(t)
	assert.Equal("test@example.com", req.Email)
}

func TestValidateSignUpRequest(t *testing.T) {
	tests := []struct {
		name string
		req  *SignUpRequest
		want error
	}{
		{
			name: "異常系: email が未指定の場合はエラー",
			req: &SignUpRequest{
				Email: "",
			},
			want: errors.New("メールアドレスを入力してください"),
		},
		{
			name: "正常系",
			req: &SignUpRequest{
				Email: "test@example.com",
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ValidateSignUpRequest(tt.req))
		})
	}
}

func TestToPasswordResetRequest(t *testing.T) {
	r := events.APIGatewayProxyRequest{
		Body: `{"email": "test@example.com"}`,
	}
	req := ToPasswordResetRequest(r)

	assert := assert.New(t)
	assert.Equal("test@example.com", req.Email)
}

func TestValidatePasswordResetRequest(t *testing.T) {
	tests := []struct {
		name string
		req  *PasswordResetRequest
		want error
	}{
		{
			name: "異常系: email が未指定の場合はエラー",
			req: &PasswordResetRequest{
				Email: "",
			},
			want: errors.New("メールアドレスを入力してください"),
		},
		{
			name: "正常系",
			req: &PasswordResetRequest{
				Email: "test@example.com",
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ValidatePasswordResetRequest(tt.req))
		})
	}
}

func TestToPasswordSetRequest(t *testing.T) {
	r := events.APIGatewayProxyRequest{
		Body: `{"token": "test-token", "password": "test-password"}`,
	}
	req := ToPasswordSetRequest(r)

	assert := assert.New(t)
	assert.Equal("test-token", req.Token)
	assert.Equal("test-password", req.Password)
}

func TestValidatePasswordSetRequest(t *testing.T) {
	tests := []struct {
		name string
		req  *PasswordSetRequest
		want error
	}{
		{
			name: "異常系: token が未指定の場合はエラー",
			req: &PasswordSetRequest{
				Token:    "",
				Password: "test-password",
			},
			want: errors.New("トークンが指定されていません"),
		},
		{
			name: "異常系: password が未指定の場合はエラー",
			req: &PasswordSetRequest{
				Token:    "test-token",
				Password: "",
			},
			want: errors.New("パスワードを入力してください"),
		},
		{
			name: "異常系: password の内容が条件を満たさない場合はエラー",
			req: &PasswordSetRequest{
				Token:    "test-token",
				Password: "test",
			},
			want: fmt.Errorf("パスワードは%d～%d文字以内にしてください", model.PasswordMinLength, model.PasswordMaxLength),
		},
		{
			name: "正常系",
			req: &PasswordSetRequest{
				Token:    "test-token",
				Password: "Password1!",
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ValidatePasswordSetRequest(tt.req))
		})
	}
}
