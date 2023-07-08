package services

import (
	"github.com/dami-pie/napi/models"
	"github.com/dami-pie/napi/src/database"
	"github.com/dami-pie/napi/src/repositories"
)

func CreateUser(user models.User) (uint64, error) {
	db, err := database.Connect()
	if err != nil {
		return 0, err
	}

	repository := repositories.NewUserRepository(db)
	userID, err := repository.Create(user)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func GetUser(id uint64) (models.User, error) {
	db, err := database.Connect()
	if err != nil {
		return models.User{}, err
	}

	repository := repositories.NewUserRepository(db)
	user, err := repository.Get(id)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func GetUserByEmail(email string) (models.User, error) {
	db, err := database.Connect()
	if err != nil {
		return models.User{}, err
	}

	repository := repositories.NewUserRepository(db)
	user, err := repository.GetByEmail(email)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func DeleteUser(id uint64) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	repository := repositories.NewUserRepository(db)
	err = repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(id uint64, newData models.User) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	repository := repositories.NewUserRepository(db)
	err = repository.Update(id, newData)

	if err != nil {
		return err
	}

	return nil
}

func UserExists(id uint64) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		return false, err
	}

	repository := repositories.NewUserRepository(db)
	return repository.Exists(id)
}
