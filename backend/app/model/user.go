package model

import (
	"time"
)

// User はユーザーの model を表す構造体です。
type User struct {
	ID        string
	Email     string
	Password  string
	Name      string
	Enabled   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
