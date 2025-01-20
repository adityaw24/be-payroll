package router

import (
	"github.com/dafiqarba/be-payroll/controller"
	"github.com/gofiber/fiber/v2"
)

type StatusRouter interface {
	StatusList(group fiber.Router, controller controller.StatusController) fiber.Router
	StatusDetail(group fiber.Router, controller controller.StatusController) fiber.Router
	StatusCreate(group fiber.Router, controller controller.StatusController) fiber.Router
	StatusDelete(group fiber.Router, controller controller.StatusController) fiber.Router
	StatusUpdate(group fiber.Router, controller controller.StatusController) fiber.Router
}

func (r *fiberRouter) StatusList(group fiber.Router, controller controller.StatusController) fiber.Router {
	return group.Get("/status-list", controller.GetStatusList())
}

func (r *fiberRouter) StatusDetail(group fiber.Router, controller controller.StatusController) fiber.Router {
	return group.Get("/status/{id:[0-9]+}", controller.GetStatusDetail())
}

func (r *fiberRouter) StatusCreate(group fiber.Router, controller controller.StatusController) fiber.Router {
	return group.Post("/status", controller.CreateStatus())
}

func (r *fiberRouter) StatusDelete(group fiber.Router, controller controller.StatusController) fiber.Router {
	return group.Delete("/status/{id:[0-9]+}", controller.DeleteStatus())
}

func (r *fiberRouter) StatusUpdate(group fiber.Router, controller controller.StatusController) fiber.Router {
	return group.Put("/status/{id:[0-9]+}", controller.UpdateStatus())
}
