package presenter

import (
	"time"

	"github.com/datsukan/attendance-plan/backend/app/port"
)

// SessionPresenter はセッションの presenter を表す構造体です。
type SessionPresenter struct {
	ID        string
	UserID    string
	ExpiresAt time.Time
	Error     error
}

// NewSessionPresenter は SessionOutputPort を生成します。
func NewSessionPresenter() port.SessionOutputPort {
	return &SessionPresenter{}
}

// GetReturn は処理結果を取得します。
func (p *SessionPresenter) GetReturn() (id string, userID string, expiresAt time.Time, err error) {
	return p.ID, p.UserID, p.ExpiresAt, p.Error
}

// SetReturnStartSession はセッション開始のレスポンスをセットします。
func (p *SessionPresenter) SetReturnStartSession(output *port.StartSessionOutputData) {
	p.ID = output.ID
	p.UserID = output.UserID
	p.ExpiresAt = output.ExpiresAt
	p.Error = output.Error
}

// SetReturnEndSession はセッション終了のレスポンスをセットします。
func (p *SessionPresenter) SetReturnEndSession(output *port.EndSessionOutputData) {
	p.Error = output.Error
}
