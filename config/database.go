package config

import (
	"github.com/spf13/viper"
)

// Database sets data base configurarion
func Database() string {
	var connectionString string
	connectionString = viper.GetString("DATABASE_CONNECTION_STRING")
	return connectionString
}
