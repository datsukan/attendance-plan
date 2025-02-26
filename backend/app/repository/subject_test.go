package repository

import (
	"fmt"
	"testing"
	"time"

	"github.com/datsukan/attendance-plan/backend/app/model"
	"github.com/datsukan/attendance-plan/backend/infrastructure"
	"github.com/guregu/dynamo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testSubjectSetup(t *testing.T) (*dynamo.DB, *dynamo.Table, error) {
	t.Helper()

	require := require.New(t)

	db := infrastructure.NewDB()
	require.NotNil(db)

	table := db.Table(subjectTableName)

	var subjects []model.Subject
	err := table.Scan().All(&subjects)
	require.NoError(err)

	for _, s := range subjects {
		err := table.Delete("ID", s.ID).Run()
		require.NoError(err)
	}

	return db, &table, nil
}

func TestSubject_ReadByUserID(t *testing.T) {
	var subjects []model.Subject
	for i := 0; i < 10; i++ {
		s := model.Subject{
			ID:        fmt.Sprintf("test-subject-%d", i),
			UserID:    "test-user-a",
			Name:      fmt.Sprintf("test-name-%d", i),
			Color:     fmt.Sprintf("test-color-%d", i),
			CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		}
		subjects = append(subjects, s)
	}

	tests := []struct {
		name   string
		userID string
		data   []model.Subject
		want   []model.Subject
	}{
		{name: "0件取得", userID: "test-user-a", data: []model.Subject{}, want: []model.Subject{}},
		{name: "1件取得", userID: "test-user-a", data: subjects[:1], want: subjects[:1]},
		{name: "2件取得", userID: "test-user-a", data: subjects[:2], want: subjects[:2]},
		{name: "10件取得", userID: "test-user-a", data: subjects, want: subjects},
		{name: "異なるユーザーID", userID: "test-user-b", data: subjects, want: []model.Subject{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			assert := assert.New(t)

			db, table, err := testSubjectSetup(t)
			require.NoError(err)
			require.NotNil(db)
			require.NotNil(table)

			for _, s := range tt.data {
				err := table.Put(s).Run()
				require.NoError(err)
			}

			repo := NewSubjectRepository(*db)
			require.NotNil(repo)

			subjects, err := repo.ReadByUserID(tt.userID)
			require.NoError(err)
			assert.Equal(tt.want, subjects)
		})
	}
}

func TestSubject_Create(t *testing.T) {
	t.Run("正常に登録できること", func(t *testing.T) {
		require := require.New(t)
		assert := assert.New(t)

		db, table, err := testSubjectSetup(t)
		require.NoError(err)
		require.NotNil(db)
		require.NotNil(table)

		subject := &model.Subject{
			ID:        "test-id",
			UserID:    "test-user-id",
			Name:      "test-name",
			Color:     "test-color",
			CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		}

		repo := NewSubjectRepository(*db)
		require.NotNil(repo)

		err = repo.Create(subject)
		require.NoError(err)

		var s *model.Subject
		err = table.Get("ID", "test-id").One(s)
		require.NoError(err)
		assert.Equal(subject, s)
	})
}

func TestSubject_Delete(t *testing.T) {
	t.Run("正常に削除できること", func(t *testing.T) {
		require := require.New(t)
		assert := assert.New(t)

		db, table, err := testSubjectSetup(t)
		require.NoError(err)
		require.NotNil(db)
		require.NotNil(table)

		subject := &model.Subject{
			ID:        "test-id",
			UserID:    "test-user-id",
			Name:      "test-name",
			Color:     "test-color",
			CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		}

		err = table.Put(subject).Run()
		require.NoError(err)

		repo := NewSubjectRepository(*db)
		require.NotNil(repo)

		err = repo.Delete("test-id")
		require.NoError(err)

		var s *model.Subject
		err = table.Get("ID", "test-id").One(s)
		require.Error(err)

		assert.Nil(s)
	})
}
