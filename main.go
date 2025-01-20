package main

import (
	"log"
	"time"

	"github.com/dafiqarba/be-payroll/controller"
	"github.com/dafiqarba/be-payroll/databases"
	"github.com/dafiqarba/be-payroll/middleware"
	"github.com/dafiqarba/be-payroll/repository"
	"github.com/dafiqarba/be-payroll/router"
	"github.com/dafiqarba/be-payroll/services"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var (
	dbRepoConn databases.DatabaseRepo = databases.NewPostgresRepo()
	db         *sqlx.DB
)

func init() {
	viper.SetConfigFile(`.env`)
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}

	dbHost := viper.GetString(`DB_HOST`)
	dbPort := viper.GetInt(`DB_PORT`)
	dbUser := viper.GetString(`DB_USER`)
	dbPass := viper.GetString(`DB_PASSWORD`)
	dbName := viper.GetString(`DB_NAME`)
	dbMigrateVersion := viper.GetUint(`DB_MIGRATE_VERSION`)
	runMigration := viper.GetBool(`DB_MIGRATE`)
	dbDriver := viper.GetString(`DB_DRIVER`)

	db, err = dbRepoConn.Connect(dbHost, dbPort, dbUser, dbPass, dbName, dbMigrateVersion, runMigration, dbDriver)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := fiber.New()

	apiVersion := viper.GetString(`API_VERSION`)
	appPort := viper.GetInt(`PORT`)
	secretKey := viper.GetString("JWT_SECRET")
	customJwt := services.NewJWTService(secretKey)
	timeoutCtx := time.Duration(viper.GetInt(`TIMEOUT_SECOND`)) * time.Second

	repoLeaveBalance := repository.NewLeaveBalanceRepo(db)
	repoLeaveRecord := repository.NewLeaveRecordRepo(db)
	repoPayrollRecord := repository.NewPayrollRecordRepo(db)
	repoUser := repository.NewUserRepo(db)
	repoPosition := repository.NewPositionRepo(db)
	repoRole := repository.NewRoleRepo(db)
	repoStatus := repository.NewStatusRepo(db)

	serviceAuth := services.NewAuthService(repoUser, timeoutCtx, db)
	serviceLeaveBalance := services.NewLeaveBalanceService(repoLeaveBalance, timeoutCtx, db)
	serviceLeaveRecord := services.NewLeaveRecordService(repoLeaveRecord, timeoutCtx, db)
	servicePayrollRecord := services.NewPayrollRecordService(repoPayrollRecord, timeoutCtx, db)
	serviceUser := services.NewUserService(repoUser, timeoutCtx, db)
	servicePosition := services.NewPositionService(repoPosition, timeoutCtx, db)
	serviceRole := services.NewRoleService(repoRole, timeoutCtx, db)
	serviceStatus := services.NewStatusService(repoStatus, timeoutCtx, db)

	controllerAuth := controller.NewAuthController(serviceAuth, customJwt, serviceUser)
	controllerLeaveBalance := controller.NewLeaveBalanceController(serviceLeaveBalance)
	controllerLeaveRecord := controller.NewLeaveRecordController(serviceLeaveRecord)
	controllerPayrollRecord := controller.NewPayrollRecordController(servicePayrollRecord)
	controllerUser := controller.NewUserController(serviceUser)
	controllerPosition := controller.NewPositionController(servicePosition)
	controllerRole := controller.NewRoleController(serviceRole)
	controllerStatus := controller.NewStatusController(serviceStatus)

	mw := middleware.InitCustomMiddleware(customJwt)

	httpRouter := router.NewFiberRouter(app)

	httpRouter.Use(mw.LoggerMiddleware(), mw.CORSMiddleware(), mw.MethodMiddleware())
	version := httpRouter.Group(apiVersion)

	version.Get("/", func(c *fiber.Ctx) error {
		log.Println(c.OriginalURL())
		return c.SendString("Hello, World!")
	}).Name("index")

	httpRouter.UserList(version, controllerUser)
	httpRouter.UserDetail(version, controllerUser)

	httpRouter.Login(version, controllerAuth)
	httpRouter.Register(version, controllerAuth)

	httpRouter.LeaveBalance(version, controllerLeaveBalance)
	httpRouter.LeaveBalanceUpdate(version, controllerLeaveBalance)
	httpRouter.LeaveRecordCreate(version, controllerLeaveRecord)
	httpRouter.LeaveRecordDetail(version, controllerLeaveRecord)
	httpRouter.LeaveRecordList(version, controllerLeaveRecord)

	httpRouter.PayrollCreate(version, controllerPayrollRecord)
	httpRouter.PayrollCreateList(version, controllerPayrollRecord)
	httpRouter.PayrollDetail(version, controllerPayrollRecord)
	httpRouter.PayrollList(version, controllerPayrollRecord)
	httpRouter.PayrollUpdate(version, controllerPayrollRecord)

	httpRouter.PositionList(version, controllerPosition)
	httpRouter.PositionCreate(version, controllerPosition)
	httpRouter.PositionUpdate(version, controllerPosition)
	httpRouter.PositionDelete(version, controllerPosition)
	httpRouter.PositionDetail(version, controllerPosition)

	httpRouter.RoleList(version, controllerRole)
	httpRouter.RoleCreate(version, controllerRole)
	httpRouter.RoleUpdate(version, controllerRole)
	httpRouter.RoleDelete(version, controllerRole)
	httpRouter.RoleDetail(version, controllerRole)

	httpRouter.StatusList(version, controllerStatus)
	httpRouter.StatusCreate(version, controllerStatus)
	httpRouter.StatusUpdate(version, controllerStatus)
	httpRouter.StatusDelete(version, controllerStatus)
	httpRouter.StatusDetail(version, controllerStatus)

	// data, _ := json.MarshalIndent(httpRouter.App().GetRoutes(true), "", "  ")
	// log.Println("routes: ", string(data))
	// log.Println("port: ", appPort)
	// log.Println("api version: ", version)

	httpRouter.Run(appPort, "be-payroll")

	// v1 := app.Group("/v1")

	// v1.Get("/user-list", controllerUser.GetUserList())
	// v1.Get("/user-detail", controllerUser.GetUserDetail())

	// v1.Post("/register", controllerAuth.Register())
	// v1.Post("/login", controllerAuth.Login())

	// v1.Get("/leave-balance", controllerLeaveBalance.GetLeaveBalance())
	// v1.Put("/update-leave-balance/{user_id:[0-9]+}", controllerLeaveBalance.UpdateLeaveBalance())

	// v1.Get("/leave-record-detail", controllerLeaveRecord.GetLeaveRecordDetail())
	// v1.Get("/leave-record-list", controllerLeaveRecord.GetLeaveRecordList())
	// v1.Post("/create-leave-record", controllerLeaveRecord.GetLeaveRecordList())

	// v1.Get("/payroll/list", controllerPayrollRecord.GetPayrollRecordList())
	// v1.Get("/payroll/detail/{id:[0-9]+}", controllerPayrollRecord.GetPayrollRecordDetail())
	// v1.Post("/payroll/create", controllerPayrollRecord.CreatePayrollRecord())
	// v1.Put("/payroll/update/{id:[0-9]+}", controllerPayrollRecord.UpdatePayrollRecord())
	// v1.Post("/payroll/create-list", controllerPayrollRecord.CreatePayrollRecordList())
}
