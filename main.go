package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func conectionDatabase() *sql.DB {
	connection := "user=postgres dbname=golang-webapp password=admin host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		log.Fatal("Connection database error: ", err.Error())
	}

	return db
}

type Procuct struct {
	id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/product", GetProducts).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetProducts(w http.ResponseWriter, r *http.Request) {

	db := conectionDatabase()

	selectProducts, err := db.Query("select * from products")

	if err != nil {
		log.Fatal("select products error: ", err.Error())
	}

	p := Procuct{}
	products := []Procuct{}

	for selectProducts.Next() {
		err = selectProducts.Scan(&p.id, &p.Name, &p.Description, &p.Price, &p.Quantity)

		if err != nil {
			log.Fatal("scan products error: ", err.Error())
		}

		products = append(products, p)
	}

	defer db.Close()

	json.NewEncoder(w).Encode(products)
}
