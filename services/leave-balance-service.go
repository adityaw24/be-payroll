package services

import (
	"context"
	"errors"
	"time"

	"github.com/dafiqarba/be-payroll/model"
	"github.com/dafiqarba/be-payroll/repository"
	"github.com/dafiqarba/be-payroll/utils"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type LeaveBalanceService interface {
	//Read
	GetLeaveBalance(ctx context.Context, id uuid.UUID, year string) (model.LeaveBalance, error)
	//Insert
	//Update
	UpdateLeaveBalance(ctx context.Context, updatedData model.UpdateLeaveBalanceModel) (int, error)
}

type leaveBalanceService struct {
	leaveBalanceRepository repository.LeaveBalanceRepo
	timeoutContext         time.Duration
	db                     *sqlx.DB
}

func NewLeaveBalanceService(leaveBalanceRepo repository.LeaveBalanceRepo, timeoutContext time.Duration, db *sqlx.DB) LeaveBalanceService {
	return &leaveBalanceService{
		leaveBalanceRepository: leaveBalanceRepo,
		timeoutContext:         timeoutContext,
		db:                     db,
	}
}

func (service *leaveBalanceService) GetLeaveBalance(ctx context.Context, id uuid.UUID, year string) (model.LeaveBalance, error) {
	ctx, cancel := context.WithTimeout(ctx, service.timeoutContext)
	defer cancel()

	var (
		leaveBalance model.LeaveBalance
		err          error
	)

	leaveBalance, err = service.leaveBalanceRepository.GetLeaveBalance(ctx, id, year)
	if err != nil {
		utils.LogError("Services", "GetLeaveBalance", err)
		return leaveBalance, err
	}
	return leaveBalance, err
}

func (service *leaveBalanceService) UpdateLeaveBalance(ctx context.Context, updatedData model.UpdateLeaveBalanceModel) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, service.timeoutContext)
	defer cancel()

	var (
		updateId int
		err      error
	)

	tx, err := service.db.Beginx()
	if err != nil {
		utils.LogError("Services", "UpdateLeaveBalance open tx", err)
		return updateId, err
	}

	// Get leave Balance Data
	leaveBalance, err := service.leaveBalanceRepository.GetLeaveBalance(ctx, updatedData.User_id, updatedData.Year)
	if err != nil {
		utils.LogError("Services", "UpdateLeaveBalance get leave balance", err)
		utils.CommitOrRollback(tx, "Services GetLeaveBalance", err)
		return updateId, err
	}

	if leaveBalance == (model.LeaveBalance{}) {
		leaveBalance, err = service.leaveBalanceRepository.CreateLeaveBalance(ctx, tx, model.LeaveBalance{
			User_id:      updatedData.User_id,
			Leave_year:   updatedData.Year,
			Cuti_tahunan: 12,
			Cuti_diambil: 0,
			Cuti_balance: 12,
			Cuti_izin:    0,
			Cuti_sakit:   0,
		})

		if err != nil {
			utils.LogError("Services", "UpdateLeaveBalance create leave balance", err)
			utils.CommitOrRollback(tx, "Services UpdateLeaveBalance", err)
			return updateId, err
		}
	}

	if updatedData.Leave_id == 1 && (leaveBalance.Cuti_tahunan <= updatedData.Amounts) {
		err = errors.New("balance cuti tahunan tidak mencukupi")
		utils.CommitOrRollback(tx, "Services UpdateLeaveBalance", err)

		return updateId, err
	} else if updatedData.Leave_id == 1 && leaveBalance.Cuti_tahunan >= updatedData.Amounts {
		updatedData.Amounts = leaveBalance.Cuti_balance - updatedData.Amounts
		updatedData.Cuti_diambil = leaveBalance.Cuti_diambil + updatedData.Amounts

		updateId, err = service.leaveBalanceRepository.UpdateLeaveBalance(ctx, tx, updatedData, "cuti_balance")

		if err != nil {
			utils.LogError("Services", "UpdateLeaveBalance", err)
			utils.CommitOrRollback(tx, "Services UpdateLeaveBalance", err)
			return updateId, err
		}
	} else if updatedData.Leave_id == 2 {
		updatedData.Amounts = leaveBalance.Cuti_izin + updatedData.Amounts

		updateId, err = service.leaveBalanceRepository.UpdateLeaveBalance(ctx, tx, updatedData, "cuti_izin")

		if err != nil {
			utils.LogError("Services", "UpdateLeaveBalance", err)
			utils.CommitOrRollback(tx, "Services UpdateLeaveBalance", err)
			return updateId, err
		}
	} else if updatedData.Leave_id == 3 {
		updatedData.Amounts = leaveBalance.Cuti_sakit + updatedData.Amounts

		updateId, err = service.leaveBalanceRepository.UpdateLeaveBalance(ctx, tx, updatedData, "cuti_sakit")

		if err != nil {
			utils.LogError("Services", "UpdateLeaveBalance", err)
			utils.CommitOrRollback(tx, "Services UpdateLeaveBalance", err)
			return updateId, err
		}
	}

	utils.CommitOrRollback(tx, "Services UpdateLeaveBalance", err)
	return updateId, err
}
