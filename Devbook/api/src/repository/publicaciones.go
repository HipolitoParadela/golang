package repository

import (
	"api/src/model"
	"database/sql"
)

// Publicaciones representa un repositorio de publicaciones
type Publicaciones struct {
	db *sql.DB
}

// Crear un repositorio de publicaciones
func NuevoRepositorioDePublicaciones(db *sql.DB) *Publicaciones {
	return &Publicaciones{db}
}

// Crear un publicacion en base de datos
func (repositorio Publicaciones) Crear(publicacion model.Publicaciones) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into publicaciones (tittlr, contenido, autor_id) values (?, ?, ?)",
	)

	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacion.Titulo, publicacion.Contenido, publicacion.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInsertado, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInsertado), nil
}

// Buscar publicación por ID
func (repository Publicaciones) BuscarPorID(ID uint64) (model.Publicaciones, error) {
	resutlUserQuery, erro := repository.db.Query(`
		select p.*, u.nick 
		from publicaciones p 
		inner join usuarios u on u.id = p.autor_id
		where id = ?
		`, ID,
	)
	if erro != nil {
		return model.Publicaciones{}, erro
	}
	defer resutlUserQuery.Close()

	var publicacion model.Publicaciones

	if resutlUserQuery.Next() {
		if erro = resutlUserQuery.Scan(
			&publicacion.ID,
			&publicacion.Titulo,
			&publicacion.Contenido,
			&publicacion.AutorID,
			&publicacion.Curtidas,
			&publicacion.Created_at,
			&publicacion.AutorNick,
		); erro != nil {
			return model.Publicaciones{}, erro
		}
	}

	return publicacion, nil
}

// Listar publicaciones de un usuario y propias
func (repository Publicaciones) ListarPublicacionesPorId(usuarioID uint64) ([]model.Publicaciones, error) {

	resutlUserQuery, erro := repository.db.Query(
		`select distinct p.*, u.nick 
		from publicaciones p 
		inner join usuarios u on u.id = p. autor_id 
		inner join seguidores s on p.autor_id
		where u.id = ? or s.seguidor_id = ?
		order by 1 desc`,
		usuarioID, usuarioID,
	)
	if erro != nil {
		return nil, erro
	}

	defer resutlUserQuery.Close()

	var publicaciones []model.Publicaciones

	for resutlUserQuery.Next() {
		var publicacion model.Publicaciones

		if erro = resutlUserQuery.Scan(
			&publicacion.ID,
			&publicacion.Titulo,
			&publicacion.Contenido,
			&publicacion.AutorID,
			&publicacion.Curtidas,
			&publicacion.Created_at,
			&publicacion.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicaciones = append(publicaciones, publicacion)
	}

	return publicaciones, nil
}

// Actualizar publicacion
func (repository Publicaciones) Actualizar(ID uint64, publicacion model.Publicaciones) error {
	statement, erro := repository.db.Prepare(
		"update publicaciones set titulo = ?, contenido = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacion.Titulo, publicacion.Contenido, ID); erro != nil {
		return erro
	}

	return nil
}

// Borrar por ID
func (repository Publicaciones) Borrar(ID uint64) error {
	statement, erro := repository.db.Prepare("delete from publicaciones where id = ?")
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
func (repository Publicaciones) BuscarPublicacionesUsuario(usuarioID uint64) ([]model.Publicaciones, error) {

	resutlUserQuery, erro := repository.db.Query(`select distinct p.*, u.nick 
		from publicaciones p 
		inner join usuarios u on u.id = p. autor_id 
		inner join seguidores s on p.autor_id
		where u.id = ?
		order by 1 desc`,
		usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer resutlUserQuery.Close()

	var publicaciones []model.Publicaciones

	for resutlUserQuery.Next() {
		var publicacion model.Publicaciones

		if erro = resutlUserQuery.Scan(
			&publicacion.ID,
			&publicacion.Titulo,
			&publicacion.Contenido,
			&publicacion.AutorID,
			&publicacion.Curtidas,
			&publicacion.Created_at,
			&publicacion.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicaciones = append(publicaciones, publicacion)
	}

	return publicaciones, nil
}

// Adiciona una curtida a la publicación
func (repository Publicaciones) CurtirPublicacion(publicacionID uint64) error {
	statement, erro := repository.db.Prepare(
		"update publicaciones set curtidas = curtidas + 1 where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacionID); erro != nil {
		return erro
	}

	return nil
}

// DEJAR DE SEGUIR UN USUARIO
func (repository Publicaciones) DescurtirPublicacion(publicacionID uint64) error {
	statement, erro := repository.db.Prepare(
		"update publicaciones set curtidas = curtidas - 1 where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacionID); erro != nil {
		return erro
	}

	return nil
}
