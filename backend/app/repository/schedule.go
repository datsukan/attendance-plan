package repository

import (
	"time"

	"github.com/datsukan/attendance-plan/backend/app/model"
	"github.com/guregu/dynamo"
)

const scheduleTableName = "AttendancePlan_Schedule"

// ScheduleRepository はスケジュールの repository を表すインターフェースです。
type ScheduleRepository interface {
	Read(id string) (*model.Schedule, error)
	ReadByUserID(userID string) ([]model.Schedule, error)
	ReadByUserIDStartsAt(userID string, startsAt time.Time) ([]model.Schedule, error)
	Create(schedule *model.Schedule) error
	Update(schedule *model.Schedule) error
	Delete(id string) error
	Exists(id string) (bool, error)
}

// ScheduleRepositoryImpl はスケジュールの repository の実装を表す構造体です。
type ScheduleRepositoryImpl struct {
	DB    dynamo.DB
	Table dynamo.Table
}

// NewScheduleRepository は ScheduleRepository を生成します。
func NewScheduleRepository(db dynamo.DB) ScheduleRepository {
	return &ScheduleRepositoryImpl{DB: db, Table: db.Table(scheduleTableName)}
}

// Read は指定された ID のスケジュールを取得します。
func (r *ScheduleRepositoryImpl) Read(id string) (*model.Schedule, error) {
	var schedule *model.Schedule
	err := r.Table.Get("ID", id).One(&schedule)
	if err != nil {
		if err == dynamo.ErrNotFound {
			return nil, NewNotFoundError()
		}

		return nil, err
	}
	return schedule, nil
}

// ReadByUserID は指定されたユーザー ID に紐づくスケジュールのリストを取得します。
func (r *ScheduleRepositoryImpl) ReadByUserID(userID string) ([]model.Schedule, error) {
	var schedules []model.Schedule
	err := r.Table.Get("UserID", userID).Index("UserID-index").Order(dynamo.Ascending).All(&schedules)
	if err != nil {
		return nil, err
	}
	return schedules, nil
}

// ReadByUserIDStartsAt は指定されたユーザー ID と開始日時に紐づくスケジュールを取得します。
func (r *ScheduleRepositoryImpl) ReadByUserIDStartsAt(userID string, startsAt time.Time) ([]model.Schedule, error) {
	var schedules []model.Schedule
	err := r.Table.Get("UserID", userID).Range("StartsAt", dynamo.Equal, startsAt).Index("UserID-index").Order(dynamo.Ascending).All(&schedules)
	if err != nil {
		return nil, err
	}
	return schedules, nil
}

// Create はスケジュールを保存します。
func (r *ScheduleRepositoryImpl) Create(schedule *model.Schedule) error {
	return r.Table.Put(schedule).Run()
}

// Update はスケジュールを更新します。
func (r *ScheduleRepositoryImpl) Update(schedule *model.Schedule) error {
	return r.Table.Put(schedule).Run()
}

// Delete はスケジュールを削除します。
func (r *ScheduleRepositoryImpl) Delete(id string) error {
	return r.Table.Delete("ID", id).Run()
}

// Exists は指定された ID のスケジュールが存在するかどうかを返します。
func (r *ScheduleRepositoryImpl) Exists(id string) (bool, error) {
	var schedule *model.Schedule
	err := r.Table.Get("ID", id).One(&schedule)
	if err != nil {
		if err == dynamo.ErrNotFound {
			return false, nil
		}
		return false, err
	}

	return schedule != nil, nil
}
