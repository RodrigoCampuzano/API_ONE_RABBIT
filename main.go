package main

import (
	"log"
	"net/http"

	"API_ONE/src/esp32/infraestructure/routes"
)

func main() {

	router := routes.NewVentaRouter()


	log.Println("API escuchando en :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}