package usecase

import (
	"time"

	"github.com/datsukan/attendance-plan/backend/app/model"
	"github.com/datsukan/attendance-plan/backend/app/port"
	"github.com/datsukan/attendance-plan/backend/app/repository"
	"golang.org/x/crypto/bcrypt"
)

type stubScheduleRepository struct{}

func (r *stubScheduleRepository) ReadByUserID(userID string) ([]*model.Schedule, error) {
	date := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	schedules := []*model.Schedule{
		{ID: "test-id-1", UserID: "test-user-id-1", Name: "test-name-1", StartsAt: date, EndsAt: date, Color: "test-color", Type: model.ScheduleTypeMaster, CreatedAt: date, UpdatedAt: date},
		{ID: "test-id-2", UserID: "test-user-id-2", Name: "test-name-2", StartsAt: date, EndsAt: date, Color: "test-color", Type: model.ScheduleTypeCustom, CreatedAt: date, UpdatedAt: date},
	}
	return schedules, nil
}

func (r *stubScheduleRepository) Read(id string) (*model.Schedule, error) {
	date := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	schedule := &model.Schedule{
		ID:        "test-id",
		UserID:    "test-user-id",
		Name:      "test-name",
		StartsAt:  date,
		EndsAt:    date,
		Color:     "test-color",
		Type:      model.ScheduleTypeMaster,
		CreatedAt: date,
		UpdatedAt: date,
	}
	return schedule, nil
}

func (r *stubScheduleRepository) Create(schedule *model.Schedule) error {
	return nil
}

func (r *stubScheduleRepository) Update(schedule *model.Schedule) error {
	return nil
}

func (r *stubScheduleRepository) Delete(id string) error {
	return nil
}

func (r *stubScheduleRepository) Exists(id string) (bool, error) {
	return true, nil
}

type stubNotFoundScheduleRepository struct{}

func (r *stubNotFoundScheduleRepository) ReadByUserID(userID string) ([]*model.Schedule, error) {
	return nil, nil
}

func (r *stubNotFoundScheduleRepository) Read(id string) (*model.Schedule, error) {
	return nil, repository.NewNotFoundError()
}

func (r *stubNotFoundScheduleRepository) Create(schedule *model.Schedule) error {
	return nil
}

func (r *stubNotFoundScheduleRepository) Update(schedule *model.Schedule) error {
	return repository.NewNotFoundError()
}

func (r *stubNotFoundScheduleRepository) Delete(id string) error {
	return repository.NewNotFoundError()
}

func (r *stubNotFoundScheduleRepository) Exists(id string) (bool, error) {
	return false, nil
}

type stubScheduleOutputPort struct {
	Output interface{}
	Result port.Result
}

func (p *stubScheduleOutputPort) GetResponse() (int, string) {
	return p.Result.StatusCode, p.Result.Message
}

func (p *stubScheduleOutputPort) SetResponseGetScheduleList(output *port.GetScheduleListOutputData, result port.Result) {
	p.Output = output
	p.Result = result
}

func (p *stubScheduleOutputPort) SetResponseGetSchedule(output *port.GetScheduleOutputData, result port.Result) {
	p.Output = output
	p.Result = result
}

func (p *stubScheduleOutputPort) SetResponseCreateSchedule(output *port.CreateScheduleOutputData, result port.Result) {
	p.Output = output
	p.Result = result
}

func (p *stubScheduleOutputPort) SetResponseUpdateSchedule(output *port.UpdateScheduleOutputData, result port.Result) {
	p.Output = output
	p.Result = result
}

func (p *stubScheduleOutputPort) SetResponseDeleteSchedule(output *port.DeleteScheduleOutputData, result port.Result) {
	p.Output = output
	p.Result = result
}

type stubUserRepository struct{}

func (r *stubUserRepository) ReadByEmail(email string) (*model.User, error) {
	ph, _ := bcrypt.GenerateFromPassword([]byte("test-password"), bcrypt.DefaultCost)
	date := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)

	return &model.User{
		ID:        "test-id",
		Email:     "test-email@example.com",
		Name:      "test name",
		Password:  string(ph),
		CreatedAt: date,
		UpdatedAt: date,
	}, nil
}

func (r *stubUserRepository) Read(id string) (*model.User, error) {
	ph, _ := bcrypt.GenerateFromPassword([]byte("test-password"), bcrypt.DefaultCost)
	date := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)

	return &model.User{
		ID:        "test-id",
		Email:     "test-email@example.com",
		Name:      "test name",
		Password:  string(ph),
		CreatedAt: date,
		UpdatedAt: date,
	}, nil
}

func (r *stubUserRepository) Create(user *model.User) error {
	return nil
}

func (r *stubUserRepository) Update(user *model.User) error {
	return nil
}

func (r *stubUserRepository) Delete(id string) error {
	return nil
}

func (r *stubUserRepository) Exists(id string) (bool, error) {
	return true, nil
}

func (r *stubUserRepository) ExistsByEmail(email string) (bool, error) {
	return false, nil
}

type stubNotFoundUserRepository struct{}

func (r *stubNotFoundUserRepository) ReadByEmail(email string) (*model.User, error) {
	return nil, repository.NewNotFoundError()
}

func (r *stubNotFoundUserRepository) Read(id string) (*model.User, error) {
	return nil, repository.NewNotFoundError()
}

func (r *stubNotFoundUserRepository) Create(user *model.User) error {
	return nil
}

func (r *stubNotFoundUserRepository) Update(user *model.User) error {
	return nil
}

func (r *stubNotFoundUserRepository) Delete(id string) error {
	return nil
}

func (r *stubNotFoundUserRepository) Exists(id string) (bool, error) {
	return false, nil
}

func (r *stubNotFoundUserRepository) ExistsByEmail(email string) (bool, error) {
	return false, nil
}

type stubUserOutputPort struct {
	Output interface{}
	Result port.Result
}

func (p *stubUserOutputPort) GetResponse() (int, string) {
	return p.Result.StatusCode, p.Result.Message
}

func (p *stubUserOutputPort) SetResponseSignIn(output *port.SignInOutputData, result port.Result) {
	p.Output = output
	p.Result = result
}

func (p *stubUserOutputPort) SetResponseSignUp(output *port.SignUpOutputData, result port.Result) {
	p.Output = output
	p.Result = result
}

type stubSessionRepository struct{}

func (r *stubSessionRepository) ReadByUserID(userID string) (*model.Session, error) {
	return &model.Session{
		ID:        "test-id",
		UserID:    "test-user-id",
		ExpiresAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	}, nil
}

func (r *stubSessionRepository) Read(id string) (*model.Session, error) {
	return &model.Session{
		ID:        "test-id",
		UserID:    "test-user-id",
		ExpiresAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	}, nil
}

func (r *stubSessionRepository) Create(session *model.Session) error {
	return nil
}

func (r *stubSessionRepository) Update(session *model.Session) error {
	return nil
}

func (r *stubSessionRepository) Delete(id string) error {
	return nil
}
