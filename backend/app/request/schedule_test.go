package request

import (
	"errors"
	"strings"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/datsukan/attendance-plan/backend/app/model"
	"github.com/stretchr/testify/assert"
)

func TestToGetScheduleListRequest(t *testing.T) {
	r := events.APIGatewayProxyRequest{
		PathParameters: map[string]string{
			"user_id": "test-user-id",
		},
	}
	req := ToGetScheduleListRequest(r)
	assert.Equal(t, "test-user-id", req.UserID)
}

func TestValidateGetScheduleListRequest(t *testing.T) {
	tests := []struct {
		name string
		req  *GetScheduleListRequest
		want error
	}{
		{
			name: "異常系: user_id が未指定の場合はエラー",
			req:  &GetScheduleListRequest{},
			want: errors.New("user_id is empty"),
		},
		{
			name: "正常系: user_id が指定されている場合はエラーなし",
			req:  &GetScheduleListRequest{UserID: "test-user-id"},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateGetScheduleListRequest(tt.req)
			assert.Equal(t, tt.want, err)
		})
	}
}

func TestToGetScheduleRequest(t *testing.T) {
	r := events.APIGatewayProxyRequest{
		PathParameters: map[string]string{
			"schedule_id": "test-schedule-id",
		},
	}
	req := ToGetScheduleRequest(r)
	assert.Equal(t, "test-schedule-id", req.ScheduleID)
}

func TestToPostScheduleRequest(t *testing.T) {
	r := events.APIGatewayProxyRequest{
		Body: `{"name":"test-name","starts_at":"2021-01-01 00:00:00","ends_at":"2021-01-01 00:00:00","color":"test-color","type":"master"}`,
	}
	req, err := ToPostScheduleRequest(r)
	assert.Nil(t, err)
	assert.Equal(t, "test-name", req.Name)
	assert.Equal(t, "2021-01-01 00:00:00", req.StartsAt)
	assert.Equal(t, "2021-01-01 00:00:00", req.EndsAt)
	assert.Equal(t, "test-color", req.Color)
	assert.Equal(t, "master", req.Type)
}

func TestValidateInputScheduleRequest(t *testing.T) {
	type Param struct {
		Name     string
		StartsAt string
		EndsAt   string
		Color    string
		Type     string
	}

	tests := []struct {
		name string
		req  Param
		want error
	}{
		{
			name: "異常系: name が未指定の場合はエラー",
			req:  Param{},
			want: errors.New("name is empty"),
		},
		{
			name: "異常系: name が50文字より多い場合はエラー",
			req:  Param{Name: strings.Repeat("a", 51)},
			want: errors.New("name must be 50 characters or less"),
		},
		{
			name: "正常系: name が50文字の場合はエラーになし",
			req:  Param{Name: strings.Repeat("a", 50), StartsAt: "2021-01-01 00:00:00", EndsAt: "2021-01-01 00:00:00", Color: "test-color", Type: model.ScheduleTypeMaster.String()},
			want: nil,
		},
		{
			name: "異常系: starts_at が未指定の場合はエラー",
			req:  Param{Name: "test-name"},
			want: errors.New("starts_at is empty"),
		},
		{
			name: "異常系: ends_at が未指定の場合はエラー",
			req:  Param{Name: "test-name", StartsAt: "2021-01-01 00:00:00"},
			want: errors.New("ends_at is empty"),
		},
		{
			name: "異常系: starts_at のフォーマットが不正な場合はエラー",
			req:  Param{Name: "test-name", StartsAt: "2021-01-01", EndsAt: "2021-01-01 00:00:00"},
			want: errors.New("starts_at is invalid"),
		},
		{
			name: "異常系: ends_at のフォーマットが不正な場合はエラー",
			req:  Param{Name: "test-name", StartsAt: "2021-01-01 00:00:00", EndsAt: "2021-01-01"},
			want: errors.New("ends_at is invalid"),
		},
		{
			name: "異常系: starts_at が ends_at より後の場合はエラー",
			req:  Param{Name: "test-name", StartsAt: "2021-01-01 00:00:01", EndsAt: "2021-01-01 00:00:00"},
			want: errors.New("starts_at must be earlier than or equal to ends_at"),
		},
		{
			name: "異常系: color が未指定の場合はエラー",
			req:  Param{Name: "test-name", StartsAt: "2021-01-01 00:00:00", EndsAt: "2021-01-01 00:00:00"},
			want: errors.New("color is empty"),
		},
		{
			name: "異常系: type が未指定の場合はエラー",
			req:  Param{Name: "test-name", StartsAt: "2021-01-01 00:00:00", EndsAt: "2021-01-01 00:00:00", Color: "test-color"},
			want: errors.New("type is empty"),
		},
		{
			name: "異常系: type が不正な場合はエラー",
			req:  Param{Name: "test-name", StartsAt: "2021-01-01 00:00:00", EndsAt: "2021-01-01 00:00:00", Color: "test-color", Type: "invalid"},
			want: errors.New("type is invalid"),
		},
		{
			name: "正常系",
			req:  Param{Name: "test-name", StartsAt: "2021-01-01 00:00:00", EndsAt: "2021-01-01 00:00:00", Color: "test-color", Type: model.ScheduleTypeMaster.String()},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateInputScheduleRequest(tt.req.Name, tt.req.StartsAt, tt.req.EndsAt, tt.req.Color, tt.req.Type)
			assert.Equal(t, tt.want, err)
		})
	}
}

