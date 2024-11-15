package schedule

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type stubScheduleRepository struct{}

func (r *stubScheduleRepository) ReadByUserID(userID string) ([]*Schedule, error) {
	date := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	schedules := []*Schedule{
		{ID: "test-id-1", UserID: "test-user-id-1", Name: "test-name-1", StartsAt: date, EndsAt: date, Color: "test-color", Type: ScheduleTypeMaster, CreatedAt: date, UpdatedAt: date},
		{ID: "test-id-2", UserID: "test-user-id-2", Name: "test-name-2", StartsAt: date, EndsAt: date, Color: "test-color", Type: ScheduleTypeCustom, CreatedAt: date, UpdatedAt: date},
	}
	return schedules, nil
}

func (r *stubScheduleRepository) Read(id string) (*Schedule, error) {
	date := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	schedule := &Schedule{
		ID:        "test-id",
		UserID:    "test-user-id",
		Name:      "test-name",
		StartsAt:  date,
		EndsAt:    date,
		Color:     "test-color",
		Type:      ScheduleTypeMaster,
		CreatedAt: date,
		UpdatedAt: date,
	}
	return schedule, nil
}

func (r *stubScheduleRepository) Create(schedule *Schedule) error {
	return nil
}

func (r *stubScheduleRepository) Update(schedule *Schedule) error {
	return nil
}

func (r *stubScheduleRepository) Delete(id string) error {
	return nil
}

func (r *stubScheduleRepository) Exists(id string) (bool, error) {
	return true, nil
}

type stubScheduleOutputPort struct {
	Output interface{}
	Result Result
}

func (p *stubScheduleOutputPort) GetResponse() (int, string) {
	return p.Result.StatusCode, p.Result.Message
}

func (p *stubScheduleOutputPort) ResponseGetScheduleList(output *GetScheduleListOutputData, result Result) {
	p.Output = output
	p.Result = result
}

func (p *stubScheduleOutputPort) ResponseGetSchedule(output *GetScheduleOutputData, result Result) {
	p.Output = output
	p.Result = result
}

func (p *stubScheduleOutputPort) ResponseCreateSchedule(output *CreateScheduleOutputData, result Result) {
	p.Output = output
	p.Result = result
}

func (p *stubScheduleOutputPort) ResponseUpdateSchedule(output *UpdateScheduleOutputData, result Result) {
	p.Output = output
	p.Result = result
}

func (p *stubScheduleOutputPort) ResponseDeleteSchedule(output *DeleteScheduleOutputData, result Result) {
	p.Output = output
	p.Result = result
}

func TestGetScheduleList(t *testing.T) {
	t.Run("スケジュールリストを取得する", func(t *testing.T) {
		require := require.New(t)
		assert := assert.New(t)

		r := &stubScheduleRepository{}
		p := &stubScheduleOutputPort{}
		i := NewScheduleInteractor(r, p)

		input := GetScheduleListInputData{UserID: "test-user-id"}
		i.GetScheduleList(input)

		date := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
		wantSchedules := []*Schedule{
			{ID: "test-id-1", UserID: "test-user-id-1", Name: "test-name-1", StartsAt: date, EndsAt: date, Color: "test-color", Type: ScheduleTypeMaster, CreatedAt: date, UpdatedAt: date},
			{ID: "test-id-2", UserID: "test-user-id-2", Name: "test-name-2", StartsAt: date, EndsAt: date, Color: "test-color", Type: ScheduleTypeCustom, CreatedAt: date, UpdatedAt: date},
		}

		output, ok := p.Output.(*GetScheduleListOutputData)
		require.True(ok)

		if !assert.NotNil(output) {
			return
		}

		assert.Len(output.Schedules, 2)

		ws1 := wantSchedules[0]
		os1 := output.Schedules[0]
		assert.Equal(ws1.ID, os1.ID)
		assert.Equal(ws1.UserID, os1.UserID)
		assert.Equal(ws1.Name, os1.Name)
		assert.Equal(ws1.StartsAt.Format(time.DateTime), os1.StartsAt)
		assert.Equal(ws1.EndsAt.Format(time.DateTime), os1.EndsAt)
		assert.Equal(ws1.Color, os1.Color)
		assert.Equal(ws1.Type.String(), os1.Type)
		assert.Equal(ws1.CreatedAt.Format(time.DateTime), os1.CreatedAt)
		assert.Equal(ws1.UpdatedAt.Format(time.DateTime), os1.UpdatedAt)

		ws2 := wantSchedules[1]
		os2 := output.Schedules[1]
		assert.Equal(ws2.ID, os2.ID)
		assert.Equal(ws2.UserID, os2.UserID)
		assert.Equal(ws2.Name, os2.Name)
		assert.Equal(ws2.StartsAt.Format(time.DateTime), os2.StartsAt)
		assert.Equal(ws2.EndsAt.Format(time.DateTime), os2.EndsAt)
		assert.Equal(ws2.Color, os2.Color)
		assert.Equal(ws2.Type.String(), os2.Type)
		assert.Equal(ws2.CreatedAt.Format(time.DateTime), os2.CreatedAt)
		assert.Equal(ws2.UpdatedAt.Format(time.DateTime), os2.UpdatedAt)

		assert.Equal(http.StatusOK, p.Result.StatusCode)
		assert.Equal("Success", p.Result.Message)
		assert.False(p.Result.HasError)
	})
}

