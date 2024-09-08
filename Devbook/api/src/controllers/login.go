package controllers

import (
	"api/src/autenticacion"
	"api/src/db"
	"api/src/model"
	"api/src/repository"
	"api/src/responses"
	"api/src/seguridad"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Función para loguearse
func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario model.Usuario
	if erro = json.Unmarshal(bodyRequest, &usuario); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	database, erro := db.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer database.Close()

	repositorio := repository.NuevoRepositorioDeUsuarios(database)
	usuarioGuardadoEnDB, erro := repositorio.BuscarPorEmail(usuario.Email)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = seguridad.VerificarPass(usuarioGuardadoEnDB.Pass, usuario.Pass); erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, _ := autenticacion.CrearToken(usuarioGuardadoEnDB.ID)
	fmt.Println(token)
	w.Write([]byte(token))
}
