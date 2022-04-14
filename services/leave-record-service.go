package services

import (
	"strconv"
	"time"

	"github.com/dafiqarba/be-payroll/dto"
	"github.com/dafiqarba/be-payroll/entity"
	"github.com/dafiqarba/be-payroll/repository"
)

type LeaveRecordService interface {
	//Read
	GetLeaveRecordDetail(req_id int, id int) (entity.LeaveRecord, error)
	GetLeaveRecordList(id int, year string) ([]entity.LeaveRecordListModel, error)
	//Insert
	CreateLeaveRecord(b dto.CreateLeaveRecordModel) (int, error)
	//InsertUser(user entity.User) (entity.User, error)
}

type leaveRecordService struct {
	leaveRecordRepository repository.LeaveRecordRepo
}

func NewLeaveRecordService(leaveRecordRepo repository.LeaveRecordRepo) LeaveRecordService {
	return &leaveRecordService{
		leaveRecordRepository: leaveRecordRepo,
	}
}

func (service *leaveRecordService) GetLeaveRecordDetail(req_id int, id int) (entity.LeaveRecord, error) {
	return service.leaveRecordRepository.GetLeaveRecordDetail(req_id, id)
}

func (service *leaveRecordService) GetLeaveRecordList(id int, year string) ([]entity.LeaveRecordListModel, error) {
	if year == "ASC" || year == "DESC" {
		return service.leaveRecordRepository.GetLeaveRecordList(id, year)
	}
	year = "ASC"
	return service.leaveRecordRepository.GetLeaveRecordList(id, year)
}

func (service *leaveRecordService) CreateLeaveRecord(b dto.CreateLeaveRecordModel) (int, error) {
	// Map dto to entity model
	var leaveRecord = entity.LeaveRecord{}
	leaveRecord.Request_on, _ = time.Parse(time.RFC3339,b.Request_on+"T00:00:00Z")
	leaveRecord.From_date,_ = time.Parse(time.RFC3339,b.From_date+"T00:00:00Z")
	leaveRecord.To_date,_ = time.Parse(time.RFC3339,b.To_date+"T00:00:00Z") 
	leaveRecord.Return_date,_ = time.Parse(time.RFC3339,b.Return_date+"T00:00:00Z")
	leaveRecord.Amount,_ = strconv.Atoi(b.Amount) 
	leaveRecord.Reason = b.Reason
	leaveRecord.Mobile = b.Mobile
	leaveRecord.Address = b.Address
	leaveRecord.Status_id,_ = strconv.Atoi(b.Status_id) 
	leaveRecord.Leave_id,_ = strconv.Atoi(b.Leave_id)
	leaveRecord.User_id = b.User_id
	// Forward to repo
	return service.leaveRecordRepository.CreateLeaveRecord(leaveRecord)
}