func TestValidatePostScheduleRequest(t *testing.T) {
	tests := []struct {
		name string
		req  *PostScheduleRequest
		want error
	}{
		{
			name: "異常系: name が未指定の場合はエラー",
			req:  &PostScheduleRequest{},
			want: errors.New("name is empty"),
		},
		{
			name: "異常系: name が50文字より多い場合はエラー",
			req:  &PostScheduleRequest{Name: strings.Repeat("a", 51)},
			want: errors.New("name must be 50 characters or less"),
		},
		{
			name: "正常系: name が50文字の場合はエラーになし",
			req:  &PostScheduleRequest{Name: strings.Repeat("a", 50), StartsAt: "2021-01-01 00:00:00", EndsAt: "2021-01-01 00:00:00", Color: "test-color", Type: model.ScheduleTypeMaster.String()},
			want: nil,
		},
		{
			name: "異常系: starts_at が未指定の場合はエラー",
			req:  &PostScheduleRequest{Name: "test-name"},
			want: errors.New("starts_at is empty"),
		},
		{
			name: "異常系: ends_at が未指定の場合はエラー",
			req:  &PostScheduleRequest{Name: "test-name", StartsAt: "2021-01-01 00:00:00"},
			want: errors.New("ends_at is empty"),
		},
		{
			name: "異常系: starts_at のフォーマットが不正な場合はエラー",
			req:  &PostScheduleRequest{Name: "test-name", StartsAt: "2021-01-01", EndsAt: "2021-01-01 00:00:00"},
			want: errors.New("starts_at is invalid"),
		},
		{
			name: "異常系: ends_at のフォーマットが不正な場合はエラー",
			req:  &PostScheduleRequest{Name: "test-name", StartsAt: "2021-01-01 00:00:00", EndsAt: "2021-01-01"},
			want: errors.New("ends_at is invalid"),
		},
		{
			name: "異常系: starts_at が ends_at より後の場合はエラー",
			req:  &PostScheduleRequest{Name: "test-name", StartsAt: "2021-01-01 00:00:01", EndsAt: "2021-01-01 00:00:00"},
			want: errors.New("starts_at must be earlier than or equal to ends_at"),
		},
		{
			name: "異常系: color が未指定の場合はエラー",
			req:  &PostScheduleRequest{Name: "test-name", StartsAt: "2021-01-01 00:00:00", EndsAt: "2021-01-01 00:00:00"},
			want: errors.New("color is empty"),
		},
		{
			name: "異常系: type が未指定の場合はエラー",
			req:  &PostScheduleRequest{Name: "test-name", StartsAt: "2021-01-01 00:00:00", EndsAt: "2021-01-01 00:00:00", Color: "test-color"},
			want: errors.New("type is empty"),
		},
		{
			name: "異常系: type が不正な場合はエラー",
			req:  &PostScheduleRequest{Name: "test-name", StartsAt: "2021-01-01 00:00:00", EndsAt: "2021-01-01 00:00:00", Color: "test-color", Type: "invalid"},
			want: errors.New("type is invalid"),
		},
		{
			name: "正常系",
			req:  &PostScheduleRequest{Name: "test-name", StartsAt: "2021-01-01 00:00:00", EndsAt: "2021-01-01 00:00:00", Color: "test-color", Type: model.ScheduleTypeMaster.String()},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePostScheduleRequest(tt.req)
			assert.Equal(t, tt.want, err)
		})
	}
}

