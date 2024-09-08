package seguridad

import "golang.org/x/crypto/bcrypt"

// HASH recibe un string y lo hashea
func Hash(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
}

// VERIFICAR recibe un password y lo compara con un password con hash
func VerificarPass(passString, passWithHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(passString), []byte(passWithHash))
}
