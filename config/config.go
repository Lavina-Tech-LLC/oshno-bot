package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Configurations exported
type configuration struct {
	Database databaseConfigurations
	Telegram telegramConfigurations
	Server   serverConfigurations
}

// DatabaseConfigurations exported
type databaseConfigurations struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     uint
	DBTimeZone string
}

type telegramConfigurations struct {
	TelegramToken string
}

type serverConfigurations struct {
	Port string
}

func Config() *configuration {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration *configuration

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	viper.SetDefault("database.dbname", "test_db")
	viper.SetDefault("server.port", ":8080")
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return configuration
}
