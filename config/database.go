package config

import (
	"fmt"
	"one-week-project-ecommerce/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB(config *Config) *gorm.DB {
	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DBHost, config.DBUsername, config.DBPassword, config.DBName, config.DBPort)

	db, err := gorm.Open(postgres.Open(dbInfo), &gorm.Config{})

	helper.ErrorPanic(err)

	fmt.Println("Connected successfully to database!")
	return db
}
