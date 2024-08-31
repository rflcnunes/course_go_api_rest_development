package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rflcnunes/course_go_api_rest_development/internal/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalities)
}
