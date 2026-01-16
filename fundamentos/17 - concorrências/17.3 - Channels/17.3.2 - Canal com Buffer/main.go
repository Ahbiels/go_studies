package main

import "fmt"

func main() {
	canal := make(chan string,3)
	canal <- "Hello, World!"
	canal <- "Hello, World!"
	canal <- "Hello, World!"
	

	mensagem1 := <- canal
	mensagem2 := <- canal
	mensagem3 := <- canal

	fmt.Println(mensagem1)
	fmt.Println(mensagem2)
	fmt.Println(mensagem3)
}
