package middlewares

import (
	"api/src/autenticacion"
	"api/src/responses"
	"log"
	"net/http"
)

// Logger escribe informacion de las solicitudes en el terminal
func Logger(avanzarProxFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		avanzarProxFunc(w, r)
	}
}

// Autenticando verifica si el usuario esta autenticado
func Autenticar(avanzarProxFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := autenticacion.ValidarToken(r); erro != nil {
			responses.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		avanzarProxFunc(w, r)
	}
}
