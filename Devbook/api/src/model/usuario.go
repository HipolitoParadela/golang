package model

import "time"

// Usuario represta un usuario de la plataforma
type Usuario struct {
	ID         uint64    `json:"id,omitempty"`
	Nome       string    `json:"nome,omitempty"`
	Nick       string    `json:"nick,omitempty"`
	Email      string    `json:"email,omitempty"`
	Pass       string    `json:"pass,omitempty"`
	Created_at time.Time `json:"created_at,omitempty"`
}
