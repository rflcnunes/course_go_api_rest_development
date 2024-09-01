package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rflcnunes/course_go_api_rest_development/internal/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalities)
}

func GetPersonality(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	for _, personality := range models.Personalities {
		if strconv.Itoa(personality.ID) == id {
			json.NewEncoder(w).Encode(personality)
			return
		}
	}

	json.NewEncoder(w).Encode(&models.Personality{})
}
