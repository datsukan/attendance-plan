package repository

import (
	"errors"

	"github.com/datsukan/attendance-plan/backend/app/model"
	"github.com/guregu/dynamo"
)

const sessionTableName = "AttendancePlan_Session"

// SessionRepository はセッションの repository を表すインターフェースです。
type SessionRepository interface {
	ReadByUserID(userID string) (*model.Session, error)
	Read(id string) (*model.Session, error)
	Create(session *model.Session) error
	Update(session *model.Session) error
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
func (r *SessionRepositoryImpl) ReadByUserID(userID string) (*model.Session, error) {
	var session *model.Session
	err := r.Table.Get("UserID", userID).Index("UserID-index").One(&session)
	if err != nil {
		if errors.Is(err, dynamo.ErrNotFound) {
			return nil, NewNotFoundError()
		}

		return nil, err
	}
	return session, nil
}

// Read は指定された ID のセッションを取得します。
func (r *SessionRepositoryImpl) Read(id string) (*model.Session, error) {
	var session *model.Session
	err := r.Table.Get("ID", id).One(&session)
	if err != nil {
		if errors.Is(err, dynamo.ErrNotFound) {
			return nil, NewNotFoundError()
		}

		return nil, err
	}
	return session, nil
}

// Create はセッションを作成します。
func (r *SessionRepositoryImpl) Create(session *model.Session) error {
	return r.Table.Put(session).Run()
}

// Update はセッションを更新します。
func (r *SessionRepositoryImpl) Update(session *model.Session) error {
	return r.Table.Put(session).Run()
}

// Delete は指定された ID のセッションを削除します。
func (r *SessionRepositoryImpl) Delete(id string) error {
	return r.Table.Delete("ID", id).Run()
}
