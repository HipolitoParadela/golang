package controllers

import (
	"api/src/autenticacion"
	"api/src/db"
	"api/src/model"
	"api/src/repository"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Crear publicacion en base de datos
func CrearPublicacion(w http.ResponseWriter, r *http.Request) {

	usuarioID, erro := autenticacion.ExtrerUsuarioID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	dataRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacion model.Publicaciones
	if erro = json.Unmarshal(dataRequest, &publicacion); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	publicacion.AutorID = usuarioID

	if erro = publicacion.Preparar(); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	database, erro := db.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer database.Close()

	repositoryPublicaciones := repository.NuevoRepositorioDePublicaciones(database)
	publicacion.ID, erro = repositoryPublicaciones.Crear(publicacion)
	if erro != nil {
		responses.Erro(w, http.StatusPreconditionFailed, erro)
		return
	}
	responses.JSON(w, http.StatusCreated, publicacion)
}

// Obtener listado de publicacion en base de datos por su id
func ObtenerPublicacion(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	publicacionID, erro := strconv.ParseUint(parametros["publicacionId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NuevoRepositorioDePublicaciones(db)
	publicacion, erro := repositorio.BuscarPorID(publicacionID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, publicacion)
}

// Obtener listado de publicacion en base de datos
func ObtenerPublicaciones(w http.ResponseWriter, r *http.Request) {

	usuarioID, erro := autenticacion.ExtrerUsuarioID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	/* nomeOrNick := strings.ToLower(r.URL.Query().Get("publicacion")) */

	db, erro := db.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NuevoRepositorioDePublicaciones(db)
	publicaciones, erro := repositorio.ListarPublicacionesPorId(usuarioID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, publicaciones)
}

// Editar un publicacion en base de datos
func EditarPublicacion(w http.ResponseWriter, r *http.Request) {

	usuarioID, erro := autenticacion.ExtrerUsuarioID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	publicacionID, erro := strconv.ParseUint(parametros["publicacionId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NuevoRepositorioDePublicaciones(db)
	publicacionGuardadaDB, erro := repositorio.BuscarPorID(publicacionID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if usuarioID != publicacionGuardadaDB.AutorID {
		responses.Erro(w, http.StatusForbidden, errors.New("no es posible actualizar un publicacion que no sea el suyo"))
		return
	}

	bodySolicitud, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacion model.Publicaciones
	if erro = json.Unmarshal(bodySolicitud, &publicacion); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = publicacion.Preparar(); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repositorio.Actualizar(publicacionID, publicacion); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// Borrar publicacion en base de datos
func BorrarPublicacion(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacion.ExtrerUsuarioID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	publicacionID, erro := strconv.ParseUint(parametros["publicacionId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NuevoRepositorioDePublicaciones(db)
	publicacionGuardadaDB, erro := repositorio.BuscarPorID(publicacionID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if usuarioID != publicacionGuardadaDB.AutorID {
		responses.Erro(w, http.StatusForbidden, errors.New("no es posible actualizar un publicacion que no sea el suyo"))
		return
	}

	if erro = repositorio.Borrar(publicacionID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// Buscar publicaciones de un usuario
func BuscarPublicacionesUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	database, erro := db.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer database.Close()

	repository := repository.NuevoRepositorioDePublicaciones(database)
	publicaciones, erro := repository.BuscarPublicacionesUsuario(usuarioID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, publicaciones)
}

// Adiciona una curtida a la publicación
func CurtirPublicacion(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	publicacionID, erro := strconv.ParseUint(parametros["publicacionId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NuevoRepositorioDePublicaciones(db)
	if erro = repository.CurtirPublicacion(publicacionID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// Adiciona una curtida a la publicación
func DescurtirPublicacion(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	publicacionID, erro := strconv.ParseUint(parametros["publicacionId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NuevoRepositorioDePublicaciones(db)
	if erro = repository.DescurtirPublicacion(publicacionID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
