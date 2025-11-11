package repository

import (
	"testing"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	secretKey := "test-secret-key"
	tokenLifeTime := 1
	r := NewSessionRepository(secretKey, tokenLifeTime)

	assert := assert.New(t)

	token, err := r.GenerateToken("test-user-id")
	assert.NoError(err)
	assert.NotEmpty(token)
}

func TestIsValidToken(t *testing.T) {
	genToken := func(secretKey string, exp int64, userID string) string {
		claims := jwt.MapClaims{
			TokenKeyValue: userID,
			TokenKeyExp:   exp,
		}
		jt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token, _ := jt.SignedString([]byte(secretKey))

		return token
	}
	exp := time.Now().Add(time.Second * 10).Unix()

	tests := []struct {
		name       string
		secretKey  string
		token      string
		wantValid  bool
		wantUserID string
	}{
		{
			name:       "有効なトークンの場合、有効と判定され正しいユーザーIDが返される",
			secretKey:  "test-secret-key",
			token:      genToken("test-secret-key", exp, "test-user-id"),
			wantValid:  true,
			wantUserID: "test-user-id",
		},
		{
			name:       "無効なトークンの場合、無効と判定される",
			secretKey:  "test-secret-key",
			token:      genToken("test-dummy-secret-key", exp, "test-user-id"),
			wantValid:  false,
			wantUserID: "",
		},
		{
			name:       "有効期限切れのトークンの場合、無効と判定される",
			secretKey:  "test-secret-key",
			token:      genToken("test-secret-key", time.Now().Add(time.Second*-10).Unix(), "test-user-id"),
			wantValid:  false,
			wantUserID: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			r := NewSessionRepository(tt.secretKey, 1)
			valid, userID := r.IsValidToken(tt.token)
			assert.Equal(tt.wantValid, valid)
			assert.Equal(tt.wantUserID, userID)
		})
	}
}
