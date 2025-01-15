package model

import "time"

// ExpirationTime はセッションの有効期限を表す定数です。
const ExpirationTime = 24 * time.Hour * 30

// Session はサービスを利用する際のセッションの model を表す構造体です。
type Session struct {
	ID        string
	UserID    string
	ExpiresAt time.Time
}

// IsExpired はセッションが有効期限切れかどうかを返します。
func (s *Session) IsExpired() bool {
	return s.ExpiresAt.Before(time.Now())
}

// ExtendExpiresAt はセッションの有効期限を延長します。
// 事前に値がセットされていない場合も使用できます。
func (s *Session) ExtendExpiresAt() {
	s.ExpiresAt = time.Now().Add(ExpirationTime)
}
