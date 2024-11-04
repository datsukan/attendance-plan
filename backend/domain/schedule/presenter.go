package schedule

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	StatusCode int
	HasError   bool
	Message    string
}

type SchedulePresenter struct {
	StatusCode int
	Body       string
}

func NewSchedulePresenter() ScheduleOutputPort {
	return &SchedulePresenter{}
}

func (p *SchedulePresenter) GetResponse() (int, string) {
	return p.StatusCode, p.Body
}

func (p *SchedulePresenter) ResponseGetScheduleList(output *GetScheduleListOutputData, result Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = result.Message
		return
	}

	b, err := json.Marshal(output)
	if err != nil {
		p.StatusCode = http.StatusInternalServerError
		p.Body = err.Error()
		return
	}

	p.Body = string(b)
}

func (p *SchedulePresenter) ResponseGetSchedule(output *GetScheduleOutputData, result Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = result.Message
		return
	}

	b, err := json.Marshal(output)
	if err != nil {
		p.StatusCode = http.StatusInternalServerError
		p.Body = err.Error()
		return
	}

	p.Body = string(b)
}

func (p *SchedulePresenter) ResponseCreateSchedule(output *CreateScheduleOutputData, result Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = result.Message
		return
	}

	b, err := json.Marshal(output)
	if err != nil {
		p.StatusCode = http.StatusInternalServerError
		p.Body = err.Error()
		return
	}

	p.Body = string(b)
}

func (p *SchedulePresenter) ResponseUpdateSchedule(output *UpdateScheduleOutputData, result Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = result.Message
		return
	}

	b, err := json.Marshal(output)
	if err != nil {
		p.StatusCode = http.StatusInternalServerError
		p.Body = err.Error()
		return
	}

	p.Body = string(b)
}

func (p *SchedulePresenter) ResponseDeleteSchedule(output *DeleteScheduleOutputData, result Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = result.Message
		return
	}

	b, err := json.Marshal(output)
	if err != nil {
		p.StatusCode = http.StatusInternalServerError
		p.Body = err.Error()
		return
	}

	p.Body = string(b)
}
