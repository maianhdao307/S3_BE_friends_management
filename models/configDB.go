package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB variable
var DB *gorm.DB

// Connect database
func Connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "123456", "friend_management")

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	db.DB()
	db.AutoMigrate(&Friend{}, &Subscription{})

	DB = db

	return err
}
