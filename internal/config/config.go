package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type (
	Config struct {
		App App
		Log Log
		DB  Database
	}

	App struct {
		Host string `yml:"host"`
		Port string `yml:"port"`
	}

	Log struct {
		Level string `yml:"level"`
	}

	Database struct {
		Host     string `yml:"host"`
		Port     string `yml:"port"`
		Username string `yml:"username"`
		Password string `yml:"password"`
		Database string `yml:"database"`
	}
)

func New() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("unable to decode config into struct: %w", err)
	}

	return &config, nil
}
