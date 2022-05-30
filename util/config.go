package util

import "github.com/spf13/viper"

// Config stores all configuration of the application
// The values are read by viper from a config file or environtment variables.
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	// Config viper
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// Check all variable in env
	viper.AutomaticEnv()

	// Find and read variable the config file
	err = viper.ReadInConfig()
	// If error
	if err != nil {
		return
	}

	// Insert value config into object viper
	err = viper.Unmarshal(&config)
	return
}
