package config

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func handleRequest() {
	http.HandleFunc("/", Home)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func Setup() {
	handleRequest()
}
