package usecase

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/datsukan/attendance-plan/backend/app/port"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSignIn(t *testing.T) {
	t.Run("サインインする", func(t *testing.T) {
		require := require.New(t)
		assert := assert.New(t)

		l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		r := &stubUserRepository{}
		sr := &stubSessionRepository{}
		p := &stubUserOutputPort{}
		i := NewUserInteractor(l, r, sr, nil, p)

		input := port.SignInInputData{
			Email:    "test-email@example.com",
			Password: "test-password",
		}
		i.SignIn(input)

		output, ok := p.Output.(*port.SignInOutputData)
		require.True(ok)

		if !assert.NotNil(output) {
			return
		}

		assert.Equal("test-id", output.ID)
		assert.Equal("test-email@example.com", output.Email)
		assert.Equal("test name", output.Name)

		date := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
		assert.Equal(date.Format(time.DateTime), output.CreatedAt)
		assert.Equal(date.Format(time.DateTime), output.UpdatedAt)

		assert.Equal("test-token", output.SessionToken)

		assert.Equal(http.StatusOK, p.Result.StatusCode)
		assert.Empty(p.Result.ErrorMessage)
		assert.False(p.Result.HasError)
	})

	t.Run("ユーザーが存在しない場合エラーを返す", func(t *testing.T) {
		require := require.New(t)
		assert := assert.New(t)

		l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		r := &stubUserRepository{}
		sr := &stubSessionRepository{}
		p := &stubUserOutputPort{}
		i := NewUserInteractor(l, r, sr, nil, p)

		input := port.SignInInputData{
			Email:    "test-not-found-email@example.com",
			Password: "test-not-found-password",
		}
		i.SignIn(input)

		output, ok := p.Output.(*port.SignInOutputData)
		require.True(ok)

		assert.Nil(output)
		assert.Equal(http.StatusUnauthorized, p.Result.StatusCode)
		assert.Equal(MsgEmailOrPasswordInvalid, p.Result.ErrorMessage)
		assert.True(p.Result.HasError)
	})
}

func TestSignUp(t *testing.T) {
	t.Run("サインアップする", func(t *testing.T) {
		assert := assert.New(t)

		l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		ur := &stubUserRepository{}
		sr := &stubSessionRepository{}
		mr := &stubEmailRepository{}
		p := &stubUserOutputPort{}
		i := NewUserInteractor(l, ur, sr, mr, p)

		ctx := context.Background()
		input := port.SignUpInputData{
			Email: "test-email@example.com",
		}
		i.SignUp(ctx, input)

		assert.Equal(http.StatusOK, p.Result.StatusCode)
		assert.Empty(p.Result.ErrorMessage)
		assert.False(p.Result.HasError)
	})
}

func TestPasswordReset(t *testing.T) {
	t.Run("パスワードリセットする", func(t *testing.T) {
		assert := assert.New(t)

		l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		ur := &stubUserRepository{}
		sr := &stubSessionRepository{}
		mr := &stubEmailRepository{}
		p := &stubUserOutputPort{}
		i := NewUserInteractor(l, ur, sr, mr, p)

		ctx := context.Background()
		input := port.PasswordResetInputData{
			Email: "test-email@example.com",
		}
		i.PasswordReset(ctx, input)

		assert.Equal(http.StatusOK, p.Result.StatusCode)
		assert.Empty(p.Result.ErrorMessage)
		assert.False(p.Result.HasError)
	})
}

func TestPasswordSet(t *testing.T) {
	t.Run("パスワードリセットする", func(t *testing.T) {
		assert := assert.New(t)

		l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		ur := &stubUserRepository{}
		sr := &stubSessionRepository{}
		mr := &stubEmailRepository{}
		p := &stubUserOutputPort{}
		i := NewUserInteractor(l, ur, sr, mr, p)

		input := port.PasswordSetInputData{
			Token:    "test-token",
			Password: "test-password",
		}
		i.PasswordSet(input)

		assert.Equal(http.StatusOK, p.Result.StatusCode)
		assert.Empty(p.Result.ErrorMessage)
		assert.False(p.Result.HasError)
	})
}

