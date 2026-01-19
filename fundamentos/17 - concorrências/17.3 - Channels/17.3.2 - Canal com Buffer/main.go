package main

import "fmt"

func main() {
	canal := make(chan string,3) //buffer com 3 canais configurados
	canal <- "Hello, World!"
	canal <- "Hello, World!"
	canal <- "Hello, World!"
	// se eu criar mais um canal, vai dar um erro de deadlock
	

	mensagem1 := <- canal
	mensagem2 := <- canal
	mensagem3 := <- canal

	fmt.Println(mensagem1)
	fmt.Println(mensagem2)
	fmt.Println(mensagem3)
}
