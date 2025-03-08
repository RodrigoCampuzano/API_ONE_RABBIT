package main

import (
	"log"
	"net/http"

	"API_ONE/src/core/db"
	"API_ONE/src/core/middleware"
	"API_ONE/src/esp32/infraestructure/routes"
)

func main() {
	db.Init()

	router := routes.NewVentaRouter()

	handler := middleware.CORS(router)

	log.Println("API escuchando en :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}