package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppName  string
	AppPort  string
	BasePath string
}

var AppConfig *Config

func LoadConfiguration() {

	viper.SetDefault("APP_NAME", "go-chat-room")
	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("BASE_PATH", "/go-chat-room")

	// Automatically read environment variables
	readConfigFile("../your-path-file")

	AppConfig = &Config{
		AppName:  viper.GetString("APP_NAME"),
		AppPort:  viper.GetString("APP_PORT"),
		BasePath: viper.GetString("BASE_PATH"),
	}
}

func readConfigFile(path string) {
	viper.AutomaticEnv()
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Warning: No .env file found, using default values")
	}
}
