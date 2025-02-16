package usecase

import (
	"net/http"
	"testing"
	"time"

	"github.com/datsukan/attendance-plan/backend/app/model"
	"github.com/datsukan/attendance-plan/backend/app/port"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetScheduleList(t *testing.T) {
	t.Run("スケジュールリストを取得する", func(t *testing.T) {
		require := require.New(t)
		assert := assert.New(t)

		r := &stubScheduleRepository{}
		p := &stubScheduleOutputPort{}
		i := NewScheduleInteractor(r, p)

		input := port.GetScheduleListInputData{UserID: "test-user-id"}
		i.GetScheduleList(input)

		startDate1 := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
		startDate2 := time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)
		startDate3 := time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)
		startDate4 := time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC)
		endDate := time.Date(2021, 1, 10, 0, 0, 0, 0, time.UTC)

		wantMasterSchedules := []model.Schedule{
			{ID: "test-id-1", UserID: "test-user-id", Name: "test-name-1", StartsAt: startDate1, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeMaster, Order: 1, CreatedAt: startDate1, UpdatedAt: startDate1},
			{ID: "test-id-3", UserID: "test-user-id", Name: "test-name-3", StartsAt: startDate1, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeMaster, Order: 2, CreatedAt: startDate1, UpdatedAt: startDate1},
			{ID: "test-id-5", UserID: "test-user-id", Name: "test-name-5", StartsAt: startDate2, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeMaster, Order: 1, CreatedAt: startDate1, UpdatedAt: startDate1},
			{ID: "test-id-7", UserID: "test-user-id", Name: "test-name-7", StartsAt: startDate2, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeMaster, Order: 2, CreatedAt: startDate1, UpdatedAt: startDate1},
			{ID: "test-id-9", UserID: "test-user-id", Name: "test-name-9", StartsAt: startDate3, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeMaster, Order: 1, CreatedAt: startDate1, UpdatedAt: startDate1},
			{ID: "test-id-11", UserID: "test-user-id", Name: "test-name-11", StartsAt: startDate3, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeMaster, Order: 2, CreatedAt: startDate1, UpdatedAt: startDate1},
		}
		wantCustomSchedules := []model.Schedule{
			{ID: "test-id-2", UserID: "test-user-id", Name: "test-name-2", StartsAt: startDate1, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeCustom, Order: 1, CreatedAt: startDate1, UpdatedAt: startDate1},
			{ID: "test-id-4", UserID: "test-user-id", Name: "test-name-4", StartsAt: startDate1, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeCustom, Order: 2, CreatedAt: startDate1, UpdatedAt: startDate1},
			{ID: "test-id-6", UserID: "test-user-id", Name: "test-name-6", StartsAt: startDate2, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeCustom, Order: 1, CreatedAt: startDate1, UpdatedAt: startDate1},
			{ID: "test-id-8", UserID: "test-user-id", Name: "test-name-8", StartsAt: startDate2, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeCustom, Order: 2, CreatedAt: startDate1, UpdatedAt: startDate1},
			{ID: "test-id-10", UserID: "test-user-id", Name: "test-name-10", StartsAt: startDate3, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeCustom, Order: 1, CreatedAt: startDate1, UpdatedAt: startDate1},
			{ID: "test-id-12", UserID: "test-user-id", Name: "test-name-12", StartsAt: startDate3, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeCustom, Order: 2, CreatedAt: startDate1, UpdatedAt: startDate1},
			{ID: "test-id-13", UserID: "test-user-id", Name: "test-name-13", StartsAt: startDate4, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeCustom, Order: 1, CreatedAt: startDate1, UpdatedAt: startDate1},
			{ID: "test-id-14", UserID: "test-user-id", Name: "test-name-14", StartsAt: startDate4, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeCustom, Order: 2, CreatedAt: startDate1, UpdatedAt: startDate1},
			{ID: "test-id-15", UserID: "test-user-id", Name: "test-name-15", StartsAt: startDate4, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeCustom, Order: 3, CreatedAt: startDate1, UpdatedAt: startDate1},
			{ID: "test-id-16", UserID: "test-user-id", Name: "test-name-16", StartsAt: startDate4, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeCustom, Order: 4, CreatedAt: startDate1, UpdatedAt: startDate1},
			{ID: "test-id-17", UserID: "test-user-id", Name: "test-name-17", StartsAt: startDate4, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeCustom, Order: 5, CreatedAt: startDate1, UpdatedAt: startDate1},
		}

		output, ok := p.Output.(*port.GetScheduleListOutputData)
		require.True(ok)

		if !assert.NotNil(output) {
			return
		}

		assert.Equal(http.StatusOK, p.Result.StatusCode)
		assert.Equal("Success", p.Result.Message)
		assert.False(p.Result.HasError)

		outputMasterLen := 0
		for _, s := range output.MasterSchedules {
			outputMasterLen += len(s.Schedules)
		}
		if !assert.Equal(len(wantMasterSchedules), outputMasterLen) {
			return
		}

		for i, s := range output.MasterSchedules {
			for j, ss := range s.Schedules {
				assert.Equal(wantMasterSchedules[i*2+j].ID, ss.ID)
				assert.Equal(wantMasterSchedules[i*2+j].UserID, ss.UserID)
				assert.Equal(wantMasterSchedules[i*2+j].Name, ss.Name)
				assert.Equal(wantMasterSchedules[i*2+j].StartsAt.Format(time.DateTime), ss.StartsAt)
				assert.Equal(wantMasterSchedules[i*2+j].EndsAt.Format(time.DateTime), ss.EndsAt)
				assert.Equal(wantMasterSchedules[i*2+j].Color, ss.Color)
				assert.Equal(wantMasterSchedules[i*2+j].Type.String(), ss.Type)
				assert.Equal(wantMasterSchedules[i*2+j].Order.Int(), ss.Order)
				assert.Equal(wantMasterSchedules[i*2+j].CreatedAt.Format(time.DateTime), ss.CreatedAt)
				assert.Equal(wantMasterSchedules[i*2+j].UpdatedAt.Format(time.DateTime), ss.UpdatedAt)
			}
		}

		outputCustomLen := 0
		for _, s := range output.CustomSchedules {
			outputCustomLen += len(s.Schedules)
		}
		if !assert.Equal(len(wantCustomSchedules), outputCustomLen) {
			return
		}

		for i, s := range output.CustomSchedules {
			for j, ss := range s.Schedules {
				assert.Equal(wantCustomSchedules[i*2+j].ID, ss.ID)
				assert.Equal(wantCustomSchedules[i*2+j].UserID, ss.UserID)
				assert.Equal(wantCustomSchedules[i*2+j].Name, ss.Name)
				assert.Equal(wantCustomSchedules[i*2+j].StartsAt.Format(time.DateTime), ss.StartsAt)
				assert.Equal(wantCustomSchedules[i*2+j].EndsAt.Format(time.DateTime), ss.EndsAt)
				assert.Equal(wantCustomSchedules[i*2+j].Color, ss.Color)
				assert.Equal(wantCustomSchedules[i*2+j].Type.String(), ss.Type)
				assert.Equal(wantCustomSchedules[i*2+j].Order.Int(), ss.Order)
				assert.Equal(wantCustomSchedules[i*2+j].CreatedAt.Format(time.DateTime), ss.CreatedAt)
				assert.Equal(wantCustomSchedules[i*2+j].UpdatedAt.Format(time.DateTime), ss.UpdatedAt)
			}
		}
	})
}

