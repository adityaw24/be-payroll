package services

import (
	"errors"
	"log"

	"github.com/dafiqarba/be-payroll/dto"
	"github.com/dafiqarba/be-payroll/entity"
	"github.com/dafiqarba/be-payroll/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredentials(d dto.UserLogin) (entity.UserLogin, error)
}

type authService struct {
	userRepo repository.UserRepo
}

func NewAuthService (userRepo repository.UserRepo) AuthService {
	return &authService {
		userRepo: userRepo,
	}
}

func (service *authService) VerifyCredentials(d dto.UserLogin) (entity.UserLogin, error) {
	//Retrieve matched entered email and password
	user, err := service.userRepo.FindByEmail(d.Email)
	if err != nil {
		log.Println("| "+err.Error())
			return user, err
	}
	//Compare password
	isValid := comparePassword(user.Password, d.Password)
	if !isValid {
		return user, errors.New("failed to login. check your credential")
	}
	//Return login user data
	return user, nil
}

// Compare plain hashed password retrieved from db against user-entered password
func comparePassword(hashedPass string, plainPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(plainPass))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		log.Println("| Incorrect password. Error "+err.Error())
		return false
	}
	log.Println("| Password Matched.")
	return true
}

