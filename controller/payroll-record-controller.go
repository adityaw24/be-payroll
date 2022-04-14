package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/dafiqarba/be-payroll/entity"
	"github.com/dafiqarba/be-payroll/services"
	"github.com/dafiqarba/be-payroll/utils"
	"github.com/gorilla/mux"
)

type PayrollRecordController interface {
	GetPayrollRecordList(response http.ResponseWriter, request *http.Request)
	GetPayrollRecordDetail(response http.ResponseWriter, request *http.Request)
	CreatePayrollRecord(response http.ResponseWriter, request *http.Request)
	CreatePayrollRecordList(response http.ResponseWriter, request *http.Request)
	UpdatePayrollRecord(response http.ResponseWriter, request *http.Request)
}

type payrollRecordController struct {
	payrollRecordService services.PayrollRecordService
}

func NewPayrollRecordController(payrollRecordServ services.PayrollRecordService) PayrollRecordController {
	return &payrollRecordController{
		payrollRecordService: payrollRecordServ,
	}
}

func (c *payrollRecordController) GetPayrollRecordList(response http.ResponseWriter, request *http.Request) {
	// v := request.URL.Query()
	// id, _ := strconv.Atoi(v.Get("user_id"))
	// year := v.Get("year")

	var payrollRecordList, err = c.payrollRecordService.GetPayrollRecordList()
	if err != nil {
		utils.BuildErrorResponse(response, http.StatusNotFound, err.Error())
		return
	}
	utils.BuildResponse(response, http.StatusOK, "success", payrollRecordList)
}

func (c *payrollRecordController) GetPayrollRecordDetail(response http.ResponseWriter, request *http.Request) {
	// v := request.URL.Query()
	// // pay_id, _ := strconv.Atoi(v.Get("pay_id"))
	// id, _ := strconv.Atoi(v.Get("user_id"))

	// id, _ := strconv.Atoi(request.URL.Query().Get("user-id"))
	// id, _ := strconv.ParseInt(request.FormValue("user-id"), 10, 64)
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])

	// params := request.URL.Query().Get("user-id")
	// id, _ := strconv.Atoi(params)

	var payrollRecordDetail, err = c.payrollRecordService.GetPayrollRecordDetail(id)

	if err != nil {
		utils.BuildErrorResponse(response, http.StatusNotFound, err.Error())
		return
	}
	utils.BuildResponse(response, http.StatusOK, "success", payrollRecordDetail)
}

func (c *payrollRecordController) CreatePayrollRecord(response http.ResponseWriter, request *http.Request) {
	var payrollRecord entity.PayrollRecord
	// var payrollRecord dto.CreatePayrollRecordModel
	errDec := json.NewDecoder(request.Body).Decode(&payrollRecord)
	if errDec != nil {
		utils.BuildErrorResponse(response, http.StatusInternalServerError, errDec.Error())
		return
	}

	newPayrollRecord, err := c.payrollRecordService.CreatePayrollRecord(payrollRecord)
	if err != nil {
		errMsg := errors.New("internal Server Error").Error()
		utils.BuildErrorResponse(response, http.StatusInternalServerError, errMsg)
		return
	}
	utils.BuildInsertResponse(response, http.StatusCreated, "success created", newPayrollRecord)
}

// Go routine to create payroll record list
func (c *payrollRecordController) CreatePayrollRecordList(response http.ResponseWriter, request *http.Request) {
	var payrollRecordList []entity.PayrollRecord
	errDec := json.NewDecoder(request.Body).Decode(&payrollRecordList)
	if errDec != nil {
		utils.BuildErrorResponse(response, http.StatusInternalServerError, errDec.Error())
		return
	}

	newPayrollRecordList, err := c.payrollRecordService.CreatePayrollRecordList(payrollRecordList)
	if err != nil {
		errMsg := errors.New("internal Server Error").Error()
		utils.BuildErrorResponse(response, http.StatusInternalServerError, errMsg)
		return
	}
	utils.BuildInsertResponse(response, http.StatusCreated, "success created", newPayrollRecordList)
}

func (c *payrollRecordController) UpdatePayrollRecord(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])

	var payrollRecord entity.PayrollRecord
	errDec := json.NewDecoder(request.Body).Decode(&payrollRecord)
	if errDec != nil {
		utils.BuildErrorResponse(response, http.StatusNotFound, errDec.Error())
		return
	}

	updatedPayrollRecord, err := c.payrollRecordService.UpdatePayrollRecord(id, payrollRecord)
	if err != nil {
		errMsg := errors.New("internal Server Error").Error()
		utils.BuildErrorResponse(response, http.StatusNotFound, errMsg)
		return
	}
	utils.BuildResponse(response, http.StatusOK, "success updated", updatedPayrollRecord)
}
