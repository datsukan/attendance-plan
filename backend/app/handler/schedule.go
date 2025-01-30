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
	config := infrastructure.GetConfig()
	ssRepo := repository.NewSessionRepository(config.SecretKey, config.TokenLifeTime)
	am := middleware.NewAuthMiddleware(ssRepo)
	userID, err := am.Auth(r)
	if err != nil {
		return response.NewError(http.StatusUnauthorized, err.Error())
	}

	req := request.ToGetScheduleListRequest(r)
	if err := request.ValidateGetScheduleListRequest(req); err != nil {
		return response.NewError(http.StatusBadRequest, err.Error())
	}

	if req.UserID != userID {
		return response.NewError(http.StatusForbidden, "forbidden")
	}

	db := infrastructure.NewDB()
	sr := repository.NewScheduleRepository(*db)
	op := presenter.NewSchedulePresenter()
	interactor := usecase.NewScheduleInteractor(sr, op)

	input := port.GetScheduleListInputData{UserID: req.UserID}
	interactor.GetScheduleList(input)

	statusCode, body := op.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
	}
	return res, nil
}

// GetSchedule はスケジュールを取得します。
func GetSchedule(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	config := infrastructure.GetConfig()
	ssRepo := repository.NewSessionRepository(config.SecretKey, config.TokenLifeTime)
	am := middleware.NewAuthMiddleware(ssRepo)
	userID, err := am.Auth(r)
	if err != nil {
		return response.NewError(http.StatusUnauthorized, err.Error())
	}

	req := request.ToGetScheduleRequest(r)
	if err := request.ValidateGetScheduleRequest(req); err != nil {
		return response.NewError(http.StatusBadRequest, err.Error())
	}

	db := infrastructure.NewDB()
	sr := repository.NewScheduleRepository(*db)

	schedule, err := sr.Read(req.ScheduleID)
	if err != nil {
		if errors.Is(err, repository.NewNotFoundError()) {
			return response.NewError(http.StatusNotFound, err.Error())
		}
		return response.NewError(http.StatusInternalServerError, err.Error())
	}

	if schedule.UserID != userID {
		return response.NewError(http.StatusForbidden, "forbidden")
	}

	op := presenter.NewSchedulePresenter()
	interactor := usecase.NewScheduleInteractor(sr, op)

	input := port.GetScheduleInputData{ScheduleID: req.ScheduleID}
	interactor.GetSchedule(input)

	statusCode, body := op.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
	}
	return res, nil
}

// PostSchedule はスケジュールを登録します。
func PostSchedule(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	config := infrastructure.GetConfig()
	ssRepo := repository.NewSessionRepository(config.SecretKey, config.TokenLifeTime)
	am := middleware.NewAuthMiddleware(ssRepo)
	userID, err := am.Auth(r)
	if err != nil {
		return response.NewError(http.StatusUnauthorized, err.Error())
	}

	req, err := request.ToPostScheduleRequest(r)
	if err != nil {
		return response.NewError(http.StatusBadRequest, err.Error())
	}

	if err := request.ValidatePostScheduleRequest(req); err != nil {
		return response.NewError(http.StatusBadRequest, err.Error())
	}

	db := infrastructure.NewDB()
	sr := repository.NewScheduleRepository(*db)
	op := presenter.NewSchedulePresenter()
	interactor := usecase.NewScheduleInteractor(sr, op)

	input := port.CreateScheduleInputData{
		Schedule: port.CreateScheduleData{
			UserID:   userID,
			Name:     req.Name,
			StartsAt: req.StartsAt,
			EndsAt:   req.EndsAt,
			Color:    req.Color,
			Type:     req.Type,
		},
	}
	interactor.CreateSchedule(input)

	statusCode, body := op.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
	}
	return res, nil
}

// PutSchedule はスケジュールを更新します。
func PutSchedule(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	config := infrastructure.GetConfig()
	ssRepo := repository.NewSessionRepository(config.SecretKey, config.TokenLifeTime)
	am := middleware.NewAuthMiddleware(ssRepo)
	userID, err := am.Auth(r)
	if err != nil {
		return response.NewError(http.StatusUnauthorized, err.Error())
	}

	req, err := request.ToPutScheduleRequest(r)
	if err != nil {
		return response.NewError(http.StatusBadRequest, err.Error())
	}

	if err := request.ValidatePutScheduleRequest(req); err != nil {
		return response.NewError(http.StatusBadRequest, err.Error())
	}

	db := infrastructure.NewDB()
	sr := repository.NewScheduleRepository(*db)

	schedule, err := sr.Read(req.ScheduleID)
	if err != nil {
		if errors.Is(err, repository.NewNotFoundError()) {
			return response.NewError(http.StatusNotFound, err.Error())
		}
		return response.NewError(http.StatusInternalServerError, err.Error())
	}

	if schedule.UserID != userID {
		return response.NewError(http.StatusForbidden, "forbidden")
	}

	op := presenter.NewSchedulePresenter()
	interactor := usecase.NewScheduleInteractor(sr, op)

	input := port.UpdateScheduleInputData{
		Schedule: port.UpdateScheduleData{
			ID:       req.ScheduleID,
			Name:     req.Name,
			StartsAt: req.StartsAt,
			EndsAt:   req.EndsAt,
			Color:    req.Color,
			Type:     req.Type,
		},
	}
	interactor.UpdateSchedule(input)

	statusCode, body := op.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
	}
	return res, nil
}

// DeleteSchedule はスケジュールを削除します。
func DeleteSchedule(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	config := infrastructure.GetConfig()
	ssRepo := repository.NewSessionRepository(config.SecretKey, config.TokenLifeTime)
	am := middleware.NewAuthMiddleware(ssRepo)
	userID, err := am.Auth(r)
	if err != nil {
		return response.NewError(http.StatusUnauthorized, err.Error())
	}

	req := request.ToDeleteScheduleRequest(r)
	if err := request.ValidateDeleteScheduleRequest(req); err != nil {
		return response.NewError(http.StatusBadRequest, err.Error())
	}

	db := infrastructure.NewDB()
	sr := repository.NewScheduleRepository(*db)

	schedule, err := sr.Read(req.ScheduleID)
	if err != nil {
		if errors.Is(err, repository.NewNotFoundError()) {
			return response.NewError(http.StatusNotFound, err.Error())
		}
		return response.NewError(http.StatusInternalServerError, err.Error())
	}

	if schedule.UserID != userID {
		return response.NewError(http.StatusForbidden, "forbidden")
	}

	op := presenter.NewSchedulePresenter()
	interactor := usecase.NewScheduleInteractor(sr, op)

	input := port.DeleteScheduleInputData{ScheduleID: req.ScheduleID}
	interactor.DeleteSchedule(input)

	statusCode, body := op.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
	}
	return res, nil
}
