package main

import (
	"fmt"
	"linea-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Punto de partida")

	aplicacao := app.Gerar()
	erro := aplicacao.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}
}
