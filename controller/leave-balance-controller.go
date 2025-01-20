package controller

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dafiqarba/be-payroll/model"
	"github.com/dafiqarba/be-payroll/services"
	"github.com/dafiqarba/be-payroll/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type LeaveBalanceController interface {
	GetLeaveBalance() fiber.Handler
	UpdateLeaveBalance() fiber.Handler
}

type leaveBalanceController struct {
	leaveBalanceService services.LeaveBalanceService
}

func NewLeaveBalanceController(leaveBalanceServ services.LeaveBalanceService) LeaveBalanceController {
	return &leaveBalanceController{
		leaveBalanceService: leaveBalanceServ,
	}
}

func (c *leaveBalanceController) GetLeaveBalance() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := uuid.MustParse(ctx.Params("user_id"))
		year := ctx.Query("year")

		var leaveBalance, err = c.leaveBalanceService.GetLeaveBalance(ctx.Context(), id, year)
		if err != nil {
			errMsg := errors.New(" the server cannot find the requested resource").Error()
			utils.BuildErrorResponse(ctx, http.StatusNotFound, errMsg)
			return err
		}
		utils.BuildResponse(ctx, http.StatusOK, "success", leaveBalance)
		return err
	}
}

func (c *leaveBalanceController) UpdateLeaveBalance() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Take url param
		param_user_id := ctx.Params("user_id")
		if param_user_id == "" {
			err := errors.New("user_id is required")
			utils.BuildErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return err
		}
		//Updated data model
		var updatedData model.UpdateLeaveBalanceModel
		updatedData.User_id = uuid.MustParse(param_user_id)
		//Decode JSON body
		err := ctx.BodyParser(&updatedData)
		if err != nil {
			utils.BuildErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return err
		}
		// Forward data to service
		updatedAmounts, err := c.leaveBalanceService.UpdateLeaveBalance(ctx.Context(), updatedData)
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
			utils.BuildErrorResponse(ctx, httpStatus, errString)
			return err
		}
		// Serve results
		utils.BuildResponse(ctx, http.StatusOK, "successfully updated new value", updatedAmounts)
		return err
	}
}
