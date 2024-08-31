package main

import (
	"fmt"

	"github.com/rflcnunes/course_go_api_rest_development/internal/models"
	"github.com/rflcnunes/course_go_api_rest_development/pkg/config"
)

func main() {
	fmt.Println("Hello Friend!")

	models.Personalities = []models.Personality{
		{Name: "John Doe", History: "Lorem Ipsum Dolor Sit Amet for John Doe"},
		{Name: "Jane Smith", History: "Lorem Ipsum Dolor Sit Amet for Jane Smith"},
		{Name: "Michael Johnson", History: "Lorem Ipsum Dolor Sit Amet for Michael Johnson"},
		{Name: "Emily Davis", History: "Lorem Ipsum Dolor Sit Amet for Emily Davis"},
	}

	config.Setup()
}
