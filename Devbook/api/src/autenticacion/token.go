package autenticacion

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

// Validar token verifica si el token es valido
func ValidarToken(r *http.Request) error {
	tokenString := extraerToken(r)
	token, erro := jwt.Parse(tokenString, retornarVerificationKey)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token inválido")
}

func ExtrerUsuarioID(r *http.Request) (uint64, error) {
	tokenString := extraerToken(r)
	token, erro := jwt.Parse(tokenString, retornarVerificationKey)
	if erro != nil {
		return 0, erro
	}

	if permisos, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permisos["usuarioId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}

		return usuarioID, nil
	}

	return 0, errors.New("token Invalid")
}

func extraerToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func retornarVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de asignatura inesperado! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
