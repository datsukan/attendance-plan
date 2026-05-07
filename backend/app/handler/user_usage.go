package handler

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/datsukan/attendance-plan/backend/app/middleware"
	"github.com/datsukan/attendance-plan/backend/app/port"
	"github.com/datsukan/attendance-plan/backend/app/presenter"
	"github.com/datsukan/attendance-plan/backend/app/repository"
	"github.com/datsukan/attendance-plan/backend/app/response"
	"github.com/datsukan/attendance-plan/backend/app/usecase"
	"github.com/datsukan/attendance-plan/backend/infrastructure"
)

// GetUserUsages は全ユーザーの利用状況リストを取得します。
func GetUserUsages(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger := infrastructure.NewLogger()
	logger.Info("start get user usages")

	config := infrastructure.GetConfig()
	sr := repository.NewSessionRepository(config.SecretKey, config.TokenLifeDays)
	am := middleware.NewAuthMiddleware(sr)
	userID, err := am.Auth(r)
	if err != nil {
		return response.NewError(http.StatusUnauthorized, usecase.MsgUnauthorized)
	}

	logger.With("user_id", userID)

	db := infrastructure.NewDB()
	ur := repository.NewUserRepository(*db)
	sbr := repository.NewSubjectRepository(*db)
	scr := repository.NewScheduleRepository(*db)

	up := presenter.NewUserUsagePresenter()
	interactor := usecase.NewUserUsageInteractor(logger, ur, sbr, scr, up)
	interactor.GetUserUsageList(port.GetUserUsageListInputData{RequesterUserID: userID})

	statusCode, body := up.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers:    response.CORSHeaders,
	}

	logger.Info("end get user usages")

	return res, nil
}
