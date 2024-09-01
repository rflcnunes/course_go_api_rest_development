package config

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rflcnunes/course_go_api_rest_development/internal/controllers"
)

func handleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.GetAllPersonalities).Methods("GET")
	r.HandleFunc("/api/personalities/{id}", controllers.GetPersonality).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func Setup() {
	handleRequest()
}
