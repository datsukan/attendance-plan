package user

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/datsukan/attendance-plan/backend/component"
	"github.com/datsukan/attendance-plan/backend/infrastructure"
)

// SignIn はサインインを行います。
func SignIn(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req := ToSignInRequest(r)
	if err := ValidateSignInRequest(req); err != nil {
		return component.NewResponseError(http.StatusBadRequest, err.Error())
	}

	db := infrastructure.NewDB()
	ur := NewUserRepository(*db)
	op := NewUserPresenter()
	interactor := NewUserInteractor(ur, op)

	input := SignInInputData{Email: req.Email, Password: req.Password}
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
	req := ToSignUpRequest(r)
	if err := ValidateSignUpRequest(req); err != nil {
		return component.NewResponseError(http.StatusBadRequest, err.Error())
	}

	db := infrastructure.NewDB()
	ur := NewUserRepository(*db)
	op := NewUserPresenter()
	interactor := NewUserInteractor(ur, op)

	input := SignUpInputData{Email: req.Email, Password: req.Password, Name: req.Name}
	interactor.SignUp(input)

	statusCode, body := op.GetResponse()
	res := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
	}
	return res, nil
}
