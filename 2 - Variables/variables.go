package main

import "fmt"

func main() {
	var variable1 string = "variable 1"
	variable2 := "Variable 2"
	fmt.Println(variable1)
	fmt.Println(variable2)

	var (
		variable3 string = "lalala"
		variable4 string = "lalala"
	)
	fmt.Println(variable3, variable4)

	variable5, variable6 := "variable 5", "variable 6"
	fmt.Println(variable5, variable6)

	const constante1 string = "constante 1"
	fmt.Println(constante1)
}