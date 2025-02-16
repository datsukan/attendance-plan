package usecase

import (
	"context"
	"time"

	"github.com/datsukan/attendance-plan/backend/app/model"
	"github.com/datsukan/attendance-plan/backend/app/port"
	"github.com/datsukan/attendance-plan/backend/app/repository"
	"golang.org/x/crypto/bcrypt"
)

type stubScheduleRepository struct{}

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
		Order:     1,
		CreatedAt: date,
		UpdatedAt: date,
	}
	return schedule, nil
}

func (r *stubScheduleRepository) ReadByUserID(userID string) ([]model.Schedule, error) {
	startDate1 := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	startDate2 := time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)
	startDate3 := time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)
	startDate4 := time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2021, 1, 10, 0, 0, 0, 0, time.UTC)

	schedules := []model.Schedule{
		{ID: "test-id-16", UserID: "test-user-id", Name: "test-name-16", StartsAt: startDate4, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeCustom, Order: 4, CreatedAt: startDate1, UpdatedAt: startDate1},
		{ID: "test-id-4", UserID: "test-user-id", Name: "test-name-4", StartsAt: startDate1, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeCustom, Order: 2, CreatedAt: startDate1, UpdatedAt: startDate1},
		{ID: "test-id-1", UserID: "test-user-id", Name: "test-name-1", StartsAt: startDate1, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeMaster, Order: 1, CreatedAt: startDate1, UpdatedAt: startDate1},
		{ID: "test-id-7", UserID: "test-user-id", Name: "test-name-7", StartsAt: startDate2, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeMaster, Order: 2, CreatedAt: startDate1, UpdatedAt: startDate1},
		{ID: "test-id-14", UserID: "test-user-id", Name: "test-name-14", StartsAt: startDate4, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeCustom, Order: 2, CreatedAt: startDate1, UpdatedAt: startDate1},
		{ID: "test-id-8", UserID: "test-user-id", Name: "test-name-8", StartsAt: startDate2, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeCustom, Order: 2, CreatedAt: startDate1, UpdatedAt: startDate1},
		{ID: "test-id-5", UserID: "test-user-id", Name: "test-name-5", StartsAt: startDate2, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeMaster, Order: 1, CreatedAt: startDate1, UpdatedAt: startDate1},
		{ID: "test-id-11", UserID: "test-user-id", Name: "test-name-11", StartsAt: startDate3, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeMaster, Order: 2, CreatedAt: startDate1, UpdatedAt: startDate1},
		{ID: "test-id-10", UserID: "test-user-id", Name: "test-name-10", StartsAt: startDate3, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeCustom, Order: 1, CreatedAt: startDate1, UpdatedAt: startDate1},
		{ID: "test-id-3", UserID: "test-user-id", Name: "test-name-3", StartsAt: startDate1, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeMaster, Order: 2, CreatedAt: startDate1, UpdatedAt: startDate1},
		{ID: "test-id-12", UserID: "test-user-id", Name: "test-name-12", StartsAt: startDate3, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeCustom, Order: 2, CreatedAt: startDate1, UpdatedAt: startDate1},
		{ID: "test-id-9", UserID: "test-user-id", Name: "test-name-9", StartsAt: startDate3, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeMaster, Order: 1, CreatedAt: startDate1, UpdatedAt: startDate1},
		{ID: "test-id-13", UserID: "test-user-id", Name: "test-name-13", StartsAt: startDate4, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeCustom, Order: 1, CreatedAt: startDate1, UpdatedAt: startDate1},
		{ID: "test-id-6", UserID: "test-user-id", Name: "test-name-6", StartsAt: startDate2, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeCustom, Order: 1, CreatedAt: startDate1, UpdatedAt: startDate1},
		{ID: "test-id-15", UserID: "test-user-id", Name: "test-name-15", StartsAt: startDate4, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeCustom, Order: 3, CreatedAt: startDate1, UpdatedAt: startDate1},
		{ID: "test-id-2", UserID: "test-user-id", Name: "test-name-2", StartsAt: startDate1, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeCustom, Order: 1, CreatedAt: startDate1, UpdatedAt: startDate1},
		{ID: "test-id-17", UserID: "test-user-id", Name: "test-name-17", StartsAt: startDate4, EndsAt: endDate, Color: "test-color", Type: model.ScheduleTypeCustom, Order: 5, CreatedAt: startDate1, UpdatedAt: startDate1},
	}
	return schedules, nil
}