func TestGetSchedule(t *testing.T) {
	t.Run("スケジュールを取得する", func(t *testing.T) {
		require := require.New(t)
		assert := assert.New(t)

		r := &stubScheduleRepository{}
		p := &stubScheduleOutputPort{}
		i := NewScheduleInteractor(r, p)

		input := GetScheduleInputData{ScheduleID: "test-id"}
		i.GetSchedule(input)

		date := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
		wantSchedule := &Schedule{
			ID:        "test-id",
			UserID:    "test-user-id",
			Name:      "test-name",
			StartsAt:  date,
			EndsAt:    date,
			Color:     "test-color",
			Type:      ScheduleTypeMaster,
			CreatedAt: date,
			UpdatedAt: date,
		}

		output, ok := p.Output.(*GetScheduleOutputData)
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
		assert.Equal(wantSchedule.CreatedAt.Format(time.DateTime), os.CreatedAt)
		assert.Equal(wantSchedule.UpdatedAt.Format(time.DateTime), os.UpdatedAt)

		assert.Equal(http.StatusOK, p.Result.StatusCode)
		assert.Equal("Success", p.Result.Message)
		assert.False(p.Result.HasError)
	})
}

func TestCreateSchedule(t *testing.T) {
	t.Run("スケジュールを作成する", func(t *testing.T) {
		assert := assert.New(t)

		r := &stubScheduleRepository{}
		p := &stubScheduleOutputPort{}
		i := NewScheduleInteractor(r, p)

		input := CreateScheduleInputData{
			Schedule: CreateScheduleData{
				Name:     "test-name",
				StartsAt: "2021-01-01 00:00:00",
				EndsAt:   "2021-01-01 00:00:00",
				Color:    "white",
				Type:     ScheduleTypeMaster.String(),
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

		input := UpdateScheduleInputData{
			Schedule: UpdateScheduleData{
				ID:       "test-id",
				Name:     "test-name",
				StartsAt: "2021-01-01 00:00:00",
				EndsAt:   "2021-01-01 00:00:00",
				Color:    "white",
				Type:     ScheduleTypeMaster.String(),
			},
		}
		i.UpdateSchedule(input)

		assert.Equal(http.StatusNoContent, p.Result.StatusCode)
		assert.Equal("Success", p.Result.Message)
		assert.False(p.Result.HasError)
	})
}

func TestDeleteSchedule(t *testing.T) {
	t.Run("スケジュールを削除する", func(t *testing.T) {
		assert := assert.New(t)

		r := &stubScheduleRepository{}
		p := &stubScheduleOutputPort{}
		i := NewScheduleInteractor(r, p)

		input := DeleteScheduleInputData{ScheduleID: "test-id"}
		i.DeleteSchedule(input)

		assert.Equal(http.StatusNoContent, p.Result.StatusCode)
		assert.Equal("Success", p.Result.Message)
		assert.False(p.Result.HasError)
	})
}
