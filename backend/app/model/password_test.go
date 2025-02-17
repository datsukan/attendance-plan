package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassword_Validate(t *testing.T) {
	tests := []struct {
		name     string
		password Password
		wantErr  error
	}{
		{
			name:     "有効",
			password: "Password1!",
			wantErr:  nil,
		},
		{
			name:     "無効: 長さが足りない",
			password: "Pass1!",
			wantErr:  ErrPasswordLength,
		},
		{
			name:     "無効: 長さが長すぎる",
			password: Password(strings.Repeat("a", 71)),
			wantErr:  ErrPasswordLength,
		},
		{
			name:     "無効: 大文字の英字が含まれていない",
			password: "password1!",
			wantErr:  ErrPasswordUppercase,
		},
		{
			name:     "無効: 小文字の英字が含まれていない",
			password: "PASSWORD1!",
			wantErr:  ErrPasswordLowercase,
		},
		{
			name:     "無効: 数字が含まれていない",
			password: "Password!",
			wantErr:  ErrPasswordNumber,
		},
		{
			name:     "無効: 記号が含まれていない",
			password: "Password1",
			wantErr:  ErrPasswordSymbol,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.password.Validate()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
