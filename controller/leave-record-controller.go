package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/dafiqarba/be-payroll/dto"
	"github.com/dafiqarba/be-payroll/services"
	"github.com/dafiqarba/be-payroll/utils"
)

type LeaveRecordController interface {
	GetLeaveRecordDetail(response http.ResponseWriter, request *http.Request)
	GetLeaveRecordList(response http.ResponseWriter, request *http.Request)
	CreateLeaveRecord(response http.ResponseWriter, request *http.Request)
}

type leaveRecordController struct {
	leaveRecordService services.LeaveRecordService
}

func NewLeaveRecordController(leaveRecordServ services.LeaveRecordService) LeaveRecordController {
	return &leaveRecordController{
		leaveRecordService: leaveRecordServ,
	}
}

func (c *leaveRecordController) GetLeaveRecordDetail(response http.ResponseWriter, request *http.Request) {
	v := request.URL.Query()
	req_id,_ := strconv.Atoi(v.Get("req_id"))
	id,_ := strconv.Atoi(v.Get("id"))

	
	var leaveRecordDetail, err = c.leaveRecordService.GetLeaveRecordDetail(req_id, id)

	if err != nil {
		errMsg := errors.New(" the server cannot find the requested resource").Error()
		utils.BuildErrorResponse(response, http.StatusNotFound, errMsg)
		return
	}
	utils.BuildResponse(response, http.StatusOK, "success", leaveRecordDetail)
}

func (c *leaveRecordController) GetLeaveRecordList(response http.ResponseWriter, request *http.Request) {
	v := request.URL.Query()
	id,_ := strconv.Atoi(v.Get("id"))
	year := v.Get("year")

	var leaveRecordList, err = c.leaveRecordService.GetLeaveRecordList(id, year)
	if err != nil {
		errMsg := errors.New("the server cannot find the requested resource").Error()
		utils.BuildErrorResponse(response, http.StatusNotFound, errMsg)
		return
	}
	utils.BuildResponse(response, http.StatusOK, "success", leaveRecordList )
}

func (c *leaveRecordController) CreateLeaveRecord(response http.ResponseWriter, request *http.Request) {
	// Reference to CreateLeaveRecord data transfer obj
	var createLeaveRecord dto.CreateLeaveRecordModel
	// Retrieve body obj from request
	errDec := json.NewDecoder(request.Body).Decode(&createLeaveRecord)
	// Error handling
	if errDec != nil {
		utils.BuildErrorResponse(response, http.StatusInternalServerError, errDec.Error())
		return
	}
	// Forwarding data to service
	var req_id, err = c.leaveRecordService.CreateLeaveRecord(createLeaveRecord)
	if err != nil {
		errMsg := errors.New("internal Server Error").Error()
		utils.BuildErrorResponse(response, http.StatusInternalServerError, errMsg)
		return 
	}
	utils.BuildInsertResponse(response, http.StatusCreated, "new leave record created", req_id)
}
