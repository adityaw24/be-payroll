package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/dafiqarba/be-payroll/entity"
	"github.com/dafiqarba/be-payroll/services"
	"github.com/dafiqarba/be-payroll/utils"
	"github.com/gorilla/mux"
)

type LeaveBalanceController interface {
	GetLeaveBalance(response http.ResponseWriter, request *http.Request)
	UpdateLeaveBalance(res http.ResponseWriter, req *http.Request)
}

type leaveBalanceController struct {
	leaveBalanceService services.LeaveBalanceService
}

func NewLeaveBalanceController(leaveBalanceServ services.LeaveBalanceService) LeaveBalanceController {
	return &leaveBalanceController{
		leaveBalanceService: leaveBalanceServ,
	}
}

func (c *leaveBalanceController) GetLeaveBalance(response http.ResponseWriter, request *http.Request) {
	v := request.URL.Query()
	id,_ := strconv.Atoi(v.Get("id"))
	year := v.Get("year")
	
	var leaveBalance, err = c.leaveBalanceService.GetLeaveBalance(id, year)
	if err != nil {
		errMsg := errors.New(" the server cannot find the requested resource").Error()
		utils.BuildErrorResponse(response, http.StatusNotFound, errMsg)
		return
	}
	utils.BuildResponse(response, http.StatusOK, "success", leaveBalance)
}

func (c *leaveBalanceController) UpdateLeaveBalance(res http.ResponseWriter, req *http.Request) {
	// Take url param
	params := mux.Vars(req)
	user_id, errConv := strconv.Atoi(params["user_id"])
	if errConv != nil {
		utils.BuildErrorResponse(res, http.StatusBadRequest, errConv.Error())
		return
	} 
	//Updated data model
	var updatedData entity.UpdateLeaveBalanceModel
	updatedData.User_id = user_id
	//Decode JSON body
	errDec := json.NewDecoder(req.Body).Decode(&updatedData)
	if errDec != nil{
		utils.BuildErrorResponse(res, http.StatusBadRequest, errDec.Error())
		return
	}
	// Forward data to service
	updatedAmounts, err := c.leaveBalanceService.UpdateLeaveBalance(updatedData)
	if err != nil {
		// convert err to str
		errString := err.Error()
		var httpStatus int
		if strings.Contains(errString, "no rows") {
			httpStatus = http.StatusNotFound
			errString = "the server cannot find the requested resource"
		} else if strings.Contains(errString, "tidak mencukupi") {
			httpStatus = http.StatusBadRequest
		}
		utils.BuildErrorResponse(res, httpStatus, errString)
		return
	}
	// Serve results
	utils.BuildUpdateResponse(res, http.StatusOK, "successfully updated new value", updatedAmounts)
}