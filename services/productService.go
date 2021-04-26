package services

import (
	"encoding/json"
	"golang-webapp/repository"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	response := repository.GetProducts()
	json.NewEncoder(w).Encode(response)
}
