package session

import (
	"errors"

	"github.com/datsukan/attendance-plan/backend/component"
	"github.com/guregu/dynamo"
)

const sessionTableName = "AttendancePlan_Session"

// SessionRepository はセッションの repository を表すインターフェースです。
type SessionRepository interface {
	ReadByUserID(userID string) (*Session, error)
	Read(id string) (*Session, error)
	Create(session *Session) error
	Update(session *Session) error
	Delete(id string) error
}

// SessionRepositoryImpl はセッションの repository の実装を表す構造体です。
type SessionRepositoryImpl struct {
	DB    dynamo.DB
	Table dynamo.Table
}

// NewSessionRepository は SessionRepository を生成します。
func NewSessionRepository(db dynamo.DB) SessionRepository {
	return &SessionRepositoryImpl{DB: db, Table: db.Table(sessionTableName)}
}

// ReadByUserID は指定されたユーザー ID のセッションを取得します。
func (r *SessionRepositoryImpl) ReadByUserID(userID string) (*Session, error) {
	var session *Session
	err := r.Table.Get("UserID", userID).Index("UserID-index").One(&session)
	if err != nil {
		if errors.Is(err, dynamo.ErrNotFound) {
			return nil, component.NewNotFoundError()
		}

		return nil, err
	}
	return session, nil
}

// Read は指定された ID のセッションを取得します。
func (r *SessionRepositoryImpl) Read(id string) (*Session, error) {
	var session *Session
	err := r.Table.Get("ID", id).One(&session)
	if err != nil {
		if errors.Is(err, dynamo.ErrNotFound) {
			return nil, component.NewNotFoundError()
		}

		return nil, err
	}
	return session, nil
}

// Create はセッションを作成します。
func (r *SessionRepositoryImpl) Create(session *Session) error {
	return r.Table.Put(session).Run()
}

// Update はセッションを更新します。
func (r *SessionRepositoryImpl) Update(session *Session) error {
	return r.Table.Put(session).Run()
}

// Delete は指定された ID のセッションを削除します。
func (r *SessionRepositoryImpl) Delete(id string) error {
	return r.Table.Delete("ID", id).Run()
}
