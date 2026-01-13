package main

import (
	"fmt"
)

func calculosMatematicos(n1, n2 int) (soma, sub int){
	soma = n1 + n2
	sub = n1 - n2
	return
}

func main(){
	// fmt.Println("Hello, World!")
	soma, sub := calculosMatematicos(10, 20)
	fmt.Printf("Soma: %v\nSubtração: %v\n\n", soma, sub)
}