func TestToPutScheduleRequest(t *testing.T) {
	r := events.APIGatewayProxyRequest{
		PathParameters: map[string]string{
			"schedule_id": "test-schedule-id",
		},
		Body: `{"name":"test-name","starts_at":"2021-01-01 00:00:00","ends_at":"2021-01-01 00:00:00","color":"test-color","type":"master"}`,
	}
	req, err := ToPutScheduleRequest(r)
	assert.Nil(t, err)
	assert.Equal(t, "test-schedule-id", req.ScheduleID)
	assert.Equal(t, "test-name", req.Name)
	assert.Equal(t, "2021-01-01 00:00:00", req.StartsAt)
	assert.Equal(t, "2021-01-01 00:00:00", req.EndsAt)
	assert.Equal(t, "test-color", req.Color)
	assert.Equal(t, "master", req.Type)
}

func TestValidatePutScheduleRequest(t *testing.T) {
	tests := []struct {
		name string
		req  *PutScheduleRequest
		want error
	}{
		{
			name: "異常系: schedule_id が未指定の場合はエラー",
			req:  &PutScheduleRequest{},
			want: errors.New("schedule_id is empty"),
		},
		{
			name: "異常系: name が未指定の場合はエラー",
			req:  &PutScheduleRequest{ScheduleID: "test-schedule-id"},
			want: errors.New("name is empty"),
		},
		{
			name: "異常系: name が50文字より多い場合はエラー",
			req:  &PutScheduleRequest{ScheduleID: "test-schedule-id", Name: strings.Repeat("a", 51)},
			want: errors.New("name must be 50 characters or less"),
		},
		{
			name: "正常系: name が50文字の場合はエラーになし",
			req:  &PutScheduleRequest{ScheduleID: "test-schedule-id", Name: strings.Repeat("a", 50), StartsAt: "2021-01-01 00:00:00", EndsAt: "2021-01-01 00:00:00", Color: "test-color", Type: model.ScheduleTypeMaster.String()},
			want: nil,
		},
		{
			name: "異常系: starts_at が未指定の場合はエラー",
			req:  &PutScheduleRequest{ScheduleID: "test-schedule-id", Name: "test-name"},
			want: errors.New("starts_at is empty"),
		},
		{
			name: "異常系: ends_at が未指定の場合はエラー",
			req:  &PutScheduleRequest{ScheduleID: "test-schedule-id", Name: "test-name", StartsAt: "2021-01-01 00:00:00"},
			want: errors.New("ends_at is empty"),
		},
		{
			name: "異常系: starts_at のフォーマットが不正な場合はエラー",
			req:  &PutScheduleRequest{ScheduleID: "test-schedule-id", Name: "test-name", StartsAt: "2021-01-01", EndsAt: "2021-01-01 00:00:00"},
			want: errors.New("starts_at is invalid"),
		},
		{
			name: "異常系: ends_at のフォーマットが不正な場合はエラー",
			req:  &PutScheduleRequest{ScheduleID: "test-schedule-id", Name: "test-name", StartsAt: "2021-01-01 00:00:00", EndsAt: "2021-01-01"},
			want: errors.New("ends_at is invalid"),
		},
		{
			name: "異常系: starts_at が ends_at より後の場合はエラー",
			req:  &PutScheduleRequest{ScheduleID: "test-schedule-id", Name: "test-name", StartsAt: "2021-01-01 00:00:01", EndsAt: "2021-01-01 00:00:00"},
			want: errors.New("starts_at must be earlier than or equal to ends_at"),
		},
		{
			name: "異常系: color が未指定の場合はエラー",
			req:  &PutScheduleRequest{ScheduleID: "test-schedule-id", Name: "test-name", StartsAt: "2021-01-01 00:00:00", EndsAt: "2021-01-01 00:00:00"},
			want: errors.New("color is empty"),
		},
		{
			name: "異常系: type が未指定の場合はエラー",
			req:  &PutScheduleRequest{ScheduleID: "test-schedule-id", Name: "test-name", StartsAt: "2021-01-01 00:00:00", EndsAt: "2021-01-01 00:00:00", Color: "test-color"},
			want: errors.New("type is empty"),
		},
		{
			name: "異常系: type が不正な場合はエラー",
			req:  &PutScheduleRequest{ScheduleID: "test-schedule-id", Name: "test-name", StartsAt: "2021-01-01 00:00:00", EndsAt: "2021-01-01 00:00:00", Color: "test-color", Type: "invalid"},
			want: errors.New("type is invalid"),
		},
		{
			name: "正常系",
			req:  &PutScheduleRequest{ScheduleID: "test-schedule-id", Name: "test-name", StartsAt: "2021-01-01 00:00:00", EndsAt: "2021-01-01 00:00:00", Color: "test-color", Type: model.ScheduleTypeMaster.String()},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePutScheduleRequest(tt.req)
			assert.Equal(t, tt.want, err)
		})
	}
}

