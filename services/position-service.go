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

type PositionService interface {
	//Insert
	CreatePosition(ctx context.Context, u model.Position) (uuid.UUID, error)
	//Read
	GetPositionList(ctx context.Context) ([]model.Position, error)
	GetPositionDetail(ctx context.Context, id uuid.UUID) (model.Position, error)
	//Update
	UpdatePosition(ctx context.Context, u model.Position) (uuid.UUID, error)
	//Delete
	DeletePosition(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
}

type positionService struct {
	repository     repository.PositionRepo
	timeoutContext time.Duration
	db             *sqlx.DB
}

func NewPositionService(repository repository.PositionRepo, timeoutContext time.Duration, db *sqlx.DB) PositionService {
	return &positionService{
		repository:     repository,
		timeoutContext: timeoutContext,
		db:             db,
	}
}

func (service *positionService) GetPositionList(ctx context.Context) ([]model.Position, error) {
	ctx, cancel := context.WithTimeout(ctx, service.timeoutContext)
	defer cancel()

	var (
		list []model.Position
		err  error
	)

	list, err = service.repository.GetPositionList(ctx)
	if err != nil {
		utils.LogError("Services", "GetPositionList", err)
		return list, err
	}

	return list, err
}

func (service *positionService) GetPositionDetail(ctx context.Context, id uuid.UUID) (model.Position, error) {
	ctx, cancel := context.WithTimeout(ctx, service.timeoutContext)
	defer cancel()

	var (
		position model.Position
		err      error
	)

	position, err = service.repository.GetPositionDetail(ctx, id)
	if err != nil {
		utils.LogError("Services", "GetPositionDetail", err)
		return position, err
	}

	return position, err
}

func (service *positionService) CreatePosition(ctx context.Context, u model.Position) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, service.timeoutContext)
	defer cancel()

	var (
		id  uuid.UUID
		err error
	)

	tx, err := service.db.Beginx()
	if err != nil {
		utils.LogError("Services", "CreatePosition open tx", err)
		return id, err
	}

	id, err = service.repository.CreatePosition(ctx, tx, u)
	if err != nil {
		utils.LogError("Service", "CreatePosition", err)
		utils.CommitOrRollback(tx, "Services CreatePosition", err)
		return id, err
	}

	utils.CommitOrRollback(tx, "Services UpdateLeaveBalance", err)
	return id, err
}

func (service *positionService) UpdatePosition(ctx context.Context, u model.Position) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, service.timeoutContext)
	defer cancel()

	var (
		id  uuid.UUID
		err error
	)

	tx, err := service.db.Beginx()
	if err != nil {
		utils.LogError("Services", "UpdatePosition open tx", err)
		return id, err
	}

	id, err = service.repository.UpdatePosition(ctx, tx, u)
	if err != nil {
		utils.LogError("Service", "UpdatePosition", err)
		utils.CommitOrRollback(tx, "Services UpdatePosition", err)
		return id, err
	}

	utils.CommitOrRollback(tx, "Services UpdateLeaveBalance", err)
	return id, err
}

func (service *positionService) DeletePosition(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, service.timeoutContext)
	defer cancel()

	var (
		err error
	)

	tx, err := service.db.Beginx()
	if err != nil {
		utils.LogError("Services", "DeletePosition open tx", err)
		return id, err
	}

	id, err = service.repository.DeletePosition(ctx, tx, id)
	if err != nil {
		utils.LogError("Service", "DeletePosition", err)
		utils.CommitOrRollback(tx, "Services DeletePosition", err)
		return id, err
	}

	utils.CommitOrRollback(tx, "Services UpdateLeaveBalance", err)
	return id, err
}
