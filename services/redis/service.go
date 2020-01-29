package redis

import (
	"Edwardz43/godevframework/config"
	"github.com/go-redis/redis/v7"
)

// Service ...
type Service struct {
	Client *redis.Client
}

// GetInstance return a redis service.
func GetInstance() *Service {
	config := config.Redis()

	return &Service{
		Client: redis.NewClient(&redis.Options{
			Addr:     config.Addr,
			Password: config.PASSWORD, // no password set
			DB:       config.INDEX,    // use default DB
		}),
	}
}
