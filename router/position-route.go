package router

import (
	"github.com/dafiqarba/be-payroll/controller"
	"github.com/gofiber/fiber/v2"
)

type PositionRouter interface {
	PositionList(group fiber.Router, controller controller.PositionController) fiber.Router
	PositionDetail(group fiber.Router, controller controller.PositionController) fiber.Router
	PositionCreate(group fiber.Router, controller controller.PositionController) fiber.Router
	PositionDelete(group fiber.Router, controller controller.PositionController) fiber.Router
	PositionUpdate(group fiber.Router, controller controller.PositionController) fiber.Router
}

func (r *fiberRouter) PositionList(group fiber.Router, controller controller.PositionController) fiber.Router {
	return group.Get("/position-list", controller.GetPositionList())
}

func (r *fiberRouter) PositionDetail(group fiber.Router, controller controller.PositionController) fiber.Router {
	return group.Get("/position/{id:[0-9]+}", controller.GetPositionDetail())
}

func (r *fiberRouter) PositionCreate(group fiber.Router, controller controller.PositionController) fiber.Router {
	return group.Post("/position", controller.CreatePosition())
}

func (r *fiberRouter) PositionDelete(group fiber.Router, controller controller.PositionController) fiber.Router {
	return group.Delete("/position/{id:[0-9]+}", controller.DeletePosition())
}

func (r *fiberRouter) PositionUpdate(group fiber.Router, controller controller.PositionController) fiber.Router {
	return group.Put("/position/{id:[0-9]+}", controller.UpdatePosition())
}