func TestGetSchedule(t *testing.T) {
	t.Run("スケジュールを取得する", func(t *testing.T) {
		require := require.New(t)
		assert := assert.New(t)

		r := &stubScheduleRepository{}
		p := &stubScheduleOutputPort{}
		i := NewScheduleInteractor(r, p)

		input := port.GetScheduleInputData{ScheduleID: "test-id"}
		i.GetSchedule(input)

		date := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
		wantSchedule := &model.Schedule{
			ID:        "test-id",
			UserID:    "test-user-id",
			Name:      "test-name",
			StartsAt:  date,
			EndsAt:    date,
			Color:     "test-color",
			Type:      model.ScheduleTypeMaster,
			Order:     1,
			CreatedAt: date,
			UpdatedAt: date,
		}

		output, ok := p.Output.(*port.GetScheduleOutputData)
		require.True(ok)

		if !assert.NotNil(output) {
			return
		}

		os := output.Schedule
		assert.Equal(wantSchedule.ID, os.ID)
		assert.Equal(wantSchedule.UserID, os.UserID)
		assert.Equal(wantSchedule.Name, os.Name)
		assert.Equal(wantSchedule.StartsAt.Format(time.DateTime), os.StartsAt)
		assert.Equal(wantSchedule.EndsAt.Format(time.DateTime), os.EndsAt)
		assert.Equal(wantSchedule.Color, os.Color)
		assert.Equal(wantSchedule.Type.String(), os.Type)
		assert.Equal(wantSchedule.Order.Int(), os.Order)
		assert.Equal(wantSchedule.CreatedAt.Format(time.DateTime), os.CreatedAt)
		assert.Equal(wantSchedule.UpdatedAt.Format(time.DateTime), os.UpdatedAt)

		assert.Equal(http.StatusOK, p.Result.StatusCode)
		assert.Equal("Success", p.Result.Message)
		assert.False(p.Result.HasError)
	})

	t.Run("スケジュールが存在しない場合はエラーを返す", func(t *testing.T) {
		assert := assert.New(t)

		r := &stubNotFoundScheduleRepository{}
		p := &stubScheduleOutputPort{}
		i := NewScheduleInteractor(r, p)

		input := port.GetScheduleInputData{ScheduleID: "not-found-id"}
		i.GetSchedule(input)

		assert.Equal(http.StatusNotFound, p.Result.StatusCode)
		assert.Equal("not found", p.Result.Message)
		assert.True(p.Result.HasError)
	})
}

