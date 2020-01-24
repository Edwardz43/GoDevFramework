package user_test

import (
	"Edwardz43/godevframework/repositories/user"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

var userRepo *user.PostgreUserRepository

func init() {
	connStr := "user=admin password=test dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	userRepo = user.GetInstance(db)
}

func TestCreate(t *testing.T) {
	tag := time.Now().Unix()
	name := fmt.Sprintf("TestUser%d", tag)
	email := fmt.Sprintf("%d@testmail.org", tag)
	isSucess, err := userRepo.Create(name, email)
	assert.Nil(t, err)
	assert.True(t, isSucess)
}

func TestFindOne(t *testing.T) {
	id := "2"
	u, err := userRepo.FindOne(id)
	assert.Nil(t, err)
	assert.NotEqual(t, u.ID, "")
}
