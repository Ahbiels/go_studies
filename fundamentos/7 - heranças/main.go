package main

import (
	"fmt"
)

type person struct {
	name string
	age uint8
	city string
}

type cell struct {
	brand string
	model string
	person
	color string
	tel int64
}

func main() {
	fmt.Println("Hello, World!")
	var p person = person{
		name: "Angelo",
		age: 21,
		city: "Carapebus",
	}
	c1 := cell {
		"Sansumg",
		"A35",
		p,
		"cinza",
		22997225096,
	}

	fmt.Println(c1)
	fmt.Printf("O usuário %v tem um celular %v %v, %v anos e mora em %v", c1.name, c1.brand, c1.color, c1.age, c1.city)
}

 