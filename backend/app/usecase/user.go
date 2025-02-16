package usecase

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/datsukan/attendance-plan/backend/app/component/id"
	"github.com/datsukan/attendance-plan/backend/app/model"
	"github.com/datsukan/attendance-plan/backend/app/port"
	"github.com/datsukan/attendance-plan/backend/app/repository"
	"github.com/datsukan/attendance-plan/backend/infrastructure"
	"golang.org/x/crypto/bcrypt"
)

// UserInteractor はユーザーのユースケースの実装を表す構造体です。
type UserInteractor struct {
	Logger            *slog.Logger
	UserRepository    repository.UserRepository
	SessionRepository repository.SessionRepository
	MailRepository    repository.EmailRepository
	OutputPort        port.UserOutputPort
}

// NewUserInteractor は UserInteractor を生成します。
func NewUserInteractor(logger *slog.Logger, userRepository repository.UserRepository, sessionRepository repository.SessionRepository, mailRepository repository.EmailRepository, outputPort port.UserOutputPort) port.UserInputPort {
	return &UserInteractor{
		Logger:            logger,
		UserRepository:    userRepository,
		SessionRepository: sessionRepository,
		MailRepository:    mailRepository,
		OutputPort:        outputPort,
	}
}

// SignUp はサインイン処理を行います。
func (i *UserInteractor) SignIn(input port.SignInInputData) {
	i.Logger.With("email", input.Email)

	user, err := i.UserRepository.ReadByEmail(input.Email, true)
	if err != nil {
		if errors.Is(err, repository.NewNotFoundError()) {
			i.Logger.Warn("user not found")
			r := port.NewResult(http.StatusUnauthorized, true, "Invalid email or password")
			i.OutputPort.SetResponseSignIn(nil, r)
			return
		}
		i.Logger.Error(err.Error())
		r := port.NewResult(http.StatusInternalServerError, true, err.Error())
		i.OutputPort.SetResponseSignIn(nil, r)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		i.Logger.Warn("password is invalid")
		r := port.NewResult(http.StatusUnauthorized, true, "Invalid email or password")
		i.OutputPort.SetResponseSignIn(nil, r)
		return
	}

	sessionToken, err := i.SessionRepository.GenerateToken(user.ID)
	if err != nil {
		i.Logger.Error(err.Error())
		r := port.NewResult(http.StatusInternalServerError, true, "Failed to start session")
		i.OutputPort.SetResponseSignIn(nil, r)
		return
	}

	o := &port.SignInOutputData{
		BaseUserData: port.BaseUserData{
			ID:        user.ID,
			Email:     user.Email,
			Name:      user.Name,
			CreatedAt: user.CreatedAt.Format(time.DateTime),
			UpdatedAt: user.UpdatedAt.Format(time.DateTime),
		},
		SessionToken: sessionToken,
	}
	r := port.NewResult(http.StatusOK, false, "Success")
	i.OutputPort.SetResponseSignIn(o, r)
}

