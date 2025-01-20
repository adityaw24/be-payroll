package router

import (
	"github.com/dafiqarba/be-payroll/controller"
	"github.com/gofiber/fiber/v2"
)

type RoleRouter interface {
	RoleList(group fiber.Router, controller controller.RoleController) fiber.Router
	RoleDetail(group fiber.Router, controller controller.RoleController) fiber.Router
	RoleCreate(group fiber.Router, controller controller.RoleController) fiber.Router
	RoleDelete(group fiber.Router, controller controller.RoleController) fiber.Router
	RoleUpdate(group fiber.Router, controller controller.RoleController) fiber.Router
}

func (r *fiberRouter) RoleList(group fiber.Router, controller controller.RoleController) fiber.Router {
	return group.Get("/role-list", controller.GetRoleList())
}

func (r *fiberRouter) RoleDetail(group fiber.Router, controller controller.RoleController) fiber.Router {
	return group.Get("/role/{id:[0-9]+}", controller.GetRoleDetail())
}

func (r *fiberRouter) RoleCreate(group fiber.Router, controller controller.RoleController) fiber.Router {
	return group.Post("/role", controller.CreateRole())
}

func (r *fiberRouter) RoleDelete(group fiber.Router, controller controller.RoleController) fiber.Router {
	return group.Delete("/role/{id:[0-9]+}", controller.DeleteRole())
}

func (r *fiberRouter) RoleUpdate(group fiber.Router, controller controller.RoleController) fiber.Router {
	return group.Put("/role/{id:[0-9]+}", controller.UpdateRole())
}
