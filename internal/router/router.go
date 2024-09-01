package router

import (
	"github.com/gorilla/mux"
	"github.com/rflcnunes/course_go_api_rest_development/internal/controllers"
)

var PATH = "/api/personalities"
var PATH_ID = "/api/personalities/{id}"

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc(PATH, controllers.GetAllPersonalities).Methods("GET")
	r.HandleFunc(PATH_ID, controllers.GetPersonality).Methods("GET")
	r.HandleFunc(PATH, controllers.CreatePersonality).Methods("POST")
	r.HandleFunc(PATH_ID, controllers.DeletePersonality).Methods("DELETE")
	r.HandleFunc(PATH_ID, controllers.EditPersonality).Methods("PUT")
	return r
}
