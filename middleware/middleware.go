package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dafiqarba/be-payroll/services"
	"github.com/dafiqarba/be-payroll/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt"
)

type CustomMiddleware interface {
	LoggerMiddleware() fiber.Handler
	CORSMiddleware() fiber.Handler
	AuthorizeJWT() fiber.Handler
	MethodMiddleware() fiber.Handler
}

type customMiddleware struct {
	customJwt services.JWTService
}

func InitCustomMiddleware(customJwt services.JWTService) CustomMiddleware {
	return &customMiddleware{
		customJwt,
	}
}

func (m *customMiddleware) LoggerMiddleware() fiber.Handler {
	// return func(c *fiber.Ctx) error {

	log.Println("Opening log file...")
	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
		// return err
	}
	log.Println("Log file opened...")

	// defer file.Close()

	return logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${ip} ${method} ${path} => ${error}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Jakarta",
		Output:     file,
	})
	// return c.Next()
	// }
}

func (m *customMiddleware) MethodMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Method() == fiber.MethodHead {
			c.Status(fiber.StatusMethodNotAllowed)
			return nil
		}
		return c.Next()
	}
}

func (m *customMiddleware) CORSMiddleware() fiber.Handler {
	return cors.New(cors.Config{
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		AllowOrigins:     "*",
		AllowHeaders:     "Content-Type,Content-Length",
		AllowCredentials: false,
	})
	// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	// c.Writer.Header().Set("Access-Control-Allow-Headers", "X_Token,Content-Type,Content-Length")
	// c.Writer.Header().Set("Access-Control-Allow-Method", "POST, GET, DELETE, PUT, OPTIONS")

	// if c.Request.Method == "OPTIONS" {
	// 	c.AbortWithStatus(204)
	// 	return
	// }

	// c.Next()

}

// AuthorizeJWT function validate the token user given, return 401 if not valid
func (m *customMiddleware) AuthorizeJWT() fiber.Handler {
	return func(c *fiber.Ctx) error {

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			utils.BuildErrorResponse(c, http.StatusUnauthorized, "no token provided")
			return nil
		}
		token, err := m.customJwt.ValidateToken(authHeader)
		if err != nil {
			log.Println("| err: ", err)
			utils.BuildErrorResponse(c, http.StatusUnauthorized, "invalid token")
			return nil
		}
		if !token.Valid {
			log.Println("| err: ", err)
			utils.BuildErrorResponse(c, http.StatusUnauthorized, "invalid token")
			return nil
		}
		claims := token.Claims.(jwt.MapClaims)
		log.Println("| claims: ", claims)
		return nil
	}
}

// AuthorizeJWT function validate the token user given, return 401 if not valid
// func AuthorizeJWT(jwtService services.JWTService) func(*gin.Context) {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		if authHeader == "" {
// 			utils.BuildErrorResponse(c, http.StatusUnauthorized, "no token provided")
// 			return
// 		}
// 		token, err := jwtService.ValidateToken(authHeader)
// 		if err != nil {
// 			log.Println("| err: ", err)
// 			utils.BuildErrorResponse(c, http.StatusUnauthorized, "invalid token")
// 			return
// 		}
// 		if token.Valid {
// 			claims := token.Claims.(jwt.MapClaims)
// 			log.Println("| claims: ", claims)
// 		}
// 	}
// }

// func AuthorizeJWT(jwtService services.JWTService) func(http.Handler) http.Handler {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
// 			authHeader := request.Header.Get("Authorization")
// 			if authHeader == "" {
// 				utils.BuildErrorResponse(response, http.StatusBadRequest, "failed to process request. no token provided")
// 				return
// 			}
// 			token, err := jwtService.ValidateToken(authHeader, request)
// 			if err != nil {
// 				log.Println("| err: ", err)
// 				utils.BuildErrorResponse(response, http.StatusBadRequest, "invalid/expired token")
// 				return
// 			}
// 			if token.Valid {
// 				claims := token.Claims.(jwt.MapClaims)
// 				log.Println("| claims: ", claims)
// 				next.ServeHTTP(response, request)
// 			}
// 		})
// 	}

// }
