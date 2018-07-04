package interfaces

import (
	"github.com/aswinda/notifyme/application/api/models"
)

type IUserRepository interface {
	GetUserDetail(userId int) (models.UserModel, error)
}
