package controller

import (
	"errors"

	"github.com/dafiqarba/be-payroll/model"
	"github.com/dafiqarba/be-payroll/services"
	"github.com/dafiqarba/be-payroll/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RoleController interface {
	//Read Operation
	GetRoleList() fiber.Handler
	GetRoleDetail() fiber.Handler
	//Create Operation
	CreateRole() fiber.Handler
	//Update Operation
	UpdateRole() fiber.Handler
	//Delete Operation
	DeleteRole() fiber.Handler
}

type roleController struct {
	service services.RoleService
}

func NewRoleController(service services.RoleService) RoleController {
	return &roleController{
		service: service,
	}
}

func (controller *roleController) GetRoleList() fiber.Handler {
	return func(c *fiber.Ctx) error {
		roles, err := controller.service.GetRoleList(c.Context())
		if err != nil {
			utils.LogError("Controller", "GetRoleList", err)
			utils.BuildErrorResponse(c, fiber.StatusInternalServerError, err.Error())
			return err
		}
		utils.BuildResponse(c, fiber.StatusOK, "success", roles)
		return err
	}
}

func (controller *roleController) GetRoleDetail() fiber.Handler {
	return func(c *fiber.Ctx) error {
		v := c.Queries()
		id := uuid.MustParse(v["id"])
		roleDetail, err := controller.service.GetRoleDetail(c.Context(), id)
		if err != nil {
			errMsg := errors.New("the server cannot find the requested resource").Error()
			utils.BuildErrorResponse(c, fiber.StatusNotFound, errMsg)
			return err
		}
		utils.BuildResponse(c, fiber.StatusOK, "success", roleDetail)
		return err
	}
}

func (controller *roleController) CreateRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var role model.Role
		err := c.BodyParser(&role)
		if err != nil {
			utils.BuildErrorResponse(c, fiber.StatusBadRequest, err.Error())
			return err
		}
		createdRole, err := controller.service.CreateRole(c.Context(), role)
		if err != nil {
			utils.BuildErrorResponse(c, fiber.StatusInternalServerError, err.Error())
			return err
		}
		utils.BuildResponse(c, fiber.StatusOK, "success", createdRole)
		return err
	}
}

func (controller *roleController) UpdateRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		params_id := c.Params("id")
		id := uuid.MustParse(params_id)

		var role model.Role

		err := c.BodyParser(&role)
		if err != nil {
			utils.BuildErrorResponse(c, fiber.StatusBadRequest, err.Error())
			return err
		}

		role.Role_id = id
		updatedRole, err := controller.service.UpdateRole(c.Context(), role)
		if err != nil {
			utils.BuildErrorResponse(c, fiber.StatusInternalServerError, err.Error())
			return err
		}
		utils.BuildResponse(c, fiber.StatusOK, "success", updatedRole)
		return err
	}
}

func (controller *roleController) DeleteRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		params_id := c.Params("id")
		id := uuid.MustParse(params_id)

		idDeleted, err := controller.service.DeleteRole(c.Context(), id)
		if err != nil {
			utils.BuildErrorResponse(c, fiber.StatusInternalServerError, err.Error())
			return err
		}

		result := map[string]interface{}{
			"id": idDeleted,
		}
		utils.BuildResponse(c, fiber.StatusOK, "success", result)
		return err
	}
}
