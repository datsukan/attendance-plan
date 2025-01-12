package user

import (
	"encoding/json"
	"net/http"

	"github.com/datsukan/attendance-plan/backend/component"
)

type UserPresenter struct {
	StatusCode int
	Body       string
}

// NewUserPresenter は UserOutputPort を生成します。
func NewUserPresenter() UserOutputPort {
	return &UserPresenter{}
}

// GetResponse はレスポンスのステータスコードとボディを取得します。
func (p *UserPresenter) GetResponse() (int, string) {
	return p.StatusCode, p.Body
}

// SetResponseSignIn はサインインのレスポンスをセットします。
func (p *UserPresenter) SetResponseSignIn(output *SignInOutputData, result component.ResponseResult) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = component.ToErrorBody(result.Message)
		return
	}

	res := ToSignInResponse(output)
	b, err := json.Marshal(res)
	if err != nil {
		p.StatusCode = http.StatusInternalServerError
		p.Body = component.ToErrorBody(err.Error())
		return
	}

	p.Body = string(b)
}

// SetResponseSignUp はサインアップのレスポンスをセットします。
func (p *UserPresenter) SetResponseSignUp(output *SignUpOutputData, result component.ResponseResult) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = component.ToErrorBody(result.Message)
		return
	}

	res := ToSignUpResponse(output)
	b, err := json.Marshal(res)
	if err != nil {
		p.StatusCode = http.StatusInternalServerError
		p.Body = component.ToErrorBody(err.Error())
		return
	}

	p.Body = string(b)
}
