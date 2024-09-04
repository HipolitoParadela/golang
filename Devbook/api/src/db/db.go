package db

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Conección con base de datos
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConeccionDB)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
