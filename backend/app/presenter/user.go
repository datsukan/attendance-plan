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
		p.Body = response.ToErrorBody(result.ErrorMessage)
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
		p.Body = response.ToErrorBody(result.ErrorMessage)
		return
	}

	// 成功時はレスポンスボディを空にする
}

// SetResponsePasswordReset はパスワードリセットのレスポンスをセットします。
func (p *UserPresenter) SetResponsePasswordReset(output *port.PasswordResetOutputData, result port.Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = response.ToErrorBody(result.ErrorMessage)
		return
	}

	// 成功時はレスポンスボディを空にする
}

// SetResponsePasswordSet はパスワード設定のレスポンスをセットします。
func (p *UserPresenter) SetResponsePasswordSet(output *port.PasswordSetOutputData, result port.Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = response.ToErrorBody(result.ErrorMessage)
		return
	}

	// 成功時はレスポンスボディを空にする
}

// SetResponseGetUser はユーザー取得のレスポンスをセットします。
func (p *UserPresenter) SetResponseGetUser(output *port.GetUserOutputData, result port.Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = response.ToErrorBody(result.ErrorMessage)
		return
	}

	res := response.ToGetUserResponse(output)
	b, err := json.Marshal(res)
	if err != nil {
		p.StatusCode = http.StatusInternalServerError
		p.Body = response.ToErrorBody(err.Error())
		return
	}

	p.Body = string(b)
}

// SetResponseUpdateUser はユーザー情報更新のレスポンスをセットします。
func (p *UserPresenter) SetResponseUpdateUser(output *port.UpdateUserOutputData, result port.Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = response.ToErrorBody(result.ErrorMessage)
		return
	}

	res := response.ToPutUserResponse(output)
	b, err := json.Marshal(res)
	if err != nil {
		p.StatusCode = http.StatusInternalServerError
		p.Body = response.ToErrorBody(err.Error())
		return
	}

	p.Body = string(b)
}

// SetResponseDeleteUser はユーザー削除のレスポンスをセットします。
func (p *UserPresenter) SetResponseDeleteUser(output *port.DeleteUserOutputData, result port.Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = response.ToErrorBody(result.ErrorMessage)
		return
	}

	// 成功時はレスポンスボディを空にする
}

// SetResponseResetEmail はメールアドレスリセットのレスポンスをセットします。
func (p *UserPresenter) SetResponseResetEmail(output *port.ResetEmailOutputData, result port.Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = response.ToErrorBody(result.ErrorMessage)
		return
	}

	// 成功時はレスポンスボディを空にする
}

// SetResponseSetEmail はメールアドレス設定のレスポンスをセットします。
func (p *UserPresenter) SetResponseSetEmail(output *port.SetEmailOutputData, result port.Result) {
	p.StatusCode = result.StatusCode

	if result.HasError {
		p.Body = response.ToErrorBody(result.ErrorMessage)
		return
	}

	// 成功時はレスポンスボディを空にする
}
