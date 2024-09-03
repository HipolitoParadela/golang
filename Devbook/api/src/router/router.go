package router

import "github.com/gorilla/mux"

// Generar va a retornar un router con las rutas configuradas
func Generar() *mux.Router {
	return mux.NewRouter()
}
