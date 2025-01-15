package repository

import (
	"errors"

	"github.com/datsukan/attendance-plan/backend/app/model"
	"github.com/guregu/dynamo"
)

const userTableName = "AttendancePlan_User"

// UserRepository はユーザーの repository を表すインターフェースです。
type UserRepository interface {
	ReadByEmail(email string) (*model.User, error)
	Read(id string) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id string) error
	Exists(id string) (bool, error)
	ExistsByEmail(email string) (bool, error)
}

// UserRepositoryImpl はユーザーの repository の実装を表す構造体です。
type UserRepositoryImpl struct {
	DB    dynamo.DB
	Table dynamo.Table
}

// NewUserRepository は UserRepository を生成します。
func NewUserRepository(db dynamo.DB) UserRepository {
	return &UserRepositoryImpl{DB: db, Table: db.Table(userTableName)}
}

// ReadByEmail は指定されたメールアドレスのユーザーを取得します。
func (r *UserRepositoryImpl) ReadByEmail(email string) (*model.User, error) {
	var user *model.User
	err := r.Table.Get("Email", email).Index("Email-index").One(&user)
	if err != nil {
		if errors.Is(err, dynamo.ErrNotFound) {
			return nil, NewNotFoundError()
		}

		return nil, err
	}
	return user, nil
}

// Read は指定された ID のユーザーを取得します。
func (r *UserRepositoryImpl) Read(id string) (*model.User, error) {
	var user *model.User
	err := r.Table.Get("ID", id).One(&user)
	if err != nil {
		if errors.Is(err, dynamo.ErrNotFound) {
			return nil, NewNotFoundError()
		}

		return nil, err
	}
	return user, nil
}

// Create はユーザーを保存します。
func (r *UserRepositoryImpl) Create(user *model.User) error {
	return r.Table.Put(user).Run()
}

// Update はユーザーを更新します。
func (r *UserRepositoryImpl) Update(user *model.User) error {
	return r.Table.Put(user).Run()
}

// Delete はユーザーを削除します。
func (r *UserRepositoryImpl) Delete(id string) error {
	return r.Table.Delete("ID", id).Run()
}

// Exists は指定された ID のユーザーが存在するかどうかを返します。
func (r *UserRepositoryImpl) Exists(id string) (bool, error) {
	var user *model.User
	err := r.Table.Get("ID", id).One(&user)
	if err != nil {
		if errors.Is(err, dynamo.ErrNotFound) {
			return false, nil
		}

		return false, err
	}

	return user != nil, nil
}

// ExistsByEmail は指定されたメールアドレスのユーザーが存在するかどうかを返します。
func (r *UserRepositoryImpl) ExistsByEmail(email string) (bool, error) {
	var user *model.User
	err := r.Table.Get("Email", email).Index("Email-index").One(&user)
	if err != nil {
		if errors.Is(err, dynamo.ErrNotFound) {
			return false, nil
		}

		return false, err
	}

	return user != nil, nil
}
