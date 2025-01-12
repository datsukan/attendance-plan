package user

import "time"

// User はユーザーの model を表す構造体です。
type User struct {
	ID        string
	Email     string
	Password  string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
