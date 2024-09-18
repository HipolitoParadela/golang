package rutas

import (
	"api/src/controllers"
	"net/http"
)

var rutasPublicaciones = []Ruta{
	{
		URI:                   "/publicaciones",
		Metodo:                http.MethodPost,
		Funcion:               controllers.CrearPublicacion,
		RequiereAutenticacion: true,
	},
	{
		URI:                   "/publicaciones",
		Metodo:                http.MethodGet,
		Funcion:               controllers.ObtenerPublicaciones,
		RequiereAutenticacion: true,
	},
	{
		URI:                   "/publicaciones/{publicacionId}",
		Metodo:                http.MethodGet,
		Funcion:               controllers.ObtenerPublicacion,
		RequiereAutenticacion: true,
	},
	{
		URI:                   "/publicaciones/{publicacionId}",
		Metodo:                http.MethodPut,
		Funcion:               controllers.EditarPublicacion,
		RequiereAutenticacion: true,
	},
	{
		URI:                   "/publicaciones/{publicacionId}",
		Metodo:                http.MethodDelete,
		Funcion:               controllers.BorrarPublicacion,
		RequiereAutenticacion: true,
	},
	{
		URI:                   "/usuarios/{usuarioId}/publicaciones",
		Metodo:                http.MethodGet,
		Funcion:               controllers.BuscarPublicacionesUsuario,
		RequiereAutenticacion: true,
	},
	{
		URI:                   "/publicaciones/{publicacionId}/curtir",
		Metodo:                http.MethodPost,
		Funcion:               controllers.CurtirPublicacion,
		RequiereAutenticacion: true,
	},
	{
		URI:                   "/publicaciones/{publicacionId}/descurtir",
		Metodo:                http.MethodPost,
		Funcion:               controllers.DescurtirPublicacion,
		RequiereAutenticacion: true,
	},
}
