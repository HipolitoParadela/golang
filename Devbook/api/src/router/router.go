package router

import (
	"api/src/router/rutas"

	"github.com/gorilla/mux"
)

// Generar va a retornar un router con las rutas configuradas
func Generar() *mux.Router {
	r := mux.NewRouter()
	return rutas.Configurar(r)

}
