package rutas

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Ruta representa todas las rutas de la API
type Ruta struct {
	URI                   string
	Metodo                string
	Funcion               func(http.ResponseWriter, *http.Request)
	RequiereAutenticacion bool
}

// Cofigurar, coloca todas las rutas dentro del router
func Configurar(r *mux.Router) *mux.Router {
	rutas := rutasUsuarios
	rutas = append(rutas, rutaLogin)
	rutas = append(rutas, rutasPublicaciones...)

	for _, ruta := range rutas {

		if ruta.RequiereAutenticacion {
			r.HandleFunc(ruta.URI,
				middlewares.Logger(middlewares.Autenticar(ruta.Funcion)),
			).Methods(ruta.Metodo)
		} else {
			r.HandleFunc(ruta.URI, middlewares.Logger(ruta.Funcion)).Methods(ruta.Metodo)
		}
	}

	return r
}
