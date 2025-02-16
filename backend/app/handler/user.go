package handler

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
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

	req := request.ToSignInRequest(r)
	if err := request.ValidateSignInRequest(req); err != nil {
		return response.NewError(http.StatusBadRequest, err.Error())
	}

	config := infrastructure.GetConfig()
	db := infrastructure.NewDB()
	ur := repository.NewUserRepository(*db)
	sr := repository.NewSessionRepository(config.SecretKey, config.TokenLifeTime)
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

	req := request.ToSignUpRequest(r)
	if err := request.ValidateSignUpRequest(req); err != nil {
		return response.NewError(http.StatusBadRequest, err.Error())
	}

	config := infrastructure.GetConfig()
	db := infrastructure.NewDB()
	ur := repository.NewUserRepository(*db)
	sr := repository.NewSessionRepository(config.SecretKey, config.TokenLifeTime)

	mc, err := infrastructure.NewMailClient(ctx, config.SESRegion)
	if err != nil {
		return response.NewError(http.StatusInternalServerError, err.Error())
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

	req := request.ToPasswordResetRequest(r)
	if err := request.ValidatePasswordResetRequest(req); err != nil {
		return response.NewError(http.StatusBadRequest, err.Error())
	}

	config := infrastructure.GetConfig()
	db := infrastructure.NewDB()
	ur := repository.NewUserRepository(*db)
	sr := repository.NewSessionRepository(config.SecretKey, config.TokenLifeTime)

	mc, err := infrastructure.NewMailClient(ctx, config.SESRegion)
	if err != nil {
		return response.NewError(http.StatusInternalServerError, err.Error())
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

	req := request.ToPasswordSetRequest(r)
	if err := request.ValidatePasswordSetRequest(req); err != nil {
		return response.NewError(http.StatusBadRequest, err.Error())
	}

	config := infrastructure.GetConfig()
	db := infrastructure.NewDB()
	ur := repository.NewUserRepository(*db)
	sr := repository.NewSessionRepository(config.SecretKey, config.TokenLifeTime)
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
