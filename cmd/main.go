package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rflcnunes/course_go_api_rest_development/config"
	"github.com/rflcnunes/course_go_api_rest_development/internal/router"
)

func main() {
	fmt.Println("Hello Friend!")

	config.Setup()
	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", r))
}
