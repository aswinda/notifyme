package repositories

import (
	"fmt"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/aswinda/notifyme/application/api/interfaces"
	"github.com/aswinda/notifyme/application/api/models"
)

type UserRepositoryWithCircuitBreaker struct {
	UserRepository interfaces.IUserRepository
}

type UserRepository struct {
	interfaces.IDbHandler
}

func (repository *UserRepositoryWithCircuitBreaker) GetUserDetail(userId int) (models.UserModel, error) {
	output := make(chan models.UserModel, 1)
	hystrix.ConfigureCommand("get_user_detail", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("get_user_detail", func() error {
		user, _ := repository.UserRepository.GetUserDetail(userId)

		output <- user
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return models.UserModel{}, err
	}
}

func (repository *UserRepository) GetUserDetail(userId int) (models.UserModel, error) {
	row, err := repository.Query(fmt.Sprintf("SELECT * FROM users WHERE id = '%d'", userId))

	if err != nil {
		return models.UserModel{}, err
	}

	var user models.UserModel

	row.Next()
	row.Scan(&user.Id, &user.Name, &user.Age)

	return user, nil
}
