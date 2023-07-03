package services

import (
	"github.com/dami-pie/napi/models"
	"github.com/dami-pie/napi/src/database"
	"github.com/dami-pie/napi/src/repositories"
)

func CreateUser(user models.User) (uint64, error) {
	db, erro := database.Connect()
	if erro != nil {
		return 0, erro
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	userID, erro := repository.Create(user)
	if erro != nil {
		return 0, erro
	}

	return userID, nil
}
