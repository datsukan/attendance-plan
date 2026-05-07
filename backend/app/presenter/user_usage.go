package presenter

import (
	"encoding/json"
	"net/http"

	"github.com/datsukan/attendance-plan/backend/app/port"
	"github.com/datsukan/attendance-plan/backend/app/response"
)

// UserUsagePresenter はユーザー利用状況の presenter を表す構造体です。
type UserUsagePresenter struct {
	StatusCode int
	Body       string
}

// NewUserUsagePresenter は UserUsageOutputPort を生成します。
func NewUserUsagePresenter() port.UserUsageOutputPort {
	return &UserUsagePresenter{}
}

// GetResponse はレスポンスのステータスコードとボディを取得します。
func (p *UserUsagePresenter) GetResponse() (int, string) {
	return p.StatusCode, p.Body
}

// SetResponseGetUserUsageList はユーザー利用状況リストを取得するレスポンスをセットします。
func (p *UserUsagePresenter) SetResponseGetUserUsageList(output *port.GetUserUsageListOutputData, result port.Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = response.ToErrorBody(result.ErrorMessage)
		return
	}

	res := response.ToGetUserUsageListResponse(output)
	b, err := json.Marshal(res)
	if err != nil {
		p.StatusCode = http.StatusInternalServerError
		p.Body = response.ToErrorBody(err.Error())
		return
	}

	p.Body = string(b)
}
