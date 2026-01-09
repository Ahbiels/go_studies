package main

import (
	"fmt"
)

// ponteiro -> salva o endereço de alguma coisa em memória

func main() {
	var value1 = 10
	var value2 = value1

	fmt.Println(value1, value2)
	value1++
	fmt.Println(value1, value2)

	var value3_pointer = &value2
	var value4_pointer *int
	value4_pointer = &value1

	value5_pointer := &value1

	value2++
	value2++

	fmt.Println(value3_pointer) //endereço na memoria
	fmt.Println(*value3_pointer) //Desreferenciação
	fmt.Printf("%T, %v\n\n",*value3_pointer,*value3_pointer) //Visualizando o tipo
	fmt.Println(*value4_pointer) //Desreferenciação
	fmt.Println(*value5_pointer) //Desreferenciação
	
}