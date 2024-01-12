package db

import (
	"fmt"
	"log"
	"oshno/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	configuration := config.Config()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=%s",
		configuration.Database.DBHost,
		configuration.Database.DBUser,
		configuration.Database.DBPassword,
		configuration.Database.DBName,
		configuration.Database.DBPort,
		configuration.Database.DBTimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Database connection error:%v", err)
	}

	return db
}
