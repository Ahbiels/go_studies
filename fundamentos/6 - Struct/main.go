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
	owner person
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
	p.name = "Angelo Gabriel"
	fmt.Printf("O usuário %v tem %v anos e mora em %v\n\n", p.name, p.age, p.city)

	c1 := cell {
		"Sansumg",
		"A35",
		person{
			"Angelo",
			21,
			"Carapebus",
		},
		"Grey",
		22997225096,
	}

	fmt.Println(c1)
	fmt.Printf("O usuário %v tem um celular %v do modelo %v e da cor %v", c1.owner.name, c1.brand ,c1.model, c1.color)

}

 