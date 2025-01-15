package port

import (
	"github.com/datsukan/attendance-plan/backend/app/model"
)

// BaseScheduleData はスケジュールの基本データを表す構造体です。
type BaseScheduleData struct {
	ID        string
	UserID    string
	Name      string
	StartsAt  string
	EndsAt    string
	Color     string
	Type      string
	CreatedAt string
	UpdatedAt string
}

// GetScheduleListInputData はスケジュールリスト取得の入力データを表す構造体です。
type GetScheduleListInputData struct {
	UserID string
}

// GetScheduleListOutputData はスケジュールリスト取得の出力データを表す構造体です。
type GetScheduleListOutputData struct {
	Schedules []BaseScheduleData
}

// GetScheduleInputData はスケジュール取得の入力データを表す構造体です。
type GetScheduleInputData struct {
	ScheduleID string
}

// GetScheduleOutputData はスケジュール取得の出力データを表す構造体です。
type GetScheduleOutputData struct {
	Schedule BaseScheduleData
}

// CreateScheduleData はスケジュール作成のスケジュールデータを表す構造体です。
type CreateScheduleData struct {
	UserID   string
	Name     string
	StartsAt string
	EndsAt   string
	Color    string
	Type     string
}

// CreateScheduleData はスケジュール作成のデータを表す構造体です。
type CreateScheduleInputData struct {
	Schedule CreateScheduleData
}

// CreateScheduleOutputData はスケジュール作成の出力データを表す構造体です。
type CreateScheduleOutputData struct {
	model.Schedule
}

// UpdateScheduleData はスケジュール更新のスケジュールデータを表す構造体です。
type UpdateScheduleData struct {
	ID       string
	Name     string
	StartsAt string
	EndsAt   string
	Color    string
	Type     string
}

// UpdateScheduleInputData はスケジュール更新の入力データを表す構造体です。
type UpdateScheduleInputData struct {
	Schedule UpdateScheduleData
}

// UpdateScheduleOutputData はスケジュール更新の出力データを表す構造体です。
type UpdateScheduleOutputData struct {
	Schedule BaseScheduleData
}

// DeleteScheduleInputData はスケジュール削除の入力データを表す構造体です。
type DeleteScheduleInputData struct {
	ScheduleID string
}

// DeleteScheduleOutputData はスケジュール削除の出力データを表す構造体です。
type DeleteScheduleOutputData struct {
	ScheduleID string
}

// ScheduleInputPort はスケジュールのユースケースを表すインターフェースです。
type ScheduleInputPort interface {
	GetScheduleList(input GetScheduleListInputData)
	GetSchedule(input GetScheduleInputData)
	CreateSchedule(input CreateScheduleInputData)
	UpdateSchedule(input UpdateScheduleInputData)
	DeleteSchedule(input DeleteScheduleInputData)
}

// ScheduleOutputPort はスケジュールのユースケースの外部出力を表すインターフェースです。
type ScheduleOutputPort interface {
	GetResponse() (int, string)
	SetResponseGetScheduleList(output *GetScheduleListOutputData, result Result)
	SetResponseGetSchedule(output *GetScheduleOutputData, result Result)
	SetResponseCreateSchedule(output *CreateScheduleOutputData, result Result)
	SetResponseUpdateSchedule(output *UpdateScheduleOutputData, result Result)
	SetResponseDeleteSchedule(output *DeleteScheduleOutputData, result Result)
}
