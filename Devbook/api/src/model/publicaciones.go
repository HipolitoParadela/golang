package model

import (
	"errors"
	"strings"
	"time"
)

type Publicaciones struct {
	ID         uint64    `json:"id,omitempty"`
	Titulo     string    `json:"titulo,omitempty"`
	Contenido  string    `json:"contenido,omitempty"`
	AutorID    uint64    `json:"autor_id,omitempty"`
	AutorNick  string    `json:"autor_nick,omitempty"`
	Curtidas   uint64    `json:"curtidas"`
	Created_at time.Time `json:"created_at,omitempty"`
}

// Prepara la informaci que venga en los datos
func (publicacion *Publicaciones) Preparar() error {
	if erro := publicacion.validar(); erro != nil {
		return erro
	}

	publicacion.formatear()
	return nil
}

// Valida que vengan con contenido los datos
func (publicacion *Publicaciones) validar() error {
	if publicacion.Titulo == "" {
		return errors.New("el título no puede estar en blanco")
	}

	if publicacion.Contenido == "" {
		return errors.New("el contenido no puede estar en blanco")
	}

	return nil
}

// Quita los espacios en blanco que puedan traer los textos
func (publicacion *Publicaciones) formatear() {
	publicacion.Titulo = strings.TrimSpace(publicacion.Titulo)
	publicacion.Contenido = strings.TrimSpace(publicacion.Contenido)
}
