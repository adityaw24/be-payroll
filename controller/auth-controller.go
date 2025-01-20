package controller

import (
	"errors"
	"net/http"

	"github.com/dafiqarba/be-payroll/model"
	"github.com/dafiqarba/be-payroll/services"
	"github.com/dafiqarba/be-payroll/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	//Read Operation
	Login() fiber.Handler
	//Create Operation
	Register() fiber.Handler
}

type authController struct {
	authServ services.AuthService
	jwtServ  services.JWTService
	userServ services.UserService
}

func NewAuthController(authServ services.AuthService, jwtServ services.JWTService, userServ services.UserService) AuthController {
	return &authController{
		authServ: authServ,
		jwtServ:  jwtServ,
		userServ: userServ,
	}
}

func (c *authController) Login() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		userLoginData := model.UserResponse{}
		userLogin := model.UserLogin{}
		err := ctx.BodyParser(&userLogin)

		// Error handling
		if err != nil {
			utils.BuildErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return err
		}

		// Forwarding data to service
		userLoginData, err = c.authServ.VerifyCredentials(ctx.Context(), userLogin)
		if err != nil {
			errMsg := errors.New("incorrect email/password").Error()
			utils.BuildErrorResponse(ctx, http.StatusUnauthorized, errMsg)
			return err
		}
		token := c.jwtServ.GenerateToken(userLoginData.User_id)
		userLoginData.Token = token

		utils.BuildResponse(ctx, http.StatusOK, "success", userLoginData)
		return err
	}
}

func (c *authController) Register() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		//Var that holds registered user data
		var regUser model.RegisterUser
		var createdUser string
		//Retrieve data from JSON
		err := ctx.BodyParser(&regUser)
		if err != nil {
			utils.BuildErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return err
		}
		//Forwarding data to user service
		createdUser, err = c.userServ.CreateUser(ctx.Context(), regUser)
		if err != nil {
			errMsg := err.Error()
			utils.BuildErrorResponse(ctx, http.StatusConflict, errMsg)
			return err
		}
		utils.BuildResponse(ctx, http.StatusOK, "success", createdUser)
		return err
	}
}
