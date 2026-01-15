package main

import "fmt"

func invertSinalComPonteiro(value *int) {
	// fmt.Println(*value) //lendo o valor da variável que está no endereço de memoria
	*value = *value * -1
}

func main() {
	value1 := 10
	value2 := new(int)
	*value2 = 20
	
	invertSinalComPonteiro(&value1) //passando o endereço na memoria onde a variável está armazenada
	invertSinalComPonteiro(value2) //passando o endereço na memoria onde a variável está armazenada
	
	fmt.Println(value1) //0xc000092020
	fmt.Println(*value2)  //0xc000092020

}
