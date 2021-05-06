package repository

import (
	"golang-webapp/db"
	"golang-webapp/entities"
	"log"
)

func GetProducts() []entities.Procuct {

	db := db.ConectionDatabase()

	selectProducts, err := db.Query("select * from products")

	if err != nil {
		log.Fatal("select products error: ", err.Error())
	}

	p := entities.Procuct{}
	products := []entities.Procuct{}

	for selectProducts.Next() {
		err = selectProducts.Scan(&p.Id, &p.Name, &p.Description, &p.Price, &p.Quantity)

		if err != nil {
			log.Fatal("scan products error: ", err.Error())
		}

		products = append(products, p)
	}

	defer db.Close()

	return products
}

func CreateProduct(product entities.Procuct) string {

	db := db.ConectionDatabase()

	insertDatabase, err := db.Prepare("Insert into products(product_name, description, price, quantity) values($1, $2, $3, $4)")

	if err != nil {
		log.Fatal("Error prepare insert product: ", err.Error())
	}

	response, err := insertDatabase.Exec(product.Name, product.Description, product.Price, product.Quantity)

	log.Fatal(response)

	defer db.Close()

	return ""
}
