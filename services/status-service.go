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

type StatusService interface {
	//Insert
	CreateStatus(ctx context.Context, u model.Status) (uuid.UUID, error)
	//Read
	GetStatusList(ctx context.Context) ([]model.Status, error)
	GetStatusDetail(ctx context.Context, id uuid.UUID) (model.Status, error)
	//Update
	UpdateStatus(ctx context.Context, u model.Status) (uuid.UUID, error)
	//Delete
	DeleteStatus(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
}

type statusService struct {
	repository     repository.StatusRepo
	timeoutContext time.Duration
	db             *sqlx.DB
}

func NewStatusService(repository repository.StatusRepo, timeoutContext time.Duration, db *sqlx.DB) StatusService {
	return &statusService{
		repository:     repository,
		timeoutContext: timeoutContext,
		db:             db,
	}
}

func (s *statusService) CreateStatus(ctx context.Context, u model.Status) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeoutContext)
	defer cancel()

	var (
		id  uuid.UUID
		err error
	)

	tx, err := s.db.Beginx()
	if err != nil {
		utils.LogError("Services", "CreateStatus open tx", err)
		return id, err
	}

	id, err = s.repository.CreateStatus(ctx, tx, u)
	if err != nil {
		utils.LogError("Service", "CreateStatus", err)
		utils.CommitOrRollback(tx, "Services CreateStatus", err)
		return id, err
	}

	utils.CommitOrRollback(tx, "Services UpdateLeaveBalance", err)
	return id, err
}

func (s *statusService) GetStatusList(ctx context.Context) ([]model.Status, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeoutContext)
	defer cancel()

	var (
		list []model.Status
		err  error
	)

	list, err = s.repository.GetStatusList(ctx)
	if err != nil {
		utils.LogError("Services", "GetStatusList", err)
		return list, err
	}

	return list, err
}

func (s *statusService) GetStatusDetail(ctx context.Context, id uuid.UUID) (model.Status, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeoutContext)
	defer cancel()

	var (
		u   model.Status
		err error
	)

	u, err = s.repository.GetStatusDetail(ctx, id)
	if err != nil {
		utils.LogError("Services", "GetStatusDetail", err)
		return u, err
	}

	return u, err
}

func (s *statusService) UpdateStatus(ctx context.Context, u model.Status) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeoutContext)
	defer cancel()

	var (
		id  uuid.UUID
		err error
	)

	tx, err := s.db.Beginx()
	if err != nil {
		utils.LogError("Services", "UpdateStatus open tx", err)
		return id, err
	}

	id, err = s.repository.UpdateStatus(ctx, tx, u)
	if err != nil {
		utils.LogError("Service", "UpdateStatus", err)
		utils.CommitOrRollback(tx, "Services UpdateStatus", err)
		return id, err
	}

	utils.CommitOrRollback(tx, "Services UpdateLeaveBalance", err)
	return id, err
}

func (s *statusService) DeleteStatus(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeoutContext)
	defer cancel()

	var (
		err error
	)

	tx, err := s.db.Beginx()
	if err != nil {
		utils.LogError("Services", "DeleteStatus open tx", err)
		return id, err
	}

	id, err = s.repository.DeleteStatus(ctx, tx, id)
	if err != nil {
		utils.LogError("Service", "DeleteStatus", err)
		utils.CommitOrRollback(tx, "Services DeleteStatus", err)
		return id, err
	}

	utils.CommitOrRollback(tx, "Services UpdateLeaveBalance", err)
	return id, err
}
