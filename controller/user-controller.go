package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/dafiqarba/be-payroll/services"
	"github.com/dafiqarba/be-payroll/utils"
)

type UserController interface {
	//Read Operation
	GetUserList(response http.ResponseWriter, request *http.Request)
	GetUserDetail(response http.ResponseWriter, request *http.Request)
}

type userController struct {
	userService services.UserService
}

func NewUserController(userServ services.UserService) UserController {
	return &userController{
		userService: userServ,
	}
}

func (c *userController) GetUserList(response http.ResponseWriter, request *http.Request) {

	var users, err = c.userService.GetUserList()
	if err != nil {
		utils.BuildErrorResponse(response, http.StatusInternalServerError, err.Error())
	}
	utils.BuildResponse(response, http.StatusOK, "success", users)
}

func (c *userController) GetUserDetail(response http.ResponseWriter, request *http.Request) {
	v := request.URL.Query()
	id,_ := strconv.Atoi(v.Get("id"))

	var userDetail, err = c.userService.GetUserDetail(id)
	if err != nil {
		errMsg := errors.New("the server cannot find the requested resource").Error()
		utils.BuildErrorResponse(response, http.StatusNotFound, errMsg)
		return
	}
	utils.BuildResponse(response, http.StatusOK, "success", userDetail)
}
