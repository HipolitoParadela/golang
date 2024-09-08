package repository

import (
	"api/src/model"
	"database/sql"
	"fmt"
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

// Listar usuarios
func (repository Usuarios) Listar(nomeOrNick string) ([]model.Usuario, error) {
	nomeOrNick = fmt.Sprintf("%%%s%%", nomeOrNick)

	resutlUserQuery, erro := repository.db.Query(
		"select id, nome, nick, email, created_at from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOrNick, nomeOrNick,
	)
	if erro != nil {
		return nil, erro
	}

	defer resutlUserQuery.Close()

	var usuarios []model.Usuario

	for resutlUserQuery.Next() {
		var usuario model.Usuario

		if erro = resutlUserQuery.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.Created_at,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// Buscar por ID
func (repository Usuarios) BuscarPorID(ID int64) (model.Usuario, error) {
	resutlUserQuery, erro := repository.db.Query(
		"select id, nome, nick, email, created_at from usuarios where id = ?",
		ID,
	)
	if erro != nil {
		return model.Usuario{}, erro
	}
	defer resutlUserQuery.Close()

	var usuario model.Usuario

	if resutlUserQuery.Next() {
		if erro = resutlUserQuery.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.Created_at,
		); erro != nil {
			return model.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Actualizar usuario
func (repository Usuarios) Actualizar(ID uint64, usuario model.Usuario) error {
	statement, erro := repository.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}

	return nil
}

// Borrar por ID
func (repository Usuarios) Borrar(ID uint64) error {
	statement, erro := repository.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// Buscar por email
func (repository Usuarios) BuscarPorEmail(email string) (model.Usuario, error) {
	result, erro := repository.db.Query("select id, pass from usuarios where email = ?", email)
	if erro != nil {
		return model.Usuario{}, erro
	}
	defer result.Close()

	var usuario model.Usuario

	if result.Next() {
		if erro = result.Scan(&usuario.ID, &usuario.Pass); erro != nil {
			return model.Usuario{}, erro
		}
	}

	return usuario, nil
}
