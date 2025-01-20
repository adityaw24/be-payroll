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

type PayrollRecordService interface {
	//Read
	GetPayrollRecordList(ctx context.Context) ([]model.PayrollRecordListModel, error)
	GetPayrollRecordDetail(ctx context.Context, id uuid.UUID) (model.PayrollRecordDetailModel, error)
	//Update
	UpdatePayrollRecord(ctx context.Context, id uuid.UUID, p model.PayrollRecord) (uuid.UUID, error)
	// UpdatePayrollRecord(ctx context.Context, id uuid.UUID, p model.PayrollRecord) (model.PayrollRecord, error)
	//Create
	// CreatePayrollRecord(ctx context.Context, p model.PayrollRecord) (model.PayrollRecord, error)
	CreatePayrollRecord(ctx context.Context, p model.PayrollRecord) (uuid.UUID, error)
	// CreatePayrollRecordList(ctx context.Context, p []model.PayrollRecord) ([]model.PayrollRecord, error)
	CreatePayrollRecordList(ctx context.Context, p []model.PayrollRecord) ([]uuid.UUID, error)
}

type payrollRecordService struct {
	payrollRecordRepo repository.PayrollRecordRepo
	timeoutContext    time.Duration
	db                *sqlx.DB
}

func NewPayrollRecordService(r repository.PayrollRecordRepo, timeoutContext time.Duration, db *sqlx.DB) PayrollRecordService {
	return &payrollRecordService{
		payrollRecordRepo: r,
		timeoutContext:    timeoutContext,
		db:                db,
	}
}

func (s *payrollRecordService) GetPayrollRecordList(ctx context.Context) ([]model.PayrollRecordListModel, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeoutContext)
	defer cancel()

	var (
		list []model.PayrollRecordListModel
		err  error
	)

	list, err = s.payrollRecordRepo.GetPayrollRecordList(ctx)

	if err != nil {
		utils.LogError("Services", "GetLeaveBalance", err)
		return list, err
	}
	return list, err
}

func (s *payrollRecordService) GetPayrollRecordDetail(ctx context.Context, id uuid.UUID) (model.PayrollRecordDetailModel, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeoutContext)
	defer cancel()

	var (
		detail model.PayrollRecordDetailModel
		err    error
	)

	detail, err = s.payrollRecordRepo.GetPayrollRecordDetail(ctx, id)

	if err != nil {
		utils.LogError("Services", "GetPayrollRecordDetail", err)
		return detail, err
	}
	return detail, err
}

// func (s *payrollRecordService) CreatePayrollRecord(ctx context.Context, p model.PayrollRecord) (model.PayrollRecord, error) {
// 	// payrollRecord := model.PayrollRecord{}
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

func (s *payrollRecordService) CreatePayrollRecord(ctx context.Context, p model.PayrollRecord) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeoutContext)
	defer cancel()

	var (
		id  uuid.UUID
		err error
	)

	tx, err := s.db.Beginx()
	if err != nil {
		utils.LogError("Services", "CreatePayrollRecord open tx", err)
		return id, err
	}

	// var payrollRecord = model.PayrollRecord{}
	// payrollRecord.Payment_period = p.Payment_period
	// payrollRecord.Payment_date, _ = p.Payment_date, time.RFC822
	// payrollRecord.Basic_salary = p.Basic_salary
	// payrollRecord.Bpjs = p.Bpjs
	// payrollRecord.Tax = p.Tax
	// payrollRecord.Total_salary = p.Total_salary
	// payrollRecord.Status_id = p.Status_id
	// payrollRecord.User_id = p.User_id

	id, err = s.payrollRecordRepo.CreatePayrollRecord(ctx, tx, p)

	if err != nil {
		utils.LogError("Services", "CreatePayrollRecord", err)
		utils.CommitOrRollback(tx, "Services CreatePayrollRecord", err)
		return id, err
	}

	utils.CommitOrRollback(tx, "Services UpdateLeaveBalance", err)
	return id, err
}

