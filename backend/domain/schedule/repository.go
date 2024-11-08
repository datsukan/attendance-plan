package schedule

import (
	"github.com/guregu/dynamo"
)

// ScheduleRepository はスケジュールの repository を表すインターフェースです。
type ScheduleRepository interface {
	ByUserID(userID string) ([]*Schedule, error)
	Read(ID string) (*Schedule, error)
	Create(schedule *Schedule) error
	Update(schedule *Schedule) error
	Delete(ID string) error
	Exists(ID string) (bool, error)
}

// ScheduleRepositoryImpl はスケジュールの repository の実装を表す構造体です。
type ScheduleRepositoryImpl struct {
	DB    dynamo.DB
	Table dynamo.Table
}

// NewScheduleRepository は ScheduleRepository を生成します。
func NewScheduleRepository(db dynamo.DB) ScheduleRepository {
	return &ScheduleRepositoryImpl{DB: db, Table: db.Table("schedule")}
}

// ByUserID は指定されたユーザー ID に紐づくスケジュールのリストを取得します。
func (r *ScheduleRepositoryImpl) ByUserID(userID string) ([]*Schedule, error) {
	var schedules []*Schedule
	err := r.Table.Get("user_id", userID).All(&schedules)
	if err != nil {
		return nil, err
	}
	return schedules, nil
}

// Read は指定された ID のスケジュールを取得します。
func (r *ScheduleRepositoryImpl) Read(ID string) (*Schedule, error) {
	var schedule Schedule
	err := r.Table.Get("id", ID).One(&schedule)
	if err != nil {
		return nil, err
	}
	return &schedule, nil
}

// Create はスケジュールを保存します。
func (r *ScheduleRepositoryImpl) Create(schedule *Schedule) error {
	err := r.Table.Put(schedule).Run()
	if err != nil {
		return err
	}
	return nil
}

// Update はスケジュールを更新します。
func (r *ScheduleRepositoryImpl) Update(schedule *Schedule) error {
	err := r.Table.Put(schedule).Run()
	if err != nil {
		return err
	}
	return nil
}

// Delete はスケジュールを削除します。
func (r *ScheduleRepositoryImpl) Delete(ID string) error {
	err := r.Table.Delete("id", ID).Run()
	if err != nil {
		return err
	}
	return nil
}

// Exists は指定された ID のスケジュールが存在するかどうかを返します。
func (r *ScheduleRepositoryImpl) Exists(ID string) (bool, error) {
	var schedule Schedule
	err := r.Table.Get("id", ID).One(&schedule)
	if err != nil {
		if err == dynamo.ErrNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
