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

// GetScheduleList はスケジュールリストを取得します。
func GetScheduleList(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger := infrastructure.NewLogger()
	logger.Info("start get schedule list")

	config := infrastructure.GetConfig()
	ssRepo := repository.NewSessionRepository(config.SecretKey, config.TokenLifeDays)
	am := middleware.NewAuthMiddleware(ssRepo)
	userID, err := am.Auth(r)
	if err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusUnauthorized, usecase.MsgUnauthorized)
	}

	logger.With("user_id", userID)

	req := request.ToGetScheduleListRequest(r)
	if err := request.ValidateGetScheduleListRequest(req); err != nil {
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

	sr := repository.NewScheduleRepository(*db)
	op := presenter.NewSchedulePresenter()
	interactor := usecase.NewScheduleInteractor(logger, sr, op)

	input := port.GetScheduleListInputData{UserID: req.UserID}
	interactor.GetScheduleList(input)

	statusCode, body := op.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers:    response.CORSHeaders,
	}

	logger.Info("end get schedule list")

	return res, nil
}

// GetSchedule はスケジュールを取得します。
func GetSchedule(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger := infrastructure.NewLogger()
	logger.Info("start get schedule")

	config := infrastructure.GetConfig()
	ssRepo := repository.NewSessionRepository(config.SecretKey, config.TokenLifeDays)
	am := middleware.NewAuthMiddleware(ssRepo)
	userID, err := am.Auth(r)
	if err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusUnauthorized, usecase.MsgUnauthorized)
	}

	logger.With("user_id", userID)

	req := request.ToGetScheduleRequest(r)
	if err := request.ValidateGetScheduleRequest(req); err != nil {
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

	sr := repository.NewScheduleRepository(*db)

	schedule, err := sr.Read(req.ScheduleID)
	if err != nil {
		if errors.Is(err, repository.NewNotFoundError()) {
			logger.Warn(err.Error())
			return response.NewError(http.StatusNotFound, usecase.MsgScheduleNotFound)
		}

		logger.Error(err.Error())
		return response.NewError(http.StatusInternalServerError, usecase.MsgInternalServerError)
	}

	if schedule.UserID != userID {
		logger.Warn("forbidden", "request_user_id", schedule.UserID)
		return response.NewError(http.StatusForbidden, usecase.MsgUserNotFound)
	}

	op := presenter.NewSchedulePresenter()
	interactor := usecase.NewScheduleInteractor(logger, sr, op)

	input := port.GetScheduleInputData{ScheduleID: req.ScheduleID}
	interactor.GetSchedule(input)

	statusCode, body := op.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers:    response.CORSHeaders,
	}

	logger.Info("end get schedule")

	return res, nil
}

// PostSchedule はスケジュールを登録します。
func PostSchedule(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger := infrastructure.NewLogger()
	logger.Info("start post schedule")

	config := infrastructure.GetConfig()
	ssRepo := repository.NewSessionRepository(config.SecretKey, config.TokenLifeDays)
	am := middleware.NewAuthMiddleware(ssRepo)
	userID, err := am.Auth(r)
	if err != nil {
		return response.NewError(http.StatusUnauthorized, usecase.MsgUnauthorized)
	}

	logger.With("user_id", userID)

	req, err := request.ToPostScheduleRequest(r)
	if err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusBadRequest, usecase.MsgRequestFormatInvalid)
	}

	if err := request.ValidatePostScheduleRequest(req); err != nil {
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

	sr := repository.NewScheduleRepository(*db)
	op := presenter.NewSchedulePresenter()
	interactor := usecase.NewScheduleInteractor(logger, sr, op)

	input := port.CreateScheduleInputData{
		Schedule: port.CreateScheduleData{
			UserID:   userID,
			Name:     req.Name,
			StartsAt: req.StartsAt,
			EndsAt:   req.EndsAt,
			Color:    req.Color,
			Type:     req.Type,
			Order:    req.Order,
		},
	}
	interactor.CreateSchedule(input)

	statusCode, body := op.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers:    response.CORSHeaders,
	}

	logger.Info("end post schedule")

	return res, nil
}

