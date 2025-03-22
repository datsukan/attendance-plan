package handler

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/datsukan/attendance-plan/backend/app/middleware"
	"github.com/datsukan/attendance-plan/backend/app/port"
	"github.com/datsukan/attendance-plan/backend/app/presenter"
	"github.com/datsukan/attendance-plan/backend/app/repository"
	"github.com/datsukan/attendance-plan/backend/app/request"
	"github.com/datsukan/attendance-plan/backend/app/response"
	"github.com/datsukan/attendance-plan/backend/app/usecase"
	"github.com/datsukan/attendance-plan/backend/infrastructure"
)

// SignIn はサインインを行います。
func SignIn(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger := infrastructure.NewLogger()
	logger.Info("start sign in")

	req, err := request.ToSignInRequest(r)
	if err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusBadRequest, usecase.MsgRequestFormatInvalid)
	}

	if err := request.ValidateSignInRequest(req); err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusBadRequest, err.Error())
	}

	config := infrastructure.GetConfig()
	db := infrastructure.NewDB()
	ur := repository.NewUserRepository(*db)
	sr := repository.NewSessionRepository(config.SecretKey, config.TokenLifeDays)
	up := presenter.NewUserPresenter()
	interactor := usecase.NewUserInteractor(logger, ur, sr, nil, up)

	input := port.SignInInputData{Email: req.Email, Password: req.Password}
	interactor.SignIn(input)

	statusCode, body := up.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers:    response.CORSHeaders,
	}

	logger.Info("end sign in")

	return res, nil
}

// SignUp はサインアップを行います。
func SignUp(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger := infrastructure.NewLogger()
	logger.Info("start sign up")

	req, err := request.ToSignUpRequest(r)
	if err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusBadRequest, usecase.MsgRequestFormatInvalid)
	}

	if err := request.ValidateSignUpRequest(req); err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusBadRequest, err.Error())
	}

	config := infrastructure.GetConfig()
	db := infrastructure.NewDB()
	ur := repository.NewUserRepository(*db)
	sr := repository.NewSessionRepository(config.SecretKey, config.TokenLifeDays)

	mc, err := infrastructure.NewMailClient(ctx, config.SESRegion)
	if err != nil {
		logger.Error(err.Error())
		return response.NewError(http.StatusInternalServerError, usecase.MsgInternalServerError)
	}

	mr := repository.NewEmailRepository(mc, config.SenderEmail, config.SenderName)
	up := presenter.NewUserPresenter()
	interactor := usecase.NewUserInteractor(logger, ur, sr, mr, up)

	input := port.SignUpInputData{Email: req.Email}
	interactor.SignUp(ctx, input)

	statusCode, body := up.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers:    response.CORSHeaders,
	}

	logger.Info("end sign up")

	return res, nil
}

// PasswordReset はパスワードリセットを行います。
func PasswordReset(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger := infrastructure.NewLogger()
	logger.Info("start sign up")

	req, err := request.ToPasswordResetRequest(r)
	if err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusBadRequest, usecase.MsgRequestFormatInvalid)
	}

	if err := request.ValidatePasswordResetRequest(req); err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusBadRequest, err.Error())
	}

	config := infrastructure.GetConfig()
	db := infrastructure.NewDB()
	ur := repository.NewUserRepository(*db)
	sr := repository.NewSessionRepository(config.SecretKey, config.TokenLifeDays)

	mc, err := infrastructure.NewMailClient(ctx, config.SESRegion)
	if err != nil {
		logger.Error(err.Error())
		return response.NewError(http.StatusInternalServerError, usecase.MsgInternalServerError)
	}

	mr := repository.NewEmailRepository(mc, config.SenderEmail, config.SenderName)
	up := presenter.NewUserPresenter()
	interactor := usecase.NewUserInteractor(logger, ur, sr, mr, up)

	input := port.PasswordResetInputData{Email: req.Email}
	interactor.PasswordReset(ctx, input)

	statusCode, body := up.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers:    response.CORSHeaders,
	}

	logger.Info("end sign up")

	return res, nil
}

