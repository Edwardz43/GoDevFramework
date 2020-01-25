package config

import (
	"log"

	"github.com/spf13/viper"
)

// Database sets data base configurarion
func Database() string {
	var connectionString string
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	connectionString = viper.GetString("DATABASE_CONNECTION_STRING")
	// if !ok {
	// 	log.Fatalf("Invalid type assertion")
	// }

	return connectionString
}