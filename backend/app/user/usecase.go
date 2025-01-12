package user

import (
	"errors"
	"net/http"
	"time"

	"github.com/datsukan/attendance-plan/backend/component"
	"golang.org/x/crypto/bcrypt"
)

// BaseUserData はユーザーの基本データを表す構造体です。
type BaseUserData struct {
	ID        string
	Email     string
	Name      string
	CreatedAt string
	UpdatedAt string
}

// SignInInputData はサインインの入力データを表す構造体です。
type SignInInputData struct {
	Email    string
	Password string
}

// SignInOutputData はサインインの出力データを表す構造体です。
type SignInOutputData BaseUserData

// SignUpInputData はサインアップの入力データを表す構造体です。
type SignUpInputData struct {
	Email    string
	Password string
	Name     string
}

// SignUpOutputData はサインアップの出力データを表す構造体です。
type SignUpOutputData BaseUserData

// UserInputPort はユーザーのユースケースを表すインターフェースです。
type UserInputPort interface {
	SignIn(input SignInInputData)
	SignUp(input SignUpInputData)
}

// UserOutputPort はユーザーのユースケースの外部出力を表すインターフェースです。
type UserOutputPort interface {
	GetResponse() (int, string)
	SetResponseSignIn(output *SignInOutputData, result component.ResponseResult)
	SetResponseSignUp(output *SignUpOutputData, result component.ResponseResult)
}

// UserInteractor はユーザーのユースケースの実装を表す構造体です。
type UserInteractor struct {
	UserRepository UserRepository
	OutputPort     UserOutputPort
}

// NewUserInteractor は UserInteractor を生成します。
func NewUserInteractor(userRepository UserRepository, outputPort UserOutputPort) UserInputPort {
	return &UserInteractor{
		UserRepository: userRepository,
		OutputPort:     outputPort,
	}
}

// SignUp はサインイン処理を行います。
func (i *UserInteractor) SignIn(input SignInInputData) {
	user, err := i.UserRepository.ReadByEmail(input.Email)
	if err != nil {
		if errors.Is(err, component.NewNotFoundError()) {
			r := component.NewResponseResult(http.StatusUnauthorized, true, "Invalid email or password")
			i.OutputPort.SetResponseSignIn(nil, r)
			return
		}
		r := component.NewResponseResult(http.StatusInternalServerError, true, err.Error())
		i.OutputPort.SetResponseSignIn(nil, r)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		r := component.NewResponseResult(http.StatusUnauthorized, true, "Invalid email or password")
		i.OutputPort.SetResponseSignIn(nil, r)
		return
	}

	// TODO セッションの保存とクッキーの設定

	o := &SignInOutputData{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.Format(time.DateTime),
		UpdatedAt: user.UpdatedAt.Format(time.DateTime),
	}
	r := component.NewResponseResult(http.StatusOK, false, "Success")
	i.OutputPort.SetResponseSignIn(o, r)
}

// SignUp はサインアップ処理を行います。
func (i *UserInteractor) SignUp(input SignUpInputData) {
	exists, err := i.UserRepository.ExistsByEmail(input.Email)
	if err != nil {
		r := component.NewResponseResult(http.StatusInternalServerError, true, err.Error())
		i.OutputPort.SetResponseSignUp(nil, r)
		return
	}
	if exists {
		r := component.NewResponseResult(http.StatusBadRequest, true, "The email is already in use")
		i.OutputPort.SetResponseSignUp(nil, r)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		r := component.NewResponseResult(http.StatusInternalServerError, true, err.Error())
		i.OutputPort.SetResponseSignUp(nil, r)
		return
	}

	user := &User{
		ID:        component.NewID(),
		Email:     input.Email,
		Password:  string(hashedPassword),
		Name:      input.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := i.UserRepository.Create(user); err != nil {
		r := component.NewResponseResult(http.StatusInternalServerError, true, err.Error())
		i.OutputPort.SetResponseSignUp(nil, r)
		return
	}

	// TODO セッションの保存とクッキーの設定

	o := &SignUpOutputData{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.Format(time.DateTime),
		UpdatedAt: user.UpdatedAt.Format(time.DateTime),
	}
	r := component.NewResponseResult(http.StatusOK, false, "Success")
	i.OutputPort.SetResponseSignUp(o, r)
}
