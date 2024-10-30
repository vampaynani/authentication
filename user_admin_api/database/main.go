package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDbConnection() *gorm.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", 
	os.Getenv("DB_HOST"), 
	os.Getenv("DB_PORT"), 
	os.Getenv("DB_USER"), 
	os.Getenv("DB_PASSWORD"), 
	os.Getenv("DB_NAME"), 
	os.Getenv("DB_SSL"))
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Fatal(err)
	}
	return db
}