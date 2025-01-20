package router

import (
	"github.com/dafiqarba/be-payroll/controller"
	"github.com/gofiber/fiber/v2"
)

type PayrollRouter interface {
	PayrollList(group fiber.Router, controller controller.PayrollRecordController) fiber.Router
	PayrollDetail(group fiber.Router, controller controller.PayrollRecordController) fiber.Router
	PayrollCreate(group fiber.Router, controller controller.PayrollRecordController) fiber.Router
	PayrollUpdate(group fiber.Router, controller controller.PayrollRecordController) fiber.Router
	PayrollCreateList(group fiber.Router, controller controller.PayrollRecordController) fiber.Router
}

func (r *fiberRouter) PayrollList(group fiber.Router, controller controller.PayrollRecordController) fiber.Router {
	return group.Get("/payroll/list", controller.GetPayrollRecordList())
}

func (r *fiberRouter) PayrollDetail(group fiber.Router, controller controller.PayrollRecordController) fiber.Router {
	return group.Get("/payroll/detail/{id:[0-9]+}", controller.GetPayrollRecordDetail())
}

func (r *fiberRouter) PayrollCreate(group fiber.Router, controller controller.PayrollRecordController) fiber.Router {
	return group.Post("/payroll/create", controller.CreatePayrollRecord())
}

func (r *fiberRouter) PayrollUpdate(group fiber.Router, controller controller.PayrollRecordController) fiber.Router {
	return group.Put("/payroll/update/{id:[0-9]+}", controller.UpdatePayrollRecord())
}

func (r *fiberRouter) PayrollCreateList(group fiber.Router, controller controller.PayrollRecordController) fiber.Router {
	return group.Post("/payroll/create-list", controller.CreatePayrollRecordList())
}
