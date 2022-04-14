package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/dafiqarba/be-payroll/config"
	"github.com/dafiqarba/be-payroll/controller"
	"github.com/dafiqarba/be-payroll/middleware"
	"github.com/dafiqarba/be-payroll/repository"
	"github.com/dafiqarba/be-payroll/services"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	// "github.com/jinzhu/gorm"
	// "gorm.io/gorm"
)

var (
	db *sql.DB = config.SetupDatabaseConnection()
	// gormDB *gorm.DB = config.InitGormDB(db)

	userRepo    repository.UserRepo       = repository.NewUserRepo(db)
	userSvc     services.UserService      = services.NewUserService(userRepo)
	userHandler controller.UserController = controller.NewUserController(userSvc)

	leaveBalanceRepo    repository.LeaveBalanceRepo       = repository.NewLeaveBalanceRepo(db)
	leaveBalanceSvc     services.LeaveBalanceService      = services.NewLeaveBalanceService(leaveBalanceRepo)
	leaveBalanceHandler controller.LeaveBalanceController = controller.NewLeaveBalanceController(leaveBalanceSvc)

	leaveRecordRepo    repository.LeaveRecordRepo       = repository.NewLeaveRecordRepo(db)
	leaveRecordSvc     services.LeaveRecordService      = services.NewLeaveRecordService(leaveRecordRepo)
	leaveRecordHandler controller.LeaveRecordController = controller.NewLeaveRecordController(leaveRecordSvc)

	payrollRecordRepo    repository.PayrollRecordRepo       = repository.NewPayrollRecordRepo(db)
	payrollRecordSvc     services.PayrollRecordService      = services.NewPayrollRecordService(payrollRecordRepo)
	payrollRecordHandler controller.PayrollRecordController = controller.NewPayrollRecordController(payrollRecordSvc)

	// adminApprovalRepo    repository.AdminApprovalRepo       = repository.NewApprovalRepo(gormDB)
	// adminApprovalSvc     services.AdminApprovalService      = services.NewAdminApprovalService(adminApprovalRepo)
	// adminApprovalHandler controller.AdminApprovalController = controller.NewAdminApprovalController(adminApprovalSvc)

	authService    services.AuthService      = services.NewAuthService(userRepo)
	jwtService     services.JWTService       = services.NewJWTService()
	authController controller.AuthController = controller.NewAuthController(authService, jwtService, userSvc)
)

func main() {
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}

	// Use gorilla mux
	router := mux.NewRouter()
	// CORS handlers
	headers := handlers.AllowedHeaders([]string{
		"X-Requested-With",
		"Content-Type",
		"Authorization",
		"Access-Control-Allow-Origin"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	credentials := handlers.AllowCredentials()

	/*-----------------------------------------------------------------------
	1.	Admin & User role
		Login					= /login
		Register				= /register

	2.	User Route:
		User Detail				= /user-detail?id=2
		Calc Leave Balance		= /update-leave-balance/:id
		Leave Balance Detail	= /leave-balance?id=1&year=2022
		Leave Record Detail		= /leave-record-detail?req_id=1&id=2
		Leave Record List		= /leave-record-list?id=2&year=ASC
		Create Leave Record		= /create-leave-record
		TO DO:
		b. payroll list

	TO DO:
	all route for admin
	 ------------------------------------------------------------------------*/
	protectR := router.Methods(http.MethodPost, http.MethodGet, http.MethodPut).Subrouter()
	protectR.HandleFunc("/user-list", userHandler.GetUserList).Methods(http.MethodGet)
	protectR.HandleFunc("/leave-balance", leaveBalanceHandler.GetLeaveBalance).Methods(http.MethodGet)
	protectR.HandleFunc("/update-leave-balance/{user_id:[0-9]+}", leaveBalanceHandler.UpdateLeaveBalance).Methods(http.MethodPut)
	protectR.HandleFunc("/leave-record-detail", leaveRecordHandler.GetLeaveRecordDetail).Methods(http.MethodGet)
	protectR.HandleFunc("/leave-record-list", leaveRecordHandler.GetLeaveRecordList).Methods(http.MethodGet)
	protectR.HandleFunc("/create-leave-record", leaveRecordHandler.CreateLeaveRecord).Methods(http.MethodPost)
	protectR.HandleFunc("/user-detail", userHandler.GetUserDetail).Methods(http.MethodGet)

	// protectR.HandleFunc("/admin-approval", adminApprovalHandler.GetAdminApprovalList).Methods(http.MethodGet)

	protectR.HandleFunc("/payroll/list", payrollRecordHandler.GetPayrollRecordList).Methods(http.MethodGet)
	protectR.HandleFunc("/payroll/detail/{id:[0-9]+}", payrollRecordHandler.GetPayrollRecordDetail).Methods(http.MethodGet)
	protectR.HandleFunc("/payroll/create", payrollRecordHandler.CreatePayrollRecord).Methods(http.MethodPost)
	protectR.HandleFunc("/payroll/update/{id:[0-9]+}", payrollRecordHandler.UpdatePayrollRecord).Methods(http.MethodPut)
	protectR.HandleFunc("/payroll/create-list", payrollRecordHandler.CreatePayrollRecordList).Methods(http.MethodPost)

	protectR.Use(middleware.AuthorizeJWT(jwtService))

	router.HandleFunc("/register", authController.Register).Methods(http.MethodPost)
	router.HandleFunc("/login", authController.Login).Methods(http.MethodPost)

	log.SetOutput(file)

	//Start server
	log.Println("| Server listening on port: 8000")
	log.Fatal(http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), handlers.CORS(headers, origins, methods, credentials)(router)))
}
