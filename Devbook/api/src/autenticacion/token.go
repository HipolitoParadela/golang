package autenticacion

import (
	"api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CREA UN TOKEN
func CrearToken(usuarioID uint64) (string, error) {
	permisos := jwt.MapClaims{}
	permisos["authorized"] = true
	permisos["exp"] = time.Now().Add(time.Hour + 6).Unix() // 12231432523
	permisos["usuarioId"] = usuarioID
	//secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permisos)

	return token.SignedString([]byte(config.SecretKey)) // secret
}
