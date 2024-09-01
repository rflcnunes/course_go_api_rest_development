package config

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rflcnunes/course_go_api_rest_development/internal/controllers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func handleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.GetAllPersonalities).Methods("GET")
	r.HandleFunc("/api/personalities/{id}", controllers.GetPersonality).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

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
	handleRequest()
}
