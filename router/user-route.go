package router

import (
	"github.com/dafiqarba/be-payroll/controller"
	"github.com/gofiber/fiber/v2"
)

type UserRouter interface {
	UserList(group fiber.Router, controller controller.UserController) fiber.Router
	UserDetail(group fiber.Router, controller controller.UserController) fiber.Router
}

func (r *fiberRouter) UserList(group fiber.Router, controller controller.UserController) fiber.Router {
	return group.Get("/user-list", controller.GetUserList())
}

func (r *fiberRouter) UserDetail(group fiber.Router, controller controller.UserController) fiber.Router {
	return group.Get("/user-detail/{id:[0-9]+}", controller.GetUserDetail())
}
