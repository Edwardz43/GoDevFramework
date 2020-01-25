package api

import (
	"Edwardz43/godevframework/config"
	"Edwardz43/godevframework/models"
	repository "Edwardz43/godevframework/repositories"
	"Edwardz43/godevframework/repositories/user"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

// Service ...
type Service struct{}

var (
	userRepo repository.UserRepository
)

// Init ..
func (s *Service) Init() {
	connStr := config.Database()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	userRepo = user.GetInstance(db)
}

// CreateNewUser ...
func (s *Service) CreateNewUser(name, email string) bool {
	bool, err := userRepo.Create(name, email)
	if err != nil {
		return false
	}
	return bool
}

// FindUserByID ...
func (s *Service) FindUserByID(id string) *models.User {
	u, err := userRepo.FindOne(id)
	if err != nil {
		return nil
	}
	return u
}
