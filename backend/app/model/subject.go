package model

import "time"

// Subject は科目の model を表す構造体です。
type Subject struct {
	ID        string
	UserID    string
	Name      string
	Color     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
