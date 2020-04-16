package config_test

import (
	"Edwardz43/godevframework/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLogConfig(t *testing.T) {
	config := config.Log()

	assert.NotEmpty(t, config.HookIndex)
	assert.NotEmpty(t, config.HookURL)
	assert.NotEmpty(t, config.Path)
}
