package services

import (
	"errors"
	"time"

	"github.com/dafiqarba/be-payroll/utils"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type JWTService interface {
	GenerateToken(userID uuid.UUID) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
}

func NewJWTService(secretKey string) JWTService {
	return &jwtService{
		secretKey: secretKey,
	}
}

func (j *jwtService) GenerateToken(userID uuid.UUID) string {
	claims := &jwtCustomClaim{
		userID.String(),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Issuer:    userID.String(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signemodelkenAsString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return signemodelkenAsString
}

func (j *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(parsemodelken *jwt.Token) (interface{}, error) {
		if method, ok := parsemodelken.Method.(*jwt.SigningMethodHMAC); !ok {
			err := errors.New("invalid signature method")
			utils.LogError("Services", "ValidateToken", err)
			return nil, err
		} else if method != jwt.SigningMethodHS256 {
			err := errors.New("invalid signature method")
			utils.LogError("Services", "ValidateToken", err)
			return nil, err
		} else {
			return []byte(j.secretKey), nil
		}
	})
	//Parsing Token Error Handling
	if err != nil {
		err := errors.New("token invalid")
		utils.LogError("Services", "ValidateToken", err)
		return nil, err
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
// 	Expiremodelken int64
// }

// type AccessDetail struct {
// 	userID     int
// 	Authorized bool
// }

// func CreateToken(userId int) (*TokenDetail, error) {
// 	td := &TokenDetail{}
// 	td.Expiremodelken = time.Now().Add(time.Minute*15).Unix()
// 	var err error
// 	atClaims := jwt.MapClaims{}
// 	atClaims["authorized"] = true
// 	atClaims["user_id"] = userId
// 	atClaims["exp"] = td.Expiremodelken

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
