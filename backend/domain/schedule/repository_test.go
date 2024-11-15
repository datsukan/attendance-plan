package schedule

import (
	"fmt"
	"testing"
	"time"

	"github.com/datsukan/attendance-plan/backend/infrastructure"
	"github.com/guregu/dynamo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testSetup(t *testing.T) (*dynamo.DB, *dynamo.Table, error) {
	t.Helper()

	type TestSchedule struct {
		ID        string       `dynamo:"ID,hash"`
		UserID    string       `dynamo:"UserID" index:"UserID-index,hash"`
		Name      string       `dynamo:"Name"`
		StartsAt  time.Time    `dynamo:"StartsAt" index:"UserID-index,range"`
		EndsAt    time.Time    `dynamo:"EndsAt"`
		Color     string       `dynamo:"Color"`
		Type      ScheduleType `dynamo:"Type"`
		CreatedAt time.Time    `dynamo:"CreatedAt"`
		UpdatedAt time.Time    `dynamo:"UpdatedAt"`
	}

	db := infrastructure.NewDB()
	require.NotNil(t, db)

	db.Table("schedule").DeleteTable().Run()
	err := db.CreateTable("schedule", TestSchedule{}).Run()
	if err != nil {
		return nil, nil, err
	}

	table := db.Table("schedule")

	return db, &table, nil
}

func TestReadByUserID(t *testing.T) {
	now := time.Now()

	var schedules []Schedule
	for i := 0; i < 10; i++ {
		s := Schedule{
			ID:        fmt.Sprintf("test-id-%d", i),
			UserID:    "test-user-id",
			Name:      "test name",
			StartsAt:  time.Date(2021, 1, 1, i, 0, 0, 0, time.UTC),
			EndsAt:    now,
			Color:     "test color",
			Type:      "master",
			CreatedAt: now,
			UpdatedAt: now,
		}
		schedules = append(schedules, s)
	}

	tests := []struct {
		name string
		data []Schedule
		want []Schedule
	}{
		{name: "0件取得", data: []Schedule{}, want: []Schedule{}},
		{name: "1件取得", data: schedules[:1], want: schedules[:1]},
		{name: "2件取得", data: schedules[:2], want: schedules[:2]},
		{name: "10件取得", data: schedules[:10], want: schedules[:10]},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			assert := assert.New(t)

			db, table, err := testSetup(t)
			require.NoError(err)
			require.NotNil(db)
			require.NotNil(table)

			repo := NewScheduleRepository(*db)

			for _, s := range tt.data {
				err := table.Put(s).Run()
				require.NoError(err)
			}

			schedules, err := repo.ReadByUserID("test-user-id")
			assert.NoError(err)

			if !assert.Len(schedules, len(tt.want)) {
				return
			}

			for i, want := range tt.want {
				assert.Equal(want.ID, schedules[i].ID)
				assert.Equal(want.UserID, schedules[i].UserID)
				assert.Equal(want.Name, schedules[i].Name)
				assert.Equal(want.StartsAt.Format(time.DateTime), schedules[i].StartsAt.Format(time.DateTime))
				assert.Equal(want.EndsAt.Format(time.DateTime), schedules[i].EndsAt.Format(time.DateTime))
				assert.Equal(want.Color, schedules[i].Color)
				assert.Equal(want.Type, schedules[i].Type)
				assert.Equal(want.CreatedAt.Format(time.DateTime), schedules[i].CreatedAt.Format(time.DateTime))
				assert.Equal(want.UpdatedAt.Format(time.DateTime), schedules[i].UpdatedAt.Format(time.DateTime))
			}
		})
	}
}

func TestRead(t *testing.T) {
	now := time.Now()

	schedule := &Schedule{
		ID:        "test-id",
		UserID:    "test-user-id",
		Name:      "test name",
		StartsAt:  now,
		EndsAt:    now,
		Color:     "test color",
		Type:      "master",
		CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: now,
	}

	tests := []struct {
		name         string
		id           string
		userID       string
		data         *Schedule
		wantData     *Schedule
		wantHasError bool
	}{
		{
			name:         "0件取得",
			id:           "test-id",
			userID:       "test-user-id",
			data:         nil,
			wantData:     nil,
			wantHasError: true,
		},
		{
			name:     "1件取得",
			id:       "test-id",
			userID:   "test-user-id",
			data:     schedule,
			wantData: schedule,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			assert := assert.New(t)

			db, table, err := testSetup(t)
			require.NoError(err)
			require.NotNil(db)
			require.NotNil(table)

			repo := NewScheduleRepository(*db)

			if tt.data != nil {
				err := table.Put(tt.wantData).Run()
				require.NoError(err)
			}

			schedule, err := repo.Read(tt.id)

			if tt.wantHasError {
				assert.Nil(schedule)
				assert.Error(err)
				return
			}

			assert.NoError(err)

			if schedule == nil {
				return
			}

			assert.Equal(tt.wantData.ID, schedule.ID)
			assert.Equal(tt.wantData.UserID, schedule.UserID)
			assert.Equal(tt.wantData.Name, schedule.Name)
			assert.Equal(tt.wantData.StartsAt.Format(time.DateTime), schedule.StartsAt.Format(time.DateTime))
			assert.Equal(tt.wantData.EndsAt.Format(time.DateTime), schedule.EndsAt.Format(time.DateTime))
			assert.Equal(tt.wantData.Color, schedule.Color)
			assert.Equal(tt.wantData.Type, schedule.Type)
			assert.Equal(tt.wantData.CreatedAt.Format(time.DateTime), schedule.CreatedAt.Format(time.DateTime))
			assert.Equal(tt.wantData.UpdatedAt.Format(time.DateTime), schedule.UpdatedAt.Format(time.DateTime))
		})
	}
}

