package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"golang-webapp/services"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/product", services.GetProducts).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
