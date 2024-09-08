package rutas

import (
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

	for _, ruta := range rutas {
		r.HandleFunc(ruta.URI, ruta.Funcion).Methods(ruta.Metodo)
	}

	return r
}