// PostBulkSchedule はスケジュールを一括登録します。
func PostBulkSchedule(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger := infrastructure.NewLogger()
	logger.Info("start post bulk schedule")

	config := infrastructure.GetConfig()
	ssRepo := repository.NewSessionRepository(config.SecretKey, config.TokenLifeDays)
	am := middleware.NewAuthMiddleware(ssRepo)
	userID, err := am.Auth(r)
	if err != nil {
		return response.NewError(http.StatusUnauthorized, usecase.MsgUnauthorized)
	}

	logger.With("user_id", userID)

	req, err := request.ToPostBulkScheduleRequest(r)
	if err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusBadRequest, usecase.MsgRequestFormatInvalid)
	}

	if err := request.ValidatePostBulkScheduleRequest(req); err != nil {
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

	sr := repository.NewScheduleRepository(*db)

	schedules := make([]port.CreateScheduleData, len(req.Schedules))
	for i, s := range req.Schedules {
		schedules[i] = port.CreateScheduleData{
			UserID:   userID,
			Name:     s.Name,
			StartsAt: s.StartsAt,
			EndsAt:   s.EndsAt,
			Color:    s.Color,
			Type:     s.Type,
			Order:    s.Order,
		}
	}

	input := port.CreateBulkScheduleInputData{Schedules: schedules}
	op := presenter.NewSchedulePresenter()
	interactor := usecase.NewScheduleInteractor(logger, sr, op)
	interactor.CreateBulkSchedule(input)

	statusCode, body := op.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers:    response.CORSHeaders,
	}

	logger.Info("end post bulk schedule")

	return res, nil
}

// PutSchedule はスケジュールを更新します。
func PutSchedule(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger := infrastructure.NewLogger()
	logger.Info("start put schedule")

	config := infrastructure.GetConfig()
	ssRepo := repository.NewSessionRepository(config.SecretKey, config.TokenLifeDays)
	am := middleware.NewAuthMiddleware(ssRepo)
	userID, err := am.Auth(r)
	if err != nil {
		return response.NewError(http.StatusUnauthorized, usecase.MsgUnauthorized)
	}

	logger.With("user_id", userID)

	req, err := request.ToPutScheduleRequest(r)
	if err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusBadRequest, usecase.MsgRequestFormatInvalid)
	}

	if err := request.ValidatePutScheduleRequest(req); err != nil {
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

	sr := repository.NewScheduleRepository(*db)

	schedule, err := sr.Read(req.ScheduleID)
	if err != nil {
		if errors.Is(err, repository.NewNotFoundError()) {
			logger.Warn(err.Error())
			return response.NewError(http.StatusNotFound, usecase.MsgScheduleNotFound)
		}

		logger.Error(err.Error())
		return response.NewError(http.StatusInternalServerError, usecase.MsgInternalServerError)
	}

	if schedule.UserID != userID {
		logger.Warn("forbidden", "request_user_id", schedule.UserID)
		return response.NewError(http.StatusForbidden, usecase.MsgUserNotFound)
	}

	op := presenter.NewSchedulePresenter()
	interactor := usecase.NewScheduleInteractor(logger, sr, op)

	input := port.UpdateScheduleInputData{
		Schedule: port.UpdateScheduleData{
			ID:       req.ScheduleID,
			Name:     req.Name,
			StartsAt: req.StartsAt,
			EndsAt:   req.EndsAt,
			Color:    req.Color,
			Type:     req.Type,
			Order:    req.Order,
		},
	}
	interactor.UpdateSchedule(input)

	statusCode, body := op.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers:    response.CORSHeaders,
	}

	logger.Info("end put schedule")

	return res, nil
}

