package middleware

import "github.com/datsukan/attendance-plan/backend/app/repository"

type stubSuccessSessionRepositoryImpl struct{}

func NewStubSuccessSessionRepository() repository.SessionRepository {
	return &stubSuccessSessionRepositoryImpl{}
}

func (r *stubSuccessSessionRepositoryImpl) GenerateToken(userID string) (string, error) {
	return "token", nil
}

func (r *stubSuccessSessionRepositoryImpl) IsValidToken(token string) (bool, string) {
	return true, "user-id"
}

type stubFailSessionRepositoryImpl struct{}

func NewStubFailSessionRepository() repository.SessionRepository {
	return &stubFailSessionRepositoryImpl{}
}

func (r *stubFailSessionRepositoryImpl) GenerateToken(userID string) (string, error) {
	return "", nil
}

func (r *stubFailSessionRepositoryImpl) IsValidToken(token string) (bool, string) {
	return false, ""
}
