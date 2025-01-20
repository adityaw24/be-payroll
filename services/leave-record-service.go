package services

import (
	"context"
	"strconv"
	"time"

	"github.com/dafiqarba/be-payroll/model"
	"github.com/dafiqarba/be-payroll/repository"
	"github.com/dafiqarba/be-payroll/utils"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type LeaveRecordService interface {
	//Read
	GetLeaveRecordDetail(ctx context.Context, req_id uuid.UUID, id uuid.UUID) (model.LeaveRecord, error)
	GetLeaveRecordList(ctx context.Context, id uuid.UUID, year string) ([]model.LeaveRecordListModel, error)
	//Insert
	CreateLeaveRecord(ctx context.Context, b model.CreateLeaveRecordModel) (uuid.UUID, error)
	//InsertUser(user model.User) (model.User, error)
}

type leaveRecordService struct {
	leaveRecordRepository repository.LeaveRecordRepo
	timeoutContext        time.Duration
	db                    *sqlx.DB
}

func NewLeaveRecordService(leaveRecordRepo repository.LeaveRecordRepo, timeoutContext time.Duration, db *sqlx.DB) LeaveRecordService {
	return &leaveRecordService{
		leaveRecordRepository: leaveRecordRepo,
		timeoutContext:        timeoutContext,
		db:                    db,
	}
}

func (service *leaveRecordService) GetLeaveRecordDetail(ctx context.Context, req_id uuid.UUID, id uuid.UUID) (model.LeaveRecord, error) {
	ctx, cancel := context.WithTimeout(ctx, service.timeoutContext)
	defer cancel()

	var (
		detail model.LeaveRecord
		err    error
	)

	detail, err = service.leaveRecordRepository.GetLeaveRecordDetail(ctx, req_id, id)
	if err != nil {
		utils.LogError("Services", "GetLeaveRecordDetail", err)
		return detail, err
	}
	return detail, err
}

func (service *leaveRecordService) GetLeaveRecordList(ctx context.Context, id uuid.UUID, year string) ([]model.LeaveRecordListModel, error) {
	ctx, cancel := context.WithTimeout(ctx, service.timeoutContext)
	defer cancel()

	var (
		list []model.LeaveRecordListModel
		err  error
	)

	if year == "ASC" || year == "DESC" {
		list, err = service.leaveRecordRepository.GetLeaveRecordList(ctx, id, year)
		return list, err
	}
	year = "ASC"

	list, err = service.leaveRecordRepository.GetLeaveRecordList(ctx, id, year)
	if err != nil {
		utils.LogError("Services", "GetLeaveRecordList", err)
		return list, err
	}
	return list, err
}

func (service *leaveRecordService) CreateLeaveRecord(ctx context.Context, b model.CreateLeaveRecordModel) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, service.timeoutContext)
	defer cancel()

	tx, err := service.db.Beginx()
	if err != nil {
		utils.LogError("Services", "CreateLeaveRecord open tx", err)
		return uuid.Nil, err
	}

	// Map model to model model
	leaveRecord := model.LeaveRecord{}
	leaveRecord.Request_on, _ = time.Parse(time.RFC3339, b.Request_on+"T00:00:00Z")
	leaveRecord.From_date, _ = time.Parse(time.RFC3339, b.From_date+"T00:00:00Z")
	leaveRecord.To_date, _ = time.Parse(time.RFC3339, b.To_date+"T00:00:00Z")
	leaveRecord.Return_date, _ = time.Parse(time.RFC3339, b.Return_date+"T00:00:00Z")
	leaveRecord.Amount, _ = strconv.Atoi(b.Amount)
	leaveRecord.Reason = b.Reason
	leaveRecord.Mobile = b.Mobile
	leaveRecord.Address = b.Address
	leaveRecord.Status_id = b.Status_id
	leaveRecord.Leave_id = b.Leave_id
	leaveRecord.User_id = b.User_id
	// Forward to repo
	record, err := service.leaveRecordRepository.CreateLeaveRecord(ctx, tx, leaveRecord)
	if err != nil {
		utils.LogError("Services", "CreateLeaveRecord", err)
		return uuid.Nil, err
	}

	utils.CommitOrRollback(tx, "Services UpdateLeaveBalance", err)
	return record, err
}
