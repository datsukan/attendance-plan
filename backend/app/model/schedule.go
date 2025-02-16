package model

import "time"

// Schedule はスケジュールの model を表す構造体です。
type Schedule struct {
	ID        string
	UserID    string
	Name      string
	StartsAt  time.Time
	EndsAt    time.Time
	Color     string
	Type      ScheduleType
	Order     Order
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ScheduleType はスケジュールの種類を表す構造体です。
type ScheduleType string

const (
	ScheduleTypeMaster ScheduleType = "master" // 学事
	ScheduleTypeCustom ScheduleType = "custom" // 受講
)

// String は ScheduleType を文字列に変換します。
func (s ScheduleType) String() string {
	switch s {
	case ScheduleTypeMaster, ScheduleTypeCustom:
		return string(s)
	default:
		return ""
	}
}

// ToScheduleType は文字列を ScheduleType に変換します。
func ToScheduleType(s string) ScheduleType {
	return ScheduleType(s)
}

// Order はスケジュールの順番を表す構造体です。
type Order int

// Empty は Order が空かどうかを返します。
func (o Order) Empty() bool {
	return o == 0
}

func (o Order) Int() int {
	return int(o)
}
