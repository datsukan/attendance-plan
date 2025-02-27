package usecase

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/datsukan/attendance-plan/backend/app/component/id"
	"github.com/datsukan/attendance-plan/backend/app/model"
	"github.com/datsukan/attendance-plan/backend/app/port"
	"github.com/datsukan/attendance-plan/backend/app/repository"
)

// ScheduleInteractor はスケジュールのユースケースの実装を表す構造体です。
type ScheduleInteractor struct {
	Logger             *slog.Logger
	ScheduleRepository repository.ScheduleRepository
	OutputPort         port.ScheduleOutputPort
}

// NewScheduleInteractor は ScheduleInteractor を生成します。
func NewScheduleInteractor(logger *slog.Logger, scheduleRepository repository.ScheduleRepository, outputPort port.ScheduleOutputPort) port.ScheduleInputPort {
	return &ScheduleInteractor{
		Logger:             logger,
		ScheduleRepository: scheduleRepository,
		OutputPort:         outputPort,
	}
}

// GetScheduleList はスケジュールリストを取得します。
func (i *ScheduleInteractor) GetScheduleList(input port.GetScheduleListInputData) {
	schedules, err := i.ScheduleRepository.ReadByUserID(input.UserID)
	if err != nil {
		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
		i.OutputPort.SetResponseGetScheduleList(nil, r)
		return
	}

	dil := model.ScheduleList(schedules).ToDateItemList()
	dilMap := dil.ToTypeMap()
	masterDateItems := dilMap[model.ScheduleTypeMaster]
	customDateItems := dilMap[model.ScheduleTypeCustom]

	appendSchedules := func(dis model.DateItemList) []port.BaseDateItemData {
		res := make([]port.BaseDateItemData, 0, len(dis))
		for _, di := range dis {
			schedules := make([]port.BaseScheduleData, 0, len(di.Schedules))
			for _, s := range di.Schedules {
				schedules = append(schedules, port.BaseScheduleData{
					ID:        s.ID,
					UserID:    s.UserID,
					Name:      s.Name,
					StartsAt:  s.StartsAt.Format(time.DateTime),
					EndsAt:    s.EndsAt.Format(time.DateTime),
					Color:     s.Color,
					Type:      s.Type.String(),
					Order:     s.Order.Int(),
					CreatedAt: s.CreatedAt.Format(time.DateTime),
					UpdatedAt: s.UpdatedAt.Format(time.DateTime),
				})
			}
			res = append(res, port.BaseDateItemData{
				Date:      di.Date.Format(model.DateFormat),
				Type:      di.Type.String(),
				Schedules: schedules,
			})
		}
		return res
	}

	masterSchedules := appendSchedules(masterDateItems)
	customSchedules := appendSchedules(customDateItems)

	o := &port.GetScheduleListOutputData{MasterSchedules: masterSchedules, CustomSchedules: customSchedules}
	r := port.NewSuccessResult(http.StatusOK)
	i.OutputPort.SetResponseGetScheduleList(o, r)
}

// GetSchedule はスケジュールを取得します。
func (i *ScheduleInteractor) GetSchedule(input port.GetScheduleInputData) {
	i.Logger.With("schedule_id", input.ScheduleID)

	schedule, err := i.ScheduleRepository.Read(input.ScheduleID)
	if err != nil {
		if errors.Is(err, repository.NewNotFoundError()) {
			i.Logger.Warn(err.Error())
			r := port.NewErrorResult(http.StatusNotFound, MsgScheduleNotFound)
			i.OutputPort.SetResponseGetSchedule(nil, r)
			return
		}

		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
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
		Order:     schedule.Order.Int(),
		CreatedAt: schedule.CreatedAt.Format(time.DateTime),
		UpdatedAt: schedule.UpdatedAt.Format(time.DateTime),
	}

	o := &port.GetScheduleOutputData{Schedule: s}
	r := port.NewSuccessResult(http.StatusOK)
	i.OutputPort.SetResponseGetSchedule(o, r)
}

