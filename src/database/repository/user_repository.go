package repository

import (
	"app/src/database"
	"app/src/database/models"
)

type userRepository struct{}

var UserRepository = userRepository{}

func (ur *userRepository) CreateUser(user *models.User) error {
	// Create a new user
	if err := database.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
