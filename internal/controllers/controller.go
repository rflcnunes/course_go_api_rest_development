package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rflcnunes/course_go_api_rest_development/config"
	"github.com/rflcnunes/course_go_api_rest_development/internal/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personality
	config.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func GetPersonality(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var p models.Personality
	config.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}
