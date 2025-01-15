package request

import (
	"errors"
	"testing"

	"github.com/aws/aws-lambda-go/events"
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
			want: errors.New("email is empty"),
		},
		{
			name: "異常系: password が未指定の場合はエラー",
			req: &SignInRequest{
				Email:    "test@example.com",
				Password: "",
			},
			want: errors.New("password is empty"),
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
		Body: `{"email": "test@example.com", "password": "test-password", "name": "test user"}`,
	}
	req := ToSignUpRequest(r)

	assert := assert.New(t)
	assert.Equal("test@example.com", req.Email)
	assert.Equal("test-password", req.Password)
	assert.Equal("test user", req.Name)
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
				Email:    "",
				Password: "test-password",
				Name:     "test user",
			},
			want: errors.New("email is empty"),
		},
		{
			name: "異常系: password が未指定の場合はエラー",
			req: &SignUpRequest{
				Email:    "test@example.com",
				Password: "",
				Name:     "test user",
			},
			want: errors.New("password is empty"),
		},
		{
			name: "異常系: name が未指定の場合はエラー",
			req: &SignUpRequest{
				Email:    "test@example.com",
				Password: "test-password",
				Name:     "",
			},
			want: errors.New("name is empty"),
		},
		{
			name: "正常系",
			req: &SignUpRequest{
				Email:    "test@example.com",
				Password: "test-password",
				Name:     "test user",
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
