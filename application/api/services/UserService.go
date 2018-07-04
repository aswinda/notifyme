package services

import (
	"github.com/aswinda/notifyme/application/api/interfaces"
)

type UserService struct {
	interfaces.IUserRepository
}

func (service *UserService) GetUserDetail(userId int) (string, error) {
	result, err := service.GetUserDetail(userId)
	if err != nil {
		// handle error
	}

	return result, nil
}