func TestCreate(t *testing.T) {
	t.Run("正常に登録できること", func(t *testing.T) {
		require := require.New(t)
		assert := assert.New(t)

		db, table, err := testSetup(t)
		require.NoError(err)
		require.NotNil(db)
		require.NotNil(table)

		repo := NewScheduleRepository(*db)

		schedule := &Schedule{
			ID:        "test-id",
			UserID:    "test-user-id",
			Name:      "test name",
			StartsAt:  time.Now(),
			EndsAt:    time.Now(),
			Color:     "test color",
			Type:      "master",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err = repo.Create(schedule)
		assert.NoError(err)

		var s *Schedule
		err = table.Get("ID", "test-id").One(&s)
		assert.NoError(err)

		assert.NotNil(s)
		if s == nil {
			return
		}

		assert.Equal(schedule.ID, s.ID)
		assert.Equal(schedule.UserID, s.UserID)
		assert.Equal(schedule.Name, s.Name)
		assert.Equal(schedule.StartsAt.Format(time.DateTime), s.StartsAt.Format(time.DateTime))
		assert.Equal(schedule.EndsAt.Format(time.DateTime), s.EndsAt.Format(time.DateTime))
		assert.Equal(schedule.Color, s.Color)
		assert.Equal(schedule.Type, s.Type)
		assert.Equal(schedule.CreatedAt.Format(time.DateTime), s.CreatedAt.Format(time.DateTime))
		assert.Equal(schedule.UpdatedAt.Format(time.DateTime), s.UpdatedAt.Format(time.DateTime))
	})
}

func TestUpdate(t *testing.T) {
	t.Run("正常に更新できること", func(t *testing.T) {
		require := require.New(t)
		assert := assert.New(t)

		db, table, err := testSetup(t)
		require.NoError(err)
		require.NotNil(db)
		require.NotNil(table)

		repo := NewScheduleRepository(*db)

		schedule := &Schedule{
			ID:        "test-id",
			UserID:    "test-user-id",
			Name:      "test name",
			StartsAt:  time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			EndsAt:    time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			Color:     "test color",
			Type:      "master",
			CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		}

		err = table.Put(schedule).Run()
		require.NoError(err)

		schedule.Name = "updated name"
		schedule.StartsAt = time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC)
		schedule.EndsAt = time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC)
		schedule.Color = "updated color"
		schedule.Type = "custom"
		schedule.UpdatedAt = time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC)

		err = repo.Update(schedule)
		assert.NoError(err)

		var s *Schedule
		err = table.Get("ID", "test-id").One(&s)
		assert.NoError(err)

		assert.NotNil(s)
		if s == nil {
			return
		}

		assert.Equal(schedule.ID, s.ID)
		assert.Equal(schedule.UserID, s.UserID)
		assert.Equal(schedule.Name, s.Name)
		assert.Equal(schedule.StartsAt.Format(time.DateTime), s.StartsAt.Format(time.DateTime))
		assert.Equal(schedule.EndsAt.Format(time.DateTime), s.EndsAt.Format(time.DateTime))
		assert.Equal(schedule.Color, s.Color)
		assert.Equal(schedule.Type, s.Type)
		assert.Equal(schedule.CreatedAt.Format(time.DateTime), s.CreatedAt.Format(time.DateTime))
		assert.Equal(schedule.UpdatedAt.Format(time.DateTime), s.UpdatedAt.Format(time.DateTime))
	})
}

func TestDelete(t *testing.T) {
	t.Run("正常に削除できること", func(t *testing.T) {
		require := require.New(t)
		assert := assert.New(t)

		db, table, err := testSetup(t)
		require.NoError(err)
		require.NotNil(db)
		require.NotNil(table)

		repo := NewScheduleRepository(*db)

		schedule := &Schedule{
			ID:        "test-id",
			UserID:    "test-user-id",
			Name:      "test name",
			StartsAt:  time.Now(),
			EndsAt:    time.Now(),
			Color:     "test color",
			Type:      "master",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err = table.Put(schedule).Run()
		require.NoError(err)

		err = repo.Delete("test-id")
		if !assert.NoError(err) {
			return
		}

		var s *Schedule
		err = table.Get("ID", "test-id").One(s)
		assert.Error(err)

		assert.Nil(s)
	})
}

func TestExists(t *testing.T) {
	now := time.Now()

	var schedules []Schedule
	for i := 0; i < 10; i++ {
		s := Schedule{
			ID:        fmt.Sprintf("test-id-%d", i),
			UserID:    "test-user-id",
			Name:      "test name",
			StartsAt:  time.Date(2021, 1, 1, i, 0, 0, 0, time.UTC),
			EndsAt:    now,
			Color:     "test color",
			Type:      "master",
			CreatedAt: now,
			UpdatedAt: now,
		}
		schedules = append(schedules, s)
	}

	tests := []struct {
		name string
		data []Schedule
		want bool
	}{
		{name: "レコード全体が0件の場合false", data: []Schedule{}, want: false},
		{name: "レコード全体が1件の場合true", data: schedules[:1], want: true},
		{name: "レコード全体が2件の場合true", data: schedules[:2], want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			assert := assert.New(t)

			db, table, err := testSetup(t)
			require.NoError(err)
			require.NotNil(db)
			require.NotNil(table)

			repo := NewScheduleRepository(*db)

			for _, s := range tt.data {
				err := table.Put(s).Run()
				require.NoError(err)
			}

			got, err := repo.Exists("test-id-0")
			assert.NoError(err)

			assert.Equal(tt.want, got)
		})
	}
}
