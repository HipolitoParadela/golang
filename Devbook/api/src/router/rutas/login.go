package rutas

import (
	"api/src/controllers"
	"net/http"
)

var rutaLogin = Ruta{
	URI:                   "/login",
	Metodo:                http.MethodPost,
	Funcion:               controllers.Login,
	RequiereAutenticacion: false,
}
