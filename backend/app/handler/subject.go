package handler

import (
	"errors"
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

// GetSubjectList は科目リストを取得します。
func GetSubjectList(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger := infrastructure.NewLogger()
	logger.Info("start get subject list")

	config := infrastructure.GetConfig()
	ssRepo := repository.NewSessionRepository(config.SecretKey, config.TokenLifeDays)
	am := middleware.NewAuthMiddleware(ssRepo)
	userID, err := am.Auth(r)
	if err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusUnauthorized, usecase.MsgUnauthorized)
	}

	logger.With("user_id", userID)

	req := request.ToGetSubjectListRequest(r)
	if err := request.ValidateGetSubjectListRequest(req); err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusBadRequest, err.Error())
	}

	if req.UserID != userID {
		logger.Warn("forbidden", "request_user_id", req.UserID)
		return response.NewError(http.StatusForbidden, usecase.MsgUserNotFound)
	}

	db := infrastructure.NewDB()
	ur := repository.NewUserRepository(*db)
	if _, err := ur.Read(userID, true); err != nil {
		if errors.Is(err, repository.NewNotFoundError()) {
			logger.Warn(err.Error())
			return response.NewError(http.StatusUnauthorized, usecase.MsgUnauthorized)
		}

		logger.Error(err.Error())
		return response.NewError(http.StatusInternalServerError, usecase.MsgInternalServerError)
	}

	sr := repository.NewSubjectRepository(*db)
	op := presenter.NewSubjectPresenter()
	interactor := usecase.NewSubjectInteractor(logger, sr, op)
	interactor.GetSubjectList(port.GetSubjectListInputData{UserID: userID})

	statusCode, body := op.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers:    response.CORSHeaders,
	}

	logger.Info("end get subject list")

	return res, nil
}

// PostSubject は科目を作成します。
func PostSubject(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger := infrastructure.NewLogger()
	logger.Info("start post subject")

	config := infrastructure.GetConfig()
	ssRepo := repository.NewSessionRepository(config.SecretKey, config.TokenLifeDays)
	am := middleware.NewAuthMiddleware(ssRepo)
	userID, err := am.Auth(r)
	if err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusUnauthorized, usecase.MsgUnauthorized)
	}

	logger.With("user_id", userID)

	req, err := request.ToPostSubjectRequest(r)
	if err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusBadRequest, err.Error())
	}

	if err := request.ValidatePostSubjectRequest(req); err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusBadRequest, err.Error())
	}

	db := infrastructure.NewDB()
	ur := repository.NewUserRepository(*db)
	if _, err := ur.Read(userID, true); err != nil {
		if errors.Is(err, repository.NewNotFoundError()) {
			logger.Warn(err.Error())
			return response.NewError(http.StatusUnauthorized, usecase.MsgUnauthorized)
		}

		logger.Error(err.Error())
		return response.NewError(http.StatusInternalServerError, usecase.MsgInternalServerError)
	}

	sr := repository.NewSubjectRepository(*db)
	op := presenter.NewSubjectPresenter()
	interactor := usecase.NewSubjectInteractor(logger, sr, op)
	interactor.CreateSubject(port.CreateSubjectInputData{
		UserID: userID,
		Name:   req.Name,
		Color:  req.Color,
	})

	statusCode, body := op.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers:    response.CORSHeaders,
	}

	logger.Info("end post subject")

	return res, nil
}

// DeleteSubject は科目を削除します。
func DeleteSubject(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger := infrastructure.NewLogger()
	logger.Info("start delete subject")

	config := infrastructure.GetConfig()
	ssRepo := repository.NewSessionRepository(config.SecretKey, config.TokenLifeDays)
	am := middleware.NewAuthMiddleware(ssRepo)
	userID, err := am.Auth(r)
	if err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusUnauthorized, usecase.MsgUnauthorized)
	}

	logger.With("user_id", userID)

	req := request.ToDeleteSubjectRequest(r)
	if err := request.ValidateDeleteSubjectRequest(req); err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusBadRequest, err.Error())
	}

	db := infrastructure.NewDB()
	ur := repository.NewUserRepository(*db)
	if _, err := ur.Read(userID, true); err != nil {
		if errors.Is(err, repository.NewNotFoundError()) {
			logger.Warn(err.Error())
			return response.NewError(http.StatusUnauthorized, usecase.MsgUnauthorized)
		}

		logger.Error(err.Error())
		return response.NewError(http.StatusInternalServerError, usecase.MsgInternalServerError)
	}

	sr := repository.NewSubjectRepository(*db)
	op := presenter.NewSubjectPresenter()
	interactor := usecase.NewSubjectInteractor(logger, sr, op)
	interactor.DeleteSubject(port.DeleteSubjectInputData{SubjectID: req.SubjectID})

	statusCode, body := op.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers:    response.CORSHeaders,
	}

	logger.Info("end delete subject")

	return res, nil
}
