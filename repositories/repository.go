package repository

import "Edwardz43/godevframework/models"

// UserRepository defines the functions of user repository.
type UserRepository interface {
	Create(name, email string) (bool, error)
	FindOne(id string) (*models.User, error)
	Update(id, name, email string) (bool, error)
}
