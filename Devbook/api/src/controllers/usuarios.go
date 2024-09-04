package controllers

import (
	"api/src/db"
	"api/src/model"
	"api/src/repository"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Crear usuario en base de datos
func CrearUsuario(w http.ResponseWriter, r *http.Request) {
	dataRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var usuario model.Usuario
	if erro = json.Unmarshal(dataRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}

	database, erro := db.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repositoryUser := repository.NuevoRepositorioDeUsuarios(database)
	usuarioId, erro := repositoryUser.Crear(usuario)
	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("Id insertado: %d", usuarioId)))
}

// Obtener listado de usuario en base de datos
func ObtenerUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Obtener usuarios"))
}

// Obtener listado de usuario en base de datos por su id
func ObtenerUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Obtener un usuarios por su id"))
}

// Editar un usuario en base de datos
func EditarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Editar usuario"))
}

// Borrar usuario en base de datos
func BorrarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Borrar usuario"))
}
