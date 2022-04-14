package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/dafiqarba/be-payroll/dto"
	"github.com/dafiqarba/be-payroll/services"
	"github.com/dafiqarba/be-payroll/utils"
)

type AuthController interface {
	//Read Operation
	Login(response http.ResponseWriter, request *http.Request)
	//Create Operation
	Register(response http.ResponseWriter, request *http.Request)
}

type authController struct {
	authServ services.AuthService
	jwtServ services.JWTService
	userServ services.UserService
}

func NewAuthController(authServ services.AuthService, jwtServ services.JWTService, userServ services.UserService) AuthController {
	return &authController{
		authServ: authServ,
		jwtServ: jwtServ,
		userServ: userServ,
	}
}

func (c *authController) Login(response http.ResponseWriter, request *http.Request) {
	var userLogin dto.UserLogin
	errDec := json.NewDecoder(request.Body).Decode(&userLogin)

	// Error handling
	if errDec != nil {
		response.WriteHeader(http.StatusBadRequest)
		utils.BuildErrorResponse(response, http.StatusBadRequest, errDec.Error())
		return
	}
	// Forwarding data to service
	var userLoginData, err = c.authServ.VerifyCredentials(userLogin)
	if err != nil {
		errMsg := errors.New("incorrect email/password").Error()
		utils.BuildErrorResponse(response, http.StatusUnauthorized, errMsg)
		return
	}
	token := c.jwtServ.GenerateToken(strconv.Itoa(userLoginData.User_id))
	userLoginData.Token = token

	utils.BuildResponse(response, http.StatusOK, "success", userLoginData);
}

func (c *authController) Register(response http.ResponseWriter, request *http.Request) {
	//Var that holds registered user data
	var regUser dto.RegisterUser
	//Retrieve data from JSON
	errDec := json.NewDecoder(request.Body).Decode(&regUser)
	if errDec != nil {
		utils.BuildErrorResponse(response, http.StatusBadRequest, errDec.Error())
		return
	}
	//Forwarding data to user service
	createdUser, err := c.userServ.CreateUser(regUser)
	if err != nil {
		errMsg := err.Error()
		utils.BuildErrorResponse(response, http.StatusConflict, errMsg)
		return
	}
	utils.BuildResponse(response, http.StatusOK, "success", createdUser);
}
