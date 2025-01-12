package schedule

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/datsukan/attendance-plan/backend/component"
	"github.com/datsukan/attendance-plan/backend/infrastructure"
)

// GetScheduleList はスケジュールリストを取得します。
func GetScheduleList(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req := ToGetScheduleListRequest(r)
	if err := ValidateGetScheduleListRequest(req); err != nil {
		return component.NewResponseError(http.StatusBadRequest, err.Error())
	}

	db := infrastructure.NewDB()
	sr := NewScheduleRepository(*db)
	op := NewSchedulePresenter()
	interactor := NewScheduleInteractor(sr, op)

	input := GetScheduleListInputData{UserID: req.UserID}
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
	req := ToGetScheduleRequest(r)
	if err := ValidateGetScheduleRequest(req); err != nil {
		return component.NewResponseError(http.StatusBadRequest, err.Error())
	}

	db := infrastructure.NewDB()
	sr := NewScheduleRepository(*db)
	op := NewSchedulePresenter()
	interactor := NewScheduleInteractor(sr, op)

	input := GetScheduleInputData{ScheduleID: req.ScheduleID}
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
	req, err := ToPostScheduleRequest(r)
	if err != nil {
		return component.NewResponseError(http.StatusBadRequest, err.Error())
	}

	if err := ValidatePostScheduleRequest(req); err != nil {
		return component.NewResponseError(http.StatusBadRequest, err.Error())
	}

	db := infrastructure.NewDB()
	sr := NewScheduleRepository(*db)
	op := NewSchedulePresenter()
	interactor := NewScheduleInteractor(sr, op)

	input := CreateScheduleInputData{
		Schedule: CreateScheduleData{
			UserID:   "1",
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
	req, err := ToPutScheduleRequest(r)
	if err != nil {
		return component.NewResponseError(http.StatusBadRequest, err.Error())
	}

	if err := ValidatePutScheduleRequest(req); err != nil {
		return component.NewResponseError(http.StatusBadRequest, err.Error())
	}

	db := infrastructure.NewDB()
	sr := NewScheduleRepository(*db)
	op := NewSchedulePresenter()
	interactor := NewScheduleInteractor(sr, op)

	input := UpdateScheduleInputData{
		Schedule: UpdateScheduleData{
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
	req := ToDeleteScheduleRequest(r)
	if err := ValidateDeleteScheduleRequest(req); err != nil {
		return component.NewResponseError(http.StatusBadRequest, err.Error())
	}

	db := infrastructure.NewDB()
	sr := NewScheduleRepository(*db)
	op := NewSchedulePresenter()
	interactor := NewScheduleInteractor(sr, op)

	input := DeleteScheduleInputData{ScheduleID: req.ScheduleID}
	interactor.DeleteSchedule(input)

	statusCode, body := op.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
	}
	return res, nil
}
