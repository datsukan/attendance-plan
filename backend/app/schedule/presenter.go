package schedule

import (
	"encoding/json"
	"net/http"

	"github.com/datsukan/attendance-plan/backend/component"
)

// SchedulePresenter はスケジュールの presenter を表す構造体です。
type SchedulePresenter struct {
	StatusCode int
	Body       string
}

// NewSchedulePresenter は ScheduleOutputPort を生成します。
func NewSchedulePresenter() ScheduleOutputPort {
	return &SchedulePresenter{}
}

// GetResponse はレスポンスのステータスコードとボディを取得します。
func (p *SchedulePresenter) GetResponse() (int, string) {
	return p.StatusCode, p.Body
}

// SetResponseGetScheduleList はスケジュールリストを取得するレスポンスをセットします。
func (p *SchedulePresenter) SetResponseGetScheduleList(output *GetScheduleListOutputData, result component.ResponseResult) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = component.ToErrorBody(result.Message)
		return
	}

	res := ToGetScheduleListResponse(output)
	b, err := json.Marshal(res)
	if err != nil {
		p.StatusCode = http.StatusInternalServerError
		p.Body = component.ToErrorBody(err.Error())
		return
	}

	p.Body = string(b)
}

// SetResponseGetSchedule はスケジュールを取得するレスポンスをセットします。
func (p *SchedulePresenter) SetResponseGetSchedule(output *GetScheduleOutputData, result component.ResponseResult) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = component.ToErrorBody(result.Message)
		return
	}

	res := ToGetScheduleResponse(output)
	b, err := json.Marshal(res)
	if err != nil {
		p.StatusCode = http.StatusInternalServerError
		p.Body = component.ToErrorBody(err.Error())
		return
	}

	p.Body = string(b)
}

// SetResponseCreateSchedule はスケジュールを作成するレスポンスをセットします。
func (p *SchedulePresenter) SetResponseCreateSchedule(output *CreateScheduleOutputData, result component.ResponseResult) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = component.ToErrorBody(result.Message)
		return
	}

	res := ToPostScheduleResponse(output)
	b, err := json.Marshal(res)
	if err != nil {
		p.StatusCode = http.StatusInternalServerError
		p.Body = component.ToErrorBody(err.Error())
		return
	}

	p.Body = string(b)
}

// SetResponseUpdateSchedule はスケジュールを更新するレスポンスをセットします。
func (p *SchedulePresenter) SetResponseUpdateSchedule(output *UpdateScheduleOutputData, result component.ResponseResult) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = component.ToErrorBody(result.Message)
		return
	}

	res := ToPutScheduleResponse(output)
	b, err := json.Marshal(res)
	if err != nil {
		p.StatusCode = http.StatusInternalServerError
		p.Body = component.ToErrorBody(err.Error())
		return
	}

	p.Body = string(b)
}

// SetResponseDeleteSchedule はスケジュールを削除するレスポンスをセットします。
func (p *SchedulePresenter) SetResponseDeleteSchedule(output *DeleteScheduleOutputData, result component.ResponseResult) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = component.ToErrorBody(result.Message)
		return
	}

	// 削除成功時はレスポンスボディを空にする
}
