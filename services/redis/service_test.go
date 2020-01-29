package redis_test

import "testing"

import "Edwardz43/godevframework/services/redis"

import "github.com/stretchr/testify/assert"

func TestConnection(t *testing.T) {
	service := redis.GetInstance()

	result, err := service.Client.Ping().Result()

	assert.Nil(t, err)
	assert.Equal(t, result, "PONG")
}
