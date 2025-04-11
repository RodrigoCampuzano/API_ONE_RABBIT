package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/rs/cors"
    "API_ONE/src/esp32/infraestructure/repositories"
    "API_ONE/src/esp32/infraestructure/routes"
)

func main() {
    ventaRepo := repositories.NewVentaRepositoryRabbitMQ()
    emailRepo := repositories.NewEmailMockRepository()

    ventaRouter := routes.NewVentaRouter(ventaRepo, emailRepo)
    emailRouter := routes.NewEmailRouter(emailRepo)

    mux := http.NewServeMux()
    mux.Handle("/email/enviar-email", emailRouter)
    mux.Handle("/ventas/enviar-venta", ventaRouter)
    
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "API funcionando")
    })

    // Configurar CORS
    handler := cors.Default().Handler(mux)

    fmt.Println("Servidor corriendo en el puerto 8080...")
    log.Fatal(http.ListenAndServe(":8080", handler))
}