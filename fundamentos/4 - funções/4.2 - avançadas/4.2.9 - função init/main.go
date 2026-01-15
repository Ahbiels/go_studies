package main

import "fmt"

var x int

func init(){ // pode ter uma função init por arquivo, e não por pacote
	fmt.Println("Executando a função Init")
	x = 10
}

func main(){
	fmt.Println("Executando a função Main")
	fmt.Println(x)
}