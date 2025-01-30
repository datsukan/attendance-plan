package middleware

import (
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/datsukan/attendance-plan/backend/app/repository"
	"github.com/pkg/errors"
)

type AuthMiddleware interface {
	Auth(r events.APIGatewayProxyRequest) (userID string, err error)
}

type AuthMiddlewareImpl struct {
	SessionRepository repository.SessionRepository
}

func NewAuthMiddleware(sessionRepository repository.SessionRepository) AuthMiddleware {
	return &AuthMiddlewareImpl{SessionRepository: sessionRepository}
}

// Auth は認証処理を行います。
func (m *AuthMiddlewareImpl) Auth(r events.APIGatewayProxyRequest) (userID string, err error) {
	authorization, ok := r.Headers["Authorization"]
	if !ok {
		return "", errors.New("unauthorized")
	}

	if authorization == "" {
		return "", errors.New("unauthorized")
	}

	if !strings.HasPrefix(authorization, "Bearer ") {
		return "", errors.New("unauthorized")
	}

	sessionToken := strings.TrimPrefix(authorization, "Bearer ")

	valid, userID := m.SessionRepository.IsValidToken(sessionToken)
	if !valid {
		return "", errors.New("unauthorized")
	}

	return userID, nil
}
