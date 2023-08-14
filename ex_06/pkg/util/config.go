package util

import (
	"github.com/spf13/viper"
	"os"
)

var (
	defaultEnvFile = ".env"
	envFileKey     = "ENV_FILE"
)

type Config struct {
	ENV           string `mapstructure:"ENV"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	envFile := os.Getenv(envFileKey)
	if envFile != "" {
		path = envFile
	} else {
		if path == "" {
			path = defaultEnvFile
		}
	}

	viper.SetConfigFile(path)

	// Automatically overrides values with the value of corresponding environment variable if they exist
	viper.AutomaticEnv()

	// Find and read that env file
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
