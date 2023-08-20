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
	MigrationURL          string        `mapstructure:"MIGRATION_URL"`
	ServerAddress         string        `mapstructure:"SERVER_ADDRESS"`
	AccessTokenExpiresIn  time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRES_IN"`
	RefreshTokenExpiresIn time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRES_IN"`
}

func LoadConfig(path string) (config Config, err error) {
	envFile := os.Getenv(envFileKey)
	if envFile == "" {
		if path != "" {
			envFile = path
		} else {
			envFile = defaultEnvFile
		}
	}

	_, err = os.Stat(envFile)
	if err == nil {
		viper.SetConfigFile(envFile)
		viper.AutomaticEnv()

		err = viper.ReadInConfig()
		if err != nil {
			return
		}

		err = viper.Unmarshal(&config)
		return
	}

	return
}
