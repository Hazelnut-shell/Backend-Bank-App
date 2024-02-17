package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"` // type is time.Duration. This type has better readability
}

// viper uses mapstructure under the hood for unmarshaling values

// LoadConfig reads configuration from file or environment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app") // file name
	viper.SetConfigType("env") // file type

	viper.AutomaticEnv() // override the values from config file with values of corresponding environment variables if they exist

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// unmarshal into struct Config
	err = viper.Unmarshal(&config)
	return
}
