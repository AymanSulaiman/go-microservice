package config

import "github.com/spf13/viper"

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	DBName   string
}

type Config struct {
	Database DatabaseConfig
}

func Load() *Config {
	// Set configuration file name and path
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	// Load configuration file
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var conf Config
	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	return &conf
}
