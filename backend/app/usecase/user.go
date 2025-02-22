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

// SignIn はサインイン処理を行います。
func (i *UserInteractor) SignIn(input port.SignInInputData) {
	i.Logger.With("email", input.Email)

	user, err := i.UserRepository.ReadByEmail(input.Email, true)
	if err != nil {
		if errors.Is(err, repository.NewNotFoundError()) {
			i.Logger.Warn("user not found")
			r := port.NewErrorResult(http.StatusUnauthorized, MsgEmailOrPasswordInvalid)
			i.OutputPort.SetResponseSignIn(nil, r)
			return
		}

		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
		i.OutputPort.SetResponseSignIn(nil, r)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		i.Logger.Warn("password is invalid")
		r := port.NewErrorResult(http.StatusUnauthorized, MsgEmailOrPasswordInvalid)
		i.OutputPort.SetResponseSignIn(nil, r)
		return
	}

	sessionToken, err := i.SessionRepository.GenerateToken(user.ID)
	if err != nil {
		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
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
	r := port.NewSuccessResult(http.StatusOK)
	i.OutputPort.SetResponseSignIn(o, r)
}

// SignUp はサインアップ処理を行います。
func (i *UserInteractor) SignUp(ctx context.Context, input port.SignUpInputData) {
	i.Logger.With("email", input.Email)

	user, err := i.UserRepository.ReadByEmail(input.Email, false)
	if err != nil && !errors.Is(err, repository.NewNotFoundError()) {
		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
		i.OutputPort.SetResponseSignUp(nil, r)
		return
	}
	if user != nil && user.Enabled {
		i.Logger.Warn("email already exists")
		r := port.NewErrorResult(http.StatusBadRequest, MsgEmailAlreadyExists)
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
			r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
			i.OutputPort.SetResponseSignUp(nil, r)
			return
		}

		i.Logger.Info("user created")
	}

	token, err := i.SessionRepository.GenerateToken(user.ID)
	if err != nil {
		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
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
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
		i.OutputPort.SetResponseSignUp(nil, r)
		return
	}

	i.Logger.Info("mail sent", "message_id", msgID)

	o := &port.SignUpOutputData{}
	r := port.NewSuccessResult(http.StatusOK)
	i.OutputPort.SetResponseSignUp(o, r)
}

// PasswordReset はパスワードリセット処理を行います。
func (i *UserInteractor) PasswordReset(ctx context.Context, input port.PasswordResetInputData) {
	i.Logger.With("email", input.Email)

	user, err := i.UserRepository.ReadByEmail(input.Email, true)
	if err != nil {
		if errors.Is(err, repository.NewNotFoundError()) {
			i.Logger.Warn("user not found by email")
			r := port.NewErrorResult(http.StatusUnauthorized, MsgEmailNotFound)
			i.OutputPort.SetResponsePasswordReset(nil, r)
			return
		}

		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
		i.OutputPort.SetResponsePasswordReset(nil, r)
		return
	}

	token, err := i.SessionRepository.GenerateToken(user.ID)
	if err != nil {
		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
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
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
		i.OutputPort.SetResponseSignUp(nil, r)
		return
	}

	i.Logger.Info("mail sent", "message_id", msgID)

	o := &port.PasswordResetOutputData{}
	r := port.NewSuccessResult(http.StatusOK)
	i.OutputPort.SetResponsePasswordReset(o, r)
}

// PasswordSet はパスワード設定処理を行います。
func (i *UserInteractor) PasswordSet(input port.PasswordSetInputData) {
	i.Logger.With("token", input.Token)

	valid, userID := i.SessionRepository.IsValidToken(input.Token)
	if !valid {
		i.Logger.Warn("token is invalid")
		r := port.NewErrorResult(http.StatusUnauthorized, MsgTokenInvalid)
		i.OutputPort.SetResponsePasswordReset(nil, r)
		return
	}

	user, err := i.UserRepository.Read(userID, false)
	if err != nil {
		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
		i.OutputPort.SetResponsePasswordReset(nil, r)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
		i.OutputPort.SetResponsePasswordReset(nil, r)
		return
	}

	user.Password = string(hashedPassword)
	user.Enabled = true
	user.UpdatedAt = time.Now()

	if err := i.UserRepository.Update(user); err != nil {
		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
		i.OutputPort.SetResponsePasswordReset(nil, r)
		return
	}

	o := &port.PasswordResetOutputData{}
	r := port.NewSuccessResult(http.StatusOK)
	i.OutputPort.SetResponsePasswordReset(o, r)
}

