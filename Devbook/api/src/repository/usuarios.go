package repository

import (
	"api/src/model"
	"database/sql"
)

// Usuarios representa un repositorio de usuarios
type Usuarios struct {
	db *sql.DB
}

// Crear un repositorio de usuarios
func NuevoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Crear un usuario en base de datos
func (repositorio Usuarios) Crear(usuario model.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, pass) values (?, ?, ?, ?)",
	)

	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Pass)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInsertado, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInsertado), nil

}
