package services

import (
	"github.com/dafiqarba/be-payroll/entity"
	"github.com/dafiqarba/be-payroll/repository"
)

type PayrollRecordService interface {
	//Read
	GetPayrollRecordList() ([]entity.PayrollRecordListModel, error)
	GetPayrollRecordDetail(id int) (entity.PayrollRecordDetailModel, error)
	//Update
	UpdatePayrollRecord(id int, p entity.PayrollRecord) (int, error)
	// UpdatePayrollRecord(id int, p entity.PayrollRecord) (entity.PayrollRecord, error)
	//Create
	// CreatePayrollRecord(p entity.PayrollRecord) (entity.PayrollRecord, error)
	CreatePayrollRecord(p entity.PayrollRecord) (int, error)
	// CreatePayrollRecordList(p []entity.PayrollRecord) ([]entity.PayrollRecord, error)
	CreatePayrollRecordList(p []entity.PayrollRecord) ([]int, error)
}

type payrollRecordService struct {
	payrollRecordRepo repository.PayrollRecordRepo
}

func NewPayrollRecordService(r repository.PayrollRecordRepo) PayrollRecordService {
	return &payrollRecordService{
		payrollRecordRepo: r,
	}
}

func (s *payrollRecordService) GetPayrollRecordList() ([]entity.PayrollRecordListModel, error) {
	return s.payrollRecordRepo.GetPayrollRecordList()
}

func (s *payrollRecordService) GetPayrollRecordDetail(id int) (entity.PayrollRecordDetailModel, error) {
	return s.payrollRecordRepo.GetPayrollRecordDetail(id)
}

// func (s *payrollRecordService) CreatePayrollRecord(p entity.PayrollRecord) (entity.PayrollRecord, error) {
// 	// payrollRecord := entity.PayrollRecord{}
// 	// payrollRecord.Payment_period = p.Payment_period
// 	// payrollRecord.Payment_date, _ = p.Payment_date, time.RFC822
// 	// payrollRecord.Basic_salary = p.Basic_salary
// 	// payrollRecord.Bpjs = p.Bpjs
// 	// payrollRecord.Tax = p.Tax
// 	// payrollRecord.Total_salary = p.Total_salary
// 	// payrollRecord.Status_id = p.Status_id
// 	// payrollRecord.User_id = p.User_id

// 	return s.payrollRecordRepo.CreatePayrollRecord(p)
// }

func (s *payrollRecordService) CreatePayrollRecord(p entity.PayrollRecord) (int, error) {
	// var payrollRecord = entity.PayrollRecord{}
	// payrollRecord.Payment_period = p.Payment_period
	// payrollRecord.Payment_date, _ = p.Payment_date, time.RFC822
	// payrollRecord.Basic_salary = p.Basic_salary
	// payrollRecord.Bpjs = p.Bpjs
	// payrollRecord.Tax = p.Tax
	// payrollRecord.Total_salary = p.Total_salary
	// payrollRecord.Status_id = p.Status_id
	// payrollRecord.User_id = p.User_id

	return s.payrollRecordRepo.CreatePayrollRecord(p)
}

// Go Routine for Form Update List Payroll
// func (s *payrollRecordService) CreatePayrollRecordList(p []entity.PayrollRecord) ([]int, error) {
// 	n := len(p) / 2
// 	channel := make(chan entity.PayrollRecord)

// 	go addList(0, n, p, channel, s)
// 	go addList(n, len(p), p, channel, s)

// 	var result []int
// 	for i := 0; i < len(p); i++ {
// 		payrollCreateList := <-channel
// 		result = append(result, payrollCreateList.Payroll_id)
// 	}

// 	return result, nil
// }

func (s *payrollRecordService) CreatePayrollRecordList(p []entity.PayrollRecord) ([]int, error) {
	n := len(p) / 2
	channel := make(chan int)

	go addList(0, n, p, channel, s)
	go addList(n, len(p), p, channel, s)

	var result []int
	for i := 0; i < len(p); i++ {
		payrollCreateList := <-channel
		result = append(result, payrollCreateList)
	}

	return result, nil
}

// func (s *payrollRecordService) CreatePayrollRecordList(p []entity.PayrollRecord) ([]entity.PayrollRecord error) {
// 	n := len(p) / 2
// 	channel := make(chan entity.PayrollRecord)

// 	go addList(0, n, p, channel, s)
// 	go addList(n, len(p), p, channel, s)

// 	var result []entity.PayrollRecord
// 	for i := 0; i < len(p); i++ {
// 		payrollCreateList := <-channel

// 		result = append(result, payrollCreateList)
// 	}

// 	return result, nil
// }

func (s *payrollRecordService) UpdatePayrollRecord(id int, p entity.PayrollRecord) (int, error) {
	// var payrollRecord = entity.PayrollRecord{}
	// payrollRecord.Payment_period = p.Payment_period
	// payrollRecord.Payment_date, _ = p.Payment_date, time.RFC822
	// payrollRecord.Basic_salary = p.Basic_salary
	// payrollRecord.Bpjs = p.Bpjs
	// payrollRecord.Tax = p.Tax
	// payrollRecord.Total_salary = p.Total_salary
	// payrollRecord.Status_id = p.Status_id
	// payrollRecord.User_id = p.User_id

	return s.payrollRecordRepo.UpdatePayrollRecord(id, p)
}

// func (s *payrollRecordService) UpdatePayrollRecord(id int, p entity.PayrollRecord) (entity.PayrollRecord, error) {
// 	// var payrollRecord = entity.PayrollRecord{}
// 	// payrollRecord.Payment_period = p.Payment_period
// 	// payrollRecord.Payment_date, _ = p.Payment_date, time.RFC822
// 	// payrollRecord.Basic_salary = p.Basic_salary
// 	// payrollRecord.Bpjs = p.Bpjs
// 	// payrollRecord.Tax = p.Tax
// 	// payrollRecord.Total_salary = p.Total_salary
// 	// payrollRecord.Status_id = p.Status_id
// 	// payrollRecord.User_id = p.User_id

// 	return s.payrollRecordRepo.UpdatePayrollRecord(id, p)
// }

// Go Routine for Form Create List Payroll
// func addList(start int, end int, createList []entity.PayrollRecord, channel chan entity.PayrollRecord, s *payrollRecordService) {
// 	for i := start; i < end; i++ {
// 		payrollList, _ := s.CreatePayrollRecord(createList[i])
// 		channel <- payrollList
// 	}
// }

func addList(start int, end int, createList []entity.PayrollRecord, channel chan int, s *payrollRecordService) {
	for i := start; i < end; i++ {
		payrollList, _ := s.CreatePayrollRecord(createList[i])
		channel <- payrollList
	}
}
