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

func testUserSetup(t *testing.T) (*dynamo.DB, *dynamo.Table, error) {
	t.Helper()

	require := require.New(t)

	db := infrastructure.NewDB()
	require.NotNil(db)

	table := db.Table(userTableName)

	var users []model.User
	err := table.Scan().All(&users)
	require.NoError(err)

	for _, s := range users {
		err := table.Delete("ID", s.ID).Run()
		require.NoError(err)
	}

	return db, &table, nil
}

func TestUser_ReadByEmail(t *testing.T) {
	db, table, err := testUserSetup(t)
	require.NoError(t, err)

	user := model.User{
		ID:    "test-id",
		Email: "test@example.com",
		Name:  "test name",
	}
	err = table.Put(user).Run()
	require.NoError(t, err)

	tests := []struct {
		name         string
		email        string
		wantHasError bool
	}{
		{name: "存在しないメールアドレス", email: "none@example.com", wantHasError: true},
		{name: "存在するメールアドレス", email: "test@example.com", wantHasError: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			repo := NewUserRepository(*db)

			got, err := repo.ReadByEmail(tt.email)
			if tt.wantHasError {
				assert.Error(err)
				return
			}

			assert.NoError(err)
			if !assert.NotNil(got) {
				return
			}

			assert.Equal(user.ID, got.ID)
			assert.Equal(user.Email, got.Email)
			assert.Equal(user.Name, got.Name)

		})
	}
}

func TestUser_Read(t *testing.T) {
	db, table, err := testUserSetup(t)
	require.NoError(t, err)

	user := model.User{
		ID:    "test-id",
		Email: "test@example.com",
		Name:  "test name",
	}
	err = table.Put(user).Run()
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
			repo := NewUserRepository(*db)

			got, err := repo.Read(tt.id)
			if tt.wantHasError {
				assert.Error(err)
				return
			}

			assert.NoError(err)
			if !assert.NotNil(got) {
				return
			}

			assert.Equal(user.ID, got.ID)
			assert.Equal(user.Email, got.Email)
			assert.Equal(user.Name, got.Name)
		})
	}
}

func TestUser_Create(t *testing.T) {
	t.Run("正常に登録できること", func(t *testing.T) {
		require := require.New(t)
		assert := assert.New(t)

		db, table, err := testUserSetup(t)
		require.NoError(err)
		require.NotNil(db)
		require.NotNil(table)

		repo := NewUserRepository(*db)

		user := &model.User{
			ID:        "test-id",
			Email:     "test@example.com",
			Password:  "test-password",
			Name:      "test name",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err = repo.Create(user)
		assert.NoError(err)

		var u *model.User
		err = table.Get("ID", user.ID).One(&u)
		assert.NoError(err)

		if !assert.NotNil(u) {
			return
		}

		assert.Equal(user.ID, u.ID)
		assert.Equal(user.Email, u.Email)
		assert.Equal(user.Password, u.Password)
		assert.Equal(user.Name, u.Name)
		assert.Equal(user.CreatedAt.Format(time.DateTime), u.CreatedAt.Format(time.DateTime))
		assert.Equal(user.UpdatedAt.Format(time.DateTime), u.UpdatedAt.Format(time.DateTime))
	})
}

func TestUser_Update(t *testing.T) {
	t.Run("正常に更新できること", func(t *testing.T) {
		require := require.New(t)
		assert := assert.New(t)

		db, table, err := testUserSetup(t)
		require.NoError(err)
		require.NotNil(db)
		require.NotNil(table)

		repo := NewUserRepository(*db)

		user := &model.User{
			ID:        "test-id",
			Email:     "test@example.com",
			Password:  "test-password",
			Name:      "test name",
			CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		}

		err = table.Put(user).Run()
		require.NoError(err)

		user.Email = "updated-test@example.com"
		user.Password = "updated-test-password"
		user.Name = "updated name"
		user.UpdatedAt = time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC)
		err = repo.Update(user)
		assert.NoError(err)

		var u *model.User
		err = table.Get("ID", user.ID).One(&u)
		assert.NoError(err)

		if !assert.NotNil(u) {
			return
		}

		assert.Equal(user.ID, u.ID)
		assert.Equal(user.Email, u.Email)
		assert.Equal(user.Password, u.Password)
		assert.Equal(user.Name, u.Name)
		assert.Equal(user.CreatedAt.Format(time.DateTime), u.CreatedAt.Format(time.DateTime))
		assert.Equal(user.UpdatedAt.Format(time.DateTime), u.UpdatedAt.Format(time.DateTime))
	})
}

