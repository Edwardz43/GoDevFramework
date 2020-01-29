package config_test

import "testing"

import "Edwardz43/godevframework/config"

import "github.com/stretchr/testify/assert"

func TestRedisConfig(t *testing.T) {
	config := config.Redis()

	assert.NotEmpty(t, config.Addr)
	// assert.NotEmpty(t, config.PASSWORD)
	assert.NotEmpty(t, config.INDEX)
}
