package config

import "github.com/spf13/viper"

// LogConfig ...
type LogConfig struct {
	IsHook    bool   `json:"is_hook"`
	HookIndex string `json:"hook_index"`
	HookURL   string `json:"hook_url"`
	Path      string `json:"path"`
}

// Log ...
func Log() *LogConfig {

	path := viper.GetString("LOG_PATH")
	hookURL := viper.GetString("ELASTICSEARCH_URL")
	hookIndex := viper.GetString("ELASTICSEARCH_INDEX")
	isHook := viper.GetBool("ELASTICSEARCH_HOOK")

	return &LogConfig{
		Path:      path,
		IsHook:    isHook,
		HookURL:   hookURL,
		HookIndex: hookIndex,
	}
}
