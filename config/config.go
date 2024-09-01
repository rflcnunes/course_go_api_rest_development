package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func setupEnv() {
	log.Println("INFO: Loading .env file")
	if err := godotenv.Load(); err != nil {
		log.Fatal("ERROR: loading .env file", err)
	}
}

func setupDB() {
	log.Println("INFO: Connecting to database")
	dsn := "host=" + os.Getenv("POSTGRES_HOST") + " user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=" + os.Getenv("POSTGRES_DB") + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("ERROR: Failed to connect to database", err)
	}

	DB = db

	log.Println("INFO: Connected to database")
}

func Setup() {
	setupEnv()
	setupDB()
}
