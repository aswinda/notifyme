package repositories

import (
	"fmt"
	"github.com/aswinda/notifyme/interfaces"
	"github.com/aswinda/notifyme/models"
)

func (repository *UserRepository) GetUserDetail(userId int) (models.UserModel, error) {
	row, err := repository.Query(fmt.Sprintf("
			SELECT
				*
			FROM
				users
			WHERE id = 	'%d'", userId
		))

		if err != nil {
			return models.UserModel{}, err
		}

		return row, nil
}