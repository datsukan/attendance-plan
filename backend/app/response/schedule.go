package response

import (
	"github.com/datsukan/attendance-plan/backend/app/port"
)

// ScheduleResponse はスケジュールのレスポンスを表す構造体です。
type ScheduleResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	StartsAt  string `json:"starts_at"`
	EndsAt    string `json:"ends_at"`
	Color     string `json:"color"`
	Type      string `json:"type"`
	Order     int    `json:"order"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ScheduleResponseDateItem struct {
	Date      string             `json:"date"`
	Type      string             `json:"type"`
	Schedules []ScheduleResponse `json:"schedules"`
}

// GetScheduleListResponse はスケジュールリスト取得のレスポンスを表す構造体です。
type GetScheduleListResponse struct {
	MasterSchedules []ScheduleResponseDateItem `json:"master_schedules"`
	CustomSchedules []ScheduleResponseDateItem `json:"custom_schedules"`
}

// GetScheduleResponse はスケジュール取得のレスポンスを表す構造体です。
type GetScheduleResponse ScheduleResponse

// PostScheduleResponse はスケジュール登録のレスポンスを表す構造体です。
type PostScheduleResponse ScheduleResponse

// PostBulkScheduleResponse はスケジュール一括登録のレスポンスを表す構造体です。
type PostBulkScheduleResponse struct {
	Schedules []ScheduleResponse `json:"schedules"`
}

// PutScheduleResponse はスケジュール更新のレスポンスを表す構造体です。
type PutScheduleResponse ScheduleResponse

type PutBulkScheduleResponse struct {
	Schedules []ScheduleResponse `json:"schedules"`
}

// ToGetScheduleListResponse はスケジュールリスト取得のレスポンスに変換します。
func ToGetScheduleListResponse(output *port.GetScheduleListOutputData) GetScheduleListResponse {
	if output == nil || (len(output.MasterSchedules) == 0 && len(output.CustomSchedules) == 0) {
		return GetScheduleListResponse{
			MasterSchedules: []ScheduleResponseDateItem{},
			CustomSchedules: []ScheduleResponseDateItem{},
		}
	}

	var ms []ScheduleResponseDateItem
	for _, s := range output.MasterSchedules {
		di := ScheduleResponseDateItem{
			Date:      s.Date,
			Type:      s.Type,
			Schedules: []ScheduleResponse{},
		}

		for _, ss := range s.Schedules {
			di.Schedules = append(di.Schedules, ScheduleResponse(ss))
		}

		ms = append(ms, di)
	}

	var cs []ScheduleResponseDateItem
	for _, s := range output.CustomSchedules {
		di := ScheduleResponseDateItem{
			Date:      s.Date,
			Type:      s.Type,
			Schedules: []ScheduleResponse{},
		}

		for _, ss := range s.Schedules {
			di.Schedules = append(di.Schedules, ScheduleResponse(ss))
		}

		cs = append(cs, di)
	}

	return GetScheduleListResponse{
		MasterSchedules: ms,
		CustomSchedules: cs,
	}
}

// ToGetScheduleResponse はスケジュール取得のレスポンスに変換します。
func ToGetScheduleResponse(output *port.GetScheduleOutputData) GetScheduleResponse {
	if output == nil {
		return GetScheduleResponse{}
	}

	return GetScheduleResponse(output.Schedule)
}

// ToPostScheduleResponse はスケジュール登録のレスポンスに変換します。
func ToPostScheduleResponse(output *port.CreateScheduleOutputData) PostScheduleResponse {
	if output == nil {
		return PostScheduleResponse{}
	}

	return PostScheduleResponse{
		ID:        output.Schedule.ID,
		UserID:    output.Schedule.UserID,
		Name:      output.Schedule.Name,
		StartsAt:  output.Schedule.StartsAt,
		EndsAt:    output.Schedule.EndsAt,
		Color:     output.Schedule.Color,
		Type:      output.Schedule.Type,
		Order:     output.Schedule.Order,
		CreatedAt: output.Schedule.CreatedAt,
		UpdatedAt: output.Schedule.UpdatedAt,
	}
}

// ToPostBulkScheduleResponse はスケジュール一括登録のレスポンスに変換します。
func ToPostBulkScheduleResponse(output *port.CreateBulkScheduleOutputData) PostBulkScheduleResponse {
	if output == nil || len(output.Schedules) == 0 {
		return PostBulkScheduleResponse{
			Schedules: []ScheduleResponse{},
		}
	}

	var ss []ScheduleResponse
	for _, s := range output.Schedules {
		ss = append(ss, ScheduleResponse(s))
	}

	return PostBulkScheduleResponse{
		Schedules: ss,
	}
}

// ToPutScheduleResponse はスケジュール更新のレスポンスに変換します。
func ToPutScheduleResponse(output *port.UpdateScheduleOutputData) PutScheduleResponse {
	if output == nil {
		return PutScheduleResponse{}
	}

	return PutScheduleResponse(output.Schedule)
}

// ToPutBulkScheduleResponse はスケジュール一括更新のレスポンスに変換します。
func ToPutBulkScheduleResponse(output *port.UpdateBulkScheduleOutputData) PutBulkScheduleResponse {
	if output == nil || len(output.Schedules) == 0 {
		return PutBulkScheduleResponse{
			Schedules: []ScheduleResponse{},
		}
	}

	var ss []ScheduleResponse
	for _, s := range output.Schedules {
		ss = append(ss, ScheduleResponse(s))
	}

	return PutBulkScheduleResponse{
		Schedules: ss,
	}
}
