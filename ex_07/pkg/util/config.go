package util

import (
	"github.com/spf13/viper"
	"os"
	"time"
)

var (
	defaultEnvFile = ".env"
	envFileKey     = "ENV_FILE"
)

type Config struct {
	ENV                   string        `mapstructure:"ENV"`
	SecretKey             string        `mapstructure:"SECRET_KEY"`
	DBSource              string        `mapstructure:"DB_SOURCE"`
	ServerAddress         string        `mapstructure:"SERVER_ADDRESS"`
	AccessTokenExpiresIn  time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRES_IN"`
	RefreshTokenExpiresIn time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRES_IN"`
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