// PasswordSet はパスワード設定を行います。
func PasswordSet(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger := infrastructure.NewLogger()
	logger.Info("start password reset")

	req, err := request.ToPasswordSetRequest(r)
	if err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusBadRequest, usecase.MsgRequestFormatInvalid)
	}

	if err := request.ValidatePasswordSetRequest(req); err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusBadRequest, err.Error())
	}

	config := infrastructure.GetConfig()
	db := infrastructure.NewDB()
	ur := repository.NewUserRepository(*db)
	sr := repository.NewSessionRepository(config.SecretKey, config.TokenLifeDays)
	up := presenter.NewUserPresenter()
	interactor := usecase.NewUserInteractor(logger, ur, sr, nil, up)

	input := port.PasswordSetInputData{Token: req.Token, Password: req.Password}
	interactor.PasswordSet(input)

	statusCode, body := up.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers:    response.CORSHeaders,
	}

	logger.Info("end password reset")

	return res, nil
}

// GetUser はユーザー情報を取得します。
func GetUser(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger := infrastructure.NewLogger()
	logger.Info("start get user")

	config := infrastructure.GetConfig()
	sr := repository.NewSessionRepository(config.SecretKey, config.TokenLifeDays)
	am := middleware.NewAuthMiddleware(sr)
	userID, err := am.Auth(r)
	if err != nil {
		return response.NewError(http.StatusUnauthorized, usecase.MsgUnauthorized)
	}

	logger.With("user_id", userID)

	req := request.ToGetUserRequest(r)
	if err := request.ValidateGetUserRequest(req); err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusBadRequest, err.Error())
	}

	if req.UserID != userID {
		logger.Warn("forbidden", "request_user_id", req.UserID)
		return response.NewError(http.StatusForbidden, usecase.MsgUserNotFound)
	}

	db := infrastructure.NewDB()
	ur := repository.NewUserRepository(*db)

	up := presenter.NewUserPresenter()
	interactor := usecase.NewUserInteractor(logger, ur, nil, nil, up)

	input := port.GetUserInputData{UserID: req.UserID}
	interactor.GetUser(input)

	statusCode, body := up.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers:    response.CORSHeaders,
	}

	logger.Info("end get user")

	return res, nil
}

// PutUser はユーザー情報を更新します。
func PutUser(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger := infrastructure.NewLogger()
	logger.Info("start update user")

	config := infrastructure.GetConfig()
	sr := repository.NewSessionRepository(config.SecretKey, config.TokenLifeDays)
	am := middleware.NewAuthMiddleware(sr)
	userID, err := am.Auth(r)
	if err != nil {
		return response.NewError(http.StatusUnauthorized, usecase.MsgUnauthorized)
	}

	logger.With("user_id", userID)

	req, err := request.ToPutUserRequest(r)
	if err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusBadRequest, usecase.MsgRequestFormatInvalid)
	}

	if err := request.ValidatePutUserRequest(req); err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusBadRequest, err.Error())
	}

	if req.UserID != userID {
		logger.Warn("forbidden", "request_user_id", req.UserID)
		return response.NewError(http.StatusForbidden, usecase.MsgUserNotFound)
	}

	db := infrastructure.NewDB()
	ur := repository.NewUserRepository(*db)

	up := presenter.NewUserPresenter()
	interactor := usecase.NewUserInteractor(logger, ur, nil, nil, up)

	input := port.UpdateUserInputData{
		UserID: req.UserID,
		Name:   req.Name,
	}
	interactor.UpdateUser(input)

	statusCode, body := up.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers:    response.CORSHeaders,
	}

	logger.Info("end update user")

	return res, nil
}

