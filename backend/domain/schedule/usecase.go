package schedule

import (
	"net/http"
	"time"

	ulid "github.com/oklog/ulid/v2"
)

// BaseScheduleData はスケジュールの基本データを表す構造体です。
type BaseScheduleData struct {
	ID        string
	UserID    string
	Name      string
	StartsAt  string
	EndsAt    string
	Color     string
	Type      string
	CreatedAt string
	UpdatedAt string
}

// GetScheduleListInputData はスケジュールリスト取得の入力データを表す構造体です。
type GetScheduleListInputData struct {
	UserID string
}

// GetScheduleListOutputData はスケジュールリスト取得の出力データを表す構造体です。
type GetScheduleListOutputData struct {
	Schedules []BaseScheduleData
}

// GetScheduleInputData はスケジュール取得の入力データを表す構造体です。
type GetScheduleInputData struct {
	ScheduleID string
}

// GetScheduleOutputData はスケジュール取得の出力データを表す構造体です。
type GetScheduleOutputData struct {
	Schedule BaseScheduleData
}

// CreateScheduleData はスケジュール作成のスケジュールデータを表す構造体です。
type CreateScheduleData struct {
	UserID   string
	Name     string
	StartsAt string
	EndsAt   string
	Color    string
	Type     string
}

// CreateScheduleData はスケジュール作成のデータを表す構造体です。
type CreateScheduleInputData struct {
	Schedule CreateScheduleData
}

// CreateScheduleOutputData はスケジュール作成の出力データを表す構造体です。
type CreateScheduleOutputData struct {
	Schedule
}

// UpdateScheduleData はスケジュール更新のスケジュールデータを表す構造体です。
type UpdateScheduleData struct {
	ID       string
	Name     string
	StartsAt string
	EndsAt   string
	Color    string
	Type     string
}

// UpdateScheduleInputData はスケジュール更新の入力データを表す構造体です。
type UpdateScheduleInputData struct {
	Schedule UpdateScheduleData
}

// UpdateScheduleOutputData はスケジュール更新の出力データを表す構造体です。
type UpdateScheduleOutputData struct {
	Schedule BaseScheduleData
}

// DeleteScheduleInputData はスケジュール削除の入力データを表す構造体です。
type DeleteScheduleInputData struct {
	ScheduleID string
}

// DeleteScheduleOutputData はスケジュール削除の出力データを表す構造体です。
type DeleteScheduleOutputData struct {
	ScheduleID string
}

// ScheduleInputPort はスケジュールのユースケースを表すインターフェースです。
type ScheduleInputPort interface {
	GetScheduleList(input GetScheduleListInputData)
	GetSchedule(input GetScheduleInputData)
	CreateSchedule(input CreateScheduleInputData)
	UpdateSchedule(input UpdateScheduleInputData)
	DeleteSchedule(input DeleteScheduleInputData)
}

// ScheduleOutputPort はスケジュールの外部出力を表すインターフェースです。
type ScheduleOutputPort interface {
	GetResponse() (int, string)
	ResponseGetScheduleList(output *GetScheduleListOutputData, result Result)
	ResponseGetSchedule(output *GetScheduleOutputData, result Result)
	ResponseCreateSchedule(output *CreateScheduleOutputData, result Result)
	ResponseUpdateSchedule(output *UpdateScheduleOutputData, result Result)
	ResponseDeleteSchedule(output *DeleteScheduleOutputData, result Result)
}

// ScheduleInteractor はスケジュールのユースケースの構造体です。
type ScheduleInteractor struct {
	ScheduleRepository ScheduleRepository
	OutputPort         ScheduleOutputPort
}

// NewScheduleInteractor は ScheduleInteractor を生成します。
func NewScheduleInteractor(scheduleRepository ScheduleRepository, outputPort ScheduleOutputPort) *ScheduleInteractor {
	return &ScheduleInteractor{ScheduleRepository: scheduleRepository, OutputPort: outputPort}
}

// GetScheduleList はスケジュールリストを取得します。
func (i *ScheduleInteractor) GetScheduleList(input GetScheduleListInputData) {
	schedules, err := i.ScheduleRepository.ReadByUserID(input.UserID)
	if err != nil {
		r := Result{StatusCode: http.StatusInternalServerError, HasError: true, Message: err.Error()}
		i.OutputPort.ResponseGetScheduleList(nil, r)
		return
	}

	var outputSchedules []BaseScheduleData
	for _, schedule := range schedules {
		s := BaseScheduleData{
			ID:        schedule.ID,
			UserID:    schedule.UserID,
			Name:      schedule.Name,
			StartsAt:  schedule.StartsAt.Format(time.DateTime),
			EndsAt:    schedule.EndsAt.Format(time.DateTime),
			Color:     schedule.Color,
			Type:      schedule.Type.String(),
			CreatedAt: schedule.CreatedAt.Format(time.DateTime),
			UpdatedAt: schedule.UpdatedAt.Format(time.DateTime),
		}
		outputSchedules = append(outputSchedules, s)
	}

	o := &GetScheduleListOutputData{Schedules: outputSchedules}
	r := Result{StatusCode: http.StatusOK, Message: "Success"}
	i.OutputPort.ResponseGetScheduleList(o, r)
}

