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
	Name      string
	StartsAt  time.Time
	EndsAt    time.Time
	Color     string
	Type      ScheduleType
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s ScheduleType) String() string {
	switch s {
	case ScheduleTypeMaster, ScheduleTypeCustom:
		return string(s)
	default:
		return ""
	}
}

func ToScheduleType(s string) ScheduleType {
	return ScheduleType(s)
}
