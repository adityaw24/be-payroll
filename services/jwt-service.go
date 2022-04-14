package services

import (
	"errors"
	"log"
	"net/http"
	"os"

	// "strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string, r *http.Request) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	return secretKey
}

func (j *jwtService) GenerateToken(userID string) string {
	claims := &jwtCustomClaim{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Issuer:    userID,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedTokenAsString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return signedTokenAsString
}

func (j *jwtService) ValidateToken(tokenString string, r *http.Request) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(parsedToken *jwt.Token) (interface{}, error) {
		if method, ok := parsedToken.Method.(*jwt.SigningMethodHMAC); !ok {
			err := errors.New("invalid signature method")
			log.Println("| err: ", err)
			return nil, err 
		} else if method != jwt.SigningMethodHS256 {
			err := errors.New("invalid signature method")
			log.Println("| err: ", err)
			return nil, err
		} else {
			return []byte(j.secretKey), nil
		}
	})
	//Parsing Token Error Handling
	if err != nil {
		log.Println("| error: ", err)
		return nil, errors.New("token invalid")
	}
	//Returns token
	return token, nil

	/* decoded with parseWithClaim

	claims := &jwtCustomClaim{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err
		}
		return nil, err
	} */
}

//

// type TokenDetail struct {
// 	AccessToken  string
// 	ExpiredToken int64
// }

// type AccessDetail struct {
// 	userID     int
// 	Authorized bool
// }

// func CreateToken(userId int) (*TokenDetail, error) {
// 	td := &TokenDetail{}
// 	td.ExpiredToken = time.Now().Add(time.Minute*15).Unix()
// 	var err error
// 	atClaims := jwt.MapClaims{}
// 	atClaims["authorized"] = true
// 	atClaims["user_id"] = userId
// 	atClaims["exp"] = td.ExpiredToken

// 	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
// 	td.AccessToken, err := at.SignedString([]byte(viper.GetString("Jwt.Secret")))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return td, nil
// }

// func ExtractToken(r *http.Request) string  {
// 	token := r.Header.Get("Authorization")
// 	strArr := strings.Split(token, " ")
// 	if len(strArr) == 2 {
// 		return strArr[1]
// 	}
// 	return ""
// }

// func VerifyToken(r *http.Request) (*jwt.Token, error)  {
// 	tokenString := ExtractToken(r)
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Wrong signature method")
// 		}
// 		return []byte(viper.GetString("Jwt.Secret")), nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	return token, nil
// }

// func TokenValid(r *http.Request) error {
// 	token, err := VerifyToken(r)
// 	if err != nil {
// 		return err
// 	}

// 	if _,ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
// 		return err
// 	}

// 	return nil
// }

// func ExtractTokenMetadata(r *http.Request) (*AccessDetail, error)  {
// 	token, err := VerifyToken(r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if ok && token.Valid {
// 		authorized, ok := claims["authorized"].(bool)
// 		if !ok {
// 			return nil, err
// 		}

// 		userId := int64(claims["user_id"].(float64))

// 		return &AccessDetail{
// 			Authorized: authorized,
// 			UserID:     userId,
// 		}, nil
// 	}

// 	return nil, err
// }
