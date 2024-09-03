package main

import "fmt"

func fibonacci(posicion uint) uint {
	if posicion <= 1 {
		return posicion
	}

	return fibonacci(posicion-2) + fibonacci(posicion-1)
}

func main() {
	fmt.Println("Funciones Recursivas")

	// 1 1 2 3 5 8 13 21

	posicion := uint(50)

	for i := uint(1); i <= posicion; i++ {
		fmt.Println(fibonacci(i))
	}
}