// GetUser はユーザー取得処理を行います。
func (i *UserInteractor) GetUser(input port.GetUserInputData) {
	i.Logger.With("user_id", input.UserID)

	user, err := i.UserRepository.Read(input.UserID, true)
	if err != nil {
		if errors.Is(err, repository.NewNotFoundError()) {
			i.Logger.Warn("user not found")
			r := port.NewErrorResult(http.StatusUnauthorized, MsgUnauthorized)
			i.OutputPort.SetResponseGetUser(nil, r)
			return
		}

		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
		i.OutputPort.SetResponseGetUser(nil, r)
		return
	}

	o := &port.GetUserOutputData{
		BaseUserData: port.BaseUserData{
			ID:        user.ID,
			Email:     user.Email,
			Name:      user.Name,
			CreatedAt: user.CreatedAt.Format(time.DateTime),
			UpdatedAt: user.UpdatedAt.Format(time.DateTime),
		},
	}
	r := port.NewSuccessResult(http.StatusOK)
	i.OutputPort.SetResponseGetUser(o, r)
}

// UpdateUser はユーザー情報更新処理を行います。
func (i *UserInteractor) UpdateUser(input port.UpdateUserInputData) {
	i.Logger.With("user_id", input.UserID)

	user, err := i.UserRepository.Read(input.UserID, true)
	if err != nil {
		if errors.Is(err, repository.NewNotFoundError()) {
			i.Logger.Warn("user not found")
			r := port.NewErrorResult(http.StatusUnauthorized, MsgUnauthorized)
			i.OutputPort.SetResponseUpdateUser(nil, r)
			return
		}

		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
		i.OutputPort.SetResponseUpdateUser(nil, r)
		return
	}

	user.Name = input.Name
	user.UpdatedAt = time.Now()

	if err := i.UserRepository.Update(user); err != nil {
		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
		i.OutputPort.SetResponseUpdateUser(nil, r)
		return
	}

	o := &port.UpdateUserOutputData{
		BaseUserData: port.BaseUserData{
			ID:        user.ID,
			Email:     user.Email,
			Name:      user.Name,
			CreatedAt: user.CreatedAt.Format(time.DateTime),
			UpdatedAt: user.UpdatedAt.Format(time.DateTime),
		},
	}
	r := port.NewSuccessResult(http.StatusOK)
	i.OutputPort.SetResponseUpdateUser(o, r)
}

// DeleteUser はユーザー削除処理を行います。
func (i *UserInteractor) DeleteUser(input port.DeleteUserInputData) {
	i.Logger.With("user_id", input.UserID)

	user, err := i.UserRepository.Read(input.UserID, true)
	if err != nil {
		if errors.Is(err, repository.NewNotFoundError()) {
			i.Logger.Warn("user not found")
			r := port.NewErrorResult(http.StatusUnauthorized, MsgUnauthorized)
			i.OutputPort.SetResponseDeleteUser(nil, r)
			return
		}

		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
		i.OutputPort.SetResponseDeleteUser(nil, r)
		return
	}

	if err := i.UserRepository.Delete(user.ID); err != nil {
		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
		i.OutputPort.SetResponseDeleteUser(nil, r)
		return
	}

	o := &port.DeleteUserOutputData{}
	r := port.NewSuccessResult(http.StatusOK)
	i.OutputPort.SetResponseDeleteUser(o, r)
}

