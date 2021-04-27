package main

import (
	"golang-webapp/routes"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8000", routes.StartRoutes()))
}
