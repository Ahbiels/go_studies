package main

import (
	"fmt"
)

func main () {
	fmt.Println("Hello, world!")

	verdadeiro, falso := true,false

	fmt.Println(verdadeiro && falso) // and
	fmt.Println(verdadeiro || falso) // or
	fmt.Println(!verdadeiro) //not
	fmt.Println(verdadeiro != falso) //xor

	var value int = 10

	value++
	value--

	value += 10
	value -= 5
	value /= 5
	value *= 3
	value %= 3

	fmt.Println(value)

}