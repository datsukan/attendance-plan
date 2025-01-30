package middleware

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/datsukan/attendance-plan/backend/app/repository"
	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	tests := []struct {
		name        string
		req         events.APIGatewayProxyRequest
		sessionRepo repository.SessionRepository
		wantError   string
	}{
		{
			name:        "認証済み",
			req:         events.APIGatewayProxyRequest{Headers: map[string]string{"Authorization": "Bearer token"}},
			sessionRepo: NewStubSuccessSessionRepository(),
		},
		{
			name:        "未認証: Authorization がない",
			req:         events.APIGatewayProxyRequest{Headers: map[string]string{}},
			sessionRepo: NewStubSuccessSessionRepository(),
			wantError:   "unauthorized",
		},
		{
			name:        "未認証: Authorization が空",
			req:         events.APIGatewayProxyRequest{Headers: map[string]string{"Authorization": ""}},
			sessionRepo: NewStubSuccessSessionRepository(),
			wantError:   "unauthorized",
		},
		{
			name:        "未認証: Authorization のセッショントークンがない",
			req:         events.APIGatewayProxyRequest{Headers: map[string]string{"Authorization": "Bearer"}},
			sessionRepo: NewStubSuccessSessionRepository(),
			wantError:   "unauthorized",
		},
		{
			name:        "未認証: セッショントークンが無効",
			req:         events.APIGatewayProxyRequest{Headers: map[string]string{"Authorization": "Bearer invalid-token"}},
			sessionRepo: NewStubFailSessionRepository(),
			wantError:   "unauthorized",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewAuthMiddleware(tt.sessionRepo)
			userID, err := m.Auth(tt.req)

			assert := assert.New(t)

			if tt.wantError != "" {
				if !assert.Error(err) {
					return
				}

				assert.Equal(tt.wantError, err.Error())
				assert.Empty(userID)
				return
			}

			assert.NoError(err)
			assert.Equal("user-id", userID)
		})
	}
}