// ResetEmail はメールアドレスリセット処理を行います。
func (i *UserInteractor) ResetEmail(input port.ResetEmailInputData) {
	i.Logger.With("user_id", input.UserID)

	user, err := i.UserRepository.Read(input.UserID, true)
	if err != nil {
		if errors.Is(err, repository.NewNotFoundError()) {
			i.Logger.Warn("user not found")
			r := port.NewErrorResult(http.StatusUnauthorized, MsgUnauthorized)
			i.OutputPort.SetResponseResetEmail(nil, r)
			return
		}

		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
		i.OutputPort.SetResponseResetEmail(nil, r)
		return
	}

	if user.Email == input.Email {
		i.Logger.Warn("email is same")
		r := port.NewErrorResult(http.StatusBadRequest, MsgEmailIsSame)
		i.OutputPort.SetResponseResetEmail(nil, r)
		return
	}

	exists, err := i.UserRepository.ExistsByEmail(input.Email, true)
	if err != nil {
		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
		i.OutputPort.SetResponseResetEmail(nil, r)
		return
	}
	if exists {
		i.Logger.Warn("email already exists")
		r := port.NewErrorResult(http.StatusBadRequest, MsgEmailAlreadyExists)
		i.OutputPort.SetResponseResetEmail(nil, r)
		return
	}

	idToken, err := i.SessionRepository.GenerateToken(user.ID)
	if err != nil {
		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
		i.OutputPort.SetResponseResetEmail(nil, r)
		return
	}

	emailToken, err := i.SessionRepository.GenerateToken(input.Email)
	if err != nil {
		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
		i.OutputPort.SetResponseResetEmail(nil, r)
		return
	}

	config := infrastructure.GetConfig()
	title := fmt.Sprintf("メールアドレスを変更してください | %s", config.ServiceName)
	baseBodyMessages := []string{
		"%sでメールアドレス変更のリクエストがありました。",
		"以下のリンクから変更ページを開いてメールアドレスを変更してください。",
		"",
		"%s/email/set?id_token=%s&email_token=%s",
	}
	baseBody := ""
	for _, m := range baseBodyMessages {
		baseBody += m + "\n"
	}

	body := fmt.Sprintf(baseBody, config.ServiceName, config.BaseUrl, idToken, emailToken)

	msgID, err := i.MailRepository.Send(context.Background(), input.Email, title, body)
	if err != nil {
		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
		i.OutputPort.SetResponseResetEmail(nil, r)
		return
	}

	i.Logger.Info("mail sent", "message_id", msgID)

	o := &port.ResetEmailOutputData{}
	r := port.NewSuccessResult(http.StatusOK)
	i.OutputPort.SetResponseResetEmail(o, r)
}

// SetEmail はメールアドレス設定処理を行います。
func (i *UserInteractor) SetEmail(input port.SetEmailInputData) {
	i.Logger.With("user_id_token", input.UserIDToken, "email_token", input.EmailToken)

	valid, userID := i.SessionRepository.IsValidToken(input.UserIDToken)
	if !valid {
		i.Logger.Warn("user_id_token is invalid")
		r := port.NewErrorResult(http.StatusUnauthorized, MsgTokenInvalid)
		i.OutputPort.SetResponseSetEmail(nil, r)
		return
	}

	valid, email := i.SessionRepository.IsValidToken(input.EmailToken)
	if !valid {
		i.Logger.Warn("email_token is invalid")
		r := port.NewErrorResult(http.StatusUnauthorized, MsgTokenInvalid)
		i.OutputPort.SetResponseSetEmail(nil, r)
		return
	}

	user, err := i.UserRepository.Read(userID, true)
	if err != nil {
		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
		i.OutputPort.SetResponseSetEmail(nil, r)
		return
	}

	user.Email = email
	user.UpdatedAt = time.Now()

	if err := i.UserRepository.Update(user); err != nil {
		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
		i.OutputPort.SetResponseSetEmail(nil, r)
		return
	}

	o := &port.SetEmailOutputData{}
	r := port.NewSuccessResult(http.StatusOK)
	i.OutputPort.SetResponseSetEmail(o, r)
}
