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

type LeaveRecordController interface {
	GetLeaveRecordDetail() fiber.Handler
	GetLeaveRecordList() fiber.Handler
	CreateLeaveRecord() fiber.Handler
}

type leaveRecordController struct {
	leaveRecordService services.LeaveRecordService
}

func NewLeaveRecordController(leaveRecordServ services.LeaveRecordService) LeaveRecordController {
	return &leaveRecordController{
		leaveRecordService: leaveRecordServ,
	}
}

func (c *leaveRecordController) GetLeaveRecordDetail() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		v := ctx.Queries()
		req_id := uuid.MustParse(v["req_id"])
		id := uuid.MustParse(v["id"])

		leaveRecordDetail, err := c.leaveRecordService.GetLeaveRecordDetail(ctx.Context(), req_id, id)

		if err != nil {
			errMsg := errors.New(" the server cannot find the requested resource").Error()
			utils.BuildErrorResponse(ctx, http.StatusNotFound, errMsg)
			return err
		}
		utils.BuildResponse(ctx, http.StatusOK, "success", leaveRecordDetail)
		return err
	}
}

func (c *leaveRecordController) GetLeaveRecordList() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		v := ctx.Queries()
		id := uuid.MustParse(v["id"])
		year := v["year"]

		leaveRecordList, err := c.leaveRecordService.GetLeaveRecordList(ctx.Context(), id, year)
		if err != nil {
			errMsg := errors.New("the server cannot find the requested resource").Error()
			utils.BuildErrorResponse(ctx, http.StatusNotFound, errMsg)
			return err
		}
		utils.BuildResponse(ctx, http.StatusOK, "success", leaveRecordList)
		return err
	}
}

func (c *leaveRecordController) CreateLeaveRecord() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Reference to CreateLeaveRecord data transfer obj
		var createLeaveRecord model.CreateLeaveRecordModel
		// Retrieve body obj from request
		err := ctx.BodyParser(&createLeaveRecord)
		// Error handling
		if err != nil {
			utils.BuildErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return err
		}
		// Forwarding data to service
		req_id, err := c.leaveRecordService.CreateLeaveRecord(ctx.Context(), createLeaveRecord)
		if err != nil {
			errMsg := errors.New("internal Server Error").Error()
			utils.BuildErrorResponse(ctx, http.StatusInternalServerError, errMsg)
			return err
		}
		utils.BuildResponse(ctx, http.StatusCreated, "new leave record created", req_id)
		return err
	}
}
