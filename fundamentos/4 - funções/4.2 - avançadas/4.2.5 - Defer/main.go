package main

import "fmt"

func func1() {
	fmt.Println("Executando a função 1")
}
func func2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média Calculada")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false
}

func main() {
	defer func1() //adiar até o último momento possível
	func2()

	alunoEstaAprovado(7,8)
}
