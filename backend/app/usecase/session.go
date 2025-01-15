package usecase

import (
	"errors"

	"github.com/datsukan/attendance-plan/backend/app/component"
	"github.com/datsukan/attendance-plan/backend/app/model"
	"github.com/datsukan/attendance-plan/backend/app/port"
	"github.com/datsukan/attendance-plan/backend/app/repository"
)

// SessionInteractor はセッションのユースケースの実装を表す構造体です。
type SessionInteractor struct {
	SessionRepository repository.SessionRepository
	UserRepository    repository.UserRepository
	OutputPort        port.SessionOutputPort
}

// NewSessionInteractor は SessionInteractor を生成します。
func NewSessionInteractor(sessionRepository repository.SessionRepository, userRepository repository.UserRepository, outputPort port.SessionOutputPort) port.SessionInputPort {
	return &SessionInteractor{
		SessionRepository: sessionRepository,
		UserRepository:    userRepository,
		OutputPort:        outputPort,
	}
}

// StartSession はセッションを開始します。
func (i *SessionInteractor) StartSession(input port.StartSessionInputDate) {
	user, err := i.UserRepository.Read(input.UserID)
	if err != nil {
		i.OutputPort.SetReturnStartSession(&port.StartSessionOutputData{
			Error: err,
		})
		return
	}

	if user == nil {
		i.OutputPort.SetReturnStartSession(&port.StartSessionOutputData{
			Error: errors.New("user not found"),
		})
		return
	}

	currentSession, err := i.SessionRepository.ReadByUserID(user.ID)
	if err != nil {
		i.OutputPort.SetReturnStartSession(&port.StartSessionOutputData{
			Error: err,
		})
		return
	}

	// セッションが存在する場合は期限を延長する
	if currentSession != nil {
		currentSession.ExtendExpiresAt()
		err = i.SessionRepository.Update(currentSession)
		if err != nil {
			i.OutputPort.SetReturnStartSession(&port.StartSessionOutputData{
				Error: err,
			})
			return
		}

		i.OutputPort.SetReturnStartSession(&port.StartSessionOutputData{
			ID:        currentSession.ID,
			UserID:    currentSession.UserID,
			ExpiresAt: currentSession.ExpiresAt,
		})
		return
	}

	// セッションが存在しない場合は新規作成する
	session := &model.Session{
		ID:     component.NewID(),
		UserID: user.ID,
	}
	session.ExtendExpiresAt()

	err = i.SessionRepository.Create(session)
	if err != nil {
		i.OutputPort.SetReturnStartSession(&port.StartSessionOutputData{
			Error: err,
		})
		return
	}

	i.OutputPort.SetReturnStartSession(&port.StartSessionOutputData{
		ID:        session.ID,
		UserID:    session.UserID,
		ExpiresAt: session.ExpiresAt,
	})
}

// EndSession はセッションを終了します。
func (i *SessionInteractor) EndSession(input port.EndSessionInputDate) {
	user, err := i.UserRepository.Read(input.UserID)
	if err != nil {
		i.OutputPort.SetReturnStartSession(&port.StartSessionOutputData{
			Error: err,
		})
		return
	}

	if user == nil {
		i.OutputPort.SetReturnStartSession(&port.StartSessionOutputData{
			Error: errors.New("user not found"),
		})
		return
	}

	session, err := i.SessionRepository.ReadByUserID(user.ID)
	if err != nil {
		i.OutputPort.SetReturnStartSession(&port.StartSessionOutputData{
			Error: err,
		})
		return
	}

	if err := i.SessionRepository.Delete(session.ID); err != nil {
		i.OutputPort.SetReturnEndSession(&port.EndSessionOutputData{
			Error: err,
		})
		return
	}

	i.OutputPort.SetReturnEndSession(&port.EndSessionOutputData{})
}
