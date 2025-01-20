package controller

import (
	"errors"

	"github.com/dafiqarba/be-payroll/model"
	"github.com/dafiqarba/be-payroll/services"
	"github.com/dafiqarba/be-payroll/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type StatusController interface {
	//Read Operation
	GetStatusList() fiber.Handler
	GetStatusDetail() fiber.Handler
	//Create Operation
	CreateStatus() fiber.Handler
	//Update Operation
	UpdateStatus() fiber.Handler
	//Delete Operation
	DeleteStatus() fiber.Handler
}

type statusController struct {
	service services.StatusService
}

func NewStatusController(service services.StatusService) StatusController {
	return &statusController{
		service: service,
	}
}

func (controller *statusController) GetStatusList() fiber.Handler {
	return func(c *fiber.Ctx) error {
		list, err := controller.service.GetStatusList(c.Context())
		if err != nil {
			utils.BuildErrorResponse(c, fiber.StatusBadRequest, err.Error())
			return err
		}
		utils.BuildResponse(c, fiber.StatusOK, "success", list)
		return err
	}
}

func (controller *statusController) GetStatusDetail() fiber.Handler {
	return func(c *fiber.Ctx) error {
		v := c.Queries()
		id := uuid.MustParse(v["id"])
		statusDetail, err := controller.service.GetStatusDetail(c.Context(), id)
		if err != nil {
			errMsg := errors.New("the server cannot find the requested resource").Error()
			utils.BuildErrorResponse(c, fiber.StatusNotFound, errMsg)
			return err
		}
		utils.BuildResponse(c, fiber.StatusOK, "success", statusDetail)
		return err
	}
}

func (controller *statusController) CreateStatus() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var status model.Status
		err := c.BodyParser(&status)
		if err != nil {
			utils.BuildErrorResponse(c, fiber.StatusBadRequest, err.Error())
			return err
		}
		createdStatus, err := controller.service.CreateStatus(c.Context(), status)
		if err != nil {
			utils.BuildErrorResponse(c, fiber.StatusBadRequest, err.Error())
			return err
		}
		utils.BuildResponse(c, fiber.StatusOK, "success", createdStatus)
		return err
	}
}

func (controller *statusController) UpdateStatus() fiber.Handler {
	return func(c *fiber.Ctx) error {
		params_id := c.Params("id")
		id := uuid.MustParse(params_id)

		var status model.Status
		err := c.BodyParser(&status)
		if err != nil {
			utils.BuildErrorResponse(c, fiber.StatusBadRequest, err.Error())
			return err
		}

		status.Status_id = id
		updatedStatus, err := controller.service.UpdateStatus(c.Context(), status)
		if err != nil {
			utils.BuildErrorResponse(c, fiber.StatusBadRequest, err.Error())
			return err
		}
		utils.BuildResponse(c, fiber.StatusOK, "success", updatedStatus)
		return err
	}
}

func (controller *statusController) DeleteStatus() fiber.Handler {
	return func(c *fiber.Ctx) error {
		params_id := c.Params("id")
		id := uuid.MustParse(params_id)
		idDeleted, err := controller.service.DeleteStatus(c.Context(), id)
		if err != nil {
			utils.BuildErrorResponse(c, fiber.StatusBadRequest, err.Error())
			return err
		}

		result := map[string]interface{}{
			"id": idDeleted,
		}
		utils.BuildResponse(c, fiber.StatusOK, "success", result)
		return err
	}
}
