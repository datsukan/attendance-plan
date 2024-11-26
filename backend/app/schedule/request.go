package schedule

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
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
}

// PutScheduleRequest はスケジュール更新のリクエストを表す構造体です。
type PutScheduleRequest struct {
	ScheduleID string `json:"id"`
	Name       string `json:"name"`
	StartsAt   string `json:"starts_at"`
	EndsAt     string `json:"ends_at"`
	Color      string `json:"color"`
	Type       string `json:"type"`
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
		return fmt.Errorf("user_id is empty")
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
		return fmt.Errorf("schedule_id is empty")
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
		return fmt.Errorf("name is empty")
	}

	// name が50文字より多い
	const upperNameLength = 50
	if len(name) > upperNameLength {
		return fmt.Errorf("name must be %d characters or less", upperNameLength)
	}

	// starts_at が空文字
	if startsAt == "" {
		return fmt.Errorf("starts_at is empty")
	}

	// ends_at が空文字
	if endsAt == "" {
		return fmt.Errorf("ends_at is empty")
	}

	// starts_at のフォーマットが正しくない
	sa, err := time.Parse(time.DateTime, startsAt)
	if err != nil {
		return fmt.Errorf("starts_at is invalid")
	}

	// ends_at のフォーマットが正しくない
	ea, err := time.Parse(time.DateTime, endsAt)
	if err != nil {
		return fmt.Errorf("ends_at is invalid")
	}

	// starts_at が ends_at より後
	if sa.After(ea) {
		return fmt.Errorf("starts_at must be earlier than or equal to ends_at")
	}

	// color が空文字
	if color == "" {
		return fmt.Errorf("color is empty")
	}

	// type が空文字
	if sType == "" {
		return fmt.Errorf("type is empty")
	}

	// type が不正
	if ScheduleType(sType).String() == "" {
		return fmt.Errorf("type is invalid")
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
		return fmt.Errorf("schedule_id is empty")
	}

	return ValidateInputScheduleRequest(req.Name, req.StartsAt, req.EndsAt, req.Color, req.Type)
}

// ToDeleteScheduleRequest は APIGatewayProxyRequest から DeleteScheduleRequest に変換します。
func ToDeleteScheduleRequest(r events.APIGatewayProxyRequest) *DeleteScheduleRequest {
	return &DeleteScheduleRequest{ScheduleID: r.PathParameters["schedule_id"]}
}

// ValidateDeleteScheduleRequest は DeleteScheduleRequest のバリデーションを行います。
func ValidateDeleteScheduleRequest(req *DeleteScheduleRequest) error {
	if req.ScheduleID == "" {
		return fmt.Errorf("schedule_id is empty")
	}
	return nil
}
