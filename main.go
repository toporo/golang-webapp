package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Procuct struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/product", GetProduct).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))

}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	products := []Procuct{
		{Name: "Notebook", Description: "Notebook Dell", Price: 1999.00, Quantity: 20},
		{Name: "Mouse", Description: "Mouse Razr", Price: 399.00, Quantity: 30},
	}
	json.NewEncoder(w).Encode(products)
}
