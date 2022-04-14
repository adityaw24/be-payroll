package services

import (
	"database/sql"
	"errors"
	"log"

	"github.com/dafiqarba/be-payroll/dto"
	"github.com/dafiqarba/be-payroll/entity"
	"github.com/dafiqarba/be-payroll/repository"
	"github.com/dafiqarba/be-payroll/utils"
)

type UserService interface {
	//Insert
	CreateUser(user dto.RegisterUser) (string, error)
	//Read
	GetUserList() ([]entity.User, error)
	GetUserDetail(id int) (entity.UserDetailModel, error)
}

type userService struct {
	userRepository repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) GetUserList() ([]entity.User, error) {
	return service.userRepository.GetUserList()
}

func (service *userService) GetUserDetail(id int) (entity.UserDetailModel, error) {
	return service.userRepository.GetUserDetail(id)
}

func (service *userService) CreateUser(d dto.RegisterUser) (string, error) {
	//Checks if the email is already registered by forwarding to FindByEmail repo
	_, err := service.userRepository.FindByEmail(d.Email)
	//If email is already registered, returns empty data and error
	if err == nil {
		log.Println("| Email already registered ")
		return "", errors.New("email address already registered")
	}
	//If error occured and the error is not because of no row returned, returns empty data and error
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return "", err
	}
	// If no error occured, map dto to entity.User
	//Hash plain password
	hashedPassword, err := utils.Hash(d.Password)
	if err != nil {
		log.Println("| Failed to hash a password " + err.Error())
	}
	registeredData := entity.User{
		Name:        d.Name,
		Username:    d.Username,
		Password:    hashedPassword,
		Email:       d.Email,
		Nik:         d.Nik,
		Role_id:     d.Role_id,
		Position_id: d.Position_id,
	}

	return service.userRepository.CreateUser(registeredData)
}
