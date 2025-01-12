package session

import (
	"errors"
	"time"

	"github.com/datsukan/attendance-plan/backend/app/user"
	"github.com/datsukan/attendance-plan/backend/component"
)

// StartSessionInputDate はセッション開始の入力データを表す構造体です。
type StartSessionInputDate struct {
	UserID string
}

// StartSessionOutputData はセッション開始の出力データを表す構造体です。
type StartSessionOutputData struct {
	ID        string
	UserID    string
	ExpiresAt time.Time
	Error     error
}

// EndSessionInputDate はセッション終了の入力データを表す構造体です。
type EndSessionInputDate struct {
	UserID string
}

// EndSessionOutputData はセッション終了の出力データを表す構造体です。
type EndSessionOutputData struct {
	Error error
}

// SessionInputPort はセッションのユースケースを表すインターフェースです。
type SessionInputPort interface {
	StartSession(input StartSessionInputDate)
	EndSession(input EndSessionInputDate)
}

// SessionOutputPort はセッションのユースケースの外部出力を表すインターフェースです。
type SessionOutputPort interface {
	SetReturnStartSession(output *StartSessionOutputData)
	SetReturnEndSession(output *EndSessionOutputData)
}

// SessionInteractor はセッションのユースケースの実装を表す構造体です。
type SessionInteractor struct {
	SessionRepository SessionRepository
	UserRepository    user.UserRepository
	OutputPort        SessionOutputPort
}

// NewSessionInteractor は SessionInteractor を生成します。
func NewSessionInteractor(sessionRepository SessionRepository, userRepository user.UserRepository, outputPort SessionOutputPort) SessionInputPort {
	return &SessionInteractor{
		SessionRepository: sessionRepository,
		UserRepository:    userRepository,
		OutputPort:        outputPort,
	}
}

// StartSession はセッションを開始します。
func (i *SessionInteractor) StartSession(input StartSessionInputDate) {
	user, err := i.UserRepository.Read(input.UserID)
	if err != nil {
		i.OutputPort.SetReturnStartSession(&StartSessionOutputData{
			Error: err,
		})
		return
	}

	if user == nil {
		i.OutputPort.SetReturnStartSession(&StartSessionOutputData{
			Error: errors.New("user not found"),
		})
		return
	}

	currentSession, err := i.SessionRepository.ReadByUserID(user.ID)
	if err != nil {
		i.OutputPort.SetReturnStartSession(&StartSessionOutputData{
			Error: err,
		})
		return
	}

	// セッションが存在する場合は期限を延長する
	if currentSession != nil {
		currentSession.ExtendExpiresAt()
		err = i.SessionRepository.Update(currentSession)
		if err != nil {
			i.OutputPort.SetReturnStartSession(&StartSessionOutputData{
				Error: err,
			})
			return
		}

		i.OutputPort.SetReturnStartSession(&StartSessionOutputData{
			ID:        currentSession.ID,
			UserID:    currentSession.UserID,
			ExpiresAt: currentSession.ExpiresAt,
		})
		return
	}

	// セッションが存在しない場合は新規作成する
	session := &Session{
		ID:     component.NewID(),
		UserID: user.ID,
	}
	session.ExtendExpiresAt()

	err = i.SessionRepository.Create(session)
	if err != nil {
		i.OutputPort.SetReturnStartSession(&StartSessionOutputData{
			Error: err,
		})
		return
	}

	i.OutputPort.SetReturnStartSession(&StartSessionOutputData{
		ID:        session.ID,
		UserID:    session.UserID,
		ExpiresAt: session.ExpiresAt,
	})
}

// EndSession はセッションを終了します。
func (i *SessionInteractor) EndSession(input EndSessionInputDate) {
	user, err := i.UserRepository.Read(input.UserID)
	if err != nil {
		i.OutputPort.SetReturnStartSession(&StartSessionOutputData{
			Error: err,
		})
		return
	}

	if user == nil {
		i.OutputPort.SetReturnStartSession(&StartSessionOutputData{
			Error: errors.New("user not found"),
		})
		return
	}

	session, err := i.SessionRepository.ReadByUserID(user.ID)
	if err != nil {
		i.OutputPort.SetReturnStartSession(&StartSessionOutputData{
			Error: err,
		})
		return
	}

	if err := i.SessionRepository.Delete(session.ID); err != nil {
		i.OutputPort.SetReturnEndSession(&EndSessionOutputData{
			Error: err,
		})
		return
	}

	i.OutputPort.SetReturnEndSession(&EndSessionOutputData{})
}