func (r *stubScheduleRepository) ReadByUserIDStartsAt(userID string, startsAt time.Time) ([]model.Schedule, error) {
	date := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	schedules := []model.Schedule{
		{ID: "test-id-3", UserID: "test-user-id", Name: "test-name-3", StartsAt: date, EndsAt: date, Color: "test-color", Type: model.ScheduleTypeMaster, Order: 2, CreatedAt: date, UpdatedAt: date},
		{ID: "test-id-2", UserID: "test-user-id", Name: "test-name-2", StartsAt: date, EndsAt: date, Color: "test-color", Type: model.ScheduleTypeCustom, Order: 1, CreatedAt: date, UpdatedAt: date},
		{ID: "test-id-4", UserID: "test-user-id", Name: "test-name-4", StartsAt: date, EndsAt: date, Color: "test-color", Type: model.ScheduleTypeCustom, Order: 2, CreatedAt: date, UpdatedAt: date},
		{ID: "test-id-1", UserID: "test-user-id", Name: "test-name-1", StartsAt: date, EndsAt: date, Color: "test-color", Type: model.ScheduleTypeMaster, Order: 1, CreatedAt: date, UpdatedAt: date},
	}
	return schedules, nil
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

func (r *stubNotFoundScheduleRepository) Read(id string) (*model.Schedule, error) {
	return nil, repository.NewNotFoundError()
}

func (r *stubNotFoundScheduleRepository) Create(schedule *model.Schedule) error {
	return nil
}

func (r *stubNotFoundScheduleRepository) ReadByUserID(userID string) ([]model.Schedule, error) {
	return nil, nil
}

func (r *stubNotFoundScheduleRepository) ReadByUserIDStartsAt(userID string, startsAt time.Time) ([]model.Schedule, error) {
	return nil, nil
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

func (p *stubScheduleOutputPort) SetResponseUpdateBulkSchedule(output *port.UpdateBulkScheduleOutputData, result port.Result) {
	p.Output = output
	p.Result = result
}

func (p *stubScheduleOutputPort) SetResponseDeleteSchedule(output *port.DeleteScheduleOutputData, result port.Result) {
	p.Output = output
	p.Result = result
}

type stubUserRepository struct{}

func (r *stubUserRepository) ReadByEmail(email string, enabledOnly bool) (*model.User, error) {
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

func (r *stubUserRepository) Read(id string, enabledOnly bool) (*model.User, error) {
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

func (r *stubUserRepository) Exists(id string, enabledOnly bool) (bool, error) {
	return true, nil
}

func (r *stubUserRepository) ExistsByEmail(email string, enabledOnly bool) (bool, error) {
	return false, nil
}

type stubNotFoundUserRepository struct{}

func (r *stubNotFoundUserRepository) ReadByEmail(email string, enabledOnly bool) (*model.User, error) {
	return nil, repository.NewNotFoundError()
}

func (r *stubNotFoundUserRepository) Read(id string, enabledOnly bool) (*model.User, error) {
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

func (r *stubNotFoundUserRepository) Exists(id string, enabledOnly bool) (bool, error) {
	return false, nil
}

func (r *stubNotFoundUserRepository) ExistsByEmail(email string, enabledOnly bool) (bool, error) {
	return false, nil
}

type stubSessionRepository struct{}

func (r *stubSessionRepository) GenerateToken(userID string) (string, error) {
	return "test-token", nil
}

func (r *stubSessionRepository) IsValidToken(token string) (bool, string) {
	return true, "test-user-id"
}

type stubEmailRepository struct{}

func (r *stubEmailRepository) Send(ctx context.Context, to, subject, body string) (string, error) {
	return "test-message-id", nil
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

func (p *stubUserOutputPort) SetResponsePasswordSet(output *port.PasswordSetOutputData, result port.Result) {
	p.Output = output
	p.Result = result
}

func (p *stubUserOutputPort) SetResponsePasswordReset(output *port.PasswordResetOutputData, result port.Result) {
	p.Output = output
	p.Result = result
}
