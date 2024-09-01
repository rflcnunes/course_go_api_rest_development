package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rflcnunes/course_go_api_rest_development/config"
	"github.com/rflcnunes/course_go_api_rest_development/internal/models"
	"github.com/rflcnunes/course_go_api_rest_development/internal/router"
)

func main() {
	fmt.Println("Hello Friend!")

	models.Personalities = []models.Personality{
		{ID: 1, Name: "John Doe", History: "Lorem Ipsum Dolor Sit Amet for John Doe"},
		{ID: 2, Name: "Jane Smith", History: "Lorem Ipsum Dolor Sit Amet for Jane Smith"},
		{ID: 3, Name: "Michael Johnson", History: "Lorem Ipsum Dolor Sit Amet for Michael Johnson"},
		{ID: 4, Name: "Emily Davis", History: "Lorem Ipsum Dolor Sit Amet for Emily Davis"},
	}

	config.Setup()
	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", r))
}