func TestGetUser(t *testing.T) {
	t.Run("ユーザー情報を取得する", func(t *testing.T) {
		assert := assert.New(t)

		l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		ur := &stubUserRepository{}
		sr := &stubSessionRepository{}
		p := &stubUserOutputPort{}
		i := NewUserInteractor(l, ur, sr, nil, p)

		input := port.GetUserInputData{
			UserID: "test-id",
		}
		i.GetUser(input)

		output, ok := p.Output.(*port.GetUserOutputData)
		assert.True(ok)

		if !assert.NotNil(output) {
			return
		}

		assert.Equal("test-id", output.ID)
		assert.Equal("test-email@example.com", output.Email)
		assert.Equal("test name", output.Name)

		date := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
		assert.Equal(date.Format(time.DateTime), output.CreatedAt)
		assert.Equal(date.Format(time.DateTime), output.UpdatedAt)

		assert.Equal(http.StatusOK, p.Result.StatusCode)
		assert.Empty(p.Result.ErrorMessage)
		assert.False(p.Result.HasError)
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("ユーザー情報を更新する", func(t *testing.T) {
		assert := assert.New(t)

		l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		ur := &stubUserRepository{}
		sr := &stubSessionRepository{}
		p := &stubUserOutputPort{}
		i := NewUserInteractor(l, ur, sr, nil, p)

		input := port.UpdateUserInputData{
			UserID: "test-id",
			Name:   "test-name",
		}
		i.UpdateUser(input)

		output, ok := p.Output.(*port.UpdateUserOutputData)
		assert.True(ok)

		if !assert.NotNil(output) {
			return
		}

		assert.Equal("test-id", output.ID)
		assert.Equal("test-email@example.com", output.Email)
		assert.Equal("test-name", output.Name)

		date := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
		assert.Equal(date.Format(time.DateTime), output.CreatedAt)
		assert.True(output.UpdatedAt >= time.Now().Format(time.DateTime))

		assert.Equal(http.StatusOK, p.Result.StatusCode)
		assert.Empty(p.Result.ErrorMessage)
		assert.False(p.Result.HasError)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("ユーザーを削除する", func(t *testing.T) {
		assert := assert.New(t)

		l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		ur := &stubUserRepository{}
		sr := &stubSessionRepository{}
		p := &stubUserOutputPort{}
		i := NewUserInteractor(l, ur, sr, nil, p)

		input := port.DeleteUserInputData{
			UserID: "test-id",
		}
		i.DeleteUser(input)

		assert.Equal(http.StatusOK, p.Result.StatusCode)
		assert.Empty(p.Result.ErrorMessage)
		assert.False(p.Result.HasError)
	})
}

func TestResetEmail(t *testing.T) {
	t.Run("メールアドレスをリセットする", func(t *testing.T) {
		assert := assert.New(t)

		l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		ur := &stubUserRepository{}
		sr := &stubSessionRepository{}
		mr := &stubEmailRepository{}
		p := &stubUserOutputPort{}
		i := NewUserInteractor(l, ur, sr, mr, p)

		input := port.ResetEmailInputData{
			UserID: "test-id",
			Email:  "test-reset-email@example.com",
		}
		i.ResetEmail(input)

		assert.Equal(http.StatusOK, p.Result.StatusCode)
		assert.Empty(p.Result.ErrorMessage)
		assert.False(p.Result.HasError)
	})
}

func TestSetEmail(t *testing.T) {
	t.Run("メールアドレスを設定する", func(t *testing.T) {
		assert := assert.New(t)

		l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		ur := &stubUserRepository{}
		sr := &stubSessionRepository{}
		mr := &stubEmailRepository{}
		p := &stubUserOutputPort{}
		i := NewUserInteractor(l, ur, sr, mr, p)

		input := port.SetEmailInputData{
			UserIDToken: "test-id-token",
			EmailToken:  "test-email-token",
		}
		i.SetEmail(input)

		assert.Equal(http.StatusOK, p.Result.StatusCode)
		assert.Empty(p.Result.ErrorMessage)
		assert.False(p.Result.HasError)
	})
}
