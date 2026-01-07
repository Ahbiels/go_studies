package main

import (
	"fmt"
)

var idade int

func main() {
	var nome string
	nome = "Gabriel"
	nome = "Angelo"
	var sobrenome string = "Gabriel"

	nome_completo := nome + " Gabriel"

	idade := 21

	fmt.Printf("%v \n", nome)
	fmt.Printf("%v \n", nome_completo)
	fmt.Printf("%v \n", idade)

	nome, sobrenome = sobrenome, nome
	println(nome)

	var (
		var1 = 1
		var2 = 2
		var3 = 3
	)

	fmt.Println(var1, var2, var3)

	var is_adult bool = true
	fmt.Println(is_adult)

	const value int16 = 233
	const value2 uint8 = 233
	fmt.Printf("%T %v \n",value, value)
	fmt.Printf("%T %v \n",value2, value2)
}