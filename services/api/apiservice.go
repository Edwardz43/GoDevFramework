package api

import (
	"Edwardz43/godevframework/config"
	"Edwardz43/godevframework/models"
	repository "Edwardz43/godevframework/repositories"
	"Edwardz43/godevframework/repositories/user"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Service ...
type Service struct {
	userRepo repository.UserRepository
}

// GetInstance ...
func GetInstance() *Service {
	connStr := config.Database()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return &Service{
		userRepo: user.GetInstance(db),
	}
}

// CreateNewUser ...
func (s *Service) CreateNewUser(name, email string) bool {
	bool, err := s.userRepo.Create(name, email)
	if err != nil {
		return false
	}
	return bool
}

// FindUserByID ...
func (s *Service) FindUserByID(id string) *models.User {
	u, err := s.userRepo.FindOne(id)
	if err != nil {
		return nil
	}
	return u
}
