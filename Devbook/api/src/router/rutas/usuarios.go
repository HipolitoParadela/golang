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
		RequiereAutenticacion: true,
	},
	{
		URI:                   "/usuarios",
		Metodo:                http.MethodGet,
		Funcion:               controllers.ObtenerUsuarios,
		RequiereAutenticacion: true,
	},
	{
		URI:                   "/usuarios/{usuarioID}",
		Metodo:                http.MethodGet,
		Funcion:               controllers.ObtenerUsuario,
		RequiereAutenticacion: true,
	},
	{
		URI:                   "/usuarios/{usuarioID}",
		Metodo:                http.MethodPut,
		Funcion:               controllers.EditarUsuario,
		RequiereAutenticacion: true,
	},
	{
		URI:                   "/usuarios/{usuarioID}",
		Metodo:                http.MethodDelete,
		Funcion:               controllers.BorrarUsuario,
		RequiereAutenticacion: true,
	},
	{
		URI:                   "/usuarios/{usuarioID}/seguir",
		Metodo:                http.MethodPost,
		Funcion:               controllers.SeguirUsuario,
		RequiereAutenticacion: true,
	},
	{
		URI:                   "/usuarios/{usuarioID}/parar-de-seguir",
		Metodo:                http.MethodPost,
		Funcion:               controllers.ParaDeSeguirUsuario,
		RequiereAutenticacion: true,
	},
	{
		URI:                   "/usuarios/{usuarioID}/seguidores",
		Metodo:                http.MethodGet,
		Funcion:               controllers.BuscarSeguidores,
		RequiereAutenticacion: true,
	},
	{
		URI:                   "/usuarios/{usuarioID}/seguidos",
		Metodo:                http.MethodGet,
		Funcion:               controllers.BuscarSeguidos,
		RequiereAutenticacion: true,
	},
	{
		URI:                   "/usuarios/{usuarioID}/actualizar-pass",
		Metodo:                http.MethodPost,
		Funcion:               controllers.ActualizarPass,
		RequiereAutenticacion: true,
	},
}
