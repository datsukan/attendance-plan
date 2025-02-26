package usecase

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/datsukan/attendance-plan/backend/app/component/id"
	"github.com/datsukan/attendance-plan/backend/app/model"
	"github.com/datsukan/attendance-plan/backend/app/port"
	"github.com/datsukan/attendance-plan/backend/app/repository"
)

// SubjectInteractor は科目のユースケースの実装を表す構造体です。
type SubjectInteractor struct {
	Logger            *slog.Logger
	SubjectRepository repository.SubjectRepository
	OutputPort        port.SubjectOutputPort
}

// NewSubjectInteractor はSubjectInteractor を生成します。
func NewSubjectInteractor(logger *slog.Logger, subjectRepository repository.SubjectRepository, outputPort port.SubjectOutputPort) port.SubjectInputPort {
	return &SubjectInteractor{
		Logger:            logger,
		SubjectRepository: subjectRepository,
		OutputPort:        outputPort,
	}
}

// GetSubjectList は科目リストを取得します。
func (i *SubjectInteractor) GetSubjectList(inputData port.GetSubjectListInputData) {
	subjects, err := i.SubjectRepository.ReadByUserID(inputData.UserID)
	if err != nil {
		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
		i.OutputPort.SetResponseGetSubjectList(nil, r)
		return
	}

	var outputSubjects []port.BaseSubjectData
	for _, subject := range subjects {
		outputSubjects = append(outputSubjects, port.BaseSubjectData{
			ID:        subject.ID,
			UserID:    subject.UserID,
			Name:      subject.Name,
			Color:     subject.Color,
			CreatedAt: subject.CreatedAt.Format(time.DateTime),
			UpdatedAt: subject.UpdatedAt.Format(time.DateTime),
		})
	}

	o := &port.GetSubjectListOutputData{Subjects: outputSubjects}
	r := port.NewSuccessResult(http.StatusOK)
	i.OutputPort.SetResponseGetSubjectList(o, r)
}

// CreateSubject は科目を作成します。
func (i *SubjectInteractor) CreateSubject(inputData port.CreateSubjectInputData) {
	s := &model.Subject{
		ID:        id.NewID(),
		UserID:    inputData.UserID,
		Name:      inputData.Name,
		Color:     inputData.Color,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	i.Logger.With("subject_id", s.ID)

	if err := i.SubjectRepository.Create(s); err != nil {
		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
		i.OutputPort.SetResponseCreateSubject(nil, r)
		return
	}

	o := &port.CreateSubjectOutputData{
		Subject: port.BaseSubjectData{
			ID:        s.ID,
			UserID:    s.UserID,
			Name:      s.Name,
			Color:     s.Color,
			CreatedAt: s.CreatedAt.Format(time.DateTime),
			UpdatedAt: s.UpdatedAt.Format(time.DateTime),
		},
	}
	r := port.NewSuccessResult(http.StatusCreated)
	i.OutputPort.SetResponseCreateSubject(o, r)
}

// DeleteSubject は科目を削除します。
func (i *SubjectInteractor) DeleteSubject(inputData port.DeleteSubjectInputData) {
	i.Logger.With("subject_id", inputData.SubjectID)

	if err := i.SubjectRepository.Delete(inputData.SubjectID); err != nil {
		i.Logger.Error(err.Error())
		r := port.NewErrorResult(http.StatusInternalServerError, MsgInternalServerError)
		i.OutputPort.SetResponseDeleteSubject(nil, r)
		return
	}

	r := port.NewSuccessResult(http.StatusNoContent)
	i.OutputPort.SetResponseDeleteSubject(nil, r)
}
