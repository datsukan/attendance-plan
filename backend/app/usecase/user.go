package usecase

import (
	"errors"
	"net/http"
	"time"

	"github.com/datsukan/attendance-plan/backend/app/component"
	"github.com/datsukan/attendance-plan/backend/app/model"
	"github.com/datsukan/attendance-plan/backend/app/port"
	"github.com/datsukan/attendance-plan/backend/app/repository"
	"golang.org/x/crypto/bcrypt"
)

// UserInteractor はユーザーのユースケースの実装を表す構造体です。
type UserInteractor struct {
	UserRepository    repository.UserRepository
	SessionRepository repository.SessionRepository
	OutputPort        port.UserOutputPort
}

// NewUserInteractor は UserInteractor を生成します。
func NewUserInteractor(userRepository repository.UserRepository, sessionRepository repository.SessionRepository, outputPort port.UserOutputPort) port.UserInputPort {
	return &UserInteractor{
		UserRepository:    userRepository,
		SessionRepository: sessionRepository,
		OutputPort:        outputPort,
	}
}

// SignUp はサインイン処理を行います。
func (i *UserInteractor) SignIn(input port.SignInInputData) {
	user, err := i.UserRepository.ReadByEmail(input.Email)
	if err != nil {
		if errors.Is(err, repository.NewNotFoundError()) {
			r := port.NewResult(http.StatusUnauthorized, true, "Invalid email or password")
			i.OutputPort.SetResponseSignIn(nil, r)
			return
		}
		r := port.NewResult(http.StatusInternalServerError, true, err.Error())
		i.OutputPort.SetResponseSignIn(nil, r)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		r := port.NewResult(http.StatusUnauthorized, true, "Invalid email or password")
		i.OutputPort.SetResponseSignIn(nil, r)
		return
	}

	// TODO セッションの保存とクッキーの設定

	o := &port.SignInOutputData{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.Format(time.DateTime),
		UpdatedAt: user.UpdatedAt.Format(time.DateTime),
	}
	r := port.NewResult(http.StatusOK, false, "Success")
	i.OutputPort.SetResponseSignIn(o, r)
}

// SignUp はサインアップ処理を行います。
func (i *UserInteractor) SignUp(input port.SignUpInputData) {
	exists, err := i.UserRepository.ExistsByEmail(input.Email)
	if err != nil {
		r := port.NewResult(http.StatusInternalServerError, true, err.Error())
		i.OutputPort.SetResponseSignUp(nil, r)
		return
	}
	if exists {
		r := port.NewResult(http.StatusBadRequest, true, "The email is already in use")
		i.OutputPort.SetResponseSignUp(nil, r)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		r := port.NewResult(http.StatusInternalServerError, true, err.Error())
		i.OutputPort.SetResponseSignUp(nil, r)
		return
	}

	user := &model.User{
		ID:        component.NewID(),
		Email:     input.Email,
		Password:  string(hashedPassword),
		Name:      input.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := i.UserRepository.Create(user); err != nil {
		r := port.NewResult(http.StatusInternalServerError, true, err.Error())
		i.OutputPort.SetResponseSignUp(nil, r)
		return
	}

	// TODO セッションの保存とクッキーの設定

	o := &port.SignUpOutputData{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.Format(time.DateTime),
		UpdatedAt: user.UpdatedAt.Format(time.DateTime),
	}
	r := port.NewResult(http.StatusOK, false, "Success")
	i.OutputPort.SetResponseSignUp(o, r)
}
