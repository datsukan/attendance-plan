package request

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/datsukan/attendance-plan/backend/app/model"
)

// GetScheduleListRequest はスケジュールリスト取得のリクエストを表す構造体です。
type GetScheduleListRequest struct {
	UserID string
}

// GetScheduleRequest はスケジュール取得のリクエストを表す構造体です。
type GetScheduleRequest struct {
	ScheduleID string
}

// PostScheduleRequest はスケジュール登録のリクエストを表す構造体です。
type PostScheduleRequest struct {
	Name     string `json:"name"`
	StartsAt string `json:"starts_at"`
	EndsAt   string `json:"ends_at"`
	Color    string `json:"color"`
	Type     string `json:"type"`
	Order    int    `json:"order"`
}

// PutScheduleRequest はスケジュール更新のリクエストを表す構造体です。
type PutScheduleRequest struct {
	ScheduleID string `json:"id"`
	Name       string `json:"name"`
	StartsAt   string `json:"starts_at"`
	EndsAt     string `json:"ends_at"`
	Color      string `json:"color"`
	Type       string `json:"type"`
	Order      int    `json:"order"`
}

type PutBulkScheduleRequest struct {
	Schedules []PutScheduleRequest `json:"schedules"`
}

// DeleteScheduleRequest はスケジュール削除のリクエストを表す構造体です。
type DeleteScheduleRequest struct {
	ScheduleID string
}

// ToGetScheduleListRequest は APIGatewayProxyRequest から GetScheduleListRequest に変換します。
func ToGetScheduleListRequest(r events.APIGatewayProxyRequest) *GetScheduleListRequest {
	return &GetScheduleListRequest{UserID: r.PathParameters["user_id"]}
}

// ValidateGetScheduleListRequest は GetScheduleListRequest のバリデーションを行います。
func ValidateGetScheduleListRequest(req *GetScheduleListRequest) error {
	if req.UserID == "" {
		return fmt.Errorf("ユーザーIDを指定してください")
	}
	return nil
}

// ToGetScheduleRequest は APIGatewayProxyRequest から GetScheduleRequest に変換します。
func ToGetScheduleRequest(r events.APIGatewayProxyRequest) *GetScheduleRequest {
	return &GetScheduleRequest{ScheduleID: r.PathParameters["schedule_id"]}
}

// ValidateGetScheduleRequest は GetScheduleRequest のバリデーションを行います。
func ValidateGetScheduleRequest(req *GetScheduleRequest) error {
	if req.ScheduleID == "" {
		return fmt.Errorf("スケジュールIDを指定してください")
	}
	return nil
}

// ToPostScheduleRequest は APIGatewayProxyRequest から PostScheduleRequest に変換します。
func ToPostScheduleRequest(r events.APIGatewayProxyRequest) (*PostScheduleRequest, error) {
	var req PostScheduleRequest
	if err := json.Unmarshal([]byte(r.Body), &req); err != nil {
		return nil, err
	}
	return &req, nil
}

// ValidateInputScheduleRequest はスケジュールの入力に対するバリデーションを行います。
func ValidateInputScheduleRequest(name, startsAt, endsAt, color, sType string) error {
	// name が空文字
	if name == "" {
		return fmt.Errorf("スケジュール名を入力してください")
	}

	// name が50文字より多い
	const upperNameLength = 50
	if len(name) > upperNameLength {
		return fmt.Errorf("スケジュール名は%d文字以内で入力してください", upperNameLength)
	}

	// starts_at が空文字
	if startsAt == "" {
		return fmt.Errorf("開始日を入力してください")
	}

	// ends_at が空文字
	if endsAt == "" {
		return fmt.Errorf("終了日を入力してください")
	}

	// starts_at のフォーマットが正しくない
	sa, err := time.Parse(time.DateTime, startsAt)
	if err != nil {
		return fmt.Errorf("開始日は yyyy-MM-dd HH:mm:ss の形式で入力してください")
	}

	// ends_at のフォーマットが正しくない
	ea, err := time.Parse(time.DateTime, endsAt)
	if err != nil {
		return fmt.Errorf("終了日は yyyy-MM-dd HH:mm:ss の形式で入力してください")
	}

	// starts_at が ends_at より後
	if sa.After(ea) {
		return fmt.Errorf("終了日は開始日以降の日付を入力してください")
	}

	// color が空文字
	if color == "" {
		return fmt.Errorf("色を指定してください")
	}

	// type が空文字
	if sType == "" {
		return fmt.Errorf("スケジュールの種類を指定してください")
	}

	// type が不正
	if model.ScheduleType(sType).String() == "" {
		return fmt.Errorf("スケジュールの種類は %s または %s を指定してください", model.ScheduleTypeMaster, model.ScheduleTypeCustom)
	}

	return nil
}

