package services

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/dafiqarba/be-payroll/model"
	"github.com/dafiqarba/be-payroll/repository"
	"github.com/dafiqarba/be-payroll/utils"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserService interface {
	//Insert
	CreateUser(ctx context.Context, u model.RegisterUser) (string, error)
	//Read
	GetUserList(ctx context.Context) ([]model.User, error)
	GetUserDetail(ctx context.Context, id uuid.UUID) (model.UserDetailModel, error)
}

type userService struct {
	userRepository repository.UserRepo
	timeoutContext time.Duration
	db             *sqlx.DB
}

func NewUserService(userRepo repository.UserRepo, timeoutContext time.Duration, db *sqlx.DB) UserService {
	return &userService{
		userRepository: userRepo,
		timeoutContext: timeoutContext,
		db:             db,
	}
}

func (service *userService) GetUserList(ctx context.Context) ([]model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, service.timeoutContext)
	defer cancel()

	var (
		list []model.User
		err  error
	)

	list, err = service.userRepository.GetUserList(ctx)
	if err != nil {
		utils.LogError("Services", "GetUserList", err)
		return list, err
	}
	return list, err
}

func (service *userService) GetUserDetail(ctx context.Context, id uuid.UUID) (model.UserDetailModel, error) {
	ctx, cancel := context.WithTimeout(ctx, service.timeoutContext)
	defer cancel()

	var (
		detail model.UserDetailModel
		err    error
	)
	detail, err = service.userRepository.GetUserDetail(ctx, id)
	if err != nil {
		utils.LogError("Services", "GetUserDetail", err)
		return detail, err
	}
	return detail, err
}

func (service *userService) CreateUser(ctx context.Context, u model.RegisterUser) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, service.timeoutContext)
	defer cancel()

	var (
		email string
		err   error
	)

	tx, err := service.db.Beginx()
	if err != nil {
		utils.LogError("Services", "CreateUser open tx", err)
		return email, err
	}

	//Checks if the email is already registered by forwarding to FindByEmail repo
	user, err := service.userRepository.FindByEmail(ctx, u.Email)

	//If email is already registered, returns empty data and error
	if user.Email != "" {
		err = errors.New("email address already registered")
		utils.LogError("Service", "CreateUser", err)
		utils.CommitOrRollback(tx, "Services CreateUser", err)
		return email, err
	}

	//If error occured and the error is not because of no row returned, returns empty data and error
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		utils.LogError("Service", "CreateUser", err)
		utils.CommitOrRollback(tx, "Services CreateUser", err)
		return email, err
	}

	// If no error occured, map model to model.User
	//Hash plain password
	hashedPassword, err := utils.Hash(u.Password)
	if err != nil {
		utils.LogError("Service", "CreateUser Hash password", err)
		utils.CommitOrRollback(tx, "Services CreateUser hash password", err)
		return email, err
	}

	registeredData := model.User{
		Name:        u.Name,
		Username:    u.Username,
		Password:    hashedPassword,
		Email:       u.Email,
		Nik:         u.Nik,
		Role_id:     u.Role_id,
		Position_id: u.Position_id,
	}

	email, err = service.userRepository.CreateUser(ctx, tx, registeredData)
	if err != nil {
		utils.LogError("Service", "CreateUser", err)
		utils.CommitOrRollback(tx, "Services CreateUser", err)
		return email, err
	}
	return email, nil
}
