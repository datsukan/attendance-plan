package usecase

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/datsukan/attendance-plan/backend/app/port"
	"github.com/datsukan/attendance-plan/backend/app/repository"
	"github.com/datsukan/attendance-plan/backend/infrastructure"
)

// UserUsageInteractor はユーザー利用状況ユースケースの実装を表す構造体です。
type UserUsageInteractor struct {
	Logger            *slog.Logger
	UserRepository    repository.UserRepository
	SubjectRepository repository.SubjectRepository
	OutputPort        port.UserUsageOutputPort
}

// NewUserUsageInteractor は UserUsageInteractor を生成します。
func NewUserUsageInteractor(logger *slog.Logger, userRepository repository.UserRepository, subjectRepository repository.SubjectRepository, outputPort port.UserUsageOutputPort) port.UserUsageInputPort {
	return &UserUsageInteractor{
		Logger:            logger,
		UserRepository:    userRepository,
		SubjectRepository: subjectRepository,
		OutputPort:        outputPort,
	}
}

// GetUserUsageList は全ユーザーの利用状況リストを取得します。
func (i *UserUsageInteractor) GetUserUsageList(inputData port.GetUserUsageListInputData) {
	requester, err := i.UserRepository.Read(inputData.RequesterUserID, true)
	if err != nil {
		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusUnauthorized, MsgUnauthorized)
		i.OutputPort.SetResponseGetUserUsageList(nil, r)
		return
	}

	config := infrastructure.GetConfig()
	if !isAdmin(requester.Email, config.AdminEmails) {
		i.Logger.Warn("forbidden: not admin", "email", requester.Email)
		r := port.NewErrorResult(http.StatusForbidden, MsgUserNotFound)
		i.OutputPort.SetResponseGetUserUsageList(nil, r)
		return
	}

	users, err := i.UserRepository.ScanAll(true)
	if err != nil {
		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
		i.OutputPort.SetResponseGetUserUsageList(nil, r)
		return
	}

	outputUsers := make([]port.UserUsageData, 0, len(users))
	for _, u := range users {
		subjects, err := i.SubjectRepository.ReadByUserID(u.ID)
		if err != nil {
			i.Logger.Error(err.Error())
			r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
			i.OutputPort.SetResponseGetUserUsageList(nil, r)
			return
		}

		lastUsedAt := u.CreatedAt
		for _, s := range subjects {
			if s.UpdatedAt.After(lastUsedAt) {
				lastUsedAt = s.UpdatedAt
			}
		}

		outputSubjects := make([]port.UserUsageSubjectData, 0, len(subjects))
		for _, s := range subjects {
			outputSubjects = append(outputSubjects, port.UserUsageSubjectData{
				ID:        s.ID,
				Name:      s.Name,
				Color:     s.Color,
				CreatedAt: s.CreatedAt.Format(time.DateTime),
				UpdatedAt: s.UpdatedAt.Format(time.DateTime),
			})
		}

		outputUsers = append(outputUsers, port.UserUsageData{
			ID:           u.ID,
			Email:        u.Email,
			Name:         u.Name,
			RegisteredAt: u.CreatedAt.Format(time.DateTime),
			LastUsedAt:   lastUsedAt.Format(time.DateTime),
			Subjects:     outputSubjects,
		})
	}

	o := &port.GetUserUsageListOutputData{
		Total: len(outputUsers),
		Users: outputUsers,
	}
	r := port.NewSuccessResult(http.StatusOK)
	i.OutputPort.SetResponseGetUserUsageList(o, r)
}

func isAdmin(email string, adminEmails []string) bool {
	for _, e := range adminEmails {
		if e == email {
			return true
		}
	}
	return false
}
