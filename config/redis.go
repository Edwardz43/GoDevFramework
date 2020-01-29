package config

import "github.com/spf13/viper"

// RedisConfig ...
type RedisConfig struct {
	Addr     string `json:"addr"`
	PASSWORD string `json:"password"`
	INDEX    int    `json:"index"`
}

// Redis ...
func Redis() *RedisConfig {

	addr := viper.GetString("REDIS_ADDRESS")
	passwd := viper.GetString("REDIS_PASSWORD")
	index := viper.GetInt("REDIS_INDEX")

	return &RedisConfig{
		Addr:     addr,
		PASSWORD: passwd,
		INDEX:    index,
	}
}
