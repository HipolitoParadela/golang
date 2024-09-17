package model

// Pass representa el formato del body de la actualización de la contraseña
type Pass struct {
	Nueva  string `json:nueva`
	Actual string `json:actual`
}
