package main

import (
	"errors"
	"fmt"
)

func divisor(x,y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("WARNING: Divisor não pode ser zero, escolha outro número")
	}
	// nil é o nulo do tipo error
	return float64(x / y), nil
}

func main() {
	fmt.Println("Testando erros")
	var x,y float64 = 50, 0
	result, err := divisor(x,y)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v dividido por %v é igual a %v\n", x,y,result)
	}
}