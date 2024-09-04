package rutas

import (
	"api/src/controllers"
	"net/http"
)

var rutasUsuarios = []Ruta{
	{
		URI:                   "/usuarios",
		Metodo:                http.MethodPost,
		Funcion:               controllers.CrearUsuario,
		RequiereAutenticacion: false,
	},
	{
		URI:                   "/usuarios",
		Metodo:                http.MethodGet,
		Funcion:               controllers.ObtenerUsuarios,
		RequiereAutenticacion: false,
	},
	{
		URI:                   "/usuarios/{usuarioId}",
		Metodo:                http.MethodGet,
		Funcion:               controllers.ObtenerUsuario,
		RequiereAutenticacion: false,
	},
	{
		URI:                   "/usuarios/{usuarioId}",
		Metodo:                http.MethodPut,
		Funcion:               controllers.EditarUsuario,
		RequiereAutenticacion: false,
	},
	{
		URI:                   "/usuarios/{usuarioId}",
		Metodo:                http.MethodDelete,
		Funcion:               controllers.BorrarUsuario,
		RequiereAutenticacion: false,
	},
}