// GetSchedule はスケジュールを取得します。
func (i *ScheduleInteractor) GetSchedule(input GetScheduleInputData) {
	schedule, err := i.ScheduleRepository.Read(input.ScheduleID)
	if err != nil {
		r := Result{StatusCode: http.StatusInternalServerError, HasError: true, Message: err.Error()}
		i.OutputPort.ResponseGetSchedule(nil, r)
		return
	}

	s := BaseScheduleData{
		ID:        schedule.ID,
		UserID:    schedule.UserID,
		Name:      schedule.Name,
		StartsAt:  schedule.StartsAt.Format(time.DateTime),
		EndsAt:    schedule.EndsAt.Format(time.DateTime),
		Color:     schedule.Color,
		Type:      schedule.Type.String(),
		CreatedAt: schedule.CreatedAt.Format(time.DateTime),
		UpdatedAt: schedule.UpdatedAt.Format(time.DateTime),
	}

	o := &GetScheduleOutputData{Schedule: s}
	r := Result{StatusCode: http.StatusOK, Message: "Success"}
	i.OutputPort.ResponseGetSchedule(o, r)
}

// CreateSchedule はスケジュールを作成します。
func (i *ScheduleInteractor) CreateSchedule(input CreateScheduleInputData) {
	startsAt, err := time.Parse(time.DateTime, input.Schedule.StartsAt)
	if err != nil {
		r := Result{StatusCode: http.StatusBadRequest, HasError: true, Message: err.Error()}
		i.OutputPort.ResponseCreateSchedule(nil, r)
		return
	}

	endsAt, err := time.Parse(time.DateTime, input.Schedule.EndsAt)
	if err != nil {
		r := Result{StatusCode: http.StatusBadRequest, HasError: true, Message: err.Error()}
		i.OutputPort.ResponseCreateSchedule(nil, r)
		return
	}

	sType := ToScheduleType(input.Schedule.Type)

	s := Schedule{
		ID:        ulid.Make().String(),
		UserID:    input.Schedule.UserID,
		Name:      input.Schedule.Name,
		StartsAt:  startsAt,
		EndsAt:    endsAt,
		Color:     input.Schedule.Color,
		Type:      sType,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := i.ScheduleRepository.Create(&s); err != nil {
		r := Result{StatusCode: http.StatusInternalServerError, HasError: true, Message: err.Error()}
		i.OutputPort.ResponseCreateSchedule(nil, r)
		return
	}

	o := &CreateScheduleOutputData{
		Schedule: s,
	}
	r := Result{StatusCode: http.StatusCreated}
	i.OutputPort.ResponseCreateSchedule(o, r)
}

// UpdateSchedule はスケジュールを更新します。
func (i *ScheduleInteractor) UpdateSchedule(input UpdateScheduleInputData) {
	startsAt, err := time.Parse(time.DateTime, input.Schedule.StartsAt)
	if err != nil {
		r := Result{StatusCode: http.StatusBadRequest, HasError: true, Message: err.Error()}
		i.OutputPort.ResponseUpdateSchedule(nil, r)
		return
	}

	endsAt, err := time.Parse(time.DateTime, input.Schedule.EndsAt)
	if err != nil {
		r := Result{StatusCode: http.StatusBadRequest, HasError: true, Message: err.Error()}
		i.OutputPort.ResponseUpdateSchedule(nil, r)
		return
	}

	sType := ToScheduleType(input.Schedule.Type)

	s := Schedule{
		ID:        input.Schedule.ID,
		Name:      input.Schedule.Name,
		StartsAt:  startsAt,
		EndsAt:    endsAt,
		Color:     input.Schedule.Color,
		Type:      sType,
		UpdatedAt: time.Now(),
	}

	if err := i.ScheduleRepository.Update(&s); err != nil {
		r := Result{StatusCode: http.StatusInternalServerError, HasError: true, Message: err.Error()}
		i.OutputPort.ResponseUpdateSchedule(nil, r)
		return
	}

	o := &UpdateScheduleOutputData{
		Schedule: BaseScheduleData{
			ID:        s.ID,
			UserID:    s.UserID,
			Name:      s.Name,
			StartsAt:  s.StartsAt.Format(time.DateTime),
			EndsAt:    s.EndsAt.Format(time.DateTime),
			Color:     s.Color,
			Type:      s.Type.String(),
			CreatedAt: s.CreatedAt.Format(time.DateTime),
			UpdatedAt: s.UpdatedAt.Format(time.DateTime),
		},
	}
	r := Result{StatusCode: http.StatusNoContent, Message: "Success"}
	i.OutputPort.ResponseUpdateSchedule(o, r)
}

// DeleteSchedule はスケジュールを削除します。
func (i *ScheduleInteractor) DeleteSchedule(input DeleteScheduleInputData) {
	if err := i.ScheduleRepository.Delete(input.ScheduleID); err != nil {
		r := Result{StatusCode: http.StatusInternalServerError, HasError: true, Message: err.Error()}
		i.OutputPort.ResponseDeleteSchedule(nil, r)
		return
	}

	o := &DeleteScheduleOutputData{ScheduleID: input.ScheduleID}
	r := Result{StatusCode: http.StatusNoContent, Message: "Success"}
	i.OutputPort.ResponseDeleteSchedule(o, r)
}
