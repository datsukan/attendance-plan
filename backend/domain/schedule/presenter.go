package schedule

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Result は処理結果を表す構造体です。
type Result struct {
	StatusCode int
	HasError   bool
	Message    string
}

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

// ResponseGetScheduleList はスケジュールリストを取得するレスポンスを生成します。
func (p *SchedulePresenter) ResponseGetScheduleList(output *GetScheduleListOutputData, result Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = toErrorBody(result.Message)
		return
	}

	b, err := json.Marshal(output)
	if err != nil {
		p.StatusCode = http.StatusInternalServerError
		p.Body = toErrorBody(err.Error())
		return
	}

	p.Body = string(b)
}

// ResponseGetSchedule はスケジュールを取得するレスポンスを生成します。
func (p *SchedulePresenter) ResponseGetSchedule(output *GetScheduleOutputData, result Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = toErrorBody(result.Message)
		return
	}

	b, err := json.Marshal(output)
	if err != nil {
		p.StatusCode = http.StatusInternalServerError
		p.Body = toErrorBody(err.Error())
		return
	}

	p.Body = string(b)
}

// ResponseCreateSchedule はスケジュールを作成するレスポンスを生成します。
func (p *SchedulePresenter) ResponseCreateSchedule(output *CreateScheduleOutputData, result Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = toErrorBody(result.Message)
		return
	}

	b, err := json.Marshal(output)
	if err != nil {
		p.StatusCode = http.StatusInternalServerError
		p.Body = toErrorBody(err.Error())
		return
	}

	p.Body = string(b)
}

// ResponseUpdateSchedule はスケジュールを更新するレスポンスを生成します。
func (p *SchedulePresenter) ResponseUpdateSchedule(output *UpdateScheduleOutputData, result Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = toErrorBody(result.Message)
		return
	}

	b, err := json.Marshal(output)
	if err != nil {
		p.StatusCode = http.StatusInternalServerError
		p.Body = toErrorBody(err.Error())
		return
	}

	p.Body = string(b)
}

// ResponseDeleteSchedule はスケジュールを削除するレスポンスを生成します。
func (p *SchedulePresenter) ResponseDeleteSchedule(output *DeleteScheduleOutputData, result Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = toErrorBody(result.Message)
		return
	}

	b, err := json.Marshal(output)
	if err != nil {
		p.StatusCode = http.StatusInternalServerError
		p.Body = toErrorBody(err.Error())
		return
	}

	p.Body = string(b)
}

func toErrorBody(message string) string {
	return fmt.Sprintf(`{"message": "%s"}`, message)
}
