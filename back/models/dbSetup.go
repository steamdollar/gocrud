package models

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// "back/models"
	"log"

	"github.com/joho/godotenv"
)

// SetupDatabase : 서버 구동시 db와 싱크
func SetupDatabase() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	dbTimezone := os.Getenv("DB_TIMEZONE")
	
	
	dsn := "host=" + dbHost + " user=" + dbUser + 
	" password=" + dbPass + " dbname=" + dbName + " port=" + dbPort +
	" sslmode=" + dbSSLMode + " TimeZone=" + dbTimezone
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	
	db.AutoMigrate(&Board{})
	
	return db
}