package schedule

import "time"

// ScheduleType はスケジュールの種類を表す構造体です。
type ScheduleType string

const (
	ScheduleTypeMaster ScheduleType = "master" // 学事
	ScheduleTypeCustom ScheduleType = "custom" // 受講
)

// Schedule はスケジュールの model を表す構造体です。
type Schedule struct {
	ID        string
	UserID    string
	Name      string
	StartsAt  time.Time
	EndsAt    time.Time
	Color     string
	Type      ScheduleType
	CreatedAt time.Time
	UpdatedAt time.Time
}

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
