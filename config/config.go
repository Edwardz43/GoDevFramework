package config

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("../../../")
	_ = viper.ReadInConfig()

	// if err != nil {
	// 	log.Fatalf("Error while reading config file %s", err)
	// }

}