func TestToPutBulkScheduleRequest(t *testing.T) {
	r := events.APIGatewayProxyRequest{
		Body: `{"schedules":[{"id":"test-schedule-id","name":"test-name","starts_at":"2021-01-01 00:00:00","ends_at":"2021-01-01 00:00:00","color":"test-color","type":"master","order":0}]}`,
	}
	req, err := ToPutBulkScheduleRequest(r)
	assert.Nil(t, err)
	assert.Len(t, req.Schedules, 1)
	assert.Equal(t, "test-schedule-id", req.Schedules[0].ScheduleID)
	assert.Equal(t, "test-name", req.Schedules[0].Name)
	assert.Equal(t, "2021-01-01 00:00:00", req.Schedules[0].StartsAt)
	assert.Equal(t, "2021-01-01 00:00:00", req.Schedules[0].EndsAt)
	assert.Equal(t, "test-color", req.Schedules[0].Color)
	assert.Equal(t, "master", req.Schedules[0].Type)
	assert.Equal(t, 0, req.Schedules[0].Order)
}

func TestValidatePutBulkScheduleRequest(t *testing.T) {
	tests := []struct {
		name string
		req  *PutBulkScheduleRequest
		want error
	}{
		{
			name: "異常系: schedules が未指定の場合はエラー",
			req:  &PutBulkScheduleRequest{},
			want: errors.New("schedules is empty"),
		},
		{
			name: "異常系: schedules が空の場合はエラー",
			req:  &PutBulkScheduleRequest{Schedules: []PutScheduleRequest{}},
			want: errors.New("schedules is empty"),
		},
		{
			name: "正常系",
			req: &PutBulkScheduleRequest{
				Schedules: []PutScheduleRequest{
					{ScheduleID: "test-schedule-id", Name: "test-name", StartsAt: "2021-01-01 00:00:00", EndsAt: "2021-01-01 00:00:00", Color: "test-color", Type: model.ScheduleTypeMaster.String()},
				},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePutBulkScheduleRequest(tt.req)
			assert.Equal(t, tt.want, err)
		})
	}
}

func TestToDeleteScheduleRequest(t *testing.T) {
	r := events.APIGatewayProxyRequest{
		PathParameters: map[string]string{
			"schedule_id": "test-schedule-id",
		},
	}
	req := ToDeleteScheduleRequest(r)
	assert.Equal(t, "test-schedule-id", req.ScheduleID)
}

func TestValidateDeleteScheduleRequest(t *testing.T) {
	tests := []struct {
		name string
		req  *DeleteScheduleRequest
		want error
	}{
		{
			name: "異常系: schedule_id が未指定の場合はエラー",
			req:  &DeleteScheduleRequest{},
			want: errors.New("schedule_id is empty"),
		},
		{
			name: "正常系",
			req:  &DeleteScheduleRequest{ScheduleID: "test-schedule-id"},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateDeleteScheduleRequest(tt.req)
			assert.Equal(t, tt.want, err)
		})
	}
}
