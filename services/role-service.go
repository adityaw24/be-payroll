package services

import (
	"context"
	"time"

	"github.com/dafiqarba/be-payroll/model"
	"github.com/dafiqarba/be-payroll/repository"
	"github.com/dafiqarba/be-payroll/utils"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type RoleService interface {
	//Insert
	CreateRole(ctx context.Context, u model.Role) (uuid.UUID, error)
	//Read
	GetRoleList(ctx context.Context) ([]model.Role, error)
	GetRoleDetail(ctx context.Context, id uuid.UUID) (model.Role, error)
	//Update
	UpdateRole(ctx context.Context, u model.Role) (uuid.UUID, error)
	//Delete
	DeleteRole(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
}

type roleService struct {
	repository     repository.RoleRepo
	timeoutContext time.Duration
	db             *sqlx.DB
}

func NewRoleService(repository repository.RoleRepo, timeoutContext time.Duration, db *sqlx.DB) RoleService {
	return &roleService{
		repository:     repository,
		timeoutContext: timeoutContext,
		db:             db,
	}
}

func (service *roleService) GetRoleList(ctx context.Context) ([]model.Role, error) {
	ctx, cancel := context.WithTimeout(ctx, service.timeoutContext)
	defer cancel()

	var (
		list []model.Role
		err  error
	)

	list, err = service.repository.GetRoleList(ctx)
	if err != nil {
		utils.LogError("Services", "GetRoleList", err)
		return list, err
	}
	return list, err
}

func (service *roleService) GetRoleDetail(ctx context.Context, id uuid.UUID) (model.Role, error) {
	ctx, cancel := context.WithTimeout(ctx, service.timeoutContext)
	defer cancel()

	var (
		detail model.Role
		err    error
	)
	detail, err = service.repository.GetRoleDetail(ctx, id)
	if err != nil {
		utils.LogError("Services", "GetRoleDetail", err)
		return detail, err
	}
	return detail, err
}

func (service *roleService) CreateRole(ctx context.Context, u model.Role) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, service.timeoutContext)
	defer cancel()

	var (
		id  uuid.UUID
		err error
	)

	tx, err := service.db.Beginx()
	if err != nil {
		utils.LogError("Services", "CreateRole open tx", err)
		return id, err
	}

	// //Checks if the email is already registered by forwarding to FindByEmail repo
	// user, err := service.repository.FindByEmail(ctx, u.Email)

	// //If email is already registered, returns empty data and error
	// if user.Email != "" {
	// 	err = errors.New("email address already registered")
	// 	utils.LogError("Service", "CreateRole", err)
	// 	utils.CommitOrRollback(tx, "Services CreateRole", err)
	// 	return email, err
	// }

	id, err = service.repository.CreateRole(ctx, tx, u)
	if err != nil {
		utils.LogError("Service", "CreateRole", err)
		utils.CommitOrRollback(tx, "Services CreateRole", err)
		return id, err
	}

	utils.CommitOrRollback(tx, "Services UpdateLeaveBalance", err)
	return id, nil
}

func (service *roleService) UpdateRole(ctx context.Context, u model.Role) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, service.timeoutContext)
	defer cancel()

	var (
		id  uuid.UUID
		err error
	)

	tx, err := service.db.Beginx()
	if err != nil {
		utils.LogError("Services", "UpdateRole open tx", err)
		return id, err
	}

	id, err = service.repository.UpdateRole(ctx, tx, u)
	if err != nil {
		utils.LogError("Service", "UpdateRole", err)
		utils.CommitOrRollback(tx, "Services UpdateRole", err)
		return id, err
	}

	utils.CommitOrRollback(tx, "Services UpdateLeaveBalance", err)
	return id, nil
}

func (service *roleService) DeleteRole(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, service.timeoutContext)
	defer cancel()

	var (
		err error
	)

	tx, err := service.db.Beginx()
	if err != nil {
		utils.LogError("Services", "DeleteRole open tx", err)
		return id, err
	}

	id, err = service.repository.DeleteRole(ctx, tx, id)
	if err != nil {
		utils.LogError("Service", "DeleteRole", err)
		utils.CommitOrRollback(tx, "Services DeleteRole", err)
		return id, err
	}

	utils.CommitOrRollback(tx, "Services UpdateLeaveBalance", err)
	return id, nil
}
