package config_test

import "testing"

import "Edwardz43/godevframework/config"

import "github.com/stretchr/testify/assert"

func TestDatabase(t *testing.T) {
	var connectionString string

	connectionString = config.Database()

	assert.NotNil(t, connectionString)
}
