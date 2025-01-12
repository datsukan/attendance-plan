package user

import (
	"net/http"
	"testing"
	"time"

	"github.com/datsukan/attendance-plan/backend/component"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

type stubUserRepository struct{}

func (r *stubUserRepository) ReadByEmail(email string) (*User, error) {
	ph, _ := bcrypt.GenerateFromPassword([]byte("test-password"), bcrypt.DefaultCost)
	date := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)

	return &User{
		ID:        "test-id",
		Email:     "test-email@example.com",
		Name:      "test name",
		Password:  string(ph),
		CreatedAt: date,
		UpdatedAt: date,
	}, nil
}

func (r *stubUserRepository) Read(id string) (*User, error) {
	ph, _ := bcrypt.GenerateFromPassword([]byte("test-password"), bcrypt.DefaultCost)
	date := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)

	return &User{
		ID:        "test-id",
		Email:     "test-email@example.com",
		Name:      "test name",
		Password:  string(ph),
		CreatedAt: date,
		UpdatedAt: date,
	}, nil
}

func (r *stubUserRepository) Create(user *User) error {
	return nil
}

func (r *stubUserRepository) Update(user *User) error {
	return nil
}

func (r *stubUserRepository) Delete(id string) error {
	return nil
}

func (r *stubUserRepository) Exists(id string) (bool, error) {
	return true, nil
}

func (r *stubUserRepository) ExistsByEmail(email string) (bool, error) {
	return false, nil
}

type stubNotFoundUserRepository struct{}

func (r *stubNotFoundUserRepository) ReadByEmail(email string) (*User, error) {
	return nil, component.NewNotFoundError()
}

func (r *stubNotFoundUserRepository) Read(id string) (*User, error) {
	return nil, component.NewNotFoundError()
}

func (r *stubNotFoundUserRepository) Create(user *User) error {
	return nil
}

func (r *stubNotFoundUserRepository) Update(user *User) error {
	return nil
}

func (r *stubNotFoundUserRepository) Delete(id string) error {
	return nil
}

func (r *stubNotFoundUserRepository) Exists(id string) (bool, error) {
	return false, nil
}

func (r *stubNotFoundUserRepository) ExistsByEmail(email string) (bool, error) {
	return false, nil
}

type stubUserOutputPort struct {
	Output interface{}
	Result component.ResponseResult
}

func (p *stubUserOutputPort) GetResponse() (int, string) {
	return p.Result.StatusCode, p.Result.Message
}

func (p *stubUserOutputPort) SetResponseSignIn(output *SignInOutputData, result component.ResponseResult) {
	p.Output = output
	p.Result = result
}

func (p *stubUserOutputPort) SetResponseSignUp(output *SignUpOutputData, result component.ResponseResult) {
	p.Output = output
	p.Result = result
}

func TestSignIn(t *testing.T) {
	t.Run("サインインする", func(t *testing.T) {
		require := require.New(t)
		assert := assert.New(t)

		r := &stubUserRepository{}
		p := &stubUserOutputPort{}
		i := NewUserInteractor(r, p)

		input := SignInInputData{
			Email:    "test-email@example.com",
			Password: "test-password",
		}
		i.SignIn(input)

		output, ok := p.Output.(*SignInOutputData)
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
		p := &stubUserOutputPort{}
		i := NewUserInteractor(r, p)

		input := SignUpInputData{
			Email:    "test-email@example.com",
			Password: "test-password",
			Name:     "test name",
		}
		i.SignUp(input)

		output, ok := p.Output.(*SignUpOutputData)
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
