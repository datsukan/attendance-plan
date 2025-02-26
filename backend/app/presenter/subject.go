package presenter

import (
	"encoding/json"
	"net/http"

	"github.com/datsukan/attendance-plan/backend/app/port"
	"github.com/datsukan/attendance-plan/backend/app/response"
)

// SubjectPresenter は科目の presenter を表す構造体です。
type SubjectPresenter struct {
	StatusCode int
	Body       string
}

// NewSubjectPresenter は SubjectOutputPort を生成します。
func NewSubjectPresenter() port.SubjectOutputPort {
	return &SubjectPresenter{}
}

// GetResponse はレスポンスのステータスコードとボディを取得します。
func (p *SubjectPresenter) GetResponse() (int, string) {
	return p.StatusCode, p.Body
}

// SetResponseGetSubjectList は科目リストを取得するレスポンスをセットします。
func (p *SubjectPresenter) SetResponseGetSubjectList(output *port.GetSubjectListOutputData, result port.Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = response.ToErrorBody(result.ErrorMessage)
		return
	}

	res := response.ToGetSubjectListResponse(output)
	b, err := json.Marshal(res)
	if err != nil {
		p.StatusCode = http.StatusInternalServerError
		p.Body = response.ToErrorBody(err.Error())
		return
	}

	p.Body = string(b)
}

// SetResponseCreateSubject は科目を作成するレスポンスをセットします。
func (p *SubjectPresenter) SetResponseCreateSubject(output *port.CreateSubjectOutputData, result port.Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = response.ToErrorBody(result.ErrorMessage)
		return
	}

	res := response.ToPostSubjectResponse(output)
	b, err := json.Marshal(res)
	if err != nil {
		p.StatusCode = http.StatusInternalServerError
		p.Body = response.ToErrorBody(err.Error())
		return
	}

	p.Body = string(b)
}

// SetResponseDeleteSubject は科目を削除するレスポンスをセットします。
func (p *SubjectPresenter) SetResponseDeleteSubject(output *port.DeleteSubjectOutputData, result port.Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = response.ToErrorBody(result.ErrorMessage)
		return
	}

	// 削除成功時はレスポンスボディを空にする
}
