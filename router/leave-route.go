package router

import (
	"github.com/dafiqarba/be-payroll/controller"
	"github.com/gofiber/fiber/v2"
)

type LeaveRouter interface {
	LeaveBalance(group fiber.Router, controller controller.LeaveBalanceController) fiber.Router
	LeaveBalanceUpdate(group fiber.Router, controller controller.LeaveBalanceController) fiber.Router
	LeaveRecordList(group fiber.Router, controller controller.LeaveRecordController) fiber.Router
	LeaveRecordDetail(group fiber.Router, controller controller.LeaveRecordController) fiber.Router
	LeaveRecordCreate(group fiber.Router, controller controller.LeaveRecordController) fiber.Router
}

func (r *fiberRouter) LeaveBalance(group fiber.Router, controller controller.LeaveBalanceController) fiber.Router {
	return group.Get("/leave-balance", controller.GetLeaveBalance())
}

func (r *fiberRouter) LeaveBalanceUpdate(group fiber.Router, controller controller.LeaveBalanceController) fiber.Router {
	return group.Put("/update-leave-balance/{user_id:[0-9]+}", controller.UpdateLeaveBalance())
}

func (r *fiberRouter) LeaveRecordList(group fiber.Router, controller controller.LeaveRecordController) fiber.Router {
	return group.Get("/leave-record-list", controller.GetLeaveRecordList())
}

func (r *fiberRouter) LeaveRecordDetail(group fiber.Router, controller controller.LeaveRecordController) fiber.Router {
	return group.Get("/leave-record-detail", controller.GetLeaveRecordDetail())
}

func (r *fiberRouter) LeaveRecordCreate(group fiber.Router, controller controller.LeaveRecordController) fiber.Router {
	return group.Post("/create-leave-record", controller.CreateLeaveRecord())
}
