package router

import (
	"github.com/gorilla/mux"
	"github.com/rflcnunes/course_go_api_rest_development/internal/controllers"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.GetAllPersonalities).Methods("GET")
	r.HandleFunc("/api/personalities/{id}", controllers.GetPersonality).Methods("GET")
	r.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("POST")
	r.HandleFunc("/api/personalities/{id}", controllers.DeletePersonality).Methods("DELETE")
	return r
}
