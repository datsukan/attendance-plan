package repository

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// SessionRepository はセッションの repository を表すインターフェースです。
type SessionRepository interface {
	GenerateToken(userID string) (token string, err error)
	IsValidToken(token string) (valid bool, userID string)
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
func (r *SessionRepositoryImpl) GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * time.Duration(r.TokenLifeTime)).Unix(),
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

	baseUserID, ok := claims["user_id"]
	if !ok {
		return false, ""
	}

	userID, ok := baseUserID.(string)
	if !ok {
		return false, ""
	}

	baseExp, ok := claims["exp"]
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

	return jt.Valid, userID
}
