package repository

import (
	"github.com/datsukan/attendance-plan/backend/app/model"
	"github.com/guregu/dynamo"
)

const subjectTableName = "AttendancePlan_Subject"

// SubjectRepository は科目の repository を表すインターフェースです。
type SubjectRepository interface {
	ReadByUserID(userID string) ([]model.Subject, error)
	Create(subject *model.Subject) error
	Delete(id string) error
}

// SubjectRepositoryImpl は科目の repository の実装を表す構造体です。
type SubjectRepositoryImpl struct {
	DB    dynamo.DB
	Table dynamo.Table
}

// NewSubjectRepository は SubjectRepository を生成します。
func NewSubjectRepository(db dynamo.DB) SubjectRepository {
	return &SubjectRepositoryImpl{DB: db, Table: db.Table(subjectTableName)}
}

// ReadByUserID は指定されたユーザー ID の科目を取得します。
func (r *SubjectRepositoryImpl) ReadByUserID(userID string) ([]model.Subject, error) {
	var subjects []model.Subject
	err := r.Table.Get("UserID", userID).Index("UserID-index").Order(dynamo.Ascending).All(&subjects)
	if err != nil {
		return nil, err
	}

	return subjects, nil
}

// Create は科目を作成します。
func (r *SubjectRepositoryImpl) Create(subject *model.Subject) error {
	return r.Table.Put(subject).Run()
}

// Delete は指定された ID の科目を削除します。
func (r *SubjectRepositoryImpl) Delete(id string) error {
	return r.Table.Delete("ID", id).Run()
}
