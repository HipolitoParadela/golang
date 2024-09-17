package controllers

import (
	"api/src/autenticacion"
	"api/src/db"
	"api/src/model"
	"api/src/repository"
	"api/src/responses"
	"api/src/seguridad"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Crear usuario en base de datos
func CrearUsuario(w http.ResponseWriter, r *http.Request) {
	dataRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario model.Usuario
	if erro = json.Unmarshal(dataRequest, &usuario); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("registro"); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	database, erro := db.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer database.Close()

	repositoryUser := repository.NuevoRepositorioDeUsuarios(database)
	usuario.ID, erro = repositoryUser.Crear(usuario)
	if erro != nil {
		responses.Erro(w, http.StatusPreconditionFailed, erro)
		return
	}
	responses.JSON(w, http.StatusCreated, usuario)
}

// Obtener listado de usuario en base de datos
func ObtenerUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOrNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, erro := db.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NuevoRepositorioDeUsuarios(db)
	usuarios, erro := repositorio.Listar(nomeOrNick)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, usuarios)
}

// Obtener listado de usuario en base de datos por su id
func ObtenerUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseInt(parametros["usuarioID"], 10, 64)
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

	repositorio := repository.NuevoRepositorioDeUsuarios(db)
	usuario, erro := repositorio.BuscarPorID(usuarioID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, usuario)
}

// Editar un usuario en base de datos
func EditarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioID"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIDenToken, erro := autenticacion.ExtrerUsuarioID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioID != usuarioIDenToken {
		responses.Erro(w, http.StatusForbidden, errors.New("no es posible actualizar un usuario que no sea el suyo"))
		return
	}

	bodySolicitud, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario model.Usuario
	if erro = json.Unmarshal(bodySolicitud, &usuario); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("edicion"); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NuevoRepositorioDeUsuarios(db)
	if erro = repositorio.Actualizar(usuarioID, usuario); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// Borrar usuario en base de datos
func BorrarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioID"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIDenToken, erro := autenticacion.ExtrerUsuarioID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioID != usuarioIDenToken {
		responses.Erro(w, http.StatusForbidden, errors.New("no es posible borrar un usuario que no sea el suyo"))
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NuevoRepositorioDeUsuarios(db)
	if erro = repositorio.Borrar(usuarioID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// Seguir usuario
func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorID, erro := autenticacion.ExtrerUsuarioID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioID"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if seguidorID == usuarioID {
		responses.Erro(w, http.StatusForbidden, errors.New("no es posible seguir a si mismo"))
		return
	}

	database, erro := db.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer database.Close()

	repository := repository.NuevoRepositorioDeUsuarios(database)
	if erro = repository.SeguirUsuario(usuarioID, seguidorID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// Parar de seguir un usuario
func ParaDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorID, erro := autenticacion.ExtrerUsuarioID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioID"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if seguidorID == usuarioID {
		responses.Erro(w, http.StatusForbidden, errors.New("no es posible dejar de seguir a si mismo"))
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NuevoRepositorioDeUsuarios(db)
	if erro = repository.ParaDeSeguirUsuario(usuarioID, seguidorID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// Buscar seguidores
func BuscarSeguidores(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioID"], 10, 64)
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

	repository := repository.NuevoRepositorioDeUsuarios(db)
	seguidores, erro := repository.BuscarSeguidores(usuarioID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, seguidores)

}

// Buscar seguidores
func BuscarSeguidos(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioID"], 10, 64)
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

	repository := repository.NuevoRepositorioDeUsuarios(db)
	seguidores, erro := repository.BuscarSeguidos(usuarioID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, seguidores)

}

// Actualizar pass
func ActualizarPass(w http.ResponseWriter, r *http.Request) {
	seguidorID, erro := autenticacion.ExtrerUsuarioID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioID"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if seguidorID != usuarioID {
		responses.Erro(w, http.StatusForbidden, errors.New("no es posible realizar esta acción"))
		return
	}

	bodySolicitud, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	var pass model.Pass
	if erro = json.Unmarshal(bodySolicitud, &pass); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NuevoRepositorioDeUsuarios(db)
	passActualUser, erro := repositorio.BuscarPassActual(usuarioID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = seguridad.VerificarPass(passActualUser, pass.Actual); erro != nil {
		responses.Erro(w, http.StatusUnauthorized, errors.New("el pass enviado no coincide con el del usuario"))
		return
	}

	passConHash, erro := seguridad.Hash(pass.Nueva)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repositorio.ActualizarPass(usuarioID, string(passConHash)); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}
