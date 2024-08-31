package config

import (
	"log"
	"net/http"

	"github.com/rflcnunes/course_go_api_rest_development/internal/controllers"
)

func handleRequest() {
	http.HandleFunc("/", controllers.Home)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func Setup() {
	handleRequest()
}
