package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// String conección conección con MYSQL
	StringConeccionDB = ""

	// Puerto de la API
	Puerto = 0
)

// CARGAR va a inicializar las varibales de ambiente
func Cargar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Puerto, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Puerto = 9000
	}

	StringConeccionDB = fmt.Sprintf("%s:$s/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"))
}
