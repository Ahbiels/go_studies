package main

import "fmt"

func recuperarPrograma() {
	if r := recover(); r != nil {
		fmt.Println("Programa recuperado")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarPrograma()
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A MÉDIA É EXATAMENTE 6")
}

func div(n1,n2 float32) {
	defer recuperarPrograma()
	if n2 == 0 {
		panic("DIVISOR É UM ZERO")
	}
	result := n1/n2
	fmt.Println(result)
}

func main() {
	// fmt.Println(alunoEstaAprovado(6, 6))
	// fmt.Println("Hello, World!!")
	div(float32(0),float32(0))
}
