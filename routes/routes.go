package routes

import (
	"golang-webapp/services"

	"github.com/gorilla/mux"
)

func StartRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/product", services.GetProducts).Methods("GET")

	return router
}
	