// CreateSchedule はスケジュールを作成します。
func (i *ScheduleInteractor) CreateSchedule(input port.CreateScheduleInputData) {
	startsAt, err := time.Parse(time.DateTime, input.Schedule.StartsAt)
	if err != nil {
		i.Logger.Warn(err.Error())
		r := port.NewErrorResult(http.StatusBadRequest, fmt.Sprintf(MsgFormatInvalid, "開始日"))
		i.OutputPort.SetResponseCreateSchedule(nil, r)
		return
	}

	endsAt, err := time.Parse(time.DateTime, input.Schedule.EndsAt)
	if err != nil {
		i.Logger.Warn(err.Error())
		r := port.NewErrorResult(http.StatusBadRequest, fmt.Sprintf(MsgFormatInvalid, "終了日"))
		i.OutputPort.SetResponseCreateSchedule(nil, r)
		return
	}

	sType := model.ToScheduleType(input.Schedule.Type)

	order := model.Order(input.Schedule.Order)
	if order.Empty() {
		someStartsAtSchedules, err := i.ScheduleRepository.ReadByUserIDStartsAt(input.Schedule.UserID, startsAt)
		if err != nil {
			i.Logger.Error(err.Error())
			r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
			i.OutputPort.SetResponseCreateSchedule(nil, r)
			return
		}

		filteredSchedules := model.ScheduleList(someStartsAtSchedules).FilterByType(sType)
		order = filteredSchedules.NextOrder()
	}

	s := model.Schedule{
		ID:        id.NewID(),
		UserID:    input.Schedule.UserID,
		Name:      input.Schedule.Name,
		StartsAt:  startsAt,
		EndsAt:    endsAt,
		Color:     input.Schedule.Color,
		Type:      sType,
		Order:     order,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	i.Logger.With("schedule_id", s.ID)

	if err := i.ScheduleRepository.Create(&s); err != nil {
		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
		i.OutputPort.SetResponseCreateSchedule(nil, r)
		return
	}

	o := &port.CreateScheduleOutputData{
		Schedule: port.BaseScheduleData{
			ID:        s.ID,
			UserID:    s.UserID,
			Name:      s.Name,
			StartsAt:  s.StartsAt.Format(time.DateTime),
			EndsAt:    s.EndsAt.Format(time.DateTime),
			Color:     s.Color,
			Type:      s.Type.String(),
			Order:     s.Order.Int(),
			CreatedAt: s.CreatedAt.Format(time.DateTime),
			UpdatedAt: s.UpdatedAt.Format(time.DateTime),
		},
	}
	r := port.NewSuccessResult(http.StatusCreated)
	i.OutputPort.SetResponseCreateSchedule(o, r)
}

// CreateBulkSchedule はスケジュールを一括作成します。
func (i *ScheduleInteractor) CreateBulkSchedule(input port.CreateBulkScheduleInputData) {
	responseSchedules := make([]port.BaseScheduleData, 0, len(input.Schedules))

	for _, s := range input.Schedules {
		startsAt, err := time.Parse(time.DateTime, s.StartsAt)
		if err != nil {
			i.Logger.Warn(err.Error())
			r := port.NewErrorResult(http.StatusBadRequest, fmt.Sprintf(MsgFormatInvalid, s.Name+"の開始日"))
			i.OutputPort.SetResponseCreateBulkSchedule(nil, r)
			return
		}

		endsAt, err := time.Parse(time.DateTime, s.EndsAt)
		if err != nil {
			i.Logger.Warn(err.Error())
			r := port.NewErrorResult(http.StatusBadRequest, fmt.Sprintf(MsgFormatInvalid, s.Name+"の終了日"))
			i.OutputPort.SetResponseCreateBulkSchedule(nil, r)
			return
		}

		sType := model.ToScheduleType(s.Type)

		order := model.Order(s.Order)
		if order.Empty() {
			someStartsAtSchedules, err := i.ScheduleRepository.ReadByUserIDStartsAt(s.UserID, startsAt)
			if err != nil {
				i.Logger.Error(err.Error())
				r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
				i.OutputPort.SetResponseCreateBulkSchedule(nil, r)
				return
			}

			filteredSchedules := model.ScheduleList(someStartsAtSchedules).FilterByType(sType)
			order = filteredSchedules.NextOrder()
		}

		s := model.Schedule{
			ID:        id.NewID(),
			UserID:    s.UserID,
			Name:      s.Name,
			StartsAt:  startsAt,
			EndsAt:    endsAt,
			Color:     s.Color,
			Type:      sType,
			Order:     order,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		i.Logger.With("schedule_id", s.ID)

		if err := i.ScheduleRepository.Create(&s); err != nil {
			i.Logger.Error(err.Error())
			r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
			i.OutputPort.SetResponseCreateBulkSchedule(nil, r)
			return
		}

		o := port.BaseScheduleData{
			ID:        s.ID,
			UserID:    s.UserID,
			Name:      s.Name,
			StartsAt:  s.StartsAt.Format(time.DateTime),
			EndsAt:    s.EndsAt.Format(time.DateTime),
			Color:     s.Color,
			Type:      s.Type.String(),
			Order:     s.Order.Int(),
			CreatedAt: s.CreatedAt.Format(time.DateTime),
			UpdatedAt: s.UpdatedAt.Format(time.DateTime),
		}

		responseSchedules = append(responseSchedules, o)
	}

	o := &port.CreateBulkScheduleOutputData{Schedules: responseSchedules}
	r := port.NewSuccessResult(http.StatusCreated)
	i.OutputPort.SetResponseCreateBulkSchedule(o, r)
}

// UpdateSchedule はスケジュールを更新します。
func (i *ScheduleInteractor) UpdateSchedule(input port.UpdateScheduleInputData) {
	i.Logger.With("schedule_id", input.Schedule.ID)

	startsAt, err := time.Parse(time.DateTime, input.Schedule.StartsAt)
	if err != nil {
		i.Logger.Warn(err.Error())
		r := port.NewErrorResult(http.StatusBadRequest, fmt.Sprintf(MsgFormatInvalid, "開始日"))
		i.OutputPort.SetResponseUpdateSchedule(nil, r)
		return
	}

	endsAt, err := time.Parse(time.DateTime, input.Schedule.EndsAt)
	if err != nil {
		i.Logger.Warn(err.Error())
		r := port.NewErrorResult(http.StatusBadRequest, fmt.Sprintf(MsgFormatInvalid, "終了日"))
		i.OutputPort.SetResponseUpdateSchedule(nil, r)
		return
	}

	sType := model.ToScheduleType(input.Schedule.Type)

	bs, err := i.ScheduleRepository.Read(input.Schedule.ID)
	if err != nil {
		if repository.IsNotFoundError(err) {
			i.Logger.Warn(err.Error())
			r := port.NewErrorResult(http.StatusNotFound, MsgScheduleNotFound)
			i.OutputPort.SetResponseUpdateSchedule(nil, r)
			return
		}

		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
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
		Order:     model.Order(input.Schedule.Order),
		CreatedAt: bs.CreatedAt,
		UpdatedAt: time.Now(),
	}

	if err := i.ScheduleRepository.Update(&s); err != nil {
		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
		i.OutputPort.SetResponseUpdateSchedule(nil, r)
		return
	}

	as, err := i.ScheduleRepository.Read(s.ID)
	if err != nil {
		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
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
			Order:     as.Order.Int(),
			CreatedAt: as.CreatedAt.Format(time.DateTime),
			UpdatedAt: as.UpdatedAt.Format(time.DateTime),
		},
	}
	r := port.NewSuccessResult(http.StatusOK)
	i.OutputPort.SetResponseUpdateSchedule(o, r)
}

// UpdateBulkSchedule はスケジュールを一括更新します。
func (i *ScheduleInteractor) UpdateBulkSchedule(input port.UpdateBulkScheduleInputData) {
	responseSchedules := make([]port.BaseScheduleData, 0, len(input.Schedules))

	for _, s := range input.Schedules {
		startsAt, err := time.Parse(time.DateTime, s.StartsAt)
		if err != nil {
			i.Logger.Warn(err.Error())
			r := port.NewErrorResult(http.StatusBadRequest, fmt.Sprintf(MsgFormatInvalid, s.Name+"の開始日"))
			i.OutputPort.SetResponseUpdateBulkSchedule(nil, r)
			return
		}

		endsAt, err := time.Parse(time.DateTime, s.EndsAt)
		if err != nil {
			i.Logger.Warn(err.Error())
			r := port.NewErrorResult(http.StatusBadRequest, fmt.Sprintf(MsgFormatInvalid, s.Name+"の終了日"))
			i.OutputPort.SetResponseUpdateBulkSchedule(nil, r)
			return
		}

		sType := model.ToScheduleType(s.Type)

		bs, err := i.ScheduleRepository.Read(s.ID)
		if err != nil {
			if repository.IsNotFoundError(err) {
				i.Logger.Warn(err.Error())
				r := port.NewErrorResult(http.StatusNotFound, MsgScheduleNotFound)
				i.OutputPort.SetResponseUpdateBulkSchedule(nil, r)
				return
			}

			i.Logger.Error(err.Error())
			r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
			i.OutputPort.SetResponseUpdateBulkSchedule(nil, r)
			return
		}

		s := model.Schedule{
			ID:        s.ID,
			UserID:    bs.UserID,
			Name:      s.Name,
			StartsAt:  startsAt,
			EndsAt:    endsAt,
			Color:     s.Color,
			Type:      sType,
			Order:     model.Order(s.Order),
			CreatedAt: bs.CreatedAt,
			UpdatedAt: time.Now(),
		}

		if err := i.ScheduleRepository.Update(&s); err != nil {
			i.Logger.Error(err.Error())
			r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
			i.OutputPort.SetResponseUpdateBulkSchedule(nil, r)
			return
		}

		as, err := i.ScheduleRepository.Read(s.ID)
		if err != nil {
			i.Logger.Error(err.Error())
			r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
			i.OutputPort.SetResponseUpdateBulkSchedule(nil, r)
			return
		}

		o := port.BaseScheduleData{
			ID:        as.ID,
			UserID:    as.UserID,
			Name:      as.Name,
			StartsAt:  as.StartsAt.Format(time.DateTime),
			EndsAt:    as.EndsAt.Format(time.DateTime),
			Color:     as.Color,
			Type:      as.Type.String(),
			Order:     as.Order.Int(),
			CreatedAt: as.CreatedAt.Format(time.DateTime),
			UpdatedAt: as.UpdatedAt.Format(time.DateTime),
		}

		responseSchedules = append(responseSchedules, o)
	}

	o := &port.UpdateBulkScheduleOutputData{Schedules: responseSchedules}
	r := port.NewSuccessResult(http.StatusOK)
	i.OutputPort.SetResponseUpdateBulkSchedule(o, r)
}

// DeleteSchedule はスケジュールを削除します。
func (i *ScheduleInteractor) DeleteSchedule(input port.DeleteScheduleInputData) {
	i.Logger.With("schedule_id", input.ScheduleID)

	if err := i.ScheduleRepository.Delete(input.ScheduleID); err != nil {
		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
		i.OutputPort.SetResponseDeleteSchedule(nil, r)
		return
	}

	o := &port.DeleteScheduleOutputData{ScheduleID: input.ScheduleID}
	r := port.NewSuccessResult(http.StatusNoContent)
	i.OutputPort.SetResponseDeleteSchedule(o, r)
}