func TestCreateSchedule(t *testing.T) {
	t.Run("スケジュールを作成する", func(t *testing.T) {
		assert := assert.New(t)

		r := &stubScheduleRepository{}
		p := &stubScheduleOutputPort{}
		i := NewScheduleInteractor(r, p)

		input := port.CreateScheduleInputData{
			Schedule: port.CreateScheduleData{
				Name:     "test-name",
				StartsAt: "2021-01-01 00:00:00",
				EndsAt:   "2021-01-01 00:00:00",
				Color:    "white",
				Type:     model.ScheduleTypeMaster.String(),
				Order:    1,
			},
		}
		i.CreateSchedule(input)

		assert.Equal(http.StatusCreated, p.Result.StatusCode)
		assert.Equal("Success", p.Result.Message)
		assert.False(p.Result.HasError)
	})
}

func TestUpdateSchedule(t *testing.T) {
	t.Run("スケジュールを更新する", func(t *testing.T) {
		assert := assert.New(t)

		r := &stubScheduleRepository{}
		p := &stubScheduleOutputPort{}
		i := NewScheduleInteractor(r, p)

		input := port.UpdateScheduleInputData{
			Schedule: port.UpdateScheduleData{
				ID:       "test-id",
				Name:     "test-name",
				StartsAt: "2021-01-01 00:00:00",
				EndsAt:   "2021-01-01 00:00:00",
				Color:    "white",
				Type:     model.ScheduleTypeMaster.String(),
				Order:    1,
			},
		}
		i.UpdateSchedule(input)

		assert.Equal(http.StatusOK, p.Result.StatusCode)
		assert.Equal("Success", p.Result.Message)
		assert.False(p.Result.HasError)
	})

	t.Run("スケジュールが存在しない場合はエラーを返す", func(t *testing.T) {
		assert := assert.New(t)

		r := &stubNotFoundScheduleRepository{}
		p := &stubScheduleOutputPort{}
		i := NewScheduleInteractor(r, p)

		input := port.UpdateScheduleInputData{
			Schedule: port.UpdateScheduleData{
				ID:       "not-found-id",
				Name:     "test-name",
				StartsAt: "2021-01-01 00:00:00",
				EndsAt:   "2021-01-01 00:00:00",
				Color:    "white",
				Type:     model.ScheduleTypeMaster.String(),
				Order:    1,
			},
		}
		i.UpdateSchedule(input)

		assert.Equal(http.StatusNotFound, p.Result.StatusCode)
		assert.Equal("not found", p.Result.Message)
		assert.True(p.Result.HasError)
	})
}

func TestUpdateBulkSchedule(t *testing.T) {
	t.Run("スケジュールを一括更新する", func(t *testing.T) {
		assert := assert.New(t)

		r := &stubScheduleRepository{}
		p := &stubScheduleOutputPort{}
		i := NewScheduleInteractor(r, p)

		input := port.UpdateBulkScheduleInputData{
			Schedules: []port.UpdateScheduleData{
				{
					ID:       "test-id-1",
					Name:     "test-name-1",
					StartsAt: "2021-01-01 00:00:00",
					EndsAt:   "2021-01-01 00:00:00",
					Color:    "white",
					Type:     model.ScheduleTypeMaster.String(),
					Order:    1,
				},
				{
					ID:       "test-id-2",
					Name:     "test-name-2",
					StartsAt: "2021-01-01 00:00:00",
					EndsAt:   "2021-01-01 00:00:00",
					Color:    "white",
					Type:     model.ScheduleTypeCustom.String(),
					Order:    1,
				},
			},
		}
		i.UpdateBulkSchedule(input)

		assert.Equal(http.StatusOK, p.Result.StatusCode)
		assert.Equal("Success", p.Result.Message)
		assert.False(p.Result.HasError)
	})

	t.Run("スケジュールが存在しない場合はエラーを返す", func(t *testing.T) {
		assert := assert.New(t)

		r := &stubNotFoundScheduleRepository{}
		p := &stubScheduleOutputPort{}
		i := NewScheduleInteractor(r, p)

		input := port.UpdateBulkScheduleInputData{
			Schedules: []port.UpdateScheduleData{
				{
					ID:       "not-found-id",
					Name:     "test-name",
					StartsAt: "2021-01-01 00:00:00",
					EndsAt:   "2021-01-01 00:00:00",
					Color:    "white",
					Type:     model.ScheduleTypeMaster.String(),
					Order:    1,
				},
			},
		}
		i.UpdateBulkSchedule(input)

		assert.Equal(http.StatusNotFound, p.Result.StatusCode)
		assert.Equal("not found", p.Result.Message)
		assert.True(p.Result.HasError)
	})
}

func TestDeleteSchedule(t *testing.T) {
	t.Run("スケジュールを削除する", func(t *testing.T) {
		assert := assert.New(t)

		r := &stubScheduleRepository{}
		p := &stubScheduleOutputPort{}
		i := NewScheduleInteractor(r, p)

		input := port.DeleteScheduleInputData{ScheduleID: "test-id"}
		i.DeleteSchedule(input)

		assert.Equal(http.StatusNoContent, p.Result.StatusCode)
		assert.Equal("Success", p.Result.Message)
		assert.False(p.Result.HasError)
	})
}
