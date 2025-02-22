package repository

import (
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	TokenKeyValue = "value"
	TokenKeyExp   = "exp"
)

// SessionRepository はセッションの repository を表すインターフェースです。
type SessionRepository interface {
	GenerateToken(value string) (token string, err error)
	IsValidToken(token string) (valid bool, value string)
}

// SessionRepositoryImpl はセッションの repository の実装を表す構造体です。
type SessionRepositoryImpl struct {
	SecretKey     string
	TokenLifeTime int
}

// NewSessionRepository は SessionRepository を生成します。
func NewSessionRepository(secretKey string, tokenLifeTime int) SessionRepository {
	return &SessionRepositoryImpl{SecretKey: secretKey, TokenLifeTime: tokenLifeTime}
}

// GenerateToken はセッショントークンを生成します。
func (r *SessionRepositoryImpl) GenerateToken(value string) (string, error) {
	claims := jwt.MapClaims{
		TokenKeyValue: value,
		TokenKeyExp:   time.Now().Add(time.Hour * time.Duration(r.TokenLifeTime)).Unix(),
	}
	jt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jt.SignedString([]byte(r.SecretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

// IsValidToken はセッショントークンが有効かどうかを判定します。
func (r *SessionRepositoryImpl) IsValidToken(token string) (bool, string) {
	jt, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(r.SecretKey), nil
	})
	if err != nil {
		return false, ""
	}

	claims, ok := jt.Claims.(jwt.MapClaims)
	if !ok {
		return false, ""
	}

	baseValue, ok := claims[TokenKeyValue]
	if !ok {
		return false, ""
	}

	value, ok := baseValue.(string)
	if !ok {
		return false, ""
	}

	baseExp, ok := claims[TokenKeyExp]
	if !ok {
		return false, ""
	}

	exp, ok := baseExp.(float64)
	if !ok {
		return false, ""
	}

	// 有効期限切れ
	if time.Now().Unix() > int64(exp) {
		return false, ""
	}

	return jt.Valid, value
}
