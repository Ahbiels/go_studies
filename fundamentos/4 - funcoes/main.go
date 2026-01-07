package main

import (
	"fmt"
)

func calculadora(x1,x2 int8) (int8, int8, int16, float32) {
	fmt.Println(x1,x2)

	sum := x1+x2
	sub := x1-x2
	mult := int16(x1) * int16(x2)
	div := float32(x1) / float32(x2)

	return sum, sub, mult, div
	
}

func main() {
	fmt.Println("Hello, World!")
	sum, sub, mult, div := calculadora(100,20)
	fmt.Printf("%T - %v\n", sum, sum)
	fmt.Printf("%T - %v\n", sub, sub)
	fmt.Printf("%T - %v\n", mult, mult)
	fmt.Printf("%T - %v\n", div, div)

	result := func(name string) string {
		return "Hello, World, " + name
	}
	
	fmt.Println(result("Angelo"))
}