// PutBulkSchedule はスケジュールを一括更新します。
func PutBulkSchedule(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger := infrastructure.NewLogger()
	logger.Info("start put bulk schedule")

	config := infrastructure.GetConfig()
	ssRepo := repository.NewSessionRepository(config.SecretKey, config.TokenLifeDays)
	am := middleware.NewAuthMiddleware(ssRepo)
	userID, err := am.Auth(r)
	if err != nil {
		return response.NewError(http.StatusUnauthorized, usecase.MsgUnauthorized)
	}

	logger.With("user_id", userID)

	req, err := request.ToPutBulkScheduleRequest(r)
	if err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusBadRequest, usecase.MsgRequestFormatInvalid)
	}

	if err := request.ValidatePutBulkScheduleRequest(req); err != nil {
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

	sr := repository.NewScheduleRepository(*db)

	schedules := make([]port.UpdateScheduleData, len(req.Schedules))
	for i, s := range req.Schedules {
		schedule, err := sr.Read(s.ScheduleID)
		if err != nil {
			if errors.Is(err, repository.NewNotFoundError()) {
				logger.Warn(err.Error())
				return response.NewError(http.StatusNotFound, usecase.MsgScheduleNotFound)
			}

			logger.Error(err.Error())
			return response.NewError(http.StatusInternalServerError, usecase.MsgInternalServerError)
		}

		if schedule.UserID != userID {
			logger.Warn("forbidden", "request_user_id", schedule.UserID)
			return response.NewError(http.StatusForbidden, usecase.MsgUserNotFound)
		}

		schedules[i] = port.UpdateScheduleData{
			ID:       s.ScheduleID,
			Name:     s.Name,
			StartsAt: s.StartsAt,
			EndsAt:   s.EndsAt,
			Color:    s.Color,
			Type:     s.Type,
			Order:    s.Order,
		}
	}

	input := port.UpdateBulkScheduleInputData{Schedules: schedules}
	op := presenter.NewSchedulePresenter()
	interactor := usecase.NewScheduleInteractor(logger, sr, op)
	interactor.UpdateBulkSchedule(input)

	statusCode, body := op.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers:    response.CORSHeaders,
	}

	logger.Info("end put bulk schedule")

	return res, nil
}

// DeleteSchedule はスケジュールを削除します。
func DeleteSchedule(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger := infrastructure.NewLogger()
	logger.Info("start delete schedule")

	config := infrastructure.GetConfig()
	ssRepo := repository.NewSessionRepository(config.SecretKey, config.TokenLifeDays)
	am := middleware.NewAuthMiddleware(ssRepo)
	userID, err := am.Auth(r)
	if err != nil {
		logger.Warn(err.Error())
		return response.NewError(http.StatusUnauthorized, usecase.MsgUnauthorized)
	}

	logger.With("user_id", userID)

	req := request.ToDeleteScheduleRequest(r)
	if err := request.ValidateDeleteScheduleRequest(req); err != nil {
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

	sr := repository.NewScheduleRepository(*db)

	schedule, err := sr.Read(req.ScheduleID)
	if err != nil {
		if errors.Is(err, repository.NewNotFoundError()) {
			logger.Warn(err.Error())
			res := events.APIGatewayProxyResponse{
				StatusCode: http.StatusNoContent,
				Headers:    response.CORSHeaders,
			}

			logger.Info("end delete schedule")

			return res, nil
		}

		logger.Error(err.Error())
		return response.NewError(http.StatusInternalServerError, usecase.MsgInternalServerError)
	}

	if schedule.UserID != userID {
		logger.Warn("forbidden", "request_user_id", schedule.UserID)
		return response.NewError(http.StatusForbidden, usecase.MsgUserNotFound)
	}

	op := presenter.NewSchedulePresenter()
	interactor := usecase.NewScheduleInteractor(logger, sr, op)

	input := port.DeleteScheduleInputData{ScheduleID: req.ScheduleID}
	interactor.DeleteSchedule(input)

	statusCode, body := op.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers:    response.CORSHeaders,
	}

	logger.Info("end delete schedule")

	return res, nil
}
