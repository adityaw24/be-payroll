package controller

import (
	"errors"

	"github.com/dafiqarba/be-payroll/model"
	"github.com/dafiqarba/be-payroll/services"
	"github.com/dafiqarba/be-payroll/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PositionController interface {
	//Read Operation
	GetPositionList() fiber.Handler
	GetPositionDetail() fiber.Handler
	//Create Operation
	CreatePosition() fiber.Handler
	//Update Operation
	UpdatePosition() fiber.Handler
	//Delete Operation
	DeletePosition() fiber.Handler
}

type positionController struct {
	service services.PositionService
}

func NewPositionController(service services.PositionService) PositionController {
	return &positionController{
		service: service,
	}
}

func (controller *positionController) GetPositionList() fiber.Handler {
	return func(c *fiber.Ctx) error {
		list, err := controller.service.GetPositionList(c.Context())
		if err != nil {
			utils.BuildErrorResponse(c, fiber.StatusBadRequest, err.Error())
			return err
		}
		utils.BuildResponse(c, fiber.StatusOK, "success", list)
		return err
	}
}

func (controller *positionController) GetPositionDetail() fiber.Handler {
	return func(c *fiber.Ctx) error {
		v := c.Queries()
		id := uuid.MustParse(v["id"])
		positionDetail, err := controller.service.GetPositionDetail(c.Context(), id)
		if err != nil {
			errMsg := errors.New("the server cannot find the requested resource").Error()
			utils.BuildErrorResponse(c, fiber.StatusNotFound, errMsg)
			return err
		}
		utils.BuildResponse(c, fiber.StatusOK, "success", positionDetail)
		return err
	}
}

func (controller *positionController) CreatePosition() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var position model.Position
		err := c.BodyParser(&position)
		if err != nil {
			utils.BuildErrorResponse(c, fiber.StatusBadRequest, err.Error())
			return err
		}
		createdPosition, err := controller.service.CreatePosition(c.Context(), position)
		if err != nil {
			utils.BuildErrorResponse(c, fiber.StatusBadRequest, err.Error())
			return err
		}
		utils.BuildResponse(c, fiber.StatusOK, "success", createdPosition)
		return err
	}
}

func (controller *positionController) UpdatePosition() fiber.Handler {
	return func(c *fiber.Ctx) error {
		params_id := c.Params("id")
		id := uuid.MustParse(params_id)

		var position model.Position
		err := c.BodyParser(&position)
		if err != nil {
			utils.BuildErrorResponse(c, fiber.StatusBadRequest, err.Error())
			return err
		}

		position.Position_id = id
		updatedPosition, err := controller.service.UpdatePosition(c.Context(), position)
		if err != nil {
			utils.BuildErrorResponse(c, fiber.StatusBadRequest, err.Error())
			return err
		}
		utils.BuildResponse(c, fiber.StatusOK, "success", updatedPosition)
		return err
	}
}

func (controller *positionController) DeletePosition() fiber.Handler {
	return func(c *fiber.Ctx) error {
		params_id := c.Params("id")
		id := uuid.MustParse(params_id)

		idDeleted, err := controller.service.DeletePosition(c.Context(), id)
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