// Go Routine for Form Update List Payroll
// func (s *payrollRecordService) CreatePayrollRecordList(ctx context.Context, p []model.PayrollRecord) ([]uuid.UUID, error) {
// 	n := len(p) / 2
// 	channel := make(chan model.PayrollRecord)

// 	go addList(0, n, p, channel, s)
// 	go addList(n, len(p), p, channel, s)

// 	var result []int
// 	for i := 0; i < len(p); i++ {
// 		payrollCreateList := <-channel
// 		result = append(result, payrollCreateList.Payroll_id)
// 	}

// 	return result, nil
// }

func (s *payrollRecordService) CreatePayrollRecordList(ctx context.Context, p []model.PayrollRecord) ([]uuid.UUID, error) {

	var (
		result []uuid.UUID
		err    error
	)

	n := len(p) / 2
	channel := make(chan uuid.UUID)

	go addList(ctx, 0, n, p, channel, s)
	go addList(ctx, n, len(p), p, channel, s)

	for i := 0; i < len(p); i++ {
		payrollCreateList := <-channel
		result = append(result, payrollCreateList)
	}

	return result, err
}

// func (s *payrollRecordService) CreatePayrollRecordList(ctx context.Context, p []model.PayrollRecord) ([]model.PayrollRecord error) {
// 	n := len(p) / 2
// 	channel := make(chan model.PayrollRecord)

// 	go addList(0, n, p, channel, s)
// 	go addList(n, len(p), p, channel, s)

// 	var result []model.PayrollRecord
// 	for i := 0; i < len(p); i++ {
// 		payrollCreateList := <-channel

// 		result = append(result, payrollCreateList)
// 	}

// 	return result, nil
// }

func (s *payrollRecordService) UpdatePayrollRecord(ctx context.Context, id uuid.UUID, p model.PayrollRecord) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeoutContext)
	defer cancel()

	var (
		idResult uuid.UUID
		err      error
	)

	tx, err := s.db.Beginx()
	if err != nil {
		utils.LogError("Services", "UpdatePayrollRecord open tx", err)
		return idResult, err
	}
	// var payrollRecord = model.PayrollRecord{}
	// payrollRecord.Payment_period = p.Payment_period
	// payrollRecord.Payment_date, _ = p.Payment_date, time.RFC822
	// payrollRecord.Basic_salary = p.Basic_salary
	// payrollRecord.Bpjs = p.Bpjs
	// payrollRecord.Tax = p.Tax
	// payrollRecord.Total_salary = p.Total_salary
	// payrollRecord.Status_id = p.Status_id
	// payrollRecord.User_id = p.User_id

	idResult, err = s.payrollRecordRepo.UpdatePayrollRecord(ctx, tx, id, p)
	if err != nil {
		utils.LogError("Services", "UpdatePayrollRecord", err)
		utils.CommitOrRollback(tx, "Services UpdatePayrollRecord", err)
		return idResult, err
	}

	utils.CommitOrRollback(tx, "Services UpdateLeaveBalance", err)
	return idResult, err
}

// func (s *payrollRecordService) UpdatePayrollRecord(ctx context.Context, id uuid.UUID, p model.PayrollRecord) (model.PayrollRecord, error) {
// 	// var payrollRecord = model.PayrollRecord{}
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
// func addList(start int, end int, createList []model.PayrollRecord, channel chan model.PayrollRecord, s *payrollRecordService) {
// 	for i := start; i < end; i++ {
// 		payrollList, _ := s.CreatePayrollRecord(createList[i])
// 		channel <- payrollList
// 	}
// }

func addList(ctx context.Context, start int, end int, createList []model.PayrollRecord, channel chan uuid.UUID, s *payrollRecordService) {

	for i := start; i < end; i++ {
		payrollList, _ := s.CreatePayrollRecord(ctx, createList[i])
		channel <- payrollList
	}
}
