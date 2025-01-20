package controller

import (
	"errors"
	"net/http"

	"github.com/dafiqarba/be-payroll/model"
	"github.com/dafiqarba/be-payroll/services"
	"github.com/dafiqarba/be-payroll/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PayrollRecordController interface {
	GetPayrollRecordList() fiber.Handler
	GetPayrollRecordDetail() fiber.Handler
	CreatePayrollRecord() fiber.Handler
	CreatePayrollRecordList() fiber.Handler
	UpdatePayrollRecord() fiber.Handler
}

type payrollRecordController struct {
	payrollRecordService services.PayrollRecordService
}

func NewPayrollRecordController(payrollRecordServ services.PayrollRecordService) PayrollRecordController {
	return &payrollRecordController{
		payrollRecordService: payrollRecordServ,
	}
}

func (c *payrollRecordController) GetPayrollRecordList() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		// v := request.URL.Query()
		// id, _ := strconv.Atoi(v.Get("user_id"))
		// year := v.Get("year")

		payrollRecordList, err := c.payrollRecordService.GetPayrollRecordList(ctx.Context())
		if err != nil {
			utils.BuildErrorResponse(ctx, http.StatusNotFound, err.Error())
			return err
		}
		utils.BuildResponse(ctx, http.StatusOK, "success", payrollRecordList)
		return err
	}
}

func (c *payrollRecordController) GetPayrollRecordDetail() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		// v := request.URL.Query()
		// // pay_id, _ := strconv.Atoi(v.Get("pay_id"))
		// id, _ := strconv.Atoi(v.Get("user_id"))

		// id, _ := strconv.Atoi(request.URL.Query().Get("user-id"))
		// id, _ := strconv.ParseInt(request.FormValue("user-id"), 10, 64)
		params_id := ctx.Params("id")
		id := uuid.MustParse(params_id)

		// params := request.URL.Query().Get("user-id")
		// id, _ := strconv.Atoi(params)

		payrollRecordDetail, err := c.payrollRecordService.GetPayrollRecordDetail(ctx.Context(), id)

		if err != nil {
			utils.BuildErrorResponse(ctx, http.StatusNotFound, err.Error())
			return err
		}
		utils.BuildResponse(ctx, http.StatusOK, "success", payrollRecordDetail)
		return err
	}
}

func (c *payrollRecordController) CreatePayrollRecord() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var payrollRecord model.PayrollRecord
		// var payrollRecord model.CreatePayrollRecordModel
		err := ctx.BodyParser(&payrollRecord)
		if err != nil {
			utils.BuildErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return err
		}

		newPayrollRecord, err := c.payrollRecordService.CreatePayrollRecord(ctx.Context(), payrollRecord)
		if err != nil {
			errMsg := errors.New("internal Server Error").Error()
			utils.BuildErrorResponse(ctx, http.StatusInternalServerError, errMsg)
			return err
		}
		utils.BuildResponse(ctx, http.StatusCreated, "success created", newPayrollRecord)
		return err
	}
}

// Go routine to create payroll record list
func (c *payrollRecordController) CreatePayrollRecordList() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var payrollRecordList []model.PayrollRecord
		err := ctx.BodyParser(&payrollRecordList)
		if err != nil {
			utils.BuildErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return err
		}

		newPayrollRecordList, err := c.payrollRecordService.CreatePayrollRecordList(ctx.Context(), payrollRecordList)
		if err != nil {
			errMsg := errors.New("internal Server Error").Error()
			utils.BuildErrorResponse(ctx, http.StatusInternalServerError, errMsg)
			return err
		}
		utils.BuildResponse(ctx, http.StatusCreated, "success created", newPayrollRecordList)
		return err
	}
}

func (c *payrollRecordController) UpdatePayrollRecord() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		params_id := ctx.Params("id")
		id := uuid.MustParse(params_id)

		var payrollRecord model.PayrollRecord
		err := ctx.BodyParser(&payrollRecord)
		if err != nil {
			utils.BuildErrorResponse(ctx, http.StatusNotFound, err.Error())
			return err
		}

		updatedPayrollRecord, err := c.payrollRecordService.UpdatePayrollRecord(ctx.Context(), id, payrollRecord)
		if err != nil {
			errMsg := errors.New("internal Server Error").Error()
			utils.BuildErrorResponse(ctx, http.StatusNotFound, errMsg)
			return err
		}
		utils.BuildResponse(ctx, http.StatusOK, "success updated", updatedPayrollRecord)
		return err
	}
}