// DeleteUser はユーザーを削除します。
func DeleteUser(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger := infrastructure.NewLogger()
	logger.Info("start delete user")

	config := infrastructure.GetConfig()
	sr := repository.NewSessionRepository(config.SecretKey, config.TokenLifeDays)
	am := middleware.NewAuthMiddleware(sr)
	userID, err := am.Auth(r)
	if err != nil {
		return response.NewError(http.StatusUnauthorized, usecase.MsgUnauthorized)
	}

	logger.With("user_id", userID)

	req := request.ToDeleteUserRequest(r)
	if err := request.ValidateDeleteUserRequest(req); err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusBadRequest, err.Error())
	}

	if req.UserID != userID {
		logger.Warn("forbidden", "request_user_id", req.UserID)
		return response.NewError(http.StatusForbidden, usecase.MsgUserNotFound)
	}

	db := infrastructure.NewDB()
	ur := repository.NewUserRepository(*db)

	up := presenter.NewUserPresenter()
	interactor := usecase.NewUserInteractor(logger, ur, nil, nil, up)

	input := port.DeleteUserInputData{UserID: req.UserID}
	interactor.DeleteUser(input)

	statusCode, body := up.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers:    response.CORSHeaders,
	}

	logger.Info("end delete user")

	return res, nil
}

// ResetEmail はメールアドレスの変更を受け付けます。
func ResetEmail(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger := infrastructure.NewLogger()
	logger.Info("start reset email")

	config := infrastructure.GetConfig()
	sr := repository.NewSessionRepository(config.SecretKey, config.TokenLifeDays)
	am := middleware.NewAuthMiddleware(sr)
	userID, err := am.Auth(r)
	if err != nil {
		return response.NewError(http.StatusUnauthorized, usecase.MsgUnauthorized)
	}

	logger.With("user_id", userID)

	req, err := request.ToResetEmailRequest(r)
	if err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusBadRequest, usecase.MsgRequestFormatInvalid)
	}

	if err := request.ValidateResetEmailRequest(req); err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusBadRequest, err.Error())
	}

	if req.UserID != userID {
		logger.Warn("forbidden", "request_user_id", req.UserID)
		return response.NewError(http.StatusForbidden, usecase.MsgUserNotFound)
	}

	db := infrastructure.NewDB()
	ur := repository.NewUserRepository(*db)

	mc, err := infrastructure.NewMailClient(context.Background(), config.SESRegion)
	if err != nil {
		logger.Error(err.Error())
		return response.NewError(http.StatusInternalServerError, usecase.MsgInternalServerError)
	}

	mr := repository.NewEmailRepository(mc, config.SenderEmail, config.SenderName)
	up := presenter.NewUserPresenter()
	interactor := usecase.NewUserInteractor(logger, ur, sr, mr, up)

	input := port.ResetEmailInputData{UserID: req.UserID, Email: req.Email}
	interactor.ResetEmail(input)

	statusCode, body := up.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers:    response.CORSHeaders,
	}

	logger.Info("end reset email")

	return res, nil
}

func SetEmail(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger := infrastructure.NewLogger()
	logger.Info("start set email")

	config := infrastructure.GetConfig()
	sr := repository.NewSessionRepository(config.SecretKey, config.TokenLifeDays)
	am := middleware.NewAuthMiddleware(sr)
	userID, err := am.Auth(r)
	if err != nil {
		return response.NewError(http.StatusUnauthorized, usecase.MsgUnauthorized)
	}

	logger.With("user_id", userID)

	req, err := request.ToSetEmailRequest(r)
	if err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusBadRequest, usecase.MsgRequestFormatInvalid)
	}

	if err := request.ValidateSetEmailRequest(req); err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusBadRequest, err.Error())
	}

	db := infrastructure.NewDB()
	ur := repository.NewUserRepository(*db)
	up := presenter.NewUserPresenter()
	interactor := usecase.NewUserInteractor(logger, ur, sr, nil, up)

	input := port.SetEmailInputData{UserIDToken: req.UserIDToken, EmailToken: req.EmailToken}
	interactor.SetEmail(input)

	statusCode, body := up.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers:    response.CORSHeaders,
	}

	logger.Info("end set email")

	return res, nil
}
