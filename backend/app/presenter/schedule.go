package presenter

import (
	"encoding/json"
	"net/http"

	"github.com/datsukan/attendance-plan/backend/app/port"
	"github.com/datsukan/attendance-plan/backend/app/response"
)

// SchedulePresenter はスケジュールの presenter を表す構造体です。
type SchedulePresenter struct {
	StatusCode int
	Body       string
}

// NewSchedulePresenter は ScheduleOutputPort を生成します。
func NewSchedulePresenter() port.ScheduleOutputPort {
	return &SchedulePresenter{}
}

// GetResponse はレスポンスのステータスコードとボディを取得します。
func (p *SchedulePresenter) GetResponse() (int, string) {
	return p.StatusCode, p.Body
}

// SetResponseGetScheduleList はスケジュールリストを取得するレスポンスをセットします。
func (p *SchedulePresenter) SetResponseGetScheduleList(output *port.GetScheduleListOutputData, result port.Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = response.ToErrorBody(result.Message)
		return
	}

	res := response.ToGetScheduleListResponse(output)
	b, err := json.Marshal(res)
	if err != nil {
		p.StatusCode = http.StatusInternalServerError
		p.Body = response.ToErrorBody(err.Error())
		return
	}

	p.Body = string(b)
}

// SetResponseGetSchedule はスケジュールを取得するレスポンスをセットします。
func (p *SchedulePresenter) SetResponseGetSchedule(output *port.GetScheduleOutputData, result port.Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = response.ToErrorBody(result.Message)
		return
	}

	res := response.ToGetScheduleResponse(output)
	b, err := json.Marshal(res)
	if err != nil {
		p.StatusCode = http.StatusInternalServerError
		p.Body = response.ToErrorBody(err.Error())
		return
	}

	p.Body = string(b)
}

// SetResponseCreateSchedule はスケジュールを作成するレスポンスをセットします。
func (p *SchedulePresenter) SetResponseCreateSchedule(output *port.CreateScheduleOutputData, result port.Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = response.ToErrorBody(result.Message)
		return
	}

	res := response.ToPostScheduleResponse(output)
	b, err := json.Marshal(res)
	if err != nil {
		p.StatusCode = http.StatusInternalServerError
		p.Body = response.ToErrorBody(err.Error())
		return
	}

	p.Body = string(b)
}

// SetResponseUpdateSchedule はスケジュールを更新するレスポンスをセットします。
func (p *SchedulePresenter) SetResponseUpdateSchedule(output *port.UpdateScheduleOutputData, result port.Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = response.ToErrorBody(result.Message)
		return
	}

	res := response.ToPutScheduleResponse(output)
	b, err := json.Marshal(res)
	if err != nil {
		p.StatusCode = http.StatusInternalServerError
		p.Body = response.ToErrorBody(err.Error())
		return
	}

	p.Body = string(b)
}

// SetResponseUpdateBulkSchedule はスケジュールを一括更新するレスポンスをセットします。
func (p *SchedulePresenter) SetResponseUpdateBulkSchedule(output *port.UpdateBulkScheduleOutputData, result port.Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = response.ToErrorBody(result.Message)
		return
	}

	res := response.ToPutBulkScheduleResponse(output)
	b, err := json.Marshal(res)
	if err != nil {
		p.StatusCode = http.StatusInternalServerError
		p.Body = response.ToErrorBody(err.Error())
		return
	}

	p.Body = string(b)
}

// SetResponseDeleteSchedule はスケジュールを削除するレスポンスをセットします。
func (p *SchedulePresenter) SetResponseDeleteSchedule(output *port.DeleteScheduleOutputData, result port.Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = response.ToErrorBody(result.Message)
		return
	}

	// 削除成功時はレスポンスボディを空にする
}