// SignUp はサインアップ処理を行います。
func (i *UserInteractor) SignUp(ctx context.Context, input port.SignUpInputData) {
	i.Logger.With("email", input.Email)

	user, err := i.UserRepository.ReadByEmail(input.Email, false)
	if err != nil && !errors.Is(err, repository.NewNotFoundError()) {
		i.Logger.Error(err.Error())
		r := port.NewResult(http.StatusInternalServerError, true, err.Error())
		i.OutputPort.SetResponseSignUp(nil, r)
		return
	}
	if user != nil && user.Enabled {
		i.Logger.Warn("user already exists")
		r := port.NewResult(http.StatusBadRequest, true, "入力されたメールアドレスはすでに登録されています")
		i.OutputPort.SetResponseSignUp(nil, r)
		return
	}

	// ユーザーが存在しない場合は新規作成
	if user == nil {
		user = &model.User{
			ID:        id.NewID(),
			Email:     input.Email,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := i.UserRepository.Create(user); err != nil {
			i.Logger.Error(err.Error())
			r := port.NewResult(http.StatusInternalServerError, true, err.Error())
			i.OutputPort.SetResponseSignUp(nil, r)
			return
		}

		i.Logger.Info("user created")
	}

	token, err := i.SessionRepository.GenerateToken(user.ID)
	if err != nil {
		i.Logger.Error(err.Error())
		r := port.NewResult(http.StatusInternalServerError, true, err.Error())
		i.OutputPort.SetResponseSignUp(nil, r)
		return
	}

	config := infrastructure.GetConfig()
	title := fmt.Sprintf("パスワードを設定してください | %s", config.ServiceName)
	baseBodyMessages := []string{
		"%sでパスワード設定のリクエストがありました。",
		"以下のリンクから設定ページを開いてパスワードを設定してください。",
		"",
		"%s/password/set?token=%s",
	}
	baseBody := ""
	for _, m := range baseBodyMessages {
		baseBody += m + "\n"
	}

	body := fmt.Sprintf(baseBody, config.ServiceName, config.BaseUrl, token)

	msgID, err := i.MailRepository.Send(ctx, user.Email, title, body)
	if err != nil {
		i.Logger.Error(err.Error())
		r := port.NewResult(http.StatusInternalServerError, true, err.Error())
		i.OutputPort.SetResponseSignUp(nil, r)
		return
	}

	i.Logger.Info("mail sent", "message_id", msgID)

	o := &port.SignUpOutputData{}
	r := port.NewResult(http.StatusOK, false, "Success")
	i.OutputPort.SetResponseSignUp(o, r)
}

// PasswordReset はパスワードリセット処理を行います。
func (i *UserInteractor) PasswordReset(ctx context.Context, input port.PasswordResetInputData) {
	i.Logger.With("email", input.Email)

	user, err := i.UserRepository.ReadByEmail(input.Email, true)
	if err != nil {
		if errors.Is(err, repository.NewNotFoundError()) {
			i.Logger.Warn("user not found")
			r := port.NewResult(http.StatusUnauthorized, true, "存在しないメールアドレスです")
			i.OutputPort.SetResponsePasswordReset(nil, r)
			return
		}

		i.Logger.Error(err.Error())
		r := port.NewResult(http.StatusInternalServerError, true, err.Error())
		i.OutputPort.SetResponsePasswordReset(nil, r)
		return
	}

	token, err := i.SessionRepository.GenerateToken(user.ID)
	if err != nil {
		i.Logger.Error(err.Error())
		r := port.NewResult(http.StatusInternalServerError, true, err.Error())
		i.OutputPort.SetResponseSignUp(nil, r)
		return
	}

	config := infrastructure.GetConfig()
	title := fmt.Sprintf("パスワードを設定してください | %s", config.ServiceName)
	baseBodyMessages := []string{
		"%sでパスワードリセットのリクエストがありました。",
		"以下のリンクから設定ページを開いてパスワードを設定してください。",
		"",
		"%s/password/set?token=%s",
	}
	baseBody := ""
	for _, m := range baseBodyMessages {
		baseBody += m + "\n"
	}

	body := fmt.Sprintf(baseBody, config.ServiceName, config.BaseUrl, token)

	msgID, err := i.MailRepository.Send(ctx, user.Email, title, body)
	if err != nil {
		i.Logger.Error(err.Error())
		r := port.NewResult(http.StatusInternalServerError, true, err.Error())
		i.OutputPort.SetResponseSignUp(nil, r)
		return
	}

	i.Logger.Info("mail sent", "message_id", msgID)

	o := &port.PasswordResetOutputData{}
	r := port.NewResult(http.StatusOK, false, "Success")
	i.OutputPort.SetResponsePasswordReset(o, r)
}

// PasswordSet はパスワード設定処理を行います。
func (i *UserInteractor) PasswordSet(input port.PasswordSetInputData) {
	i.Logger.With("token", input.Token)

	valid, userID := i.SessionRepository.IsValidToken(input.Token)
	if !valid {
		i.Logger.Warn("token is invalid")
		r := port.NewResult(http.StatusUnauthorized, true, "トークンが無効もしくは期限切れです")
		i.OutputPort.SetResponsePasswordReset(nil, r)
		return
	}

	user, err := i.UserRepository.Read(userID, false)
	if err != nil {
		i.Logger.Error(err.Error())
		r := port.NewResult(http.StatusInternalServerError, true, err.Error())
		i.OutputPort.SetResponsePasswordReset(nil, r)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		i.Logger.Error(err.Error())
		r := port.NewResult(http.StatusInternalServerError, true, err.Error())
		i.OutputPort.SetResponsePasswordReset(nil, r)
		return
	}

	user.Password = string(hashedPassword)
	user.Enabled = true
	user.UpdatedAt = time.Now()

	if err := i.UserRepository.Update(user); err != nil {
		i.Logger.Error(err.Error())
		r := port.NewResult(http.StatusInternalServerError, true, err.Error())
		i.OutputPort.SetResponsePasswordReset(nil, r)
		return
	}

	o := &port.PasswordResetOutputData{}
	r := port.NewResult(http.StatusOK, false, "Success")
	i.OutputPort.SetResponsePasswordReset(o, r)
}
