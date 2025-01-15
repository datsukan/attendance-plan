package repository

import (
	"testing"
	"time"

	"github.com/datsukan/attendance-plan/backend/app/model"
	"github.com/datsukan/attendance-plan/backend/infrastructure"
	"github.com/guregu/dynamo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testSessionSetup(t *testing.T) (*dynamo.DB, *dynamo.Table, error) {
	t.Helper()

	require := require.New(t)

	db := infrastructure.NewDB()
	require.NotNil(db)

	table := db.Table(sessionTableName)

	var session []model.Session
	err := table.Scan().All(&session)
	require.NoError(err)

	for _, s := range session {
		err := table.Delete("ID", s.ID).Run()
		require.NoError(err)
	}

	return db, &table, nil
}

func TestSession_ReadByUserID(t *testing.T) {
	db, table, err := testSessionSetup(t)
	require.NoError(t, err)

	session := model.Session{
		ID:        "test-id",
		UserID:    "test-user-id",
		ExpiresAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	err = table.Put(session).Run()
	require.NoError(t, err)

	tests := []struct {
		name         string
		userID       string
		wantHasError bool
	}{
		{name: "存在しないユーザー ID", userID: "none-user-id", wantHasError: true},
		{name: "存在するユーザー ID", userID: "test-user-id", wantHasError: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			repo := NewSessionRepository(*db)

			got, err := repo.ReadByUserID(tt.userID)
			if tt.wantHasError {
				assert.Error(err)
				return
			}

			assert.NoError(err)
			if !assert.NotNil(got) {
				return
			}

			assert.Equal(session.ID, got.ID)
			assert.Equal(session.UserID, got.UserID)
			assert.Equal(session.ExpiresAt, got.ExpiresAt)
		})
	}
}

func TestSession_Read(t *testing.T) {
	db, table, err := testSessionSetup(t)
	require.NoError(t, err)

	session := model.Session{
		ID:        "test-id",
		UserID:    "test-user-id",
		ExpiresAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	err = table.Put(session).Run()
	require.NoError(t, err)

	tests := []struct {
		name         string
		id           string
		wantHasError bool
	}{
		{name: "存在しない ID", id: "none-id", wantHasError: true},
		{name: "存在する ID", id: "test-id", wantHasError: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			repo := NewSessionRepository(*db)

			got, err := repo.Read(tt.id)
			if tt.wantHasError {
				assert.Error(err)
				return
			}

			assert.NoError(err)
			if !assert.NotNil(got) {
				return
			}

			assert.Equal(session.ID, got.ID)
			assert.Equal(session.UserID, got.UserID)
			assert.Equal(session.ExpiresAt, got.ExpiresAt)
		})
	}
}

func TestSession_Create(t *testing.T) {
	t.Run("正常に登録できること", func(t *testing.T) {
		require := require.New(t)
		assert := assert.New(t)

		db, table, err := testSessionSetup(t)
		require.NoError(err)
		require.NotNil(db)
		require.NotNil(table)

		repo := NewSessionRepository(*db)

		session := &model.Session{
			ID:        "test-id",
			UserID:    "test-user-id",
			ExpiresAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		}

		err = repo.Create(session)
		assert.NoError(err)

		var s *model.Session
		err = table.Get("ID", session.ID).One(&s)
		assert.NoError(err)

		if !assert.NotNil(s) {
			return
		}

		assert.Equal(session.ID, s.ID)
		assert.Equal(session.UserID, s.UserID)
		assert.Equal(session.ExpiresAt, s.ExpiresAt)
	})
}

func TestSession_Update(t *testing.T) {
	t.Run("正常に更新できること", func(t *testing.T) {
		require := require.New(t)
		assert := assert.New(t)

		db, table, err := testSessionSetup(t)
		require.NoError(err)
		require.NotNil(db)
		require.NotNil(table)

		repo := NewSessionRepository(*db)

		session := &model.Session{
			ID:        "test-id",
			UserID:    "test-user-id",
			ExpiresAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		}
		err = table.Put(session).Run()
		require.NoError(err)

		session.ExpiresAt = time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
		err = repo.Update(session)
		assert.NoError(err)

		var s *model.Session
		err = table.Get("ID", session.ID).One(&s)
		assert.NoError(err)

		if !assert.NotNil(s) {
			return
		}

		assert.Equal(session.ID, s.ID)
		assert.Equal(session.UserID, s.UserID)
		assert.Equal(session.ExpiresAt, s.ExpiresAt)
	})
}

func TestSession_Delete(t *testing.T) {
	t.Run("正常に削除できること", func(t *testing.T) {
		require := require.New(t)
		assert := assert.New(t)

		db, table, err := testSessionSetup(t)
		require.NoError(err)
		require.NotNil(db)
		require.NotNil(table)

		repo := NewSessionRepository(*db)

		session := &model.Session{
			ID:        "test-id",
			UserID:    "test-user-id",
			ExpiresAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		}
		err = table.Put(session).Run()
		require.NoError(err)

		err = repo.Delete(session.ID)
		assert.NoError(err)

		var s *model.Session
		err = table.Get("ID", session.ID).One(&s)
		assert.Error(err)
		assert.Nil(s)
	})
}
