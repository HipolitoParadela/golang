package model

import (
	"api/src/seguridad"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuario represta un usuario de la plataforma
type Usuario struct {
	ID         uint64    `json:"id,omitempty"`
	Nome       string    `json:"nome,omitempty"`
	Nick       string    `json:"nick,omitempty"`
	Email      string    `json:"email,omitempty"`
	Pass       string    `json:"pass,omitempty"`
	Created_at time.Time `json:"created_at,omitempty"`
}

// Preparar va a llamar los metodos para validar y formatear el usuario recibido
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatear(etapa); erro != nil {
		return erro
	}

	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("el nombre no puede estar en blanco")
	}

	if usuario.Nick == "" {
		return errors.New("el Nick no puede estar en blanco")
	}

	if usuario.Email == "" {
		return errors.New("el email no puede estar en blanco")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("el email no tiene un formato válido")
	}

	if etapa == "registro" && usuario.Pass == "" {
		return errors.New("la contraseña no puede estar en blanco")
	}

	return nil
}

func (usuario *Usuario) formatear(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "registro" {
		passWithHash, erro := seguridad.Hash(usuario.Pass)
		if erro != nil {
			return erro
		}

		usuario.Pass = string(passWithHash)
	}

	return nil
}