// ValidatePostScheduleRequest は PostScheduleRequest のバリデーションを行います。
func ValidatePostScheduleRequest(req *PostScheduleRequest) error {
	return ValidateInputScheduleRequest(req.Name, req.StartsAt, req.EndsAt, req.Color, req.Type)
}

// ToPutScheduleRequest は APIGatewayProxyRequest から PutScheduleRequest に変換します。
func ToPutScheduleRequest(r events.APIGatewayProxyRequest) (*PutScheduleRequest, error) {
	var req PutScheduleRequest
	if err := json.Unmarshal([]byte(r.Body), &req); err != nil {
		return nil, err
	}

	req.ScheduleID = r.PathParameters["schedule_id"]

	return &req, nil
}

// ValidatePutScheduleRequest は PutScheduleRequest のバリデーションを行います。
func ValidatePutScheduleRequest(req *PutScheduleRequest) error {
	// ID が空文字
	if req.ScheduleID == "" {
		return fmt.Errorf("スケジュールIDを指定してください")
	}

	return ValidateInputScheduleRequest(req.Name, req.StartsAt, req.EndsAt, req.Color, req.Type)
}

// ToPutBulkScheduleRequest は APIGatewayProxyRequest から PutBulkScheduleRequest に変換します。
func ToPutBulkScheduleRequest(r events.APIGatewayProxyRequest) (*PutBulkScheduleRequest, error) {
	var req PutBulkScheduleRequest
	if err := json.Unmarshal([]byte(r.Body), &req); err != nil {
		return nil, err
	}

	return &req, nil
}

// ValidatePutBulkScheduleRequest は PutBulkScheduleRequest のバリデーションを行います。
func ValidatePutBulkScheduleRequest(req *PutBulkScheduleRequest) error {
	if len(req.Schedules) == 0 {
		return fmt.Errorf("スケジュールを指定してください")
	}

	for i, schedule := range req.Schedules {
		// ID が空文字
		if schedule.ScheduleID == "" {
			return fmt.Errorf("スケジュールIDを指定してください: %d番目", i+1)
		}

		if err := ValidateInputScheduleRequest(schedule.Name, schedule.StartsAt, schedule.EndsAt, schedule.Color, schedule.Type); err != nil {
			return fmt.Errorf("%s: %d番目", err.Error(), i+1)
		}
	}
	return nil
}

// ToDeleteScheduleRequest は APIGatewayProxyRequest から DeleteScheduleRequest に変換します。
func ToDeleteScheduleRequest(r events.APIGatewayProxyRequest) *DeleteScheduleRequest {
	return &DeleteScheduleRequest{ScheduleID: r.PathParameters["schedule_id"]}
}

// ValidateDeleteScheduleRequest は DeleteScheduleRequest のバリデーションを行います。
func ValidateDeleteScheduleRequest(req *DeleteScheduleRequest) error {
	if req.ScheduleID == "" {
		return fmt.Errorf("スケジュールIDを指定してください")
	}
	return nil
}
