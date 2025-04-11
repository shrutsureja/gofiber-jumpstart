package repository

import (
	"app/src/database"
	"app/src/database/models"
)

type userRepository struct{}

var UserRepository = userRepository{}

// Create a new user
func (ur *userRepository) CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}
