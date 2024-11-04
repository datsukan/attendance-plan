package schedule

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

type GetScheduleListRequest struct {
	UserID string
}

type GetScheduleRequest struct {
	ScheduleID string
}

type PostScheduleRequest struct {
	Name     string `json:"name"`
	StartsAt string `json:"starts_at"`
	EndsAt   string `json:"ends_at"`
	Color    string `json:"color"`
	Type     string `json:"type"`
}

type PutScheduleRequest struct {
	ScheduleID string `json:"id"`
	Name       string `json:"name"`
	StartsAt   string `json:"starts_at"`
	EndsAt     string `json:"ends_at"`
	Color      string `json:"color"`
	Type       string `json:"type"`
}

func ToGetScheduleListRequest(r events.APIGatewayProxyRequest) *GetScheduleListRequest {
	return &GetScheduleListRequest{UserID: r.PathParameters["user_id"]}
}

func ValidateGetScheduleListRequest(req *GetScheduleListRequest) error {
	if req.UserID == "" {
		return fmt.Errorf("user_id is empty")
	}
	return nil
}

func ToGetScheduleRequest(r events.APIGatewayProxyRequest) *GetScheduleRequest {
	return &GetScheduleRequest{ScheduleID: r.PathParameters["schedule_id"]}
}

func ValidateGetScheduleRequest(req *GetScheduleRequest) error {
	if req.ScheduleID == "" {
		return fmt.Errorf("schedule_id is empty")
	}
	return nil
}

func ToPostScheduleRequest(r events.APIGatewayProxyRequest) (*PostScheduleRequest, error) {
	var req PostScheduleRequest
	if err := json.Unmarshal([]byte(r.Body), &req); err != nil {
		return nil, err
	}
	return &req, nil
}

func ValidateInputScheduleRequest(name, startsAt, endsAt, color, sType string) error {
	// name が空文字
	if name == "" {
		return fmt.Errorf("name is empty")
	}

	// name が50文字より大きい
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

func ValidatePostScheduleRequest(req *PostScheduleRequest) error {
	return ValidateInputScheduleRequest(req.Name, req.StartsAt, req.EndsAt, req.Color, req.Type)
}

func ToPutScheduleRequest(r events.APIGatewayProxyRequest) (*PutScheduleRequest, error) {
	var req PutScheduleRequest
	if err := json.Unmarshal([]byte(r.Body), &req); err != nil {
		return nil, err
	}

	req.ScheduleID = r.PathParameters["schedule_id"]

	return &req, nil
}

func ValidatePutScheduleRequest(req *PutScheduleRequest) error {
	// ID が空文字
	if req.ScheduleID == "" {
		return fmt.Errorf("schedule_id is empty")
	}

	return ValidateInputScheduleRequest(req.Name, req.StartsAt, req.EndsAt, req.Color, req.Type)
}

func ToDeleteScheduleRequest(r events.APIGatewayProxyRequest) *GetScheduleRequest {
	return &GetScheduleRequest{ScheduleID: r.PathParameters["schedule_id"]}
}

func ValidateDeleteScheduleRequest(req *GetScheduleRequest) error {
	if req.ScheduleID == "" {
		return fmt.Errorf("schedule_id is empty")
	}
	return nil
}
