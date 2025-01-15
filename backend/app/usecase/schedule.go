package usecase

import (
	"errors"
	"net/http"
	"time"

	"github.com/datsukan/attendance-plan/backend/app/model"
	"github.com/datsukan/attendance-plan/backend/app/port"
	"github.com/datsukan/attendance-plan/backend/app/repository"
	ulid "github.com/oklog/ulid/v2"
)

// ScheduleInteractor はスケジュールのユースケースの実装を表す構造体です。
type ScheduleInteractor struct {
	ScheduleRepository repository.ScheduleRepository
	OutputPort         port.ScheduleOutputPort
}

// NewScheduleInteractor は ScheduleInteractor を生成します。
func NewScheduleInteractor(scheduleRepository repository.ScheduleRepository, outputPort port.ScheduleOutputPort) *ScheduleInteractor {
	return &ScheduleInteractor{ScheduleRepository: scheduleRepository, OutputPort: outputPort}
}

// GetScheduleList はスケジュールリストを取得します。
func (i *ScheduleInteractor) GetScheduleList(input port.GetScheduleListInputData) {
	schedules, err := i.ScheduleRepository.ReadByUserID(input.UserID)
	if err != nil {
		r := port.Result{StatusCode: http.StatusInternalServerError, HasError: true, Message: err.Error()}
		i.OutputPort.SetResponseGetScheduleList(nil, r)
		return
	}

	var outputSchedules []port.BaseScheduleData
	for _, schedule := range schedules {
		s := port.BaseScheduleData{
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

	o := &port.GetScheduleListOutputData{Schedules: outputSchedules}
	r := port.Result{StatusCode: http.StatusOK, Message: "Success"}
	i.OutputPort.SetResponseGetScheduleList(o, r)
}

// GetSchedule はスケジュールを取得します。
func (i *ScheduleInteractor) GetSchedule(input port.GetScheduleInputData) {
	schedule, err := i.ScheduleRepository.Read(input.ScheduleID)
	if err != nil {
		if errors.Is(err, repository.NewNotFoundError()) {
			r := port.Result{StatusCode: http.StatusNotFound, HasError: true, Message: err.Error()}
			i.OutputPort.SetResponseGetSchedule(nil, r)
			return
		}

		r := port.Result{StatusCode: http.StatusInternalServerError, HasError: true, Message: err.Error()}
		i.OutputPort.SetResponseGetSchedule(nil, r)
		return
	}

	s := port.BaseScheduleData{
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

	o := &port.GetScheduleOutputData{Schedule: s}
	r := port.Result{StatusCode: http.StatusOK, Message: "Success"}
	i.OutputPort.SetResponseGetSchedule(o, r)
}

// CreateSchedule はスケジュールを作成します。
func (i *ScheduleInteractor) CreateSchedule(input port.CreateScheduleInputData) {
	startsAt, err := time.Parse(time.DateTime, input.Schedule.StartsAt)
	if err != nil {
		r := port.Result{StatusCode: http.StatusBadRequest, HasError: true, Message: err.Error()}
		i.OutputPort.SetResponseCreateSchedule(nil, r)
		return
	}

	endsAt, err := time.Parse(time.DateTime, input.Schedule.EndsAt)
	if err != nil {
		r := port.Result{StatusCode: http.StatusBadRequest, HasError: true, Message: err.Error()}
		i.OutputPort.SetResponseCreateSchedule(nil, r)
		return
	}

	sType := model.ToScheduleType(input.Schedule.Type)

	s := model.Schedule{
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
		r := port.Result{StatusCode: http.StatusInternalServerError, HasError: true, Message: err.Error()}
		i.OutputPort.SetResponseCreateSchedule(nil, r)
		return
	}

	o := &port.CreateScheduleOutputData{
		Schedule: s,
	}
	r := port.Result{StatusCode: http.StatusCreated, Message: "Success"}
	i.OutputPort.SetResponseCreateSchedule(o, r)
}

// UpdateSchedule はスケジュールを更新します。
func (i *ScheduleInteractor) UpdateSchedule(input port.UpdateScheduleInputData) {
	startsAt, err := time.Parse(time.DateTime, input.Schedule.StartsAt)
	if err != nil {
		r := port.Result{StatusCode: http.StatusBadRequest, HasError: true, Message: err.Error()}
		i.OutputPort.SetResponseUpdateSchedule(nil, r)
		return
	}

	endsAt, err := time.Parse(time.DateTime, input.Schedule.EndsAt)
	if err != nil {
		r := port.Result{StatusCode: http.StatusBadRequest, HasError: true, Message: err.Error()}
		i.OutputPort.SetResponseUpdateSchedule(nil, r)
		return
	}

	sType := model.ToScheduleType(input.Schedule.Type)

	bs, err := i.ScheduleRepository.Read(input.Schedule.ID)
	if err != nil {
		if repository.IsNotFoundError(err) {
			r := port.Result{StatusCode: http.StatusNotFound, HasError: true, Message: err.Error()}
			i.OutputPort.SetResponseUpdateSchedule(nil, r)
			return
		}

		r := port.Result{StatusCode: http.StatusInternalServerError, HasError: true, Message: err.Error()}
		i.OutputPort.SetResponseUpdateSchedule(nil, r)
		return
	}

	s := model.Schedule{
		ID:        input.Schedule.ID,
		UserID:    bs.UserID,
		Name:      input.Schedule.Name,
		StartsAt:  startsAt,
		EndsAt:    endsAt,
		Color:     input.Schedule.Color,
		Type:      sType,
		CreatedAt: bs.CreatedAt,
		UpdatedAt: time.Now(),
	}

	if err := i.ScheduleRepository.Update(&s); err != nil {
		r := port.Result{StatusCode: http.StatusInternalServerError, HasError: true, Message: err.Error()}
		i.OutputPort.SetResponseUpdateSchedule(nil, r)
		return
	}

	as, err := i.ScheduleRepository.Read(s.ID)
	if err != nil {
		r := port.Result{StatusCode: http.StatusInternalServerError, HasError: true, Message: err.Error()}
		i.OutputPort.SetResponseUpdateSchedule(nil, r)
		return
	}

	o := &port.UpdateScheduleOutputData{
		Schedule: port.BaseScheduleData{
			ID:        as.ID,
			UserID:    as.UserID,
			Name:      as.Name,
			StartsAt:  as.StartsAt.Format(time.DateTime),
			EndsAt:    as.EndsAt.Format(time.DateTime),
			Color:     as.Color,
			Type:      as.Type.String(),
			CreatedAt: as.CreatedAt.Format(time.DateTime),
			UpdatedAt: as.UpdatedAt.Format(time.DateTime),
		},
	}
	r := port.Result{StatusCode: http.StatusOK, Message: "Success"}
	i.OutputPort.SetResponseUpdateSchedule(o, r)
}

// DeleteSchedule はスケジュールを削除します。
func (i *ScheduleInteractor) DeleteSchedule(input port.DeleteScheduleInputData) {
	if err := i.ScheduleRepository.Delete(input.ScheduleID); err != nil {
		r := port.Result{StatusCode: http.StatusInternalServerError, HasError: true, Message: err.Error()}
		i.OutputPort.SetResponseDeleteSchedule(nil, r)
		return
	}

	o := &port.DeleteScheduleOutputData{ScheduleID: input.ScheduleID}
	r := port.Result{StatusCode: http.StatusNoContent, Message: "Success"}
	i.OutputPort.SetResponseDeleteSchedule(o, r)
}
