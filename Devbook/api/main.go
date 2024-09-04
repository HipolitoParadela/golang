package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Cargar()
	r := router.Generar()

	fmt.Printf("Ejecutando en puerto %d", config.Puerto)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Puerto), r))
}
