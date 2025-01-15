package response

import (
	"time"

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
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// GetScheduleListResponse はスケジュールリスト取得のレスポンスを表す構造体です。
type GetScheduleListResponse struct {
	Schedules []ScheduleResponse `json:"schedules"`
}

// GetScheduleResponse はスケジュール取得のレスポンスを表す構造体です。
type GetScheduleResponse ScheduleResponse

// PostScheduleResponse はスケジュール登録のレスポンスを表す構造体です。
type PostScheduleResponse ScheduleResponse

// PutScheduleResponse はスケジュール更新のレスポンスを表す構造体です。
type PutScheduleResponse ScheduleResponse

// ToGetScheduleListResponse はスケジュールリスト取得のレスポンスに変換します。
func ToGetScheduleListResponse(output *port.GetScheduleListOutputData) GetScheduleListResponse {
	if output == nil || len(output.Schedules) == 0 {
		return GetScheduleListResponse{
			Schedules: []ScheduleResponse{},
		}
	}

	var ss []ScheduleResponse
	for _, s := range output.Schedules {
		ss = append(ss, ScheduleResponse(s))
	}

	return GetScheduleListResponse{
		Schedules: ss,
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
		ID:        output.ID,
		UserID:    output.UserID,
		Name:      output.Name,
		StartsAt:  output.StartsAt.Format(time.DateTime),
		EndsAt:    output.EndsAt.Format(time.DateTime),
		Color:     output.Color,
		Type:      output.Type.String(),
		CreatedAt: output.CreatedAt.Format(time.DateTime),
		UpdatedAt: output.UpdatedAt.Format(time.DateTime),
	}
}

// ToPutScheduleResponse はスケジュール更新のレスポンスに変換します。
func ToPutScheduleResponse(output *port.UpdateScheduleOutputData) PutScheduleResponse {
	if output == nil {
		return PutScheduleResponse{}
	}

	return PutScheduleResponse(output.Schedule)
}