func TestUser_Delete(t *testing.T) {
	t.Run("正常に削除できること", func(t *testing.T) {
		require := require.New(t)
		assert := assert.New(t)

		db, table, err := testUserSetup(t)
		require.NoError(err)
		require.NotNil(db)
		require.NotNil(table)

		repo := NewUserRepository(*db)

		user := &model.User{
			ID:        "test-id",
			Email:     "test@example.com",
			Password:  "test-password",
			Name:      "test name",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err = table.Put(user).Run()
		require.NoError(err)

		err = repo.Delete(user.ID)
		assert.NoError(err)

		var u *model.User
		err = table.Get("ID", user.ID).One(&u)
		assert.Error(err)
		assert.Nil(u)
	})
}

func TestUser_Exists(t *testing.T) {
	now := time.Now()

	var users []model.User
	for i := 0; i < 2; i++ {
		user := model.User{
			ID:        fmt.Sprintf("test-id-%d", i),
			Email:     fmt.Sprintf("test-%d@example.com", i),
			Password:  fmt.Sprintf("test-password-%d", i),
			Name:      fmt.Sprintf("test name %d", i),
			CreatedAt: now,
			UpdatedAt: now,
		}
		users = append(users, user)
	}

	tests := []struct {
		name string
		data []model.User
		want bool
	}{
		{name: "レコード全体が0件の場合", data: []model.User{}, want: false},
		{name: "レコード全体が1件の場合", data: users[:1], want: true},
		{name: "レコード全体が2件の場合", data: users[:2], want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			assert := assert.New(t)

			db, table, err := testUserSetup(t)
			require.NoError(err)
			require.NotNil(db)
			require.NotNil(table)

			for _, user := range tt.data {
				err := table.Put(user).Run()
				require.NoError(err)
			}

			repo := NewUserRepository(*db)

			got, err := repo.Exists("test-id-0")
			assert.NoError(err)

			assert.Equal(tt.want, got)
		})
	}
}

func TestUser_ExistsByEmail(t *testing.T) {
	now := time.Now()

	var users []model.User
	for i := 0; i < 2; i++ {
		user := model.User{
			ID:        fmt.Sprintf("test-id-%d", i),
			Email:     fmt.Sprintf("test-%d@example.com", i),
			Password:  fmt.Sprintf("test-password-%d", i),
			Name:      fmt.Sprintf("test name %d", i),
			CreatedAt: now,
			UpdatedAt: now,
		}
		users = append(users, user)
	}

	tests := []struct {
		name string
		data []model.User
		want bool
	}{
		{name: "レコード全体が0件の場合", data: []model.User{}, want: false},
		{name: "レコード全体が1件の場合", data: users[:1], want: true},
		{name: "レコード全体が2件の場合", data: users[:2], want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			assert := assert.New(t)

			db, table, err := testUserSetup(t)
			require.NoError(err)
			require.NotNil(db)
			require.NotNil(table)

			for _, user := range tt.data {
				err := table.Put(user).Run()
				require.NoError(err)
			}

			repo := NewUserRepository(*db)

			got, err := repo.ExistsByEmail("test-0@example.com")
			assert.NoError(err)

			assert.Equal(tt.want, got)
		})
	}
}
