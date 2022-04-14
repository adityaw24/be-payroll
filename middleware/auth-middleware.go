package middleware

import (
	"log"
	"net/http"

	"github.com/dafiqarba/be-payroll/services"
	"github.com/dafiqarba/be-payroll/utils"
	"github.com/golang-jwt/jwt"
)

//AuthorizeJWT function validate the token user given, return 401 if not valid
func AuthorizeJWT(jwtService services.JWTService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
			authHeader := request.Header.Get("Authorization")
			if authHeader == "" {
				utils.BuildErrorResponse(response, http.StatusBadRequest, "failed to process request. no token provided")
				return
			}
			token, err := jwtService.ValidateToken(authHeader, request)
			if err != nil {
				log.Println("| err: ", err)
				utils.BuildErrorResponse(response, http.StatusBadRequest, "invalid/expired token")
				return
			}
			if token.Valid {
				claims := token.Claims.(jwt.MapClaims)
				log.Println("| claims: ", claims)
				next.ServeHTTP(response, request)
			}
		})
	}

}
