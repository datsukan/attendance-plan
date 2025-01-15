package request

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func ToSignInRequest(r events.APIGatewayProxyRequest) *SignInRequest {
	req := &SignInRequest{}
	json.Unmarshal([]byte(r.Body), req)
	return req
}

func ValidateSignInRequest(req *SignInRequest) error {
	if req.Email == "" {
		return fmt.Errorf("email is empty")
	}

	if req.Password == "" {
		return fmt.Errorf("password is empty")
	}

	return nil
}

func ToSignUpRequest(r events.APIGatewayProxyRequest) *SignUpRequest {
	req := &SignUpRequest{}
	json.Unmarshal([]byte(r.Body), req)
	return req
}

func ValidateSignUpRequest(req *SignUpRequest) error {
	if req.Email == "" {
		return fmt.Errorf("email is empty")
	}

	if req.Password == "" {
		return fmt.Errorf("password is empty")
	}

	if req.Name == "" {
		return fmt.Errorf("name is empty")
	}

	return nil
}
