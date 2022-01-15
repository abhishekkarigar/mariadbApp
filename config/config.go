package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type dbConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

type AppConfig struct {
	Db dbConfig
}

func ReadEnv() *AppConfig {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	var appConfig AppConfig

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return nil
	}
	err := viper.Unmarshal(&appConfig)
	if err != nil {
		fmt.Printf("unable to decode into struct %v", err)
		return nil
	}
	return &appConfig
}
