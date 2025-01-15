package handler

import (
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
	req := request.ToSignInRequest(r)
	if err := request.ValidateSignInRequest(req); err != nil {
		return response.NewError(http.StatusBadRequest, err.Error())
	}

	db := infrastructure.NewDB()
	ur := repository.NewUserRepository(*db)
	sr := repository.NewSessionRepository(*db)
	op := presenter.NewUserPresenter()
	interactor := usecase.NewUserInteractor(ur, sr, op)

	input := port.SignInInputData{Email: req.Email, Password: req.Password}
	interactor.SignIn(input)

	statusCode, body := op.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
	}
	return res, nil
}

// SignUp はサインアップを行います。
func SignUp(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req := request.ToSignUpRequest(r)
	if err := request.ValidateSignUpRequest(req); err != nil {
		return response.NewError(http.StatusBadRequest, err.Error())
	}

	db := infrastructure.NewDB()
	ur := repository.NewUserRepository(*db)
	sr := repository.NewSessionRepository(*db)
	op := presenter.NewUserPresenter()
	interactor := usecase.NewUserInteractor(ur, sr, op)

	input := port.SignUpInputData{Email: req.Email, Password: req.Password, Name: req.Name}
	interactor.SignUp(input)

	statusCode, body := op.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
	}
	return res, nil
}
