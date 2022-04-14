package services

import (
	"errors"

	"github.com/dafiqarba/be-payroll/entity"
	"github.com/dafiqarba/be-payroll/repository"
)

type LeaveBalanceService interface {
	//Read
	GetLeaveBalance(id int, year string) (entity.LeaveBalance, error)
	//Insert
	//Update
	UpdateLeaveBalance(updatedData entity.UpdateLeaveBalanceModel) (int, error)
}

type leaveBalanceService struct {
	leaveBalanceRepository repository.LeaveBalanceRepo
}

func NewLeaveBalanceService(leaveBalanceRepo repository.LeaveBalanceRepo) LeaveBalanceService {
	return &leaveBalanceService{
		leaveBalanceRepository: leaveBalanceRepo,
	}
}

func (service *leaveBalanceService) GetLeaveBalance(id int, year string) (entity.LeaveBalance,error) {
	return service.leaveBalanceRepository.GetLeaveBalance(id, year)
}

func (service *leaveBalanceService) UpdateLeaveBalance(updatedData entity.UpdateLeaveBalanceModel) (int, error) {
	// Get leave Balance Data
	var leaveBalance, err = service.leaveBalanceRepository.GetLeaveBalance(updatedData.User_id, updatedData.Year)
	if err != nil {
		return 0, err
	}

	if updatedData.Leave_id == 1 && (leaveBalance.Cuti_tahunan <= updatedData.Amounts) {
		return 0, errors.New("balance cuti tahunan tidak mencukupi")
	} else if updatedData.Leave_id == 1 && leaveBalance.Cuti_tahunan >= updatedData.Amounts {
		updatedData.Amounts = leaveBalance.Cuti_balance - updatedData.Amounts
		updatedData.Cuti_diambil = leaveBalance.Cuti_diambil + updatedData.Amounts
		return service.leaveBalanceRepository.UpdateLeaveBalance(updatedData, "cuti_balance")
	} else if updatedData.Leave_id == 2 {
		updatedData.Amounts = leaveBalance.Cuti_izin + updatedData.Amounts
		return service.leaveBalanceRepository.UpdateLeaveBalance(updatedData, "cuti_izin")
	} else if updatedData.Leave_id == 3 {
		updatedData.Amounts = leaveBalance.Cuti_sakit + updatedData.Amounts
		return service.leaveBalanceRepository.UpdateLeaveBalance(updatedData, "cuti_sakit")
	}
	return 0, err
}