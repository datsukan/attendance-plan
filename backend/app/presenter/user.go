package presenter

import (
	"encoding/json"
	"net/http"

	"github.com/datsukan/attendance-plan/backend/app/port"
	"github.com/datsukan/attendance-plan/backend/app/response"
)

type UserPresenter struct {
	StatusCode int
	Body       string
}

// NewUserPresenter は UserOutputPort を生成します。
func NewUserPresenter() port.UserOutputPort {
	return &UserPresenter{}
}

// GetResponse はレスポンスのステータスコードとボディを取得します。
func (p *UserPresenter) GetResponse() (int, string) {
	return p.StatusCode, p.Body
}

// SetResponseSignIn はサインインのレスポンスをセットします。
func (p *UserPresenter) SetResponseSignIn(output *port.SignInOutputData, result port.Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = response.ToErrorBody(result.Message)
		return
	}

	res := response.ToSignInResponse(output)
	b, err := json.Marshal(res)
	if err != nil {
		p.StatusCode = http.StatusInternalServerError
		p.Body = response.ToErrorBody(err.Error())
		return
	}

	p.Body = string(b)
}

// SetResponseSignUp はサインアップのレスポンスをセットします。
func (p *UserPresenter) SetResponseSignUp(output *port.SignUpOutputData, result port.Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = response.ToErrorBody(result.Message)
		return
	}

	// 成功時はレスポンスボディを空にする
}

// SetResponsePasswordReset はパスワードリセットのレスポンスをセットします。
func (p *UserPresenter) SetResponsePasswordReset(output *port.PasswordResetOutputData, result port.Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = response.ToErrorBody(result.Message)
		return
	}

	// 成功時はレスポンスボディを空にする
}

// SetResponsePasswordSet はパスワード設定のレスポンスをセットします。
func (p *UserPresenter) SetResponsePasswordSet(output *port.PasswordSetOutputData, result port.Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = response.ToErrorBody(result.Message)
		return
	}

	// 成功時はレスポンスボディを空にする
}
