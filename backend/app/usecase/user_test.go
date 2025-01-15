package usecase

import (
	"net/http"
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

		r := &stubUserRepository{}
		s := &stubSessionRepository{}
		p := &stubUserOutputPort{}
		i := NewUserInteractor(r, s, p)

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

		assert.Equal(http.StatusOK, p.Result.StatusCode)
		assert.Equal("Success", p.Result.Message)
		assert.False(p.Result.HasError)
	})
}

func TestSignUp(t *testing.T) {
	t.Run("サインアップする", func(t *testing.T) {
		require := require.New(t)
		assert := assert.New(t)

		r := &stubUserRepository{}
		s := &stubSessionRepository{}
		p := &stubUserOutputPort{}
		i := NewUserInteractor(r, s, p)

		input := port.SignUpInputData{
			Email:    "test-email@example.com",
			Password: "test-password",
			Name:     "test name",
		}
		i.SignUp(input)

		output, ok := p.Output.(*port.SignUpOutputData)
		require.True(ok)

		if !assert.NotNil(output) {
			return
		}

		assert.NotEmpty(output.ID)
		assert.Equal("test-email@example.com", output.Email)
		assert.Equal("test name", output.Name)
		assert.NotEqual(time.Time{}, output.CreatedAt)
		assert.NotEqual(time.Time{}, output.UpdatedAt)

		assert.Equal(http.StatusOK, p.Result.StatusCode)
		assert.Equal("Success", p.Result.Message)
		assert.False(p.Result.HasError)
	})